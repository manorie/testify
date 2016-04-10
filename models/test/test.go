package test

type Test struct {
	Id             uint64 `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Path           string `json:"-"`
	SecretKey      string `json:"-"`
	AuthorEmail    string `json:"author_email"`
	IsPublished    bool   `json:"is_published"`
	TimeLimit      uint8  `json:"time_limit"`
	AnswerSize     uint8  `json:"answer_size"`
	CreationDate   int64  `json:"creation_date"`
	ExpirationDate int64  `json:"expiration_date"`
}
