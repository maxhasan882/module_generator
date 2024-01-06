package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"template/domain"
)

type ModelTestResult struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CorrectAnswer   *int32             `json:"correct_answer,omitempty" bson:"correct_answer,omitempty"`
	IncorrectAnswer *int32             `json:"incorrect_answer,omitempty" bson:"incorrect_answer,omitempty"`
	TotalQuestion   int32              `json:"total_question,omitempty" bson:"total_question,omitempty"`
	IsPassed        bool               `json:"is_passed,omitempty" bson:"is_passed,omitempty"`
	MarksObtained   float64            `json:"marks_obtained,omitempty" bson:"marks_obtained,omitempty"`
	TotalMarks      float64            `json:"total_marks,omitempty" bson:"total_marks,omitempty"`
	TimeSpent       *float64           `json:"time_spent,omitempty" bson:"time_spent,omitempty"`
}

func (m *ModelTestResult) GetData(d *domain.ModelTestResult) *ModelTestResult {
	return &ModelTestResult{
		CorrectAnswer:   d.CorrectAnswer,
		IncorrectAnswer: d.IncorrectAnswer,
		TotalQuestion:   d.TotalQuestion,
		IsPassed:        d.IsPassed,
		MarksObtained:   d.MarksObtained,
		TotalMarks:      d.TotalMarks,
		TimeSpent:       d.TimeSpent,
	}
}
