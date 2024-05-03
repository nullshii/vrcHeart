package uiManager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/nullshii/vrcHeart/internal/settingsManager"
)

var OnApplicationQuit func()
var fApp fyne.App

func InitUI(settings *settingsManager.Settings) *fyne.App {
	fApp = app.New()

	setupMainWindow()
	setupSettingsWindow(settings)
	initTray()

	fMainWindow.Show()
	return &fApp
}

func initTray() {
	if desk, ok := fApp.(desktop.App); ok {
		m := fyne.NewMenu("System tray",
			fyne.NewMenuItem("Show", func() {
				fMainWindow.Show()
			}),
			fyne.NewMenuItem("Hide", func() {
				fMainWindow.Hide()
			}),
		)
		desk.SetSystemTrayMenu(m)
	}
}
