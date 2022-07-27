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
	if question == "q1" || question == "question1" {
		questions.QuestionOne()
	}

	if question == "q2" || question == "question2" {
		questions.QuestionTwo()
	}
}
