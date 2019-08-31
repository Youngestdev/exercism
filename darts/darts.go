package darts

import "math"

func Score(x, y float64) int {
	var coordinatesoffset float64 = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	switch true {
	case coordinatesoffset > 10:
		return 0
	case coordinatesoffset > 5:
		return 1
	case coordinatesoffset > 1:
		return 5
	default:
		return 10
	}
}