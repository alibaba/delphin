package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/quota/v1"
	"math"
	"net/http"
	"strconv"
	"time"
)

func Succeed(c *gin.Context, obj interface{}) {
	c.JSONP(http.StatusOK, gin.H{
		"code": "200",
		"data": obj,
	})
}

func Failed(c *gin.Context, msg string) {
	c.JSONP(http.StatusOK, gin.H{
		"code": "300",
		"data": msg,
	})
	c.Set("failed", true)
}

// ComputePodSpecResourceRequest returns the requested resource of the PodSpec
func ComputePodSpecResourceRequest(spec *v1.PodSpec) v1.ResourceList {
	result := v1.ResourceList{}
	for _, container := range spec.Containers {
		result = quota.Add(result, container.Resources.Requests)
	}
	// take max_resource(sum_pod, any_init_container)
	for _, container := range spec.InitContainers {
		result = quota.Max(result, container.Resources.Requests)
	}
	// If Overhead is being utilized, add to the total requests for the pod
	if spec.Overhead != nil {
		result = quota.Add(result, spec.Overhead)
	}
	return result
}

func GetTimeDiffer(startTime time.Time, endTime time.Time) (differ string) {
	seconds := endTime.Sub(startTime).Seconds()
	var buffer bytes.Buffer
	hours := math.Floor(seconds / 3600)
	if hours > 0 {
		buffer.WriteString(strconv.FormatFloat(hours, 'g', -1, 64))
		buffer.WriteString("h")
		seconds = seconds - 3600*hours
	}
	minutes := math.Floor(seconds / 60)
	if minutes > 0 {
		buffer.WriteString(strconv.FormatFloat(minutes, 'g', -1, 64))
		buffer.WriteString("m")
		seconds = seconds - 60*minutes
	}
	buffer.WriteString(strconv.FormatFloat(seconds, 'g', 2, 64))
	buffer.WriteString("s")
	return buffer.String()
}


func generateTunableParametersYAML(params []morphlingv1alpha1.ParameterCategory) string {
    var result strings.Builder
    result.WriteString("  tunableParameters:\n")
    
    for _, category := range params {
        result.WriteString(fmt.Sprintf("    - category: \"%s\"\n", category.Category))
        result.WriteString("      parameters:\n")
        
        for _, param := range category.Parameters {
            result.WriteString(fmt.Sprintf("        - parameterType: %s\n", param.ParameterType))
            result.WriteString(fmt.Sprintf("          name: \"%s\"\n", param.Name))
            result.WriteString("          feasibleSpace:\n")
            
            if len(param.FeasibleSpace.List) > 0 {
                result.WriteString("            list:\n")
                for _, item := range param.FeasibleSpace.List {
                    result.WriteString(fmt.Sprintf("              - \"%s\"\n", item))
                }
            } else {
                if param.FeasibleSpace.Max != "" {
                    result.WriteString(fmt.Sprintf("            max: \"%s\"\n", param.FeasibleSpace.Max))
                }
                if param.FeasibleSpace.Min != "" {
                    result.WriteString(fmt.Sprintf("            min: \"%s\"\n", param.FeasibleSpace.Min))
                }
                if param.FeasibleSpace.Step != "" {
                    result.WriteString(fmt.Sprintf("            step: \"%s\"\n", param.FeasibleSpace.Step))
                }
            }
        }
    }
    
    return result.String()
}