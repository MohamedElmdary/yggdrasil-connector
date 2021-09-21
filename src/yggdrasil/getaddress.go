package yggdrasil

import (
	"fmt"
	"os/exec"

	"github.com/MohamedElmdary/yggdrasil-connector/src/constants"
)

func GetAddress() string {
	cmd := exec.Command("yggdrasil", "-useconffile", constants.Path, "-address")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return "Couldn't load address"
	}

	return string(output)
}
