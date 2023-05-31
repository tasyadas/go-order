package models

import "time"

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `gorm:"not null;type:varchar(191)" json:"itemCode"`
	Description string `gorm:"type:text" json:"description"`
	Quantity    int    `gorm:"not null;" json:"quantity"`
	OrderID     uint
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
