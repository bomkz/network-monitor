package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/iliasgal/network-monitor/network"
)

func InitUi() {
	a = app.New()
	w = a.NewWindow("Hello Wordle")
	w.Resize(fyne.NewSize(900, 600))
	w.SetContent(widget.NewLabel("wordles r us"))
	ChooseNetDev()
	w.ShowAndRun()

}

// Place this outside the function, at package level
type verticalSepLayout struct{}

func (v *verticalSepLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(2, 0) // enforce 2px width, height comes from parent
}
func (v *verticalSepLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, o := range objects {
		o.Resize(size) // fill whatever size the parent gives
	}
}

type smartScrollContainer struct {
	widget.BaseWidget
	label     *widget.Label
	scroll    *container.Scroll
	maxHeight float32
}

func newSmartScroll(maxHeight float32) *smartScrollContainer {
	lbl := widget.NewLabel("")
	s := &smartScrollContainer{
		label:     lbl,
		scroll:    container.NewScroll(lbl),
		maxHeight: maxHeight,
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *smartScrollContainer) SetText(text string) {
	s.label.SetText(text)
	s.Refresh()
}

func (s *smartScrollContainer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(s.scroll)
}

func (s *smartScrollContainer) MinSize() fyne.Size {
	if s.label.Text == "" {
		return fyne.NewSize(0, 0) // collapse when empty
	}
	natural := s.label.MinSize().Height
	if natural <= s.maxHeight {
		return fyne.NewSize(s.label.MinSize().Width, natural) // no overflow → no scroll chrome visible
	}
	return fyne.NewSize(0, s.maxHeight) // overflow → capped + scroll shown
}

func ChooseNetDev() {
	var netList []string
	for _, x := range network.NetDevs {
		netList = append(netList, x.Description)
	}

	interfaceNameTextLabel := widget.NewLabel("Interface Name:")
	interfaceAddressTextLabel := widget.NewLabel("Interface Addresses:")

	// Declare scrolls first so the closure can capture them
	nameScroll := newSmartScroll(40)
	addrScroll := newSmartScroll(100)

	netDevChangeFunc := func(s string) {
		for _, x := range network.NetDevs {
			if x.Description == s {
				network.NetDev = x
				nameScroll.SetText(x.Name)
				addressText := ""
				for _, y := range x.Addresses {
					addressText += y.IP.String() + "\n"
				}
				addrScroll.SetText(addressText)
			}
		}
	}

	NetDevForm := widget.NewForm(
		widget.NewFormItem("Network Interface", widget.NewSelect(netList, netDevChangeFunc)),
	)

	separatorRect := canvas.NewRectangle(theme.DisabledColor())
	separatorContainer := container.New(&verticalSepLayout{}, separatorRect)

	leftLabels := container.NewVBox(interfaceNameTextLabel, interfaceAddressTextLabel)
	leftSide := container.NewHBox(leftLabels, separatorContainer)

	rightValues := container.NewVBox(nameScroll, addrScroll)
	infoArea := container.NewBorder(nil, nil, leftSide, nil, rightValues)

	topContent := container.NewVBox(NetDevForm)
	mainContainer := container.NewBorder(topContent, nil, nil, nil, infoArea)

	modal := widget.NewModalPopUp(container.NewPadded(mainContainer), w.Canvas())
	modal.Resize(fyne.NewSize(400, 400))
	modal.Show()
}
