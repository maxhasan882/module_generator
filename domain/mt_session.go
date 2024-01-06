package domain

type ModelTestResult struct {
	Id              string   `json:"id,omitempty"`
	CorrectAnswer   *int32   `json:"correct_answer,omitempty"`
	IncorrectAnswer *int32   `json:"incorrect_answer,omitempty"`
	TotalQuestion   int32    `json:"total_question,omitempty"`
	IsPassed        bool     `json:"is_passed,omitempty"`
	MarksObtained   float64  `json:"marks_obtained,omitempty"`
	TotalMarks      float64  `json:"total_marks,omitempty"`
	TimeSpent       *float64 `json:"time_spent,omitempty"`
}

func (m *ModelTestResult) Generate(Id string, correctAnswer *int32, incorrectAnswer *int32, totalQuestion int32, isPassed bool, marksObtained float64, totalMarks float64, timeSpent *float64) *ModelTestResult {
	return &ModelTestResult{
		Id:              Id,
		CorrectAnswer:   correctAnswer,
		IncorrectAnswer: incorrectAnswer,
		TotalQuestion:   totalQuestion,
		IsPassed:        isPassed,
		MarksObtained:   marksObtained,
		TotalMarks:      totalMarks,
		TimeSpent:       timeSpent,
	}
}