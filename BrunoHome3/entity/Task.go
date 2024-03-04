package entity

import "fmt"

type Task struct {
	id                int
	DescriptionOfTask string
	IsDone            bool
	DeadlineInHours   int
}
type ListOfTasks struct {
	Tasks []Task
}

func (lt *ListOfTasks) AddTask(id int, DescriptionOfTask string, IsDone bool, DeadlineInHours int) {
	newTask := Task{id: id, DescriptionOfTask: DescriptionOfTask, IsDone: IsDone, DeadlineInHours: DeadlineInHours}
	newTask.IsDone = false
	lt.Tasks = append(lt.Tasks, newTask)
}

func (lt *ListOfTasks) TaskIsDone(id int) {
	for i, task := range lt.Tasks {
		if task.id == id {
			lt.Tasks[i].IsDone = true
		}
	}

}
func (lt *ListOfTasks) DeleeteTask(id int) {
	for i, task := range lt.Tasks {
		if task.id == id {
			lt.Tasks[i] = lt.Tasks[len(lt.Tasks)-1]
			lt.Tasks = lt.Tasks[:len(lt.Tasks)-1]
		}
	}
}

func (lt *ListOfTasks) TaskList() {
	fmt.Println("Спсисок всех", len(lt.Tasks), "задач:")
	for _, task := range lt.Tasks {
		var status string
		if task.IsDone == false {
			status = "Невыполнено"
		} else {
			status = "Выполнено"
		}
		fmt.Println("id:", task.id, "Описание задачи:", task.DescriptionOfTask, "Статус задачи:", status, "Дэдлайн через:", task.DeadlineInHours, "часлв")
	}
}
