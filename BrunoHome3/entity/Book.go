package entity

import (
	"fmt"
	"time"
)

type Book struct {
	id            int
	NameOfBook    string
	AuthoorOfBook string
	Availability  bool
	ReservedUntil time.Time
}

type Libryary struct {
	Books []Book
}

func (l *Libryary) AddBook(id int, NameOfBook string, AuthoorOfBook string, Availability bool,
	ReservedUntil time.Time) {
	newBook := Book{id: id, NameOfBook: NameOfBook, AuthoorOfBook: AuthoorOfBook,
		Availability: Availability, ReservedUntil: ReservedUntil}
	l.Books = append(l.Books, newBook)
}

func (l *Libryary) DeleteBook(NameOfBook string) {
	for i, book := range l.Books {
		if book.NameOfBook == NameOfBook {
			l.Books[i] = l.Books[len(l.Books)-1]
			l.Books = l.Books[:len(l.Books)-1]
		}
	}
}

func (l *Libryary) CheckBook(id int) {
	for _, book := range l.Books {
		if book.id == id {
			if book.Availability == true {
				fmt.Println("Книга Доступна")
			} else {
				fmt.Println("Забранировано до:", book.ReservedUntil.Format("2006-01-02"))
			}
		}
	}
}

func (l *Libryary) ListOfBool() {
	fmt.Println("Вот список из", len(l.Books), "книг библиотеки:")
	for _, book := range l.Books {
		var status string
		if book.Availability == false {
			status = "Недоступна"
		} else {
			status = "Доступна"
		}
		fmt.Println("id:", book.id, "Название книги:", "Автор:", book.AuthoorOfBook, book.NameOfBook, "Наличие:", status, "Забронирована до:", book.ReservedUntil)

	}
}
