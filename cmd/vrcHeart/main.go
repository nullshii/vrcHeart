package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/nullshii/vrcHeart/internal/mathExtensions"
	"github.com/nullshii/vrcHeart/internal/randomExtensions"
	"github.com/nullshii/vrcHeart/internal/settingsManager"
	"github.com/nullshii/vrcHeart/internal/splash"
)

var bpm int
var isShuttingDown = false

func main() {
	splash.PrintSplash()
	settingsManager.InitSettings()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		isShuttingDown = true

		if settingsManager.SettingsInstance.SaveLastRate {
			settingsManager.SettingsInstance.StartRate = bpm
			settingsManager.SaveSettins()
		}

		fmt.Println("\nShutting down...")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	client := osc.NewClient(settingsManager.SettingsInstance.Address, settingsManager.SettingsInstance.Port)
	bpm = settingsManager.SettingsInstance.StartRate

	for {
		if isShuttingDown {
			return
		}
		sendHeartBeat(client, &settingsManager.SettingsInstance)
		time.Sleep(time.Duration(settingsManager.SettingsInstance.SendFrequency) * time.Millisecond)
	}
}

func sendHeartBeat(client *osc.Client, s *settingsManager.Settings) {
	if settingsManager.SettingsInstance.NumberType == settingsManager.NUMBER_TYPE_RAMDOM {
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
	} else if settingsManager.SettingsInstance.NumberType == settingsManager.NUMBER_TYPE_INCREMENT {
		bpm++
	} else if settingsManager.SettingsInstance.NumberType == settingsManager.NUMBER_TYPE_DECREMEMT {
		bpm--
	}

	bpm = mathExtensions.Clamp(bpm, s.MinRate, s.MaxRate)

	text := fmt.Sprintf("%s%s%v%s%s", s.TextAbove, s.LeftEmoji, bpm, s.RightEmoji, s.TextBelow)

	msg := osc.NewMessage("/chatbox/input")
	msg.Append(text)
	msg.Append(true)
	msg.Append(false)
	client.Send(msg)

	fmt.Printf("[%v] Sending: %s\n", time.Now().Format("15:04:05"), strings.ReplaceAll(text, "\n", "   "))
}
