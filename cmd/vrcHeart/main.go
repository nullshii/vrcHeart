package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/nullshii/vrcHeart/internal/mathExtensions"
	"github.com/nullshii/vrcHeart/internal/randomExtensions"
	"github.com/nullshii/vrcHeart/internal/settingsManager"
	"github.com/nullshii/vrcHeart/internal/uiManager"
)

var bpm int
var isRunning = false
var settings = settingsManager.Settings{}

func main() {
	settingsManager.InitSettings(&settings)

	uiManager.OnApplicationQuit = onApplicationQuit
	uiManager.OnSettingsSave = onSettingsSave

	uiManager.OnStartSending = onStartSending
	uiManager.OnStopSending = onStopSending

	fApp := uiManager.InitUI(&settings)

	client := osc.NewClient(settings.Address, settings.Port)
	bpm = settings.StartRate

	go func() {
		for range time.Tick(time.Duration(settings.SendFrequency) * time.Millisecond) {
			if isRunning {
				sendHeartBeat(client)
			}
		}
	}()

	(*fApp).Run()
}

func onStartSending() {
	isRunning = true
}

func onStopSending() {
	isRunning = false
}

func onApplicationQuit() {
	isRunning = false

	if settings.SaveOnQuit {
		if settings.SaveLastRate {
			settings.StartRate = bpm
		}

		onSettingsSave()
	}

	fmt.Println("Shitting down")
}

func onSettingsSave() {
	if settings.SaveLastRate {
		settings.StartRate = bpm
	}

	settingsManager.SaveSettins(settings)
	fmt.Println("Saving settings")
}

func sendHeartBeat(client *osc.Client) {
	if settings.NumberType == settingsManager.NUMBER_TYPE_RAMDOM {
		randAdd := randomExtensions.RandRange(1, 3)
		randNum := randomExtensions.RandRange(1, 5)

		if randNum == 2 { //Add big chunk
			randAdd += randomExtensions.RandRange(6, 18)
		}

		if randNum == 3 { // add small amount
			bpm += randAdd
		} else {
			bpm -= randAdd
		}
	} else if settings.NumberType == settingsManager.NUMBER_TYPE_INCREMENT {
		bpm++
	} else if settings.NumberType == settingsManager.NUMBER_TYPE_DECREMEMT {
		bpm--
	}

	bpm = mathExtensions.Clamp(bpm, settings.MinRate, settings.MaxRate)

	ta := settings.TextAbove
	if ta != "" && !strings.HasSuffix(ta, "\n") {
		ta += "\n"
	}

	tb := settings.TextBelow
	if tb != "" && !strings.HasSuffix(tb, "\n") {
		tb = "\n" + tb
	}

	text := fmt.Sprintf("%s%s%v%s%s", ta, settings.LeftEmoji, bpm, settings.RightEmoji, tb)

	msg := osc.NewMessage("/chatbox/input")
	msg.Append(text)
	msg.Append(true)
	msg.Append(false)
	client.Send(msg)

	fmt.Printf("[%v] Sending: %s\n", time.Now().Format("15:04:05"), strings.ReplaceAll(text, "\n", "   "))
}
