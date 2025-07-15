package model

import (
	"time"
)

type ItemLookup struct {
	ID int64 `gorm:"primarykey"`

	ItemCode     string `gorm:"column:item_code"`
	ItemName     string `gorm:"column:item_name"`
	Unit         string `gorm:"column:unit"`
	ItemGroup    string `gorm:"column:item_group"`
	ItemCategory string `gorm:"column:item_category"`

	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (ItemLookup) TableName() string {
	return "item_lookups"
}
