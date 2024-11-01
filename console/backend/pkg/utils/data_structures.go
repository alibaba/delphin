package utils

import (
	morphlingv1alpha1 "github.com/alibaba/morphling/api/v1alpha1"
	"time"
)

type MorphlingConfig struct {
	Namespace       string   `json:"namespace"`
	HttpClientImage string   `json:"http-client-image"`
	HsfClientImage  string   `json:"hsf-client-image"`
	HttpClientYaml  string   `json:"http-client-yaml"`
	HsfClientYaml   string   `json:"hsf-client-yaml"`
	HttpServiceYaml string   `json:"http-service-yaml"`
	HsfServiceYaml  string   `json:"hsf-service-yaml"`
	AlgorithmNames  []string `json:"algorithm-names"`
}

type ClusterTotalResources struct {
	TotalCPU    int64 `json:"totalCPU"`
	TotalMemory int64 `json:"totalMemory"`
	TotalGPU    int64 `json:"totalGPU"`
}

type ClusterRequestResource struct {
	RequestCPU    int64 `json:"requestCPU"`
	RequestMemory int64 `json:"requestMemory"`
	RequestGPU    int64 `json:"requestGPU"`
}

type NodeInfo struct {
	NodeName      string `json:"nodeName"`
	InstanceType  string `json:"instanceType"`
	GPUType       string `json:"gpuType"`
	TotalCPU      int64  `json:"totalCPU"`
	TotalMemory   int64  `json:"totalMemory"`
	TotalGPU      int64  `json:"totalGPU"`
	RequestCPU    int64  `json:"requestCPU"`
	RequestMemory int64  `json:"requestMemory"`
	RequestGPU    int64  `json:"requestGPU"`
}

type NodeInfoList struct {
	Items []NodeInfo `json:"items,omitempty"`
}

type ProfilingExperimentInfo struct {
	Name               string                                   `json:"name"`
	ExperimentUserID   string                                   `json:"UserId,omitempty"`
	ExperimentUserName string                                   `json:"UserName,omitempty"`
	ExperimentStatus   morphlingv1alpha1.ProfilingConditionType `json:"peStatus"`
	Namespace          string                                   `json:"namespace"`
	CreateTime         string                                   `json:"createTime"`
	EndTime            string                                   `json:"endTime"`
	DurationTime       string                                   `json:"durationTime"`
}

type ParameterSpec struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Space    string `json:"space"`
}

//type ParameterSample struct {
//	Name  string `json:"name"`
//	Value string `json:"value"`
//}

type TrialSpec struct {
	Name             string            `json:"name"`
	Status           string            `json:"Status"`
	ObjectiveName    string            `json:"objectiveName"`
	ObjectiveValue   string            `json:"objectiveValue"`
	ParameterSamples map[string]string `json:"parameterSamples"`
	CreateTime       string            `json:"createTime"`
}

type CurrentOptimalTrial struct {
	ObjectiveName    string            `json:"objectiveName"`
	ObjectiveValue   string            `json:"objectiveValue"`
	ParameterSamples map[string]string `json:"parameterSamples"`
}

type ProfilingExperimentDetail struct {
	Name               string                                   `json:"name"`
	ExperimentUserID   string                                   `json:"UserId,omitempty"`
	ExperimentUserName string                                   `json:"UserName,omitempty"`
	ExperimentStatus   morphlingv1alpha1.ProfilingConditionType `json:"peStatus"`
	Namespace          string                                   `json:"namespace"`
	CreateTime         string                                   `json:"createTime"`
	EndTime            string                                   `json:"endTime"`
	DurationTime       string                                   `json:"durationTime"`

	TrialsTotal          int32                 `json:"trialsTotal"`
	TrialsSucceeded      int32                 `json:"trialsSucceeded"`
	AlgorithmName        string                `json:"algorithmName"`
	MaxNumTrials         int32                 `json:"maxNumTrials"`
	Objective            string                `json:"objective"`
	Parallelism          int32                 `json:"parallelism"`
	Parameters           []ParameterSpec       `json:"parameters"`
	Trials               []TrialSpec           `json:"trials"`
	CurrentOptimalTrials []CurrentOptimalTrial `json:"currentOptimalTrials"`
}

type Query struct {
	Name      string
	Namespace string
	//Region     string
	Status     morphlingv1alpha1.ProfilingConditionType
	StartTime  time.Time
	EndTime    time.Time
	Pagination *QueryPagination
}

type QueryPagination struct {
	PageNum  int
	PageSize int
	Count    int
}

type GitHubRepoInfo struct {
	Owner         string `json:"owner"`
	Repo          string `json:"repo"`
	Branch        string `json:"branch"`
	AccessToken   string `json:"accessToken"`
}


type LLMServiceVersion struct {
	ModelName string `json:"modelName"`
	Version   string `json:"version"`
	CreationTime string `json:"creationTime"`
	AssociatedExperimentSpec morphlingv1alpha1.ProfilingExperimentSpec `json:"associatedExperimentSpec"`
}

type LLMServiceVersionRequest struct {
	GitHubRepoInfo 	GitHubRepoInfo 	`json:"gitHubRepoInfo"`
	LLMServiceVersion 	LLMServiceVersion `json:"llmServiceVersion"`
}
