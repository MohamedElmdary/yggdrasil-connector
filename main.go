package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MohamedElmdary/yggdrasil-connector/src/helpers"
	"github.com/MohamedElmdary/yggdrasil-connector/src/yggdrasil"
)

func loadUI(peers map[string][]string, countries []string, onSelectPeer func(bool, string)) fyne.CanvasObject {
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

		if helpers.FindIndex(countries, country) > -1 {
			checkbox.Checked = true
		}

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
	// check if yggdrasil exists on system
	yggdrasil.CheckYggdrasil()

	// get all peers
	peers := yggdrasil.GetPeers()

	// Stop yggdrasil service if it's working
	yggdrasil.Kill()

	// Load inital selected countries
	countries := helpers.LoadCountries()

	// Init conf file
	yggdrasil.UpdatePeers(peers, countries)

	application := app.New()
	window := application.NewWindow("Yggdrasil Connector")
	window.Resize(fyne.NewSize(300, 600))
	window.SetContent(loadUI(peers, countries, yggdrasil.UpdatePeersHandler(peers, countries)))
	window.ShowAndRun()
}
