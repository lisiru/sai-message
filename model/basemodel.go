package model

import "time"

type BaseColumn struct {
	ID        uint64    `json:"id,omitempty" gorm:"primary_key;AUTO_INCRMENT;column:id"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
