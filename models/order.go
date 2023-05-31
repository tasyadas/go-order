package models

import "time"

type Order struct {
	OrderID      uint       `gorm:"primaryKey"`
	CustomerName string     `gorm:"not null;type:varchar(191)" json:"customerName"`
	Items        []Item     `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
	OrderedAt    *time.Time `json:"orderedAt"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}
