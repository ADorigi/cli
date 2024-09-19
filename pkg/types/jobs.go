package types

import "time"

type GetAsyncQueryRunJobStatusResponse struct {
	JobId          uint      `json:"job_id"`
	QueryId        string    `json:"query_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedBy      string    `json:"created_by"`
	JobStatus      string    `json:"job_status"`
	FailureMessage string    `json:"failure_message"`
}

type GetComplianceJobStatusResponse struct {
	JobId           uint              `json:"job_id"`
	IntegrationInfo []IntegrationInfo `json:"integration_info"`
	JobStatus       string            `json:"job_status"`
	BenchmarkId     string            `json:"benchmark_id"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

type ListJobsByTypeItem struct {
	JobId     string    `json:"job_id"`
	JobType   string    `json:"job_type"`
	JobStatus string    `json:"job_status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetDescribeJobStatusResponse struct {
	JobId           uint            `json:"job_id"`
	IntegrationInfo IntegrationInfo `json:"integration_info"`
	JobStatus       string          `json:"job_status"`
	DiscoveryType   string          `json:"discovery_type"`
	ResourceType    string          `json:"resource_type"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type GetAnalyticsJobStatusResponse struct {
	JobId     uint      `json:"job_id"`
	JobStatus string    `json:"job_status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
