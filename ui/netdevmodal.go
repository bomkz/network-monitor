package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/iliasgal/network-monitor/network"
)

func (v *verticalSepLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(2, 0) // enforce 2px width, height comes from parent
}
func (v *verticalSepLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, o := range objects {
		o.Resize(size) // fill whatever size the parent gives
	}
}

func netDevModal() {
	var netList []string
	for _, x := range network.NetDevs {
		netList = append(netList, x.Description)
	}

	interfaceNameTextLabel := widget.NewLabel("Interface Name:")
	interfaceNameValueLabel := widget.NewLabel("")

	interfaceAddressTextLabel := widget.NewLabel("Interface Addresses:")
	interfaceAddressLabel := widget.NewLabel("")
	// Declare modal here so the button closure can reference it
	var modal *widget.PopUp

	netDevChangeFunc := func(s string) {
		for _, x := range network.NetDevs {
			if x.Description == s {
				network.NetDev = x
				interfaceNameValueLabel.SetText(x.Name)
				addressText := ""
				for _, y := range x.Addresses {
					addressText += y.IP.String() + "\n"
				}
				interfaceAddressLabel.SetText(addressText)
			}
		}
	}

	separatorRect := canvas.NewRectangle(color.Gray{Y: 100})
	separatorContainer := container.New(&verticalSepLayout{}, separatorRect)

	selectButtonWidget := widget.NewButton("Select", func() {
		if network.NetDev.Name == "" {
			return
		}
		modal.Hide()
		updateInfo <- true
	})

	leftLabels := container.NewVBox(
		widget.NewLabel("Network Interface"),
		interfaceNameTextLabel,
		interfaceAddressTextLabel,
		layout.NewSpacer(),
		selectButtonWidget,
	)
	leftSide := container.NewHBox(leftLabels, separatorContainer)

	rightValues := container.NewVBox(
		widget.NewSelect(netList, netDevChangeFunc),
		interfaceNameValueLabel,
		interfaceAddressLabel,
	)

	infoArea := container.NewBorder(nil, nil, leftSide, nil, rightValues)
	paddedContainer := container.NewPadded(infoArea) // pass content directly, no separate Add()

	modal = widget.NewModalPopUp(paddedContainer, w.Canvas())
	modal.Resize(fyne.NewSize(700, 400))
	modal.Show()
}
