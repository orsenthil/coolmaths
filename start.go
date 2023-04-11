package main

import (
	"fmt"
	"math/rand"
	"time"
)

func start() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	var answer int
	var count int

	for i := 0; i < 10; i++ {
		n1 := r.Intn(10)
		n2 := r.Intn(10)

		fmt.Printf("What is %d x %d? ", n1, n2)
		fmt.Scanln(&answer)

		if answer == n1*n2 {
			fmt.Println("Correct!")
			count += 1
		} else {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("You got %d out of 10 correct!\n\nYour score is %f%%\n", count, float64(count)/10*100)
}

func main() {
	start()
}
