package main

import (
	"fmt"
)

// target area: x=155..182, y=-117..-67
// target area: x=20..30, y=-10..-5

// type to store output
type velocity struct {
	x     int
	y     int
	highy int
}

func main() {
	x1 := 155
	x2 := 182
	y1 := -117
	y2 := -67

	var velocities []velocity

	for row := 0; row <= x2; row++ {
		for col := x1; col >= y1; col-- {
			out, v := testTrajectory(x1, x2, y1, y2, row, col)
			if out {
				velocities = append(velocities, v)
			}
		}
	}
	fmt.Println("Total distinct velocities ", len(velocities))
	for _, v := range velocities {
		fmt.Println(v.x, v.y, v.highy)
	}

}

func testTrajectory(x1, x2, y1, y2, a, b int) (bool, velocity) {
	var x, y int
	var highy int
	v := velocity{a, b, 0}
	// repeat the loop until within target
	for !hasOvershot(x2, y1, x, y) {
		// fmt.Println("current speed ", a, b)
		x = x + a
		y = y + b
		// fmt.Println("Reached point ", x, y)
		if y >= highy {
			highy = y
		}
		if x <= x2 && x >= x1 && y <= y2 && y >= y1 {
			fmt.Println("Reached target area with values ", v.x, v.y)
			fmt.Println("Highest y point reached ", highy)
			v.highy = highy
			return true, v
		}
		if a != 0 {
			a--
		}
		b--
	}
	return false, v
}

func hasOvershot(x2, y2, x, y int) bool {
	if x <= x2 && y >= y2 {
		return false
	}
	// fmt.Println("Probe has overshot", x, y)
	return true
}
