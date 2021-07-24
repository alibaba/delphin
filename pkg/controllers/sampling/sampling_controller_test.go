/*
Copyright 2021 The Alibaba Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sampling

import (
	"encoding/json"
	stdlog "log"
	"os"
	"path/filepath"

	"sync"
	"testing"
	"time"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	morphlingv1alpha1 "github.com/alibaba/morphling/api/v1alpha1"
	"github.com/alibaba/morphling/pkg/controllers/consts"
)

var c client.Client

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "foo", Namespace: "default"}}
var depKey = types.NamespacedName{Name: "foo-grid", Namespace: "default"}

const timeout = time.Second * 50

var cfg *rest.Config

func TestMain(m *testing.M) {
	t := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "..", "config", "crd", "bases"),
			filepath.Join("..", "..", "..", "config", "crd", "patches"),
		},
	}
	morphlingv1alpha1.AddToScheme(scheme.Scheme)

	var err error
	if cfg, err = t.Start(); err != nil {
		stdlog.Fatal(err)
	}

	code := m.Run()
	t.Stop()
	os.Exit(code)
}

// SetupTestReconcile returns a reconcile.Reconcile implementation that delegates to inner and
// writes the request to requests after Reconcile is finished.
func SetupTestReconcile(inner reconcile.Reconciler) (reconcile.Reconciler, chan reconcile.Request) {
	requests := make(chan reconcile.Request)
	fn := reconcile.Func(func(req reconcile.Request) (reconcile.Result, error) {
		result, err := inner.Reconcile(req)
		requests <- req
		return result, err
	})
	return fn, requests
}

// StartTestManager adds recFn
func StartTestManager(mgr manager.Manager, g *gomega.GomegaWithT) (chan struct{}, *sync.WaitGroup) {
	stop := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		g.Expect(mgr.Start(stop)).NotTo(gomega.HaveOccurred())
	}()
	return stop, wg
}
func init() {
	logf.SetLogger(logf.ZapLogger(true))
}

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &morphlingv1alpha1.Sampling{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: morphlingv1alpha1.SamplingSpec{
			Algorithm: morphlingv1alpha1.AlgorithmSpec{
				AlgorithmName:     "grid",
				AlgorithmSettings: []morphlingv1alpha1.AlgorithmSetting{morphlingv1alpha1.AlgorithmSetting{Name: "initpoints", Value: "5"}},
			},
			//Objective: morphlingv1alpha1.ObjectiveSpec{
			//	Type:                morphlingv1alpha1.ObjectiveTypeMaximize,
			//	ObjectiveMetricName: "qps",
			//},
			NumSamplingsRequested: 1,
		},
	}
	configMap := newConfigMapInstance()

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	c = mgr.GetClient()

	recFn, requests := SetupTestReconcile(newReconciler(mgr))
	g.Expect(add(mgr, recFn)).NotTo(gomega.HaveOccurred())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	// Create the Sampling object and expect the Reconcile and Deployment to be created
	err = c.Create(context.TODO(), instance)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	err = c.Create(context.TODO(), configMap)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	// The instance object may not be a valid object because it might be missing some required fields.
	// Please modify the instance object by adding required fields and then remove the following if statement.
	if apierrors.IsInvalid(err) {
		t.Logf("failed to create object, got an invalid object error: %v", err)
		return
	}
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer c.Delete(context.TODO(), instance)
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	deploy := &appsv1.Deployment{}
	g.Eventually(func() error { return c.Get(context.TODO(), depKey, deploy) }, timeout).
		Should(gomega.Succeed())

	// Delete the Deployment and expect Reconcile to be called for Deployment deletion
	g.Expect(c.Delete(context.TODO(), deploy)).NotTo(gomega.HaveOccurred())
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))
	g.Eventually(func() error { return c.Get(context.TODO(), depKey, deploy) }, timeout).
		Should(gomega.Succeed())

	// Manually delete Deployment since GC isn't enabled in the test control plane
	g.Eventually(func() error { return c.Delete(context.TODO(), deploy) }, timeout).
		Should(gomega.MatchError("deployments.apps \"foo-grid\" not found"))

}

func newConfigMapInstance() *corev1.ConfigMap {
	samplingConfig := map[string]map[string]string{
		"grid": {"image": consts.ImageSamplingAlgorithmRandom},
	}
	b, _ := json.Marshal(samplingConfig)
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      consts.MorphlingConfigMapName,
			Namespace: consts.DefaultMorphlingNamespace,
		},
		Data: map[string]string{
			"sampling": string(b),
		},
	}
}
