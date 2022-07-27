package main

import (
	"flag"
	"fmt"

	"exercise2.tests/questions"
	"github.com/fatih/color"
)

var (
	question = flag.String("question", "q0", "Question you want to test. Must be entered like this : q1 or question1")
)

func main() {
	fmt.Println(color.GreenString("[TESTS INITIALIZATION]"))

	// We get the value of the question flag passed in the call of the executable
	// For exemple : tests -question=q1
	flag.Parse()

	questionNumber := checkQuestionsExistence(*question)

	if questionNumber == -1 {
		fmt.Println(color.RedString("[ERROR]"), "Wrong question number the question must be entered like this : q1 or question1")
		return
	}

	fmt.Println(color.GreenString("[STARTING TESTS]"))

	if questionNumber == 1 {
		questions.QuestionOne()
	}

	fmt.Println(color.GreenString("[END OF TESTS]"))
}

func checkQuestionsExistence(question string) int {
	if question == "q1" || question == "question1" {
		return 1
	}

	return -1
}
