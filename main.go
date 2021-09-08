package main

import (
	"fmt"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MohamedElmdary/yggdrasil-connector/src/helpers"
	"github.com/MohamedElmdary/yggdrasil-connector/src/yggdrasil"
)

func loadUI(peers map[string][]string, onSelectPeer func(bool, string)) fyne.CanvasObject {
	items := []fyne.CanvasObject{}
	tempItems := []*widget.Check{}

	for country, values := range peers {
		if len(values) == 0 {
			continue
		}
		c := country
		checkbox := widget.NewCheck(country, func(checked bool) {
			onSelectPeer(checked, c)
		})
		items = append(items, checkbox)
		tempItems = append(tempItems, checkbox)
	}

	countriesList := container.NewVScroll(
		container.New(
			layout.NewVBoxLayout(),
			items...,
		),
	)
	countriesList.SetMinSize(fyne.NewSize(1, 600))
	btnConnect := helpers.CreateConnectionBtn(tempItems)
	btnConnect.Resize(fyne.NewSize(1, 100))
	layout1 := layout.NewVBoxLayout()
	contApp := container.New(layout1, countriesList, btnConnect)
	return contApp
}

func main() {
	peers := yggdrasil.GetPeers()

	// Stop yggdrasil service if it's working
	cmd := exec.Command("systemctl", "stop", "yggdrasil")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	// Init conf file
	yggdrasil.UpdatePeers(peers, []string{})

	application := app.New()
	window := application.NewWindow("Yggdrasil Connector")
	window.Resize(fyne.NewSize(300, 600))
	window.SetContent(loadUI(peers, yggdrasil.UpdatePeersHandler(peers)))
	window.ShowAndRun()
}
