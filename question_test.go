package question

import (
	"fmt"
	"testing"
)

func newQuestionOrFatal(t *testing.T, number uint16, description string, answers *[]Answer) *Question {
	question, err := NewQuestion(number, description, answers)
	if err != nil {
		t.Fatalf("new question: %v", err)
	}
	return question
}

func TestNewQuestion(t *testing.T) {
	number, description := uint16(1), "How many edges are there in a square?"
	answers := make([]Answer, 0, 4)

	question := newQuestionOrFatal(t, number, description, &answers)
	if question.number != number {
		t.Errorf("expected number %d, got %d", number, question.number)
	}
	if question.description != description {
		t.Errorf("expected description %q, got %q", description, question.description)
	}
	if question.AnswerSize() != 0 {
		t.Errorf("expected answer size %d, got %d", 0, question.AnswerSize())
	}
}

func TestNewQuestionEmptyDescription(t *testing.T) {
	answers := make([]Answer, 0, 4)
	_, err := NewQuestion(uint16(1), "", &answers)
	if err == nil {
		t.Errorf("expected %v error, got %v", fmt.Errorf("empty description"))
	}
}

func TestNewQuestionNonPositiveNumber(t *testing.T) {
	answers := make([]Answer, 0, 4)
	_, err := NewQuestion(uint16(0), "How many edges are there in a triangle?", &answers)
	if err == nil {
		t.Errorf("expected %v error, got %v", fmt.Errorf("question number has to be > 0"))
	}
}

func newQuestion(number uint16, description string) *Question {
	return &Question{number, description, make([]Answer, 0, 4)}
}

func TestAddNewAnswer(t *testing.T) {
	question := newQuestion(uint16(1), "How many edges are there in a square?")
	err := question.AddNewAnswer("", true)

	if err == nil {
		t.Errorf("expected empty answer error but nothing raised")
	}
	if err = question.AddNewAnswer("3", false); err != nil {
		t.Errorf("expected no error but returned an error")
	}
	if question.AnswerSize() != 1 {
		t.Errorf("expected answer size of 1 but returned %v", question.AnswerSize())
	}
}
