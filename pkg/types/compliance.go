package types

import "time"

type ComplianceSummaryOfIntegrationRequest struct {
	BenchmarkId string                `json:"benchmark_id"'`
	Integration IntegrationFilterInfo `json:"integration"`
	ShowTop     int                   `json:"show_top"`
}

type TopFiledRecordV2 struct {
	Field  string `json:"field"`
	Key    string `json:"key"`
	Issues int    `json:"issues"`
}

type ComplianceSummaryOfIntegrationResponse struct {
	BenchmarkID                string                             `json:"benchmark_id"`
	Integration                IntegrationInfo                    `json:"integration"`
	ComplianceScore            float64                            `json:"compliance_score"`
	SeveritySummaryByControl   BenchmarkControlsSeverityStatusV2  `json:"severity_summary_by_control"`
	SeveritySummaryByResource  BenchmarkResourcesSeverityStatusV2 `json:"severity_summary_by_resource"`
	FindingsSummary            ConformanceStatusSummaryV2         `json:"findings_summary"`
	IssuesCount                int                                `json:"issues_count"`
	TopResourcesWithIssues     []TopFiledRecordV2                 `json:"top_resources_with_issues"`
	TopResourceTypesWithIssues []TopFiledRecordV2                 `json:"top_resource_types_with_issues"`
	TopControlsWithIssues      []TopFiledRecordV2                 `json:"top_controls_with_issues"`
	LastEvaluatedAt            *time.Time                         `json:"last_evaluated_at"`
	LastJobStatus              string                             `json:"last_job_status"`
	LastJobId                  string                             `json:"last_job_id"`
}

type ComplianceSummaryOfBenchmarkRequest struct {
	Benchmarks []string `json:"benchmarks"`
	IsRoot     *bool    `json:"is_root"`
	ShowTop    int      `json:"show_top"`
}

type ComplianceSummaryOfBenchmarkResponse struct {
	BenchmarkID                string                             `json:"benchmark_id"`
	ComplianceScore            float64                            `json:"compliance_score"`
	SeveritySummaryByControl   BenchmarkControlsSeverityStatusV2  `json:"severity_summary_by_control"`
	SeveritySummaryByResource  BenchmarkResourcesSeverityStatusV2 `json:"severity_summary_by_resource"`
	FindingsSummary            ConformanceStatusSummaryV2         `json:"findings_summary"`
	IssuesCount                int                                `json:"issues_count"`
	TopIntegrations            []TopIntegration                   `json:"top_integrations"`
	TopResourcesWithIssues     []TopFiledRecordV2                 `json:"top_resources_with_issues"`
	TopResourceTypesWithIssues []TopFiledRecordV2                 `json:"top_resource_types_with_issues"`
	TopControlsWithIssues      []TopFiledRecordV2                 `json:"top_controls_with_issues"`
	LastEvaluatedAt            *time.Time                         `json:"last_evaluated_at"`
	LastJobStatus              string                             `json:"last_job_status"`
	LastJobId                  string                             `json:"last_job_id"`
}

type TopIntegration struct {
	IntegrationInfo IntegrationInfo `json:"integration_info"`
	Issues          int             `json:"issues"`
}

type ConformanceStatusSummaryV2 struct {
	TotalCount  int `json:"total_count"`
	PassedCount int `json:"passed"`
	FailedCount int `json:"failed"`
}

type BenchmarkStatusResultV2 struct {
	TotalCount  int `json:"total"`
	PassedCount int `json:"passed"`
	FailedCount int `json:"failed"`
}

type BenchmarkControlsSeverityStatusV2 struct {
	Total BenchmarkStatusResultV2 `json:"total"`

	Critical BenchmarkStatusResultV2 `json:"critical"`
	High     BenchmarkStatusResultV2 `json:"high"`
	Medium   BenchmarkStatusResultV2 `json:"medium"`
	Low      BenchmarkStatusResultV2 `json:"low"`
	None     BenchmarkStatusResultV2 `json:"none"`
}

type BenchmarkResourcesSeverityStatusV2 struct {
	Total BenchmarkStatusResultV2 `json:"total"`

	Critical BenchmarkStatusResultV2 `json:"critical"`
	High     BenchmarkStatusResultV2 `json:"high"`
	Medium   BenchmarkStatusResultV2 `json:"medium"`
	Low      BenchmarkStatusResultV2 `json:"low"`
	None     BenchmarkStatusResultV2 `json:"none"`
}
