package http

import (
	"encoding/json"
	"errors"
	"time"
)

type CompleteBookDTO struct {
	Complete bool
}

type BookDTO struct {
	Title  string
	Author string
	Pages  int
}

func (t BookDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Author == "" {
		return errors.New("author is empty")
	}
	if t.Pages == 0 {
		return errors.New("pages is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
