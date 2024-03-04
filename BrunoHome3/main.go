package main

import (
	"BrunoHome3/entity"
	"time"
)

func main() {
	//Easy()
	//Midle()
	Hard()
}

func Easy() {
	school := entity.School{}

	school.CreateNewSubject(1, "Физика", "Иванов", 43)
	school.CreateNewSubject(2, "Физика", "Иванов", 32)
	school.CreateNewSubject(3, "Физика", "Иванов", 14)
	school.CreateNewSubject(4, "Физика", "Иванов", 12)

	school.ListOfSubjects()

	school.DeleteSubject(1)

	school.ListOfSubjects()

}

func Midle() {
	tasks := entity.ListOfTasks{}

	tasks.AddTask(1, "Задача 1", false, 5)
	tasks.AddTask(2, "Задача 2", true, 32)
	tasks.AddTask(3, "Задача 3", false, 1)
	tasks.AddTask(4, "Задача 4", false, 6)
	tasks.AddTask(5, "Задача 5", false, 9)

	tasks.TaskList()

	tasks.TaskIsDone(5)
	tasks.DeleeteTask(2)

	tasks.TaskList()
}

func Hard() {
	libryary := entity.Libryary{}

	libryary.AddBook(1, "Война и Мир", "Толстой", true, time.Now())
	libryary.AddBook(2, "Текст", "Глуховский", false, time.Date(
		2024, time.March, 23, 12, 0, 0, 0, time.UTC))

	libryary.AddBook(3, "Пост", "Глуховский", false, time.Date(
		2024, time.April, 12, 10, 0, 0, 0, time.UTC))
	libryary.AddBook(4, "Война и Пир", "Толстой", true, time.Now())

	libryary.ListOfBool()

	libryary.CheckBook(1)
	libryary.CheckBook(2)

	libryary.DeleteBook("Война и Мир")

	libryary.ListOfBool()
}
