package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MohamedElmdary/yggdrasil-connector/src/helpers"
	"github.com/MohamedElmdary/yggdrasil-connector/src/yggdrasil"
)

func loadUI(peers map[string][]string, connectionHandler func(*widget.Button) *widget.Button, onSelectPeer func(bool, string)) fyne.CanvasObject {
	items := []fyne.CanvasObject{}

	for country, values := range peers {
		if len(values) == 0 {
			continue
		}
		c := country
		checkbox := widget.NewCheck(country, func(checked bool) {
			onSelectPeer(checked, c)
		})
		items = append(items, checkbox)
	}

	countriesList := container.NewVScroll(
		container.New(
			layout.NewVBoxLayout(),
			items...,
		),
	)
	countriesList.SetMinSize(fyne.NewSize(1, 600))
	btnConnect := connectionHandler(widget.NewButton("Connect", func() {}))
	btnConnect.Resize(fyne.NewSize(1, 100))
	layout1 := layout.NewVBoxLayout()
	contApp := container.New(layout1, countriesList, btnConnect)
	return contApp
	// return container.NewVSplit(countriesList, btnConnect)
}

func updatePeers(peers map[string][]string, countries []string) {
	// create conf file if not exists
	if _, err := os.Stat("/tmp/ygg.conf"); err != nil {
		os.Create("/tmp/ygg.conf")
	}

	// Generate new configs
	cmd := exec.Command("yggdrasil", "-genconf")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	configs := string(output)
	configs = strings.Replace(configs, "Peers: []", "Peers: "+yggdrasil.ConcatPeers(peers, countries), 1)

	// // Save new configs
	f, err := os.Create("/tmp/ygg.conf")
	if err != nil {
		return
	}
	defer f.Close()

	io.WriteString(f, configs)

}

func removeCountry(countries []string, country string) []string {
	idx := helpers.FindIndex(countries, country)
	return append(countries[:idx], countries[idx+1:]...)
}

func main() {
	peers := yggdrasil.GetPeers()
	selectedCountries := []string{}
	connectionHandler := func(btn *widget.Button) *widget.Button {
		var cmd *exec.Cmd

		btn.OnTapped = func() {
			if cmd != nil {
				btn.SetText("Connect")
				cmd.Process.Kill()
				cmd = nil
			} else {
				go func() {
					btn.SetText("Disconnect")
					cmd = exec.Command("yggdrasil", "-useconffile", "/tmp/ygg.conf")
					_, err := cmd.Output()
					if err != nil {
						fmt.Println(err)
					}
				}()
			}
		}
		return btn
	}

	// init conf file
	updatePeers(peers, []string{})

	application := app.New()
	window := application.NewWindow("Yggdrasil Connector")
	window.Resize(fyne.NewSize(300, 600))

	window.SetContent(loadUI(peers, connectionHandler, func(checked bool, country string) {
		if checked {
			selectedCountries = append(selectedCountries, country)
		} else {
			selectedCountries = removeCountry(selectedCountries, country)
		}

		updatePeers(peers, selectedCountries)
	}))
	window.ShowAndRun()
}
