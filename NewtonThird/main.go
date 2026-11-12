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
// В функции main мы задаем значение действия и выводим как действие, так и реакцию на экран.
// Пример вывода:
// Action: 1000 N
// Reaction: -1000 N
// Таким образом, мы демонстрируем принцип третьего закона Ньютона в действии.
