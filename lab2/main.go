package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. Определение четного или нечетного числа
func CheckEvenOdd() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

// 2. Определение Positive/Negative/Zero
func CheckNumberType(n int) string {
	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func HandleCheckNumberType() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	result := CheckNumberType(number)
	fmt.Println("Результат:", result)
}

// 3. Вывод чисел от 1 до 10
func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 4. Длина строки

func HandleStringLength() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scan(&input)

	length := utf8.RuneCountInString(input)
	fmt.Println("Длина строки:", length)
}

// 5. Структура Rectangle и метод для вычисления площади
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func HandleRectangleArea() {
	var width, height float64
	fmt.Println("Введите ширину и высоту прямоугольника:")
	fmt.Scan(&width, &height)

	rect := Rectangle{width: width, height: height}
	fmt.Println("Площадь прямоугольника:", rect.Area())
}

// 6. Среднее значение двух чисел
func Average(a, b int) float64 {
	return float64(a+b) / 2
}

func HandleAverage() {
	var x, y int
	fmt.Print("Введите два целых числа: ")
	fmt.Scan(&x, &y)

	avg := Average(x, y)
	fmt.Println("Среднее значение:", avg)
}

func main() {
	for {
		fmt.Println("Выбор задания:")
		fmt.Println("1 - Определение четного или нечетного числа")
		fmt.Println("2 - Определение Positive/Negative/Zero")
		fmt.Println("3 - Вывод чисел от 1 до 10")
		fmt.Println("4 - Длина строки")
		fmt.Println("5 - Площадь прямоугольника")
		fmt.Println("6 - Среднее значение двух чисел")
		fmt.Println("7 - Выход")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			CheckEvenOdd()
		case 2:
			HandleCheckNumberType()
		case 3:
			PrintNumbers()
		case 4:
			HandleStringLength()
		case 5:
			HandleRectangleArea()
		case 6:
			HandleAverage()
		case 7:
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println("Неверный выбор!")
		}
	}
}
