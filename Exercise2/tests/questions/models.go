package questions

import "os/exec"

type questionTwoCmdModel struct {
	cmd    *exec.Cmd
	result string
}

type questionThreeCmdModel struct {
	cmd       *exec.Cmd
	firstInt  string
	secondInt string
	thirdInt  string
	result    string
}
