package entity // Package entity Объявление пакета entity.

import (
	"fmt"  // Импорт пакета fmt для форматированного вывода.
	"time" // Импорт пакета time для работы с временем.
)

// Book Структура представляет собой описание книги в библиотеке.
type Book struct {
	id            int       // Идентификатор книги.
	NameOfBook    string    // Название книги.
	AuthoorOfBook string    // Автор книги.
	Availability  bool      // Доступность книги.
	ReservedUntil time.Time // Дата зарезервированности книги.
}

// Libryary Структура "Библиотека" содержит список книг.
type Libryary struct {
	Books []Book // Список книг в библиотеке.
}

// AddBook добавляет новую книгу в библиотеку.
func (l *Libryary) AddBook(id int, NameOfBook string, AuthoorOfBook string, Availability bool,
	ReservedUntil time.Time) {
	newBook := Book{id: id, NameOfBook: NameOfBook, AuthoorOfBook: AuthoorOfBook,
		Availability: Availability, ReservedUntil: ReservedUntil} // Создание новой книги.
	l.Books = append(l.Books, newBook) // Добавление новой книги в список книг библиотеки.
}

// DeleteBook удаляет книгу из библиотеки по её названию.
func (l *Libryary) DeleteBook(NameOfBook string) {
	for i, book := range l.Books { // Перебор всех книг в списке.
		if book.NameOfBook == NameOfBook { // Проверка, соответствует ли название книги заданному.
			l.Books[i] = l.Books[len(l.Books)-1] // Перемещение последней книги в место удаляемой.
			l.Books = l.Books[:len(l.Books)-1]   // Удаление последней книги из списка.
		}
	}
}

// CheckBook проверяет доступность книги по её ID.
func (l *Libryary) CheckBook(id int) {
	for _, book := range l.Books { // Перебор всех книг в списке.
		if book.id == id { // Проверка, соответствует ли идентификатор книги заданному.
			if book.Availability == true { // Проверка доступности книги.
				fmt.Println("Книга Доступна") // Вывод сообщения о доступности.
			} else {
				fmt.Println("Забранировано до:", book.ReservedUntil.Format("2006-01-02")) // Вывод даты зарезервированности.
			}
		}
	}
}

// ListOfBool выводит список всех книг в библиотеке.
func (l *Libryary) ListOfBool() {
	fmt.Println("Вот список из", len(l.Books), "книг библиотеки:") // Вывод сообщения о количестве книг в библиотеке.
	for _, book := range l.Books {                                 // Перебор всех книг в списке.
		var status string               // Переменная для хранения статуса книги.
		if book.Availability == false { // Проверка доступности книги.
			status = "Недоступна" // Установка статуса "Недоступна".
		} else {
			status = "Доступна" // Установка статуса "Доступна".
		}
		fmt.Println("id:", book.id, "Название книги:", "Автор:", book.AuthoorOfBook, book.NameOfBook,
			"Наличие:", status, "Забронирована до:", book.ReservedUntil) // Вывод информации о книге.
	}
}
