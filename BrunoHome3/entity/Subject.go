package entity // Package entity Объявление пакета entity.

import "fmt" // Импорт пакета fmt для форматированного вывода.

// Subject Структура "Предмет" представляет собой описание предмета в школе.
type Subject struct {
	id              int    // Идентификатор предмета.
	subjectName     string // Название предмета.
	teatcher        string // Преподаватель предмета.
	durationInHours int    // Продолжительность занятий по предмету в часах.
}

// School Структура "Школа" содержит список предметов.
type School struct {
	Subjects []Subject // Список предметов в школе.
}

// CreateNewSubject Метод  создает новый предмет в школе.
func (sch *School) CreateNewSubject(id int, subjectName string, teatcher string, durationInHours int) {
	newSubject := Subject{id: id, subjectName: subjectName, teatcher: teatcher, durationInHours: durationInHours} // Создание нового предмета.
	sch.Subjects = append(sch.Subjects, newSubject)                                                               // Добавление нового предмета в список предметов школы.
}

// DeleteSubject Метод удаляет предмет из школы по его идентификатору.
func (sch *School) DeleteSubject(id int) {
	for i, subjects := range sch.Subjects { // Перебор всех предметов в списке.
		if subjects.id == id { // Проверка, соответствует ли идентификатор предмета заданному.
			sch.Subjects[i] = sch.Subjects[len(sch.Subjects)-1] // Перемещение последнего предмета в место удаляемого.
			sch.Subjects = sch.Subjects[:len(sch.Subjects)-1]   // Удаление последнего предмета из списка.
		}
	}
}

// ListOfSubjects Метод выводит список всех предметов в школе.
func (sch *School) ListOfSubjects() {
	fmt.Println("Список всех", len(sch.Subjects), "предметов") // Вывод сообщения о количестве предметов в школе.
	for _, subject := range sch.Subjects {                     // Перебор всех предметов в списке.
		fmt.Println("id:", subject.id, "Наименование предмет:", subject.subjectName,
			"Преподаватель:", subject.teatcher, "Кол-во часов:", subject.durationInHours) // Вывод информации о предмете.
	}

}
