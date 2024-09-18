package types

type Parameter struct {
	Key      string `json:"key"`
	Required bool   `json:"required"`
}

type Query struct {
	PrimaryTable string      `json:"primary_table"`
	ListOfTables []string    `json:"list_of_tables"`
	Parameters   []Parameter `json:"parameters"`
}

type Tags struct {
	ScoreServiceName []string `json:"score_service_name"`
	ScoreTags        []string `json:"score_tags"`
}

type Control struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Connector       []string `json:"connector"`
	Severity        string   `json:"severity"`
	Tags            Tags     `json:"tags"`
	Query           Query    `json:"query"`
	FindingsSummary *string  `json:"findings_summary"`
}

type GetControlsResponse struct {
	Items      []Control `json:"items"`
	TotalCount int64     `json:"total_count"`
}

type FindingFilters struct {
	BenchmarkID []string `json:"benchmarkID"`
}

type GetControlsPayload struct {
	Cursor         int            `json:"cursor"`
	PerPage        int            `json:"per_page"`
	FindingFilters FindingFilters `json:"finding_filters"`
}
