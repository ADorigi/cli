package types

type GetAsyncQueryRunResultResponse struct {
	RunId       string           `json:"runID"`
	QueryID     string           `json:"queryID"`
	Parameters  []QueryParameter `json:"parameters"`
	ColumnNames []string         `json:"columnNames"`
	CreatedBy   string           `json:"createdBy"`
	TriggeredAt int64            `json:"triggeredAt"`
	EvaluatedAt int64            `json:"evaluatedAt"`
	Result      [][]string       `json:"result"`
}

type QueryParameter struct {
	Key      string `json:"key" example:"key"`
	Required bool   `json:"required" example:"true"`
}
