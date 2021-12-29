package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	var prev int
	var inc int
	for i := 0; i < len(input); i++ {
		j := i + 1
		k := i + 2
		if j < len(input) && k < len(input) {
			a, err := strconv.Atoi(input[i])
			if err != nil {
				fmt.Printf("error converting i to int: %s\n", err)
			}

			b, err := strconv.Atoi(input[j])
			if err != nil {
				fmt.Printf("error converting j to int: %s\n", err)
			}

			c, err := strconv.Atoi(input[k])
			if err != nil {
				fmt.Printf("error converting k to int: %s\n", err)
			}

			if i == 0 {
				prev = a + b + c
				continue
			}
			sum := a + b + c

			if prev < sum {
				inc++
			}
			prev = sum
		}

	}
	fmt.Println(inc)
}

func getInput() []string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s\n", err)
	}

	str := strings.Trim(string(content[:]), "\n")
	return strings.Split(str, "\n")
}
