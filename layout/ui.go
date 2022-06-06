package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Ui struct {
	topmenu, bottommenu, members *widget.List
	Win                          fyne.Window
}

const (
	leftMenu = 3
)

func (u *Ui) MakeUI(w fyne.Window) fyne.CanvasObject {
	u.topmenu = widget.NewList(
		func() int {
			return leftMenu
		},
		func() fyne.CanvasObject {
			img := &canvas.Image{}
			img.SetMinSize(fyne.NewSize(theme.IconInlineSize()*2, theme.IconInlineSize()*2))
			return img
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			switch id {
			case 0:
				co.(*canvas.Image).Resource = theme.AccountIcon()
			case 1:
				co.(*canvas.Image).Resource = theme.MailSendIcon()
			case 2:
				co.(*canvas.Image).Resource = theme.MoreHorizontalIcon()
			}
			co.Refresh()
		})
	u.topmenu.OnSelected = func(id widget.ListItemID) {

	}
	u.members = widget.NewList(
		// members count
		func() int {
			return 10
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {

		})

	return container.NewBorder(nil, nil, u.topmenu, nil, u.members)
}
