package project

import (
	"encoding/csv"

	"fmt"
	"log"
	"os"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, w := range strings.Fields(s) {
		wc[w]++
	}
	return wc
}
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type question struct {
	problem string
	result  string
}

type quiz struct {
	questions []question
	score     int
}

func (q *quiz) ask() {
	q.score = 0
	for _, question := range q.questions {
		fmt.Println(question.problem)
		var answer string
		fmt.Scanln(&answer)
		if answer == question.result {
			q.score++
		}
	}
}

func readCSV(p string) []question {
	qfile, err := os.Open(p)
	if err != nil {
		log.Fatalf("could not open file: %v\n", err)
	}
	defer qfile.Close()
	csvrd := csv.NewReader(qfile)
	records, err := csvrd.ReadAll()
	if err != nil {
		log.Fatalf("could not parse .csv: %v\n", err)
	}
	var questions []question
	for _, record := range records {
		q := question{problem: record[0], result: record[1]}
		questions = append(questions, q)
	}
	return questions
}

func QuizExercise() {

	questions := readCSV("quiz.csv")
	qz := quiz{questions: questions, score: 0}
	resultCh := make(chan quiz)
	go func() {
		qz.ask()
		resultCh <- qz
	}()

	<-resultCh

	fmt.Println("Score: ", qz.score)
	fmt.Println("Amount of questions: ", len(qz.questions))
}
