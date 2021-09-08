package helpers

import (
	"fmt"
	"os/exec"

	"fyne.io/fyne/v2/widget"
)

const (
	CONNECT    = "Connect"
	DISCONNECT = "Disconnect"
)

func CreateConnectionBtn(checks []*widget.Check) *widget.Button {
	var cmd *exec.Cmd
	var btn *widget.Button

	btn = widget.NewButton("Connect", func() {
		if cmd != nil {
			btn.SetText(CONNECT)
			cmd.Process.Kill()
			cmd = nil
			enableList(checks)
			return
		}

		go func() {
			disableList(checks)
			btn.SetText(DISCONNECT)
			cmd = exec.Command("yggdrasil", "-useconffile", "/tmp/ygg.conf")
			_, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
		}()
	})
	return btn
}

func enableList(checks []*widget.Check) {
	for _, check := range checks {
		check.Enable()
	}
}

func disableList(checks []*widget.Check) {
	for _, check := range checks {
		check.Disable()
	}
}
