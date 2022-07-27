package questions

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func QuestionOne() {
	fmt.Println(color.GreenString("		[STARTING HELLO WORLD TEST]"))
	cmd := exec.Command("../build/questionOne.exe")

	so, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd.Start()

	stdout := bufio.NewScanner(so)

	stdout.Scan()

	answer := stdout.Text()

	passedTests := 0

	if answer == "Hello World !!" {
		fmt.Println(color.GreenString("			✓ Success :"), "Hello world executable.")
		passedTests += 1
	} else {
		fmt.Println(color.RedString("			✕ Failed :"), "Hello world executable.")
		fmt.Printf("				Expected 'Hello World !!' but got '%s'\n", answer)
	}

	fmt.Printf("		Hello World : Passed %d/1 test\n", passedTests)
	fmt.Println(color.GreenString("		[ENDING HELLO WORLD TEST]"))
}
