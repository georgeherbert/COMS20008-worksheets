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
func ask(s score, question <-chan question, sc chan<- score) {
	currentQuestion := <-question
	fmt.Println(currentQuestion.q)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter answer: ")
	scanner.Scan()
	text := scanner.Text()
	if strings.Compare(text, currentQuestion.a) == 0 {
		fmt.Println("Correct!\n")
		s++
	} else {
		fmt.Println("Incorrect :-(\n")
	}
	sc <- s
}

func main() {
	sc := make(chan score)
	qs := make(chan question)
	s := score(0)

	timeOut := time.After(5 * time.Second)
	questionsLoop:
		for _, q := range questions()[1:] {
			go ask(s, qs, sc)
			qs <- q
			select {
			case sNew := <-sc:
				s = sNew
			case <-timeOut:
				fmt.Println("\n\nTime's up!")
				break questionsLoop
			}
		}

	fmt.Println("Final score", s)
}