package settingsManager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Settings struct {
	Address       string
	Port          int
	SendFrequency int
	MinRate       int
	MaxRate       int
	StartRate     int
	LeftEmoji     string
	RightEmoji    string
	TextAbove     string
	TextBelow     string
}

var SettingsInstance Settings = Settings{
	Address:       "127.0.0.1",
	Port:          9000,
	SendFrequency: 1500,
	MinRate:       70,
	MaxRate:       190,
	StartRate:     120,
	LeftEmoji:     "♥",
	RightEmoji:    "♥",
	TextAbove:     "",
	TextBelow:     "",
}

func InitSettings() {
	settingsFilePath := "settings.json"

	if _, err := os.Stat(settingsFilePath); errors.Is(err, os.ErrNotExist) {

		data, err := json.MarshalIndent(SettingsInstance, "", "\t")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Can't find settings file... Creating new one.\n\n")

		err = os.WriteFile(settingsFilePath, data, 0644)
		if err != nil {
			panic(err)
		}

	} else {
		data, err := os.ReadFile(settingsFilePath)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(data, &SettingsInstance)
		if err != nil {
			panic(err)
		}
	}
}

func FixFormatting() {
	if SettingsInstance.TextAbove != "" {
		SettingsInstance.TextAbove += "\n"
	}

	if SettingsInstance.TextBelow != "" {
		SettingsInstance.TextBelow += "\n"
	}

	if SettingsInstance.LeftEmoji != "" {
		SettingsInstance.LeftEmoji = SettingsInstance.LeftEmoji + " "
	}

	if SettingsInstance.RightEmoji != "" {
		SettingsInstance.RightEmoji = " " + SettingsInstance.RightEmoji
	}
}
