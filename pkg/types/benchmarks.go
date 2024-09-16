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
	Metadata Metadata    `json:"metadata"`
	Findings interface{} `json:"findings"` // need to update with actual type
}
