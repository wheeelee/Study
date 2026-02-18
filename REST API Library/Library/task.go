package Library

import "time"

type Book struct {
	Title     string
	Author    string //description
	Pages     int
	Completed bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title string, author string, pages int) Book {
	return Book{
		Title:     title,
		Author:    author,
		Completed: false,
		Pages:     pages,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Book) Complete() {
	completeTime := time.Now()

	t.Completed = true
	t.CompletedAt = &completeTime
}

func (t *Book) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}
