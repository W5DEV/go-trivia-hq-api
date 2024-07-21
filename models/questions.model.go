package models

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/datatypes"
)

type Questions struct {
	ID            uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Question      string            `gorm:"uniqueIndex;not null" json:"question,omitempty"`
	Answers       datatypes.JSON    `gorm:"uniqueIndex;not null" json:"answers,omitempty"`
	Source		  string            `json:"source,omitempty"`
	Type	      string            `json:"type,omitempty"`
	Tags		  datatypes.JSON    `json:"tags,omitempty"`
	CreatedAt     time.Time         `json:"created_at,omitempty"`
	UpdatedAt     time.Time         `json:"updated_at,omitempty"`
}

type CreateQuestionsRequest struct {
	Question        string            `json:"question"  binding:"required"`
	Answers         datatypes.JSON    `json:"answers"  binding:"required"`
	Source   		string            `json:"source,omitempty"`
	Type      		string            `json:"type,omitempty"`
	Tags		  	datatypes.JSON    `json:"tags,omitempty"`
	CreatedAt     	time.Time         `json:"created_at,omitempty"`
	UpdatedAt     	time.Time         `json:"updated_at,omitempty"`
}

type UpdateQuestions struct {
	Question        string             `json:"question,omitempty"`
	Answers         datatypes.JSON     `json:"answers,omitempty"`
	Source  		string             `json:"source,omitempty"`
	Type     		string             `json:"type,omitempty"`
	Tags	      	datatypes.JSON     `json:"tags,omitempty"`
	CreateAt      	time.Time          `json:"created_at,omitempty"`
	UpdatedAt     	time.Time          `json:"updated_at,omitempty"`
}