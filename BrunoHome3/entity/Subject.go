package entity

import "fmt"

type Subject struct {
	id              int
	subjectName     string
	teatcher        string
	durationInHours int
}

type School struct {
	Subjects []Subject
}

func (sch *School) CreateNewSubject(id int, subjectName string, teatcher string, durationInHours int) {
	newSubject := Subject{id: id, subjectName: subjectName, teatcher: teatcher, durationInHours: durationInHours}
	sch.Subjects = append(sch.Subjects, newSubject)
}

func (sch *School) DeleteSubject(id int) {
	for i, subjects := range sch.Subjects {
		if subjects.id == id {
			sch.Subjects[i] = sch.Subjects[len(sch.Subjects)-1]
			sch.Subjects = sch.Subjects[:len(sch.Subjects)-1]
		}
	}
}

func (sch *School) ListOfSubjects() {
	fmt.Println("Список всех", len(sch.Subjects), "предметов")
	for _, subject := range sch.Subjects {
		fmt.Println("id:", subject.id, "Наименование предмет:", subject.subjectName,
			"Преподаватель:", subject.teatcher, "Кол-во часов:", subject.durationInHours)
	}

}
