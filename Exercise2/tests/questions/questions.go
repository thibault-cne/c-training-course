package questions

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func QuestionOne() {
	fmt.Println(color.BlueString("	[STARTING HELLO WORLD TEST]"))
	cmd := exec.Command("../build/questionOne.exe")

	_, stdout := startCommand(cmd)

	stdout.Scan()

	answer := stdout.Text()

	passedTests := 0

	if answer == "Hello World !!" {
		fmt.Println(color.GreenString("		✓ Success :"), "Hello world executable.")
		passedTests += 1
	} else {
		fmt.Println(color.RedString("		✕ Failed :"), "Hello world executable.")
		fmt.Printf("				Expected 'Hello World !!' but got '%s'\n", answer)
	}

	fmt.Printf("	Hello World : Passed %d/1 test\n", passedTests)
	fmt.Println(color.BlueString("	[ENDING HELLO WORLD TEST]"))
}

func QuestionTwo() {
	fmt.Println(color.BlueString("	[STARTING MAX3 TEST]"))

	cmds := make([]questionTwoCmdModel, 0)

	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "1", "10", "11"), result: "11"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "10", "10", "10"), result: "10"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "-1", "10", "2"), result: "10"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "10", "100", "30"), result: "100"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "-1", "-30", "-2"), result: "-1"})

	passedTests := 0

	for _, cmd := range cmds {
		_, stdout := startCommand(cmd.cmd)

		stdout.Scan()

		answer := stdout.Text()

		if answer == cmd.result {
			fmt.Println(color.GreenString("		✓ Success :"), "Max executable.")
			passedTests += 1
		} else {
			fmt.Println(color.RedString("		✕ Failed :"), "Max executable.")
			fmt.Printf("			Expected '11' but got '%s'\n", answer)
		}
	}

	fmt.Printf("		Max : Passed %d/5 test\n", passedTests)
	fmt.Println(color.BlueString("	[ENDING MAX3 TEST]"))
}
