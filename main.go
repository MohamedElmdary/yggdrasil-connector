package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/MohamedElmdary/yggdrasil-connector/src/yggdrasil"
)

func loadUI(peers map[string][]string) fyne.CanvasObject {
	items := []fyne.CanvasObject{}

	for country, values := range peers {
		if len(values) == 0 {
			continue
		}

		countryLabel := canvas.NewText(country, color.White)
		countryLabel.Alignment = fyne.TextAlignCenter
		countryLabel.TextStyle = fyne.TextStyle{Bold: true}
		countryLabel.TextSize = 20

		items = append(items, countryLabel)

		for _, value := range values {
			checkbox := widget.NewCheck(value, func(checked bool) {
				fmt.Println(checked)
			})
			items = append(items, checkbox)
		}
	}

	return container.NewVScroll(
		container.New(
			layout.NewHBoxLayout(),
			container.New(
				layout.NewVBoxLayout(),
				items...,
			),
		),
	)
}

func main() {
	peers := yggdrasil.GetPeers()

	application := app.New()
	window := application.NewWindow("Hello world")
	window.Resize(fyne.NewSize(250, 600))

	window.SetContent(loadUI(peers))
	window.ShowAndRun()
}
