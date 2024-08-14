package models

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type GeneralQuestions struct {
	ID            		uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Question      		string            `gorm:"not null" json:"question,omitempty"`
	AnswerOne			string			  `json:"answer_one,omitempty"`
	AnswerTwo			string			  `json:"answer_two,omitempty"`
	AnswerThree			string			  `json:"answer_three,omitempty"`
	AnswerFour			string			  `json:"answer_four,omitempty"`
	Difficulty	  		float32           `json:"difficulty"`
	AmountSeen			float32           `json:"amount_seen,omitempty"`
	AmountCorrect		float32           `json:"amount_correct,omitempty"`
	Likes				int               `json:"likes"`
	Dislikes			int               `json:"dislikes"`
	CorrectAnswer 		string            `json:"correct_answer,omitempty"`
	Topic				string       	  `json:"topic"`
	CreatedAt     		time.Time         `json:"created_at,omitempty"`
	UpdatedAt     		time.Time         `json:"updated_at,omitempty"`
}

type CreateGeneralQuestionsRequest struct {
	Question      		string            `gorm:"not null" json:"question,omitempty"`
	AnswerOne			string			  `json:"answer_one,omitempty"`
	AnswerTwo			string			  `json:"answer_two,omitempty"`
	AnswerThree			string			  `json:"answer_three,omitempty"`
	AnswerFour			string			  `json:"answer_four,omitempty"`
	Difficulty	  		float32           `json:"difficulty"`
	AmountSeen			float32           `json:"amount_seen,omitempty"`
	AmountCorrect		float32           `json:"amount_correct,omitempty"`
	Likes				int               `json:"likes"`
	Dislikes			int               `json:"dislikes"`
	CorrectAnswer 		string            `json:"correct_answer,omitempty"`
	Topic				string       	  `json:"topic"`
	CreatedAt     		time.Time         `json:"created_at,omitempty"`
	UpdatedAt     		time.Time         `json:"updated_at,omitempty"`
}

type UpdateGeneralQuestions struct {
	Question      		string            `gorm:"not null" json:"question,omitempty"`
	AnswerOne			string			  `json:"answer_one,omitempty"`
	AnswerTwo			string			  `json:"answer_two,omitempty"`
	AnswerThree			string			  `json:"answer_three,omitempty"`
	AnswerFour			string			  `json:"answer_four,omitempty"`
	CorrectAnswer 		string            `json:"correct_answer,omitempty"`
	Topic				string       	  `json:"topic"`
	UpdatedAt     		time.Time         `json:"updated_at,omitempty"`
}

type GQUpdateTopic struct {
	QuestionId     uuid.UUID          `json:"question_id,omitempty"`
	Topic 		   string             `json:"topic,omitempty"`
}