package dto

import "time"

type ItemDto struct {
	ItemCode     string `json:"item_code"`
	ItemName     string `json:"item_name"`
	Unit         string `json:"unit"`
	ItemGroup    string `json:"item_group"`
	ItemCategory string `json:"item_category"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
