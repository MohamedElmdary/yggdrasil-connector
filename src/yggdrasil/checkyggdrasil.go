package yggdrasil

import (
	"image/color"
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CheckYggdrasil() {
	cmd := exec.Command("which", "yggdrasil")
	output, _ := cmd.Output()
	if string(output) == "" {
		app := app.New()
		w := app.NewWindow("Error!")
		w.Resize(fyne.NewSize(400, 150))

		title := canvas.NewText("Yggdrasil not found.", color.White)
		title.TextStyle = fyne.TextStyle{Bold: true}

		w.SetContent(
			container.New(
				layout.NewVBoxLayout(),
				title,
				canvas.NewText("Yggdrasil was not found on system consider install it from offical website.", color.White),
				canvas.NewText("https://yggdrasil-network.github.io/installation.html", color.RGBA{R: 63, G: 81, B: 181}),

				container.New(
					layout.NewHBoxLayout(),
					layout.NewSpacer(),
					widget.NewButton("Ok", func() {
						log.Fatal("Yggdrasil not found.")
					}),
					layout.NewSpacer(),
				),
			),
		)

		w.ShowAndRun()
	}
}
