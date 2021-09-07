package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MohamedElmdary/yggdrasil-connector/src/helpers"
	"github.com/MohamedElmdary/yggdrasil-connector/src/yggdrasil"
)

func loadUI(peers map[string][]string, onSelectPeer func(checked bool, peer string)) fyne.CanvasObject {
	items := []fyne.CanvasObject{}

	for country, values := range peers {
		if len(values) == 0 {
			continue
		}

		countryLabel := canvas.NewText(country, color.White)
		countryLabel.Alignment = fyne.TextAlignCenter
		countryLabel.TextStyle = fyne.TextStyle{Bold: true}
		countryLabel.TextSize = 18

		checkbox := widget.NewCheck(country, func(checked bool) {
			// onSelectPeer(checked, country)
		})
		items = append(items, checkbox)

		// items = append(items, countryLabel)

		// for _, value := range values {
		// 	checkbox := widget.NewCheck(value, func(checked bool) {
		// 		onSelectPeer(checked, value)
		// 	})
		// 	items = append(items, checkbox)
		// }
	}

	return container.New(
		layout.NewHBoxLayout(),
		container.NewVScroll(
			container.New(
				layout.NewVBoxLayout(),
				items...,
			),
		),
		container.New(
			layout.NewVBoxLayout(),
			widget.NewButton("Connect", func() {
				fmt.Println("clicked")
			}),
		),
	)
}

func updatePeers(peers map[string][]string, countries []string) {
	// Generate new configs
	cmd := exec.Command("yggdrasil", "-genconf")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	configs := string(output)
	configs = strings.Replace(configs, "Peers: []", "Peers: "+yggdrasil.ConcatPeers(peers, countries), 1)

	// // Remove old configs
	cmd = exec.Command("rm /tmp/ygg.conf")
	cmd.Output()

	// // Save new configs
	f, err := os.Create("/tmp/ygg.conf")
	if err != nil {
		return
	}
	defer f.Close()

	io.WriteString(f, configs)

	cmd = exec.Command("sudo", "yggdrasil", "-useconffile", "/tmp/ygg.conf")
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	// cmd.Process.Kill()
}

func removeCountry(countries []string, country string) []string {
	idx := helpers.FindIndex(countries, country)
	return append(countries[:idx], countries[idx+1:]...)
}

func main() {
	peers := yggdrasil.GetPeers()

	application := app.New()
	window := application.NewWindow("Yggdrasil Connector")
	window.Resize(fyne.NewSize(1, 600))

	selectedCountries := []string{}
	window.SetContent(loadUI(peers, func(checked bool, country string) {
		if checked {
			selectedCountries = append(selectedCountries, country)
		} else {
			selectedCountries = removeCountry(selectedCountries, country)
		}

		updatePeers(peers, selectedCountries)
	}))
	window.ShowAndRun()
}
