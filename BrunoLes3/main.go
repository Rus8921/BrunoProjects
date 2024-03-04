package main

import (
	"fmt"
	"reflect"
)

//import (
//	"BrunoLes3/Entity"
//)
//
//func main() {
//	person := entity.Person{
//		FirstName: "Artem",
//		LastName:  "Antonyan",
//		Age:       21,
//	}
//
//	//group := entity.Group{Persons: person}
//	entity.PrintInfo(person)
//
//}

func main() {
	var myDict map[string]interface{}
	myDict = make(map[string]interface{})

	myDict["строка 1"] = true
	myDict["строка 2"] = "Правда"
	myDict["строка 3"] = 41
	myDict["строка 4"] = 43.3

	for key, val := range myDict {
		fmt.Println(key, val, reflect.TypeOf(val))
	}
}
