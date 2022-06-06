package main

import (
	"github.com/LimJiAn/fyne-practice/layout"
	"github.com/LimJiAn/fyne-practice/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	winTitle = "Practice"
)

func main() {
	app := app.New()
	// app.Settings().SetTheme(theme.LightTheme()) // if delete, defaultTheme
	utils.LogLifecycle(app)
	win := app.NewWindow(winTitle)
	u := layout.Ui{Win: win}
	win.SetContent(u.MakeUI(win))
	win.Resize(fyne.NewSize(400, 800))
	win.ShowAndRun()
}
