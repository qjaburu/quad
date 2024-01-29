package piscine

import (
	"github.com/01-edu/z01"
)

func QuadE(x, y int) {
	// check for zero or negative values of x and y
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				// 90deg => Top left conner 
				z01.PrintRune('A')
			} else if i == y-1 && j == 0 {
				// 270deg => Bottom left conner
				z01.PrintRune('C')
			} else if i == 0 && j == x-1 {
				// 90deg => Top right conner
				z01.PrintRune('C')
			} else if i == y-1 && j == x-1 {
				// 180deg => Bottom right
				z01.PrintRune('A')
			} else if i == 0 || i == y-1 {
				// top or bottom line
				// the horizontal characters
				z01.PrintRune('B')
			} else if j == 0 || j == x-1 {
				// left and right margins
				// the vertical middle characters
				z01.PrintRune('B')
			} else {
				// middle spaces
				z01.PrintRune(' ')
			}
		}
		// line per line spacing
		z01.PrintRune('\n')
	}
}
