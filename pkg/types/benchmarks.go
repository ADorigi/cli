package types

import "time"

type Metadata struct {
	ID               string              `json:"id"`
	Title            string              `json:"title"`
	Description      string              `json:"description"`
	Connectors       []string            `json:"connectors"`
	NumberOfControls int                 `json:"number_of_controls"`
	Enabled          bool                `json:"enabled"`
	TrackDriftEvents bool                `json:"track_drift_events"`
	PrimaryTables    []string            `json:"primary_tables"`
	Tags             map[string][]string `json:"tags"`
	CreatedAt        time.Time           `json:"created_at"`
	UpdatedAt        time.Time           `json:"updated_at"`
}

type BenchMark struct {
	Metadata Metadata    `json:"benchmark"`
	Findings interface{} `json:"findings"` // need to update with actual type
}

type GetBenchmarksResponse struct {
	Items      []BenchMark `json:"items"`
	TotalCount int64       `json:"total_count"`
}

type GetBenchmarkPayload struct {
	Cursor                 int  `json:"cursor"`
	PerPage                int  `json:"per_page"`
	OnlyRootBenchmark      bool `json:"root"`
	IncludeFindingsSummary bool `json:"finding_summary"`
}
