package main

import (
	"flag"
	"fmt"

	"exercise2.tests/questions"
	"github.com/fatih/color"
)

var (
	question = flag.Bool("questions", false, "Must be present to provide tests.")
)

func main() {
	// We get the value of the question flag passed in the call of the executable
	// For exemple : tests -question=q1
	flag.Parse()

	if !*question {
		fmt.Println(color.RedString("[ERROR]"), "Questions flag must be enabled to provide tests.")
		return
	}

	fmt.Println(color.YellowString("[STARTING TESTS]"))

	questions := flag.Args()

	for _, q := range questions {
		launchQuestion(q)
	}

	fmt.Println(color.YellowString("[END OF TESTS]"))
}

func launchQuestion(question string) {
	if question == "q11" || question == "question11" {
		questions.QuestionOneOne()
	}

	if question == "q22" || question == "question22" {
		questions.QuestionTwoTwo()
	}

	if question == "q23" || question == "question23" {
		questions.QuestionTwoThree()
	}

	if question == "q32" || question == "question32" {
		questions.QuestionThreeTwo()
	}
}
