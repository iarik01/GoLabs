package main

import "fmt"

func CalculateTask5() {
	var x, y float64
	fmt.Println("Введите два числа с плавающей запятой:")
	fmt.Scan(&x, &y)

	sum := x + y
	diff := x - y
	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", diff)
}
