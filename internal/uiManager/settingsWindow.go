package uiManager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nullshii/vrcHeart/internal/fyneExtensions"
	"github.com/nullshii/vrcHeart/internal/settingsManager"
)

var OnSettingsSave func()
var OnSettingsLoad func()

var fSettingsWindow fyne.Window

func setupSettingsWindow(settings *settingsManager.Settings) {
	fSettingsWindow = fApp.NewWindow("VRC Heart Settings")

	fSettingsWindow.Resize(fyne.NewSize(500, fSettingsWindow.Canvas().Size().Height))

	fSettingsWindow.SetCloseIntercept(func() {
		fSettingsWindow.Hide()
	})

	fSettingsWindow.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			createStringSettingsEntry("Address", binding.BindString(&settings.Address)),
			createIntSettingsEntry("Port", binding.BindInt(&settings.Port)),
			createIntSettingsEntry("Send frequency", binding.BindInt(&settings.SendFrequency)),
			createIntSettingsEntry("Minimum rate", binding.BindInt(&settings.MinRate)),
			createIntSettingsEntry("Maximum rate", binding.BindInt(&settings.MaxRate)),
			createIntSettingsEntry("Start rate", binding.BindInt(&settings.StartRate)),
			createStringSettingsEntry("Left emoji", binding.BindString(&settings.LeftEmoji)),
			createStringSettingsEntry("Right emoji", binding.BindString(&settings.RightEmoji)),
			createStringSettingsEntry("Text above", binding.BindString(&settings.TextAbove)),
			createStringSettingsEntry("Text below", binding.BindString(&settings.TextBelow)),
			createNumberTypeEntry(binding.BindInt(&settings.NumberType)),
			widget.NewCheckWithData("Save last rate", binding.BindBool(&settings.SaveLastRate)),
			widget.NewCheckWithData("Save on quit", binding.BindBool(&settings.SaveOnQuit)),
			widget.NewButton("Save", OnSettingsSave),
		),
	)
}

func createStringSettingsEntry(name string, data binding.String) *fyne.Container {
	return container.New(
		layout.NewFormLayout(),
		widget.NewLabel(name+":"),
		widget.NewEntryWithData(data),
	)
}

func createIntSettingsEntry(name string, data binding.Int) *fyne.Container {
	return container.New(
		layout.NewFormLayout(),
		widget.NewLabel(name+":"),
		fyneExtensions.NewNumericalEntryWithData(data),
	)
}

func createNumberTypeEntry(data binding.Int) *fyne.Container {
	s := widget.NewSelect([]string{"Random", "Increment", "Decrement"}, func(s string) {
		switch s {
		case "Random":
			data.Set(0)
		case "Increment":
			data.Set(1)
		case "Decrement":
			data.Set(2)
		}
	})

	i, err := data.Get()
	if err != nil {
		panic(err)
	}

	switch i {
	case 0:
		s.SetSelected("Random")
	case 1:
		s.SetSelected("Increment")
	case 2:
		s.SetSelected("Decrement")
	}

	return container.New(
		layout.NewFormLayout(),
		widget.NewLabel("Number type:"),
		s,
	)
}
