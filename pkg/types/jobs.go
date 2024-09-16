package types

import "time"

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
