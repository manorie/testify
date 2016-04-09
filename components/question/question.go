package question

import "fmt"

type Question struct {
	number      uint16
	description string
	answers     []Answer
	photos      []Photo
}

type Photo struct {
	link string
}

type Answer struct {
	descripton string
	isCorrect  bool
	photos     []Photo
}

func NewQuestion(number uint16, description string) (*Question, error) {
	if description == "" {
		return nil, fmt.Errorf("empty description")
	}
	if number < 1 {
		return nil, fmt.Errorf("question number has to be > 0")
	}
	answers := make([]Answer, 0, 4)
	photos := make([]Photo, 0, 1)
	return &Question{number, description, answers, photos}, nil
}

func NewAnswer(description string, isCorrect bool) (*Answer, error) {
	if description == "" {
		return nil, fmt.Errorf("empty description")
	}
	photos := make([]Photo, 0, 1)
	return &Answer{description, isCorrect, photos}, nil
}

func (q *Question) AnswerSize() uint16 {
	return uint16(len(q.answers))
}

func (q *Question) AddNewAnswer(description string, isCorrect bool) error {
	if description == "" {
		return fmt.Errorf("Answer description can not be empty")
	}
	answer, _ := NewAnswer(description, isCorrect)
	q.answers = append(q.answers, *answer)
	return nil
}

func (q *Question) RemoveAnswers() {
	q.answers = nil
}

//adding photos has to be url-safe
func (q *Question) AddNewQuestionPhoto(link string) error {
	if link == "" {
		return fmt.Errorf("Photo link can not be empty")
	}
	q.photos = append(q.photos, Photo{link})
	return nil
}

func (q *Question) QuestionPhotoCount() uint16 {
	return uint16(len(q.photos))
}

func (a *Answer) AddNewAnswerPhoto(link string) error {
	if link == "" {
		return fmt.Errorf("Photo link can not be empty")
	}
	a.photos = append(a.photos, Photo{link})
	return nil
}

func (a *Answer) AnswerPhotoCount() uint16 {
	return uint16(len(a.photos))
}
