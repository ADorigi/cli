package types

import "time"

type Metadata struct {
	ID               string            `json:"id"`
	Title            string            `json:"title"`
	Description      string            `json:"description"`
	Enabled          bool              `json:"enabled"`
	TrackDriftEvents bool              `json:"track_drift_events"`
	PrimaryTables    []string          `json:"primary_tables"`
	Tags             map[string]string `json:"tags"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

type BenchMark struct {
	Metadata Metadata    `json:"metadata"`
	Findings interface{} `json:"findings"` // need to update with actual type
}

type GetBenchmarkSummaryV2Request struct {
	Integration          []IntegrationFilterInfo `json:"integration"`
	TopIntegrationsCount int                     `json:"top_integrations_count"`
}

type GetBenchmarkSummaryV2Response struct {
	ComplianceScore           float64                          `json:"compliance_score"`
	SeveritySummaryByControl  BenchmarkControlsSeverityStatus  `json:"severity_summary_by_control"`
	SeveritySummaryByResource BenchmarkResourcesSeverityStatus `json:"severity_summary_by_resource"`
	TopIntegrations           []TopIntegration                 `json:"top_connections"`
	FindingsSummary           ConformanceStatusSummary         `json:"findings_summary"`
	EvaluatedAt               *time.Time                       `json:"evaluatedAt" example:"2020-01-01T00:00:00Z"`
	LastJobStatus             string                           `json:"lastJobStatus" example:"success"`
}

type TopIntegration struct {
	IntegrationInfo IntegrationInfo `json:"integration_info"`
	Issues          int             `json:"issues"`
}

type ConformanceStatusSummary struct {
	PassedCount int `json:"passed"`
	FailedCount int `json:"failed"`
}

type BenchmarkStatusResult struct {
	PassedCount int `json:"passed"`
	TotalCount  int `json:"total"`
}

type BenchmarkResourcesSeverityStatus struct {
	Total BenchmarkStatusResult `json:"total"`

	Critical BenchmarkStatusResult `json:"critical"`
	High     BenchmarkStatusResult `json:"high"`
	Medium   BenchmarkStatusResult `json:"medium"`
	Low      BenchmarkStatusResult `json:"low"`
	None     BenchmarkStatusResult `json:"none"`
}

type BenchmarkControlsSeverityStatus struct {
	Total BenchmarkStatusResult `json:"total"`

	Critical BenchmarkStatusResult `json:"critical"`
	High     BenchmarkStatusResult `json:"high"`
	Medium   BenchmarkStatusResult `json:"medium"`
	Low      BenchmarkStatusResult `json:"low"`
	None     BenchmarkStatusResult `json:"none"`
}
