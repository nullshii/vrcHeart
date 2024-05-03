package settingsManager

import (
	"encoding/json"
	"errors"
	"os"
)

var settingsFilePath = "settings.json"

const (
	NUMBER_TYPE_RAMDOM    = iota
	NUMBER_TYPE_INCREMENT = iota
	NUMBER_TYPE_DECREMEMT = iota
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
	NumberType    int
	SaveLastRate  bool
	SaveOnQuit    bool
}

func InitSettings(settings *Settings) {
	defaultSettings := Settings{
		Address:       "127.0.0.1",
		Port:          9000,
		SendFrequency: 1500,
		MinRate:       70,
		MaxRate:       190,
		StartRate:     120,
		LeftEmoji:     "♥ ",
		RightEmoji:    " ♥",
		TextAbove:     "",
		TextBelow:     "",
		NumberType:    NUMBER_TYPE_RAMDOM,
		SaveLastRate:  false,
		SaveOnQuit:    true,
	}

	if _, err := os.Stat(settingsFilePath); errors.Is(err, os.ErrNotExist) {
		SaveSettins(defaultSettings)

	} else {
		LoadSettings(settings)
	}
}

func LoadSettings(settings *Settings) {
	data, err := os.ReadFile(settingsFilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, settings)
	if err != nil {
		panic(err)
	}
}

func SaveSettins(settings Settings) {
	data, err := json.MarshalIndent(settings, "", "\t")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(settingsFilePath, data, 0644)
	if err != nil {
		panic(err)
	}
}
