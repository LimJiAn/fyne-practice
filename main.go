package main

import (
	"fyne-practice/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

const (
	winTitle = "Practice"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme()) // if delete, defaultTheme
	utils.LogLifecycle(app)
	win := app.NewWindow(winTitle)
	// win.SetContent()
	win.Resize(fyne.NewSize(400, 800))
	win.ShowAndRun()
}
