package mathExtensions

import "math"

func Clamp(value int, min int, max int) int {
	return int(math.Min(math.Max(float64(value), float64(min)), float64(max)))
}
