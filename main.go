package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	// "github.com/MohamedElmdary/yggdrasil-connector/src/yggdrasil"
)

func main() {
	// peers := yggdrasil.GetPeers()
	// fmt.Println(peers)

	application := app.New()
	window := application.NewWindow("Hello world")

	window.SetContent(widget.NewLabel("Hello"))
	window.ShowAndRun()

}
