package main

import "fmt"

func NewtonThird(force float64) float64 {
	return -force
}

func main() {
	action := 1000.0
	reaction := NewtonThird(action)

	fmt.Println("Action:", action, "N")
	fmt.Println("Reaction:", reaction, "N")
}

// Эта программа демонстрирует третий закон Ньютона, который гласит, что на каждое действие есть равное и противоположное противодействие. В данном случае, если на объект действует сила в 1000 Н, то объект оказывает равную по величине, но противоположную по направлению силу в -1000 Н.
// Вывод программы:
// Action: 1000 N
// Reaction: -1000 N
