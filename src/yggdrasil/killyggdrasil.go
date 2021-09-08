package yggdrasil

import (
	"fmt"
	"os/exec"
)

func Kill() {
	cmd := exec.Command("killall", "yggdrasil")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}
