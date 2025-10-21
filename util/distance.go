package util

import (
	"fmt"
	"math"
)

func FormatDistance(meter float64) string {
	if meter < 1000 {
		return fmt.Sprintf("%.0fm", meter)
	}
	km := meter / 1000
	roundedKm := math.Round(km*10) / 10
	return fmt.Sprintf("%.1fkm ", roundedKm)
}
