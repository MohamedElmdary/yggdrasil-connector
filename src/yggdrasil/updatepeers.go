package yggdrasil

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/MohamedElmdary/yggdrasil-connector/src/constants"
)

func UpdatePeers(peers map[string][]string, countries []string) {

	// Generate new configs
	cmd := exec.Command("yggdrasil", "-genconf")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	configs := string(output)
	configs = strings.Replace(configs, "Peers: []", "Peers: "+ConcatPeers(peers, countries), 1)

	// Re-Write the file
	if _, err := os.Stat(constants.Path); err == nil {
		os.Remove(constants.Path)
	}

	f, err := os.Create(constants.Path)
	if err != nil {
		fmt.Println("Error!", err)
	}
	defer f.Close()

	fmt.Fprintln(f, configs)
}
