package uiManager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var fMainWindow fyne.Window

var OnStartSending func()
var OnStopSending func()

func setupMainWindow() {
	fMainWindow = fApp.NewWindow("VRC Heart")

	fMainWindow.Resize(fyne.NewSize(300, 200))

	fMainWindow.SetCloseIntercept(func() {
		OnApplicationQuit()
		fApp.Quit()
	})

	fMainWindow.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			widget.NewButton("Start", OnStartSending),
			widget.NewButton("Stop", OnStopSending),
			layout.NewSpacer(),
			widget.NewButton("Settings", func() {
				fSettingsWindow.Show()
			}),
			layout.NewSpacer(),
			widget.NewButton("Hide to system tray", func() {
				fMainWindow.Hide()
				fSettingsWindow.Hide()
			}),
		),
	)

	fMainWindow.SetMaster()
}
