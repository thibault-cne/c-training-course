package questions

import "os/exec"

type commandResultModel struct {
	cmd    *exec.Cmd
	result string
}

type commandThreeStdinArgsModel struct {
	cmd       *exec.Cmd
	firstInt  string
	secondInt string
	thirdInt  string
	result    string
}

type commandStdinArgsModel struct {
	cmd    *exec.Cmd
	args   []string
	result string
}
