package main

import "fmt"

func CalculateAverageTask6() {
	var a, b, c float64
	fmt.Println("Введите три числа:")
	fmt.Scan(&a, &b, &c)

	average := (a + b + c) / 3
	fmt.Println("Среднее значение:", average)
}
