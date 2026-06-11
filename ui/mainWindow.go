package ui

import (
	"fyne.io/fyne/v2/container"
)

func mainWindow() {
	tabWindow := container.NewAppTabs()

	packetTab := packetCapTab()
	tabWindow.Append(packetTab)

	w.SetContent(tabWindow)
}
