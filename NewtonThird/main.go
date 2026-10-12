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

// Законы Ньютона: Третий закон Ньютона утверждает, что на каждое действие есть равное и противоположное противодействие.
// В данном коде функция NewtonThird принимает силу (action) и возвращает противоположную силу (reaction).
