package util

import (
	"os/exec"
)

func CallScript(script string) error {
	command := exec.Command("/bin/bash", script)
	err := command.Start()
	if err != nil {
		return err
	}
	return command.Wait()
}