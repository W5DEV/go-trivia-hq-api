package models

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Sources struct {
	ID            		uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Order				int				  `json:"order,omitempty"`
	Source      		string            `json:"source,omitempty"`
	Citation			string			  `json:"citation,omitempty"`
	Topic				string			  `json:"topic,omitempty"`
	Status				string			  `json:"status,omitempty"`
	CreatedAt     		time.Time         `json:"created_at,omitempty"`
	UpdatedAt     		time.Time         `json:"updated_at,omitempty"`
}

type CreateSourcesRequest struct {
	Order				int				  `json:"order,omitempty"`
	Source      		string            `json:"source" binding:"required"`
	Citation			string			  `json:"citation" binding:"required"`
	Topic				string			  `json:"topic" binding:"required"`
	Status				string			  `json:"status,omitempty"`
	CreatedAt	 		time.Time         `json:"created_at,omitempty"`
	UpdatedAt	 		time.Time         `json:"updated_at,omitempty"`
}

type UpdateSources struct {
	Order				int				  `json:"order,omitempty"`
	Source      		string            `json:"source,omitempty" binding:"required"`
	Citation			string			  `json:"citation" binding:"required"`
	Topic				string			  `json:"topic" binding:"required"`
	Status				string			  `json:"status,omitempty"`
	UpdatedAt	 		time.Time         `json:"updated_at,omitempty"`
}

