package models

import "time"

type Tenant struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `json:"email" gorm:"not null;unique;default:null"`
	Password  string `json:"password" gorm:"not null;unique;default:null"`
	Name      string `json:"name" gorm:"not null;unique;default:null"`
	CreatedAt time.Time
}
