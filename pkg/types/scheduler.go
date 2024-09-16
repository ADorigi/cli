package types

import (
	"strings"
	"time"
)

type RunBenchmarkByIdRequest struct {
	IntegrationInfo []IntegrationFilterInfo `json:"integration_info"`
}

type RunBenchmarkResponse struct {
	JobId           uint              `json:"job_id"`
	BenchmarkId     string            `json:"benchmark_id"`
	IntegrationInfo []IntegrationInfo `json:"integration_info"`
}

type IntegrationFilterInfo struct {
	Integration        *string `json:"integration"`
	Type               *string `json:"type"`
	ID                 *string `json:"id"`
	IDName             *string `json:"id_name"`
	IntegrationTracker *string `json:"integration_tracker"`
}

type IntegrationInfo struct {
	Integration        string `json:"integration"`
	Type               string `json:"type"`
	ID                 string `json:"id"`
	IDName             string `json:"id_name"`
	IntegrationTracker string `json:"integration_tracker"`
}

func ParseIntegrationInfo(infoString string) IntegrationFilterInfo {
	info := IntegrationFilterInfo{}
	// Split the input by commas
	pairs := strings.Split(infoString, ",")

	// Parse each key=value pair
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key := kv[0]
		value := kv[1]

		// Assign values to the appropriate struct fields
		switch key {
		case "integration":
			info.Integration = &value
		case "type":
			info.Type = &value
		case "id":
			info.ID = &value
		case "id_name":
			info.IDName = &value
		case "integration_tracker":
			info.IntegrationTracker = &value
		}
	}

	return info
}

type RunDiscoveryRequest struct {
	ResourceTypes   []string                `json:"resource_types"`
	ForceFull       bool                    `json:"force_full"` // force full discovery. only matters if ResourceTypes is empty
	IntegrationInfo []IntegrationFilterInfo `json:"integration_info"`
}

type RunDiscoveryResponse struct {
	JobId           uint            `json:"job_id"`
	ResourceType    string          `json:"resource_type"`
	Status          string          `json:"status"`
	FailureReason   string          `json:"failure_reason"`
	IntegrationInfo IntegrationInfo `json:"integration_info"`
}

type ListJobsByTypeRequest struct {
	JobType         string                  `json:"job_type"`
	Selector        string                  `json:"selector"`
	JobId           []string                `json:"job_id"`
	IntegrationInfo []IntegrationFilterInfo `json:"integration_info"`
	Status          []string                `json:"status"`
	Benchmark       []string                `json:"benchmark"`
	SortBy          JobSort                 `json:"sort_by"`
	SortOrder       JobSortOrder            `json:"sort_order"`
	UpdatedAt       struct {
		From *int64 `json:"from"`
		To   *int64 `json:"to"`
	} `json:"updated_at"`
	CreatedAt struct {
		From *int64 `json:"from"`
		To   *int64 `json:"to"`
	} `json:"created_at"`
	Cursor  *int64 `json:"cursor"`
	PerPage *int64 `json:"per_page"`
}

type JobSort string

const (
	JobSort_ByJobID        = "id"
	JobSort_ByJobType      = "job_type"
	JobSort_ByConnectionID = "connection_id"
	JobSort_ByStatus       = "status"
)

type JobSortOrder string

const (
	JobSortOrder_ASC  = "ASC"
	JobSortOrder_DESC = "DESC"
)

type ListJobsByTypeResponse struct {
	JobId     string    `json:"job_id"`
	JobType   string    `json:"job_type"`
	JobStatus string    `json:"job_status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
