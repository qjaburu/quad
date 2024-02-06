package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	// exit if either or both x and y are 0
	if x <= 0 || y <= 0 {
		return
	} else {
		// top to bottom outer for loop ranges from i to y-1
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				// avoid getting into the if statement whenever x less than 0 or y is less than 0

				// x and y are both greater than 0
				// test for the top(i == 0) and bottom(i == y-1) line condition
				if i == 0 || i == (y-1) {
					// either we are in the top or bottom line
					//  test for conners
					if i == 0 && j == 0 {
						// top-left conner
						z01.PrintRune('/')
						// check for the special input combination (1, 1)
						// append a new line after
						if x == 1 && y == 1 {
							z01.PrintRune('\n')
						} else if x == 1 {
							// if x is 1 but y is any other number
							// skip to the next line
							z01.PrintRune('\n')
						}
					} else if i == 0 && j == (x-1) {
						// top-right
						z01.PrintRune('\\')
						z01.PrintRune('\n')
					} else if i == (y-1) && j == (x-1) {
						// bottom-right
						z01.PrintRune('/')
						z01.PrintRune('\n')
					} else if i == (y-1) && j == 0 {
						// bottom-left conner
						z01.PrintRune('\\')
					} else {
						// the horrizontal characters in the top or bottom lines
						// j > 0 and  j < x
						z01.PrintRune('*')
					}
				} else {
					// we are in the lines where i > 0 and i < (y-1)
					if j == 0 {
						z01.PrintRune('*')
						if x == 1 {
							z01.PrintRune('\n')
						}
					} else if j == (x - 1) {
						z01.PrintRune('|')
						z01.PrintRune('\n')
					} else {
						// we are in positions that need spaces
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
