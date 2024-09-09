package model

import "time"

// ArgoCDInstance ArgoCD instance table
type ArgoCDInstance struct {
	ID        uint       `json:"id" gorm:"primaryKey;index"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Endpoint  string     `json:"endpoint" gorm:"type:varchar(255);not null"`
	AuthToken string     `json:"auth_token" gorm:"type:varchar(255);not null"`
}
