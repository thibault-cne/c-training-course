package questions

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func QuestionOneOne() {
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

func QuestionTwoTwo() {
	fmt.Println(color.BlueString("	[STARTING MAX3 ARGV TEST]"))

	var (
		cmds        = make([]questionTwoCmdModel, 0)
		passedTests = 0
		totalTest   = 0
	)

	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "1", "10", "11"), result: "11"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "10", "10", "10"), result: "10"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "-1", "10", "2"), result: "10"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "10", "100", "30"), result: "100"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "-1", "-30", "-2"), result: "-1"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "-1", "-30"), result: "null"})
	cmds = append(cmds, questionTwoCmdModel{cmd: exec.Command("../build/questionThree.exe", "-1", "-30", "-2", "4"), result: "null"})

	for _, cmd := range cmds {
		totalTest++
		_, stdout := startCommand(cmd.cmd)

		stdout.Scan()

		answer := stdout.Text()

		if cmd.result == "null" && answer == "" {
			fmt.Println(color.GreenString("		✓ Success :"), "Max argv executable.")
			passedTests += 1
			continue
		}

		if answer == cmd.result {
			fmt.Println(color.GreenString("		✓ Success :"), "Max argv executable.")
			passedTests += 1
		} else {
			fmt.Println(color.RedString("		✕ Failed :"), "Max argv executable.")
			fmt.Printf("			Expected '%s' but got '%s'\n", cmd.result, answer)
		}
	}

	fmt.Printf("		Max : Passed %d/%d test\n", passedTests, totalTest)
	fmt.Println(color.BlueString("	[ENDING MAX3 ARGV TEST]"))
}

func QuestionTwoThree() {
	fmt.Println(color.BlueString("	[STARTING MAX3 STDIN TEST]"))

	var (
		cmds        = make([]questionThreeCmdModel, 0)
		passedTests = 0
		totalTest   = 0
	)

	cmds = append(cmds, questionThreeCmdModel{
		cmd:      exec.Command("../build/questionFour.exe"),
		firstInt: "10", secondInt: "12", thirdInt: "15", result: "15"})
	cmds = append(cmds, questionThreeCmdModel{
		cmd:      exec.Command("../build/questionFour.exe"),
		firstInt: "-1", secondInt: "12", thirdInt: "0", result: "12"})
	cmds = append(cmds, questionThreeCmdModel{
		cmd:      exec.Command("../build/questionFour.exe"),
		firstInt: "-10", secondInt: "-1", thirdInt: "0", result: "0"})
	cmds = append(cmds, questionThreeCmdModel{
		cmd:      exec.Command("../build/questionFour.exe"),
		firstInt: "0", secondInt: "0", thirdInt: "0", result: "0"})
	cmds = append(cmds, questionThreeCmdModel{
		cmd:      exec.Command("../build/questionFour.exe"),
		firstInt: "-100", secondInt: "-90", thirdInt: "-100", result: "-90"})

	for _, cmd := range cmds {
		totalTest++
		stdin, stdout := startCommand(cmd.cmd)

		stdin.Write([]byte(cmd.firstInt + "\n"))
		stdin.Write([]byte(cmd.secondInt + "\n"))
		stdin.Write([]byte(cmd.thirdInt + "\n"))

		stdout.Scan()

		answer := stdout.Text()

		if answer == cmd.result {
			fmt.Println(color.GreenString("		✓ Success :"), "Max stdin executable.")
			passedTests += 1
		} else {
			fmt.Println(color.RedString("		✕ Failed :"), "Max stdin executable.")
			fmt.Printf("			Expected '%s' but got '%s'\n", cmd.result, answer)
		}
	}

	fmt.Printf("		Max : Passed %d/%d test\n", passedTests, totalTest)
	fmt.Println(color.BlueString("	[ENDING MAX3 STDIN TEST]"))
}
