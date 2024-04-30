package randomExtensions

import (
	"math/rand"
)

func RandRange(min int, max int) int {
	if max <= min || max == 0 {
		return 0
	}

	return rand.Intn(max-min) + min
}
