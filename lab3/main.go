package main

import (
	"GoLabs/lab3/mathutils"
	"GoLabs/lab3/stringutils"
	"fmt"
)

func main() {

	var number int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&number)

	result := mathutils.Factorial(number)
	fmt.Println("Факториал:", result)

	var input string
	fmt.Print("Введите строку для переворота: ")
	fmt.Scan(&input)

	reversed := stringutils.Reverse(input)
	fmt.Println("Перевернутая строка:", reversed)

	createAndPrintArray()

	manipulateSlice()

	findLongestString()
}

func createAndPrintArray() {
	var arr [5]int
	fmt.Println("Введите 5 целых чисел:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Массив:", arr)
}

func manipulateSlice() {
	var arr = [5]int{1, 2, 3, 4, 5}
	slice := arr[:]

	slice = append(slice, 6)
	fmt.Println("Срез после добавления элемента:", slice)

	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Срез после удаления элемента:", slice)
}

func findLongestString() {
	strSlice := []string{"Go", "Programming", "Language", "is", "awesome"}
	var longest string

	for _, str := range strSlice {
		if len(str) > len(longest) {
			longest = str
		}
	}
	fmt.Println("Самая длинная строка:", longest)
}
