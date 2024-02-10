package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(isChetnoe())
	//fmt.Println(MidleArifm())
	fmt.Println(AriaByPoints())

}

func isChetnoe() bool {
	var a int
	fmt.Println("Введите число на проверку четности")
	fmt.Scan(&a)
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func MidleArifm() (int, int, float64) {
	var list []int

	// Бесконечный цикл для получения ввода от пользователя
	for {
		fmt.Println("Введите число (для выхода введите 'escape'):")

		// Получение ввода от пользователя
		reader := bufio.NewReader(os.Stdin)   //  функция создает объект введеный с ведеными с клавиатуры данными
		input, err := reader.ReadString('\n') // читаем все до переноса строки(данные можем записывать только в одну стрноку)
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			continue
		} // при нормальной работе err - всегда nil, но как только это не так выводим информацию об ошибке

		// Удаление символа новой строки из ввода
		input = strings.TrimSpace(input)

		// Проверка, нужно ли выйти из цикла
		if input == "escape" {
			break
		}

		// Преобразование введенной строки в число
		num, err := strconv.Atoi(input) // функция переводит из str в инт
		if err != nil {
			fmt.Println("Ошибка: введите корректное число или 'escape' для выхода")
			continue
		}

		// Добавление числа в список
		list = append(list, num)
	}

	if len(list) == 0 {
		fmt.Println("Список пуст")
		return 0, 0, 0
	} // выкидываем информацию о пустом списке

	min := 200
	max := 0
	sum := 0
	for _, num := range list {
		sum += num
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	middle := float64(sum) / float64(len(list))

	return min, max, middle
}

type Point struct {
	x, y float64
}

func AriaByPoints() float64 {

	var points []Point // создаем список для точек

	// Ввод координат точек
	fmt.Println("Введите координаты четырех точек на плоскости:")
	for i := 0; i < 4; i++ {
		var x, y float64
		fmt.Println("Точка %d", i+1)
		fmt.Scan(&x, &y)
		points = append(points, Point{x, y})
	}

	isRectangle := isRectangle(points)
	if isRectangle {
		fmt.Println("Прямоугольник")
		area := rectangleArea(points)
		return area
	} else {
		fmt.Println("Треугольник")
		area := triangleArea(points)
		return area
	}

}
func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2)) // Pow возводит разницы кординат в квадрат, Sqrt  - берет корень
} // функция по поиску расстояния между двумя точками

func rectangleArea(points []Point) float64 {
	a := distance(points[0], points[1]) // Длина
	b := distance(points[1], points[2]) // ширина
	return a * b
}

func triangleArea(points []Point) float64 {
	a := distance(points[0], points[1]) // длины сторон
	b := distance(points[1], points[2])
	c := distance(points[2], points[0])
	s := (a + b + c) / 2                              // полупериметр
	return math.Sqrt(s * (s - a) * (s - b) * (s - c)) // формула Герона для поиска площадли
}

func isRectangle(point []Point) bool {
	d1 := distance(point[0], point[1])
	d2 := distance(point[1], point[2])
	d3 := distance(point[2], point[3])
	d4 := distance(point[3], point[0])
	d5 := distance(point[0], point[2])
	d6 := distance(point[1], point[3])

	return d1 == d3 && d2 == d4 && d5 == d6
}
