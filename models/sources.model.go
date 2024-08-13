package models

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Sources struct {
	ID            		uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Source      		string            `json:"source,omitempty"`
	Citation			string			  `json:"citation,omitempty"`
	Active				bool			  `json:"active,omitempty"`
	Topic				string			  `json:"topic,omitempty"`
	Completed			bool			  `json:"completed,omitempty"`
	CreatedAt     		time.Time         `json:"created_at,omitempty"`
	UpdatedAt     		time.Time         `json:"updated_at,omitempty"`
}

type CreateSourcesRequest struct {
	Source      		string            `json:"source" binding:"required"`
	Citation			string			  `json:"citation"  binding:"required"`
	Active				bool			  `json:"active,omitempty"`
	Topic				string			  `json:"topic"  binding:"required"`
	Completed			bool			  `json:"completed,omitempty"`
	CreatedAt	 		time.Time         `json:"created_at,omitempty"`
	UpdatedAt	 		time.Time         `json:"updated_at,omitempty"`
}

type UpdateSources struct {
	Source      		string            `json:"source,omitempty" binding:"required"`
	Citation			string			  `json:"citation" binding:"required"`
	Active				bool			  `json:"active,omitempty"`
	Topic				string			  `json:"topic" binding:"required"`
	Completed			bool			  `json:"completed,omitempty"`
	UpdatedAt	 		time.Time         `json:"updated_at,omitempty"`
}