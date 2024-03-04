package entity

import (
	"errors"
	"strings"
)

// Если мы будем писать имена переменных и функций с маленькой буквы, они будут только внутри этого файла и чтобы их достать нам будут нужны геттеры

func init() {
	// функция для инициализации пакетов, проигрывается всегда первой во всем файле/ пакете
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson() {

}

type Group struct {
	Persons Person
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) EditName() { // Если мы не добавим * к персон, то мы будем работать просто с копией, а не оригинал не изменится
	p.FirstName = "Anton"
}

type User struct {
	UserName string
	Email    string
	Password string
}

func (u *User) ChangePassword() string {
	u.Password = "newpassword"
	return u.Password
}

func (u *User) CheckEmain() (bool, error) {
	if !strings.Contains(u.Email, "a") {
		return false, errors.New("not valid email")
	}
	return true, nil
}

type Printer interface {
	FullName() string
}

func PrintInfo(p Printer) {
	p.FullName()
}
