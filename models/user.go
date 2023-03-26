package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"not null;unique;type:varchar(100)" json:"email"`
	Products  []Product `json:"products"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
