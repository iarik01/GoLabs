package main

import (
	"fmt"
	"strings"
)

// 1. Создание карты с именами людей и их возрастами, добавление нового человека и вывод всех записей
func createAndPrintPeople() {
	people := map[string]int{
		"Иван":   25,
		"Анна":   30,
		"Сергей": 40,
	}

	// Добавляем нового человека
	people["Мария"] = 22

	// Выводим всех людей
	fmt.Println("Записи в карте:")
	for name, age := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
	}
}

// 2. Функция для вычисления среднего возраста людей в карте
func averageAge(people map[string]int) float64 {
	totalAge := 0
	count := 0

	for _, age := range people {
		totalAge += age
		count++
	}

	if count == 0 {
		return 0
	}

	return float64(totalAge) / float64(count)
}

func calculateAverageAge() {
	people := map[string]int{
		"Иван":   25,
		"Анна":   30,
		"Сергей": 40,
		"Мария":  22,
	}

	avg := averageAge(people)
	fmt.Printf("Средний возраст: %.2f\n", avg)
}

// 3. Программа для удаления записи из карты по имени
func removePerson(people map[string]int, name string) {
	delete(people, name) // Удаляем запись по ключу (имени)
}

func handleRemovePerson() {
	people := map[string]int{
		"Иван":   25,
		"Анна":   30,
		"Сергей": 40,
		"Мария":  22,
	}

	var nameToRemove string
	fmt.Print("Введите имя для удаления: ")
	fmt.Scan(&nameToRemove)

	removePerson(people, nameToRemove)

	fmt.Println("Оставшиеся записи:")
	for name, age := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
	}
}

// 4. Программа для считывания строки и вывода её в верхнем регистре
func toUpperCase() {
	var input string
	fmt.Print("Введите строку: ")
	fmt.Scan(&input)

	upper := strings.ToUpper(input)
	fmt.Println("Строка в верхнем регистре:", upper)
}

// 5. Программа для считывания нескольких чисел и вывода их суммы
func sumOfNumbers() {
	var n, num, sum int

	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&num)
		sum += num
	}

	fmt.Println("Сумма чисел:", sum)
}

// 6. Программа для считывания массива целых чисел и вывода их в обратном порядке
func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func handleReverseArray() {
	var n int

	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Числа в обратном порядке:", reverseArray(arr))
}

func main() {
	createAndPrintPeople()

	calculateAverageAge()

	handleRemovePerson()

	toUpperCase()

	sumOfNumbers()

	handleReverseArray()
}
