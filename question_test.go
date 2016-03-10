package question

import (
	"fmt"
	"testing"
)

func newQuestionOrFatal(t *testing.T, number uint16, description string) *Question {
	question, err := NewQuestion(number, description)
	if err != nil {
		t.Fatalf("new question: %v", err)
	}
	return question
}

func TestNewQuestion(t *testing.T) {
	number, description := uint16(1), "How many edges are there in a square?"

	question := newQuestionOrFatal(t, number, description)
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
	_, err := NewQuestion(uint16(1), "")
	if err == nil {
		t.Errorf("expected %v error, got %v", fmt.Errorf("empty description"))
	}
}

func TestNewQuestionNonPositiveNumber(t *testing.T) {
	_, err := NewQuestion(uint16(0), "How many edges are there in a triangle?")
	if err == nil {
		t.Errorf("expected %v error, got %v", fmt.Errorf("question number has to be > 0"))
	}
}

func newQuestion(number uint16, description string) *Question {
	answers, photos := make([]Answer, 0, 4), make([]Photo, 0, 1)
	return &Question{number, description, answers, photos}
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

func TestAddNewQuestionPhoto(t *testing.T) {
	question := newQuestion(uint16(1), "How many edges are there in a triangle?")
	err := question.AddNewQuestionPhoto("")

	if err == nil {
		t.Errorf("expected empty link error but nothing raised")
	}
	if err = question.AddNewQuestionPhoto("/xxx/y.png"); err != nil {
		t.Errorf("expected no error but returned an error")
	}
	if question.QuestionPhotoCount() != 1 {
		t.Errorf("expected photo count of 1 but returned %v", question.QuestionPhotoCount())
	}
	if question.photos[0].link != "/xxx/y.png" {
		t.Errorf("expected link is not matching %v", question.photos[0].link)
	}
}

func TestAddNewAnswerPhoto(t *testing.T) {
	answer, _ := NewAnswer("Triangle has 3 sides", true)
	err := answer.AddNewAnswerPhoto("")

	if err == nil {
		t.Errorf("expexted empty link error but nothing raised")
	}
	if err = answer.AddNewAnswerPhoto("/zzz/a.jpg"); err != nil {
		t.Errorf("expected no error but returned an error")
	}
	if answer.AnswerPhotoCount() != 1 {
		t.Errorf("expected photo count of 1 but returned %v", answer.AnswerPhotoCount())
	}
	if answer.photos[0].link != "/zzz/a.jpg" {
		t.Errorf("expected link is not matching %v", answer.photos[0].link)
	}
}
