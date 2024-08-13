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
	Active				string			  `json:"active,omitempty"`
	Topic				string			  `json:"topic,omitempty"`
	Completed			string			  `json:"completed,omitempty"`
	CreatedAt     		time.Time         `json:"created_at,omitempty"`
	UpdatedAt     		time.Time         `json:"updated_at,omitempty"`
}

type CreateSourcesRequest struct {
	Source      		string            `json:"source" binding:"required"`
	Citation			string			  `json:"citation" binding:"required"`
	Active				string			  `json:"active"`
	Topic				string			  `json:"topic" binding:"required"`
	Completed			string			  `json:"completed"`
	Order				int				  `json:"order,omitempty"`
	CreatedAt	 		time.Time         `json:"created_at,omitempty"`
	UpdatedAt	 		time.Time         `json:"updated_at,omitempty"`
}

type UpdateSources struct {
	Source      		string            `json:"source,omitempty" binding:"required"`
	Citation			string			  `json:"citation" binding:"required"`
	Active				string			  `json:"active"`
	Topic				string			  `json:"topic" binding:"required"`
	Completed			string			  `json:"completed"`
	UpdatedAt	 		time.Time         `json:"updated_at,omitempty"`
}

