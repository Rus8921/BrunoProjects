package main // Объявление пакета main.

import (
	"BrunoHome3/entity" // Импорт пакета entity, предполагается, что путь к пакету указан корректно.
	"time"              // Импорт пакета time для работы с временем.
)

func main() {
	//Easy()
	//Midle()
	Hard()
}

// Easy Функция  представляет простой сценарий использования для школы.
func Easy() {
	school := entity.School{} // Создание экземпляра школы.

	school.CreateNewSubject(1, "Физика", "Иванов", 43) // Добавление предметов.
	school.CreateNewSubject(2, "Физика", "Иванов", 32)
	school.CreateNewSubject(3, "Физика", "Иванов", 14)
	school.CreateNewSubject(4, "Физика", "Иванов", 12)

	school.ListOfSubjects() // Вывод списка предметов.

	school.DeleteSubject(1) // Удаление предмета.

	school.ListOfSubjects() // Вывод обновленного списка предметов.

}

// Midle представляет средний сценарий использования для списка задач.
func Midle() {
	tasks := entity.ListOfTasks{} // Создание экземпляра списка задач.

	tasks.AddTask("Задача 1", false, 5) // Добавление задач.
	tasks.AddTask("Задача 2", true, 32)
	tasks.AddTask("Задача 3", false, 1)
	tasks.AddTask("Задача 4", false, 6)
	tasks.AddTask("Задача 5", false, 9)

	tasks.TaskList() // Вывод списка задач.

	tasks.TaskIsDone(5)  // Пометка задачи как выполненной.
	tasks.DeleeteTask(2) // Удаление задачи.

	tasks.TaskList() // Вывод обновленного списка задач.
}

// Hard представляет сложный сценарий использования для библиотеки.
func Hard() {
	libryary := entity.Libryary{} // Создание экземпляра библиотеки.

	libryary.AddBook(1, "Война и Мир", "Толстой", true, time.Now()) // Добавление книг.
	libryary.AddBook(2, "Текст", "Глуховский", false, time.Date(
		2024, time.March, 23, 12, 0, 0, 0, time.UTC))

	libryary.AddBook(3, "Пост", "Глуховский", false, time.Date(
		2024, time.April, 12, 10, 0, 0, 0, time.UTC))
	libryary.AddBook(4, "Война и Пир", "Толстой", true, time.Now())

	libryary.ListOfBool() // Вывод списка книг.

	libryary.CheckBook(1) // Проверка доступности книги.
	libryary.CheckBook(2)

	libryary.DeleteBook("Война и Мир") // Удаление книги.

	libryary.ListOfBool() // Вывод обновленного списка книг.
}
