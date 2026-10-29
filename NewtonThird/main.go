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

// This program demonstrates Newton's Third Law of Motion, which states that for every action, there is an equal and opposite reaction. The `NewtonThird` function takes a force as input and returns the negative of that force, representing the reaction force. In the `main` function, we define an action force of 1000 N and calculate the corresponding reaction force using the `NewtonThird` function, then print both values.
// Таким образом, программа иллюстрирует принцип действия и противодействия в физике, показывая, что сила действия и сила реакции равны по величине, но противоположны по направлению.
