package main

import (
	"fmt"
)

func main() {
	//maps()
	//ptr()
	//slicer()
	//zadanie1()
	//ifElse(87)
	//switchCase("")
	//Vosrast()
	//SummaInMassive()
	//SummaTsifr(18)
}

func ptr() {
	var ptr *int // создаем возможность ссылки на именно позицию в памяти
	fmt.Println(ptr)
	var num int = 10
	ptr = &num //  оперсант показывает что мы ссылаемся не на значение, а на позицию
	fmt.Println(ptr)
}

func maps() {
	numbers := make(map[string]int, 100) // по сути словари из Python
	numbers["Коля"] = 1800
	numbers["Коля"]++
	val, exists := numbers["Коля"]
	val2, exists2 := numbers["Вова"]
	_, exists3 := numbers["Леха"]
	fmt.Println(numbers, val, exists, numbers["Вова"], val2, exists2, numbers["Леха"], exists3)
}

func modifySlice(slice []int) []int {
	slice[0] = 99
	slice = append(slice, 4, 5, 6, 7)
	return slice
}

func slicer() {
	numbers := make([]int, 3, 6) // слайс макс загруженностью 6
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers)
	numbers = modifySlice(numbers)
	fmt.Println(numbers)
}

func zadanie1() {
	products := make(map[string]int, 10)
	products["Помидоры"] = 200
	products["Огурцы"] = 300
	products["Свекла"] = 400
	products["Морковь"] = 100
	fmt.Println("Price one of the product -", products["Свекла"])
	delete(products, "Свекла")
	fmt.Println("Весь список продуктов -", products)
}
func ifElse(a int) {
	if (a > 0) && (a < 41) {
		fmt.Println("E")
	}
	if (a >= 41) && (a < 61) {
		fmt.Println("C")
	}
	if (a >= 61) && (a < 90) {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}
func switchCase(b string) {
	fmt.Println("Введите день недели")
	fmt.Scan(&b)
	switch b {
	case "Md":
		fmt.Println(1)
	case "Tu":
		fmt.Println(2)
	case "We":
		fmt.Println(3)
	case "Thr":
		fmt.Println(4)
	case "Fr":
		fmt.Println(5)
	case "Su":
		fmt.Println(6)
	case "St":
		fmt.Println(7)
	default:
		fmt.Println("error")
	}
}

func Vosrast() {
	vosrast := make([]int, 5)
	for i := 0; i < 5; i++ {
		vosrast[i] = i + 1
	}
	fmt.Println(vosrast)

	for _, val := range vosrast {
		fmt.Println(val)

	}
}

func SummaInMassive() {
	massive := []int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i < len(massive); i++ {
		sum += massive[i]
		fmt.Println(sum)
	}
}

func SummaTsifr(a int) {
	sum := 0
	//for {  // бесконечный цикл
	//	sum += a % 10
	//	a /= 10
	//	if a == 0 {
	//		break
	//	}
	for a != 0 { // бесконечный цикл/ аналог while
		sum += a % 10
		a /= 10
	}
	fmt.Println(sum)
}
