package models

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/datatypes"
)

type Questions struct {
	ID            		uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Question      		string            `gorm:"not null" json:"question,omitempty"`
	Answers       		datatypes.JSON    `gorm:"not null" json:"answers,omitempty"`
	Source		  		string            `json:"source,omitempty"`
	Type	      		string            `json:"type,omitempty"`
	Tags		  		datatypes.JSON    `json:"tags,omitempty"`
	Difficulty	  		float32           `json:"difficulty"`
	AmountSeen			float32           `json:"amount_seen,omitempty"`
	AmountCorrect		float32           `json:"amount_correct,omitempty"`
	Likes				int               `json:"likes"`
	Dislikes			int               `json:"dislikes"`
	CorrectAnswer 		string            `json:"correct_answer,omitempty"`
	Completed     		string            `json:"completed,omitempty"`
	QuestionOrigin 		string            `gorm:"column:question_origin" json:"question_origin,omitempty"`
	Topic				string       	  `json:"topic"`
	CreatedAt     		time.Time         `json:"created_at,omitempty"`
	UpdatedAt     		time.Time         `json:"updated_at,omitempty"`
}

type CreateQuestionsRequest struct {
	Question        string            `json:"question"  binding:"required"`
	Answers         datatypes.JSON    `json:"answers"  binding:"required"`
	Source   		string            `json:"source" binding:"required"`
	Type      		string            `json:"type" binding:"required"`
	Tags		  	datatypes.JSON    `json:"tags" binding:"required"`
	CorrectAnswer 	string            `json:"correct_answer" binding:"required"`
	Completed     	string            `json:"completed" binding:"required"`
	QuestionOrigin 	string            `json:"question_origin" binding:"required"`
	Topic 			string            `json:"topic" binding:"required"`
	CreatedAt     	time.Time         `json:"created_at,omitempty"`
	UpdatedAt     	time.Time         `json:"updated_at,omitempty"`
}

type UpdateQuestions struct {
	Question        string             `json:"question,omitempty"`
	Answers         datatypes.JSON     `json:"answers,omitempty"`
	Source  		string             `json:"source,omitempty"`
	Type     		string             `json:"type,omitempty"`
	Tags	      	datatypes.JSON     `json:"tags,omitempty"`
	CorrectAnswer 	string             `json:"correct_answer,omitempty"`
	Completed     	string             `json:"completed,omitempty"`
	QuestionOrigin 	string             `json:"question_origin,omitempty"`
	CreateAt      	time.Time          `json:"created_at,omitempty"`
	UpdatedAt     	time.Time          `json:"updated_at,omitempty"`
}

type RecordAnswer struct {
	QuestionId     uuid.UUID          `json:"question_id,omitempty"`
	IsCorrect      bool               `json:"is_correct,omitempty"`
}

type RecordLike struct {
	QuestionId     uuid.UUID          `json:"question_id,omitempty"`
	IsLiked        bool               `json:"is_liked,omitempty"`
}

type RecordDislike struct {
	QuestionId     uuid.UUID          `json:"question_id,omitempty"`
	IsDisliked     bool               `json:"is_disliked,omitempty"`
}

type UpdateTopic struct {
	QuestionId     uuid.UUID          `json:"question_id,omitempty"`
	Topic 		   string             `json:"topic,omitempty"`
}