package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// question struct stores a single question and its corresponding answer.
type question struct {
	q, a string
}

type score int

// check handles a potential error.
// It stops execution of the program ("panics") if an error has happened.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// questions reads in questions and corresponding answers from a CSV file into a slice of question structs.
func questions() []question {
	f, err := os.Open("quiz-questions.csv")
	check(err)
	reader := csv.NewReader(f)
	table, err := reader.ReadAll()
	check(err)
	var questions []question
	for _, row := range table {
		questions = append(questions, question{q: row[0], a: row[1]})
	}
	return questions
}

// ask asks a question and returns an updated score depending on the answer.
func ask(s chan score, question <-chan question) {
	time.Sleep(5 * time.Millisecond)
	currentQuestion := <-question
	fmt.Println(currentQuestion.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	currentScore := <-s
	if strings.Compare(text, currentQuestion.a) == 0 {
		fmt.Println("Correct!\n")
		currentScore++
	} else {
		fmt.Println("Incorrect :-(\n")
	}
	s <- currentScore
}

func main() {
	s := make(chan score)
	qs := make(chan question)

	scoreInitialised := false

	timeOut := time.After(5 * time.Second)
	questionsLoop:
		for _, q := range(questions()) {
			select {
			case <-timeOut:
				fmt.Println("5 seconds are up.")
				break questionsLoop
			default:
				go ask(s, qs)
				qs <- q
				if scoreInitialised {
					s <- (<-s)
				} else {
					s <- score(0)
					scoreInitialised = true
				}
			}
		}

	finalScore := <-s
	fmt.Println("Final score", finalScore)

}
