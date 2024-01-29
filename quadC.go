package piscine

import (
	"github.com/01-edu/z01"
)

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				z01.PrintRune('A') // TOP LEFT
			} else if i == y-1 && j == 0 {
				z01.PrintRune('C') // BOTTOM LEFT
			} else if i == 0 && j == x-1 {
				z01.PrintRune('A') // TOP RIGHT
			} else if i == y-1 && j == x-1 {
				z01.PrintRune('C') // BOTTOM RIGHT
			} else if i == 0 || i == y-1 {
				z01.PrintRune('B') // PRINTS HORIZONTAL LINE
			} else if j == 0 || j == x-1 {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
