package main

import (
	"fmt"
)

func main() {
	fmt.Println("Выбор задания:")
	fmt.Println("1 - Вывести текущее время и дату")
	fmt.Println("2 - Переменные различных типов")
	fmt.Println("3 - Краткая форма объявления переменных")
	fmt.Println("4 - Арифметические операции")
	fmt.Println("5 - Сумма и разность двух чисел")
	fmt.Println("6 - Среднее значение трёх чисел")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		ShowCurrentTimeTask1() // task1.go
	case 2:
		ShowVariablesTask2() // task2.go
	case 3:
		ShortVariablesTask3() // task3.go
	case 4:
		CalculateTask4() // task4.go
	case 5:
		CalculateTask5() // task5.go
	case 6:
		CalculateAverageTask6() //  task6.go
	default:
		fmt.Println("Неверный выбор!")
	}
}
