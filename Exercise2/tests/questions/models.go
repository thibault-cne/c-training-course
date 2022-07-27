package questions

import "os/exec"

type questionTwoCmdModel struct {
	cmd    *exec.Cmd
	result string
}
