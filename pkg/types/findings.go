package types

type FindingRequestFilters struct {
	BenchmarkIDs []string                `json:"benchmark_id"`
	Integrations []IntegrationFilterInfo `json:"integration"`
}

type FindingsRequestPayload struct {
	Filters FindingRequestFilters `json:"filters"`
}

type FindingsResponse struct {
	Findings   []Findings `json:"findings"`
	TotalCount int64      `json:"totalCount"`
}

type Findings struct {
	ID                        string      `json:"id"`
	BenchmarkID               string      `json:"benchmarkID"`
	ControlID                 string      `json:"controlID"`
	ConnectionID              string      `json:"connectionID"`
	EvaluatedAt               int64       `json:"evaluatedAt"`
	StateActive               bool        `json:"stateActive"`
	ConformanceStatus         string      `json:"conformanceStatus"`
	Severity                  string      `json:"severity"`
	Evaluator                 string      `json:"evaluator"`
	Connector                 string      `json:"connector"`
	KaytuResourceID           string      `json:"kaytuResourceID"`
	ResourceID                string      `json:"resourceID"`
	ResourceName              string      `json:"resourceName"`
	ResourceLocation          string      `json:"resourceLocation"`
	ResourceType              string      `json:"resourceType"`
	Reason                    string      `json:"reason"`
	CostOptimization          interface{} `json:"costOptimization"`
	ComplianceJobID           int64       `json:"complianceJobID"`
	ParentComplianceJobID     int64       `json:"parentComplianceJobID"`
	ParentBenchmarkReferences []string    `json:"parentBenchmarkReferences"`
	ParentBenchmarks          []string    `json:"parentBenchmarks"`
	LastEvent                 string      `json:"lastEvent"`
	ResourceTypeName          string      `json:"resourceTypeName"`
	ParentBenchmarkNames      []string    `json:"parentBenchmarkNames"`
	ControlTitle              string      `json:"controlTitle"`
	ProviderConnectionID      string      `json:"providerConnectionID"`
	ProviderConnectionName    string      `json:"providerConnectionName"`
}
