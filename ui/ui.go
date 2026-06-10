package ui

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func InitUi() {
	a = app.New()
	w = a.NewWindow("Hello Wordle")
	w.Resize(fyne.NewSize(900, 600))
	w.SetContent(widget.NewLabel("wordles r us"))
	ChooseNetDev()

	var ok bool
	if desk, ok = a.(desktop.App); ok {
		desktopMode = true
		tray = fyne.NewMenu("NetDiag",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}), fyne.NewMenuItem("Quit", func() { os.Exit(0) }))
		desk.SetSystemTrayMenu(tray)
		w.SetCloseIntercept(func() {
			w.Hide()
		})
	}
	w.ShowAndRun()
}
