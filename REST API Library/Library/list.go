package Library

import "sync"

type List struct {
	Books map[string]Book
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		Books: make(map[string]Book),
	}
}

func (l *List) AddBook(book Book) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.Books[book.Title]; ok {
		return ErrBookAlreadyExists
	}

	l.Books[book.Title] = book

	return nil
}

func (l *List) GetBook(title string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.Books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	return task, nil
}

func (l *List) ListofBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Book, len(l.Books))

	for k, v := range l.Books {
		tmp[k] = v
	}

	return tmp
}

func (l *List) ListofUncompletedBooks() map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	uncompletedBooks := make(map[string]Book)

	for title, task := range l.Books {
		if task.Completed == false {
			uncompletedBooks[title] = task
		}
	}

	return uncompletedBooks
}

func (l *List) CompletedBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.Books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	task.Complete()

	l.Books[title] = task

	return task, nil
}

func (l *List) UncompletedBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book, ok := l.Books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}

	book.Uncomplete()

	l.Books[title] = book

	return book, nil
}

func (l *List) DeleteBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.Books[title]
	if !ok {
		return ErrBookNotFound
	}

	delete(l.Books, title)

	return nil
}
