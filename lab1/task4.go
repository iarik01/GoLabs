package main

import "fmt"

func CalculateTask4() {
	var x, y int
	fmt.Println("Введите два целых числа:")
	fmt.Scan(&x, &y)

	fmt.Println("Сумма:", x+y)
	fmt.Println("Разность:", x-y)
	fmt.Println("Произведение:", x*y)
	fmt.Println("Частное:", x/y)
}
