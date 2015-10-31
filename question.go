package question

import "fmt"

type Question struct {
	number      uint16
	description string
	answers     []Answer
}

type Answer struct {
	descripton string
	Correct    bool
}

func NewQuestion(number uint16, description string, answers *[]Answer) (*Question, error) {
	if description == "" {
		return nil, fmt.Errorf("empty description")
	}
	if number < 1 {
		return nil, fmt.Errorf("question number has to be > 0")
	}
	return &Question{number, description, *answers}, nil
}

func (q *Question) AnswerSize() uint16 {
	return uint16(len(q.answers))
}

func (q *Question) AddNewAnswer(description string, IsCorrect bool) error {
	if description == "" {
		return fmt.Errorf("Answer description can not be empty")
	}
	q.answers = append(q.answers, Answer{description, IsCorrect})
	return nil
}
