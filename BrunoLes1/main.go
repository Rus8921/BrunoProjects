package main

import (
	"fmt"
)

func main() {
	//var a, b int
	//fmt.Println("Введите перменные:")
	//fmt.Scan(&a, &b)
	//sum(a, b)
	//
	//fmt.Println("Введите для Normal: сторона 1, сторона 2, угол между ними")
	//var x, y, z, p float64
	//fmt.Scan(&x, &y, &z)
	//p = x / 2 * y / 2 * math.Sin(z)
	//fmt.Println(p)
	//fmt.Println(PyphfagorTrue(int(x), int(y), int(z)))
	//
	//fmt.Println("введите для хард: число для факториала")
	//var n int
	//fmt.Scan(&n)
	//result := factorial(n)
	//fmt.Printf("Факториал %d равен %d\n", n, result)
	//
	//arr := [3]int{1, 2, 3}
	//arr2 := [6]int{1, 2, 5}
	//fmt.Println(arr, arr2)

	var slice []int
	slice = append(slice, 5, 5, 4, 32, 3)
	slice2 := make([]int, 6, 9)
	slice2 = append(slice2, 43, 32, 21, 23)

	fmt.Println(slice, len(slice), cap(slice))
	slice3 := slice[1:]
	slice3 = append(slice3, 14)
	fmt.Println(slice3, len(slice3), cap(slice3))
	slice[2] = 100

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))
}

func sum(a, b int) {
	v := a + b
	c := a == b
	fmt.Println(a, b, c, v)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func PyphfagorTrue(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	} else {
		return false
	}
}
