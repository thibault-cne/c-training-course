package questions

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func QuestionOneOne() {
	fmt.Println(color.BlueString("	[STARTING HELLO WORLD TEST]"))
	cmd := exec.Command("../build/questionOneOne.exe")

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
		cmds        = make([]commandResultModel, 0)
		passedTests = 0
		totalTest   = 0
	)

	cmds = append(cmds, commandResultModel{cmd: exec.Command("../build/questionTwoTwo.exe", "1", "10", "11"), result: "11"})
	cmds = append(cmds, commandResultModel{cmd: exec.Command("../build/questionTwoTwo.exe", "10", "10", "10"), result: "10"})
	cmds = append(cmds, commandResultModel{cmd: exec.Command("../build/questionTwoTwo.exe", "-1", "10", "2"), result: "10"})
	cmds = append(cmds, commandResultModel{cmd: exec.Command("../build/questionTwoTwo.exe", "10", "100", "30"), result: "100"})
	cmds = append(cmds, commandResultModel{cmd: exec.Command("../build/questionTwoTwo.exe", "-1", "-30", "-2"), result: "-1"})
	cmds = append(cmds, commandResultModel{cmd: exec.Command("../build/questionTwoTwo.exe", "-1", "-30"), result: "null"})
	cmds = append(cmds, commandResultModel{cmd: exec.Command("../build/questionTwoTwo.exe", "-1", "-30", "-2", "4"), result: "null"})

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

	fmt.Printf("		Max argv : Passed %d/%d test\n", passedTests, totalTest)
	fmt.Println(color.BlueString("	[ENDING MAX3 ARGV TEST]"))
}

func QuestionTwoThree() {
	fmt.Println(color.BlueString("	[STARTING MAX3 STDIN TEST]"))

	var (
		cmds        = make([]commandThreeStdinArgsModel, 0)
		passedTests = 0
		totalTest   = 0
	)

	cmds = append(cmds, commandThreeStdinArgsModel{
		cmd:      exec.Command("../build/questionTwoThree.exe"),
		firstInt: "10", secondInt: "12", thirdInt: "15", result: "15"})
	cmds = append(cmds, commandThreeStdinArgsModel{
		cmd:      exec.Command("../build/questionTwoThree.exe"),
		firstInt: "-1", secondInt: "12", thirdInt: "0", result: "12"})
	cmds = append(cmds, commandThreeStdinArgsModel{
		cmd:      exec.Command("../build/questionTwoThree.exe"),
		firstInt: "-10", secondInt: "-1", thirdInt: "0", result: "0"})
	cmds = append(cmds, commandThreeStdinArgsModel{
		cmd:      exec.Command("../build/questionTwoThree.exe"),
		firstInt: "0", secondInt: "0", thirdInt: "0", result: "0"})
	cmds = append(cmds, commandThreeStdinArgsModel{
		cmd:      exec.Command("../build/questionTwoThree.exe"),
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

	fmt.Printf("		Max stdin : Passed %d/%d test\n", passedTests, totalTest)
	fmt.Println(color.BlueString("	[ENDING MAX3 STDIN TEST]"))
}

func QuestionThreeTwo() {
	fmt.Println(color.BlueString("	[STARTING DAYOFBIRTH ARGV TEST]"))

	var (
		cmds        = make([]commandResultModel, 0)
		passedTests = 0
		totalTest   = 0
	)

	cmds = append(cmds, commandResultModel{
		cmd:    exec.Command("../build/questionThreeTwo.exe", "Thibault", "C.", "23/03/2001"),
		result: "Thibault C. was born on a Friday.",
	})
	cmds = append(cmds, commandResultModel{
		cmd:    exec.Command("../build/questionThreeTwo.exe", "Gerald", "O.", "19/03/1978"),
		result: "Gerald O. was born on a Sunday.",
	})
	cmds = append(cmds, commandResultModel{
		cmd:    exec.Command("../build/questionThreeTwo.exe", "John", "D.", "23/12/1970"),
		result: "John D. was born on a Wednesday.",
	})
	cmds = append(cmds, commandResultModel{
		cmd:    exec.Command("../build/questionThreeTwo.exe", "John", "C.", "30/14/1978"),
		result: "The date 30/14/1978 don't exist.",
	})
	cmds = append(cmds, commandResultModel{
		cmd:    exec.Command("../build/questionThreeTwo.exe", "Thibault", "C."),
		result: "",
	})

	for _, cmd := range cmds {
		totalTest++
		_, stdout := startCommand(cmd.cmd)

		stdout.Scan()

		answer := stdout.Text()

		if answer == cmd.result {
			fmt.Println(color.GreenString("		✓ Success :"), "Dayofbirth argv executable.")
			passedTests += 1
		} else {
			fmt.Println(color.RedString("		✕ Failed :"), "Dayofbirth argv executable.")
			fmt.Printf("			Expected '%s' but got '%s'\n", cmd.result, answer)
		}
	}

	fmt.Printf("		Dayofbirth argv : Passed %d/%d test\n", passedTests, totalTest)
	fmt.Println(color.BlueString("	[ENDING DAYOFBIRTH ARGV TEST]"))
}

func QuestionThreeThree() {
	fmt.Println(color.BlueString("	[STARTING DAYOFBIRTH STDIN TEST]"))

	var (
		cmds        = make([]commandStdinArgsModel, 0)
		passedTests = 0
		totalTest   = 0
		executable  = "../build/questionThreeThree.exe"
	)

	cmds = append(cmds, commandStdinArgsModel{
		cmd:    exec.Command(executable),
		args:   []string{"Thibault C.", "23/03/2001"},
		result: "Thibault C. was born on a Friday.",
	})
	cmds = append(cmds, commandStdinArgsModel{
		cmd:    exec.Command(executable),
		args:   []string{"Gerald O.", "19/03/1978"},
		result: "Gerald O. was born on a Sunday.",
	})
	cmds = append(cmds, commandStdinArgsModel{
		cmd:    exec.Command(executable),
		args:   []string{"John D.", "23/12/1970"},
		result: "John D. was born on a Wednesday.",
	})
	cmds = append(cmds, commandStdinArgsModel{
		cmd:    exec.Command(executable),
		args:   []string{"John C.", "30/14/1978"},
		result: "The date 30/14/1978 don't exist.",
	})

	for _, cmd := range cmds {
		totalTest++
		stdin, stdout := startCommand(cmd.cmd)

		for _, arg := range cmd.args {
			stdin.Write([]byte(arg + "\n"))
		}

		stdout.Scan()

		answer := stdout.Text()

		if answer == cmd.result {
			fmt.Println(color.GreenString("		✓ Success :"), "Dayofbirth stdin executable.")
			passedTests += 1
		} else {
			fmt.Println(color.RedString("		✕ Failed :"), "Dayofbirth stdin executable.")
			fmt.Printf("			Expected '%s' but got '%s'\n", cmd.result, answer)
		}
	}

	fmt.Printf("		Dayofbirth stdin : Passed %d/%d test\n", passedTests, totalTest)
	fmt.Println(color.BlueString("	[ENDING DAYOFBIRTH STDIN TEST]"))
}
