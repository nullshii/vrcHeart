package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/nullshii/vrcHeart/internal/mathExtensions"
	"github.com/nullshii/vrcHeart/internal/randomExtensions"
	"github.com/nullshii/vrcHeart/internal/settingsManager"
	"github.com/nullshii/vrcHeart/internal/splash"
)

func main() {
	splash.PrintSplash()
	settingsManager.InitSettings()
	settingsManager.FixFormatting()

	client := osc.NewClient(settingsManager.SettingsInstance.Address, settingsManager.SettingsInstance.Port)

	for {
		sendHeartBeat(client, &settingsManager.SettingsInstance)
		time.Sleep(time.Duration(settingsManager.SettingsInstance.SendFrequency) * time.Millisecond)
	}
}

func sendHeartBeat(client *osc.Client, s *settingsManager.Settings) {
	bpm := s.StartRate

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

	bpm = mathExtensions.Clamp(bpm, s.MinRate, s.MaxRate)

	text := fmt.Sprintf("%s%s%v%s%s", s.TextAbove, s.LeftEmoji, bpm, s.RightEmoji, s.TextBelow)

	msg := osc.NewMessage("/chatbox/input")
	msg.Append(text)
	msg.Append(true)
	msg.Append(false)
	client.Send(msg)

	fmt.Printf("[%v] Sending: %s\n", time.Now().Format("15:04:05"), strings.ReplaceAll(text, "\n", "   "))
}
