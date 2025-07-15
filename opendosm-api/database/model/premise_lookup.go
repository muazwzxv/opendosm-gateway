package model

import "time"

type PremiseLookup struct {
	ID int64 `gorm:"primarykey"`

	PremiseCode int64  `gorm:"column:premise_code"`
	PremiseName string `gorm:"column:premise_name"`
	Address     string `gorm:"column:address"`
	Type        string `gorm:"column:type"`
	District    string `gorm:"column:district"`
	State       string `gorm:"column:state"`

	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (PremiseLookup) TableName() string {
	return "premise_lookups"
}
