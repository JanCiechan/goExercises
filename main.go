package main

import (
	"fmt"

	"example.com/go-exercises/project"
)

func main() {
	//pierwsze zadanko
	fmt.Println(project.WordCount("liczymy sobie slowa slowa jest fajnie fajnie fajnie"))
	//drugie zadanko
	f := project.Fibonacci()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f(), f(), f(), f(), f(), f())
	//trzecie zadanko
	project.QuizExercise()
}
