package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MathewKostiuk/adventhelpers"
)

type Point struct {
	x int
	y int
}

func main() {
	input := adventhelpers.GetInput()
	vents := make(map[Point]int)

	for _, line := range input {
		s := strings.Split(strings.Replace(line, " -> ", ",", -1), ",")
		mapPoints(s[0], s[1], s[2], s[3], vents)
	}
	overlaps := countOverlaps(vents)
	fmt.Println(overlaps)
}

func mapPoints(x1, y1, x2, y2 string, vents map[Point]int) {
	x1i, err := strconv.Atoi(x1)
	if err != nil {
		fmt.Printf("error converting string to int: %s\n", err)
	}
	x2i, err := strconv.Atoi(x2)
	if err != nil {
		fmt.Printf("error converting string to int: %s\n", err)
	}
	y1i, err := strconv.Atoi(y1)
	if err != nil {
		fmt.Printf("error converting string to int: %s\n", err)
	}
	y2i, err := strconv.Atoi(y2)
	if err != nil {
		fmt.Printf("error converting string to int: %s\n", err)
	}

	var smallX, largeX, smallY, largeY int
	if x1i <= x2i {
		smallX, largeX = x1i, x2i
	} else {
		smallX, largeX = x2i, x1i
	}

	if y1i <= y2i {
		smallY, largeY = y1i, y2i
	} else {
		smallY, largeY = y2i, y1i
	}

	xShift := largeX - smallX
	yShift := largeY - smallY

	if xShift == 0 {
		for i := smallY; i <= largeY; i++ {
			vents[Point{smallX, i}]++
		}
	}

	if yShift == 0 {
		for i := smallX; i <= largeX; i++ {
			vents[Point{i, smallY}]++
		}
	}

	if xShift == yShift {
		negativeX := x2i-x1i < 0
		negativeY := y2i-y1i < 0
		for i := 0; i <= xShift; i++ {
			var newX, newY int
			if negativeX {
				newX = x1i - i
			} else {
				newX = x1i + i
			}

			if negativeY {
				newY = y1i - i
			} else {
				newY = y1i + i
			}
			vents[Point{newX, newY}]++
		}
	}

}

func countOverlaps(vents map[Point]int) int {
	var overlaps int
	for _, value := range vents {
		if value > 1 {
			overlaps++
		}
	}
	return overlaps
}
