package main

import (
	"fmt"
	"strconv"

	"github.com/MathewKostiuk/adventhelpers"
)

func main() {
	input := adventhelpers.GetInput()
	bl := len(input[0])

	calculateGammaAndEpsilon(input, bl)
	o, err := strconv.ParseUint(calculateOxygen(input, 0), 2, bl)
	if err != nil {
		fmt.Printf("error parsing uint: %s", err)
	}
	fmt.Println(o)

	c, err := strconv.ParseUint(calculateCarbon(input, 0), 2, bl)
	if err != nil {
		fmt.Printf("error parsing uint: %s", err)
	}
	fmt.Println(c)

	fmt.Printf("Part 2 result is: %d\n", o*c)
}

func calculateGammaAndEpsilon(input []string, total int) {
	var gs, es string

	for i := 0; i < total; i++ {
		var mc = map[string]int{
			"0": 0,
			"1": 0,
		}

		for _, b := range input {
			mc[string(b[i])]++
		}
		if mc["0"] < mc["1"] {
			gs += "1"
			es += "0"
		} else {
			gs += "0"
			es += "1"
		}
	}

	gamma, err := strconv.ParseUint(gs, 2, total)
	if err != nil {
		fmt.Printf("error during parseuint: %s", err)
	}

	epsilon, err := strconv.ParseUint(es, 2, total)
	if err != nil {
		fmt.Printf("error during parseuint: %s", err)
	}

	fmt.Printf("Part 1 result is: %d\n", gamma*epsilon)
}

func calculateOxygen(input []string, index int) string {
	if len(input) == 1 {
		return input[0]
	}

	mc := make(map[string][]string)

	for _, b := range input {
		mc[string(b[index])] = append(mc[string(b[index])], b)
	}

	if len(mc["0"]) <= len(mc["1"]) {
		return calculateOxygen(mc["1"], index+1)
	} else {
		return calculateOxygen(mc["0"], index+1)
	}
}

func calculateCarbon(input []string, index int) string {
	if len(input) == 1 {
		return input[0]
	}

	mc := make(map[string][]string)

	for _, b := range input {
		mc[string(b[index])] = append(mc[string(b[index])], b)
	}

	if len(mc["0"]) <= len(mc["1"]) {
		return calculateCarbon(mc["0"], index+1)
	} else {
		return calculateCarbon(mc["1"], index+1)
	}
}
