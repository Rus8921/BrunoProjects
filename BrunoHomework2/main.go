package main

import (
	"errors"
	"fmt"
)

//Easy: написать функцию, которая принимает число-возраст и  определяет
//совершеннолетний ли человек и возвращает результат
//
//Normal: написать функцию, которая принимает набор чисел-возрастов,
//и определяет есть ли хоть один совершеннолетний в наборе, возвращает «Да»
//или «Нет»
//
//Hard: написать функцию, которая принимает карту, где имя-ключ,
//возраст-значение. Функция должна возвращать имена несовершеннолетних
//и ошибку с произвольным текстом, если такие имеются (errors.New())

func main() {
	//fmt.Println(easy(18))
	//massive := []int{1, 2, 18, 4, 5}
	//normal(massive)
	ageMap := map[string]int{
		"Alice":   25,
		"Bob":     12,
		"Charlie": 20,
	}
	fmt.Println(hard(ageMap))
}

func easy(a int) bool {
	if a >= 18 {
		return true
	} else {
		return false
	}
}

func normal(arr []int) {
	result := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 18 {
			result++
		}
	}
	if result > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func hard(maps map[string]int) ([]string, error) {
	underage := []string{}
	for name, age := range maps {
		if age < 18 {
			underage = append(underage, name)
		}
	}
	if len(underage) > 0 {
		return underage, errors.New("Найдены несовершенолетние")
	}
	return underage, nil
}
