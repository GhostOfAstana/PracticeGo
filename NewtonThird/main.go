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

// Третий закон Ньютона гласит, что на каждое действие есть равное и противоположное противодействие.
// В данном коде функция NewtonThird принимает силу (action) и возвращает противоположную силу (reaction).
// В функции main мы задаем значение действия (action) и вызываем функцию NewtonThird, чтобы получить противоположную силу (reaction).
// Затем мы выводим значения действия и противодействия на экран.
