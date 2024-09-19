package types

import (
	"time"
)

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

type QueryRunnerJob struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	QueryId   string
	CreatedBy string
	Status    string
}
