package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var a fyne.App

var w fyne.Window

type verticalSepLayout struct{}

var desktopMode bool

var desk desktop.App
var tray *fyne.Menu
