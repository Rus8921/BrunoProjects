package iternal

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strconv"
)

type Task struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty" binding:"required"`
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status"`
	Priority    uint8  `json:"priority,omitempty"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"` //  тут мы получаем все задачи из БД в виде массива Task с помощью json.Marshal и json.Unmarshal и сохраняем в переменную Tasks
	// []Task для дальнейшего использования
	Total int `json:"total"` // количество задач в БД   тут мы получаем количество задач в БД в виде переменной total
}

func (t *TasksResponse) createTask(c *gin.Context) {
	var task Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	//task.ID = uint(len(tasks) + 1)
	//tasks[task.ID] = task
	task.ID = uuid.New().String()
	t.Tasks = append(t.Tasks, task)
	t.Total++
	c.JSON(http.StatusOK, gin.H{"message": "task created"})
	err = saveData(t.Tasks)
	if err != nil {
		c.JSON(http.StatusMultiStatus, gin.H{"error": err.Error()})
	}

}

func (t *TasksResponse) getAllTasks(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=3600")
	c.JSON(http.StatusOK, t.Tasks)
}

func (t *TasksResponse) getFilterAllTasks(c *gin.Context) {
	statusStr, okStatus := c.GetQuery("status")
	priorityStr, okPriority := c.GetQuery("priority")
	if !okStatus && !okPriority {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not valid data"})
		return
	}
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	priority, err := strconv.Atoi(priorityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filterTasks := make([]Task, 0, len(t.Tasks))

	for _, task := range t.Tasks {
		if (task.Status == status) && (task.Priority == uint8(priority)) {
			filterTasks = append(filterTasks, task)
		}
	}
	c.JSON(http.StatusOK, filterTasks)
}

func (t *TasksResponse) updateTask(c *gin.Context) {
	id := c.Param("id")
	ok, idx, task := t.getTaskBYID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t.Tasks[idx] = task

	c.JSON(http.StatusOK, task)
	err = saveData(t.Tasks)
	if err != nil {
		c.JSON(http.StatusMultiStatus, gin.H{"error": err.Error()})
	}
}

func (t *TasksResponse) getTaskBYID(id string) (bool, int, Task) {
	for idx, task := range t.Tasks {
		if task.ID == id {
			return true, idx, task
		}
	}
	return false, -1, Task{}
}

func (t *TasksResponse) deleteTaskByID(id string) bool {
	ok, idx, _ := t.getTaskBYID(id)
	if ok {
		t.Tasks = append(t.Tasks[:idx], t.Tasks[idx+1:]...) // remove task from array by index
		t.Total--
		return true
	}
	return false // task not found
}

func (t *TasksResponse) deleteTask(c *gin.Context) {
	id := c.Param("id")
	ok := t.deleteTaskByID(id)
	if ok {
		c.JSON(http.StatusOK, gin.H{"message": "task delete"})
		err := saveData(t.Tasks)
		if err != nil {
			c.JSON(http.StatusMultiStatus, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "task not found"})
}

func saveData(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *TasksResponse) loadTasksFromFile() error {

	data, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		_, err = os.Create("tasks.json")
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &t.Tasks)
	if err != nil {
		return err
	}
	t.Total = len(t.Tasks) // total tasks count in file
	return nil

}

func (t *TasksResponse) listTasks(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	low := (page - 1) * 10
	high := page * 10
	if high > len(t.Tasks) {
		high = len(t.Tasks)
	}
	res := t.Tasks[low:high]

	total := len(t.Tasks)

	c.JSON(http.StatusOK, gin.H{"tasks": res, "total": total}) // return tasks list	to clientf
}

//var tasks = make(map[uint]Task)
