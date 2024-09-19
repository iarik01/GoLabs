package main

import (
	"fmt"
	"time"
)

func ShowCurrentTimeTask1() {
	currentTime := time.Now()
	fmt.Println("Текущее время и дата:", currentTime.Format("02.01.2006 15:04:05"))
}
