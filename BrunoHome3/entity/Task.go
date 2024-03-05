package entity // Package entity Объявление пакета entity.

import "fmt" // Импорт пакета fmt для форматированного вывода.

// Task Структура "Задача" представляет собой описание задачи.
type Task struct {
	id                uint   // Идентификатор задачи.
	DescriptionOfTask string // Описание задачи.
	IsDone            bool   // Статус выполнения задачи.
	DeadlineInHours   int    // Срок выполнения задачи в часах.
}

// ListOfTasks Структура "Список задач" содержит список задач.
type ListOfTasks struct {
	Tasks []Task // Список задач.
}

var counterId uint // Глобальная переменная для учета идентификаторов задач.

// AddTask Методдобавляет новую задачу в список.
func (lt *ListOfTasks) AddTask(DescriptionOfTask string, IsDone bool, DeadlineInHours int) {
	counterId++                                                                                                            // Увеличение счетчика идентификаторов.
	newTask := Task{id: counterId, DescriptionOfTask: DescriptionOfTask, IsDone: IsDone, DeadlineInHours: DeadlineInHours} // Создание новой задачи.
	newTask.IsDone = false                                                                                                 // Установка начального статуса "Невыполнено".
	lt.Tasks = append(lt.Tasks, newTask)                                                                                   // Добавление новой задачи в список задач.
}

// TaskIsDone Метод помечает задачу с заданным идентификатором как выполненную.
func (lt *ListOfTasks) TaskIsDone(id uint) {
	for i, task := range lt.Tasks { // Перебор всех задач в списке.
		if task.id == id { // Проверка, соответствует ли идентификатор задачи заданному.
			lt.Tasks[i].IsDone = true // Установка статуса "Выполнено".
		}
	}

}

// DeleeteTask удаляет задачу из списка по её идентификатору.
func (lt *ListOfTasks) DeleeteTask(id uint) {
	for i, task := range lt.Tasks { // Перебор всех задач в списке.
		if task.id == id { // Проверка, соответствует ли идентификатор задачи заданному.
			lt.Tasks[i] = lt.Tasks[len(lt.Tasks)-1] // Перемещение последней задачи в место удаляемой.
			lt.Tasks = lt.Tasks[:len(lt.Tasks)-1]   // Удаление последней задачи из списка.
		}
	}
}

// TaskList Метод выводит список всех задач.
func (lt *ListOfTasks) TaskList() {
	fmt.Println("Спсисок всех", len(lt.Tasks), "задач:") // Вывод сообщения о количестве задач в списке.
	for _, task := range lt.Tasks {                      // Перебор всех задач в списке.
		var status string         // Переменная для хранения статуса задачи.
		if task.IsDone == false { // Проверка статуса выполнения задачи.
			status = "Невыполнено" // Установка статуса "Невыполнено".
		} else {
			status = "Выполнено" // Установка статуса "Выполнено".
		}
		fmt.Println("id:", task.id, "Описание задачи:", task.DescriptionOfTask, "Статус задачи:", status, "Дэдлайн через:", task.DeadlineInHours, "часлв") // Вывод информации о задаче.
	}
}
