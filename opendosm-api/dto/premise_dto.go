package dto

import "time"

type PremiseDto struct {
	PremiseCode int64  `json:"premise_code"`
	PremiseName string `json:"premise_name"`
	Address     string `json:"address"`
	Type        string `json:"type"`
	District    string `json:"district"`
	State       string `json:"state"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
