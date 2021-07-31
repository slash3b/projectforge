package util

import (
	"fmt"
)

func MicrosToMillis(i int) string {
	div := 1000
	min := 20

	ms := i / div
	if ms >= min {
		return fmt.Sprintf("%dms", ms)
	}

	x := float64(ms) + (float64(i%div) / float64(div))
	return fmt.Sprintf("%.3f", x) + "ms"
}
