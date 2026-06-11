package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/iliasgal/network-monitor/network"
)

func packetCapTab() (packetTab *container.TabItem) {

	packetAppContainer := container.NewAppTabs()

	settingsTab := packetCapSettingsTab()
	packetAppContainer.Append(settingsTab)

	packetTab = container.NewTabItem("Packet Capture", packetAppContainer)

	return

}

func packetCapSettingsTab() (settingsTab *container.TabItem) {

	capSettingsTabContainer := container.NewHBox()

	netDevNameBaseText := "Name:      \n"
	netDevName := widget.NewLabel(netDevNameBaseText + network.NetDev.Name)
	netDevSelectionBaseText := "Interface: \n"
	netDevSelection := widget.NewLabel(netDevSelectionBaseText + network.NetDev.Description)

	selectButton := widget.NewButton("Select Interface", func() {
		netDevModal()
		go func() {
			<-updateInfo
			fyne.DoAndWait(func() {
				netDevName.SetText(netDevNameBaseText + network.NetDev.Name)
				netDevSelection.SetText(netDevSelectionBaseText + network.NetDev.Description)
			})
		}()

	})

	netDevSection := container.NewVBox(netDevName, netDevSelection, layout.NewSpacer(), selectButton)
	capSettingsTabContainer.Add(netDevSection)
	settingsTab = container.NewTabItem("Settings", capSettingsTabContainer)
	return
}
