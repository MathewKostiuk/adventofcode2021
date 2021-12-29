package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var horizontal, depth, aim int
	input := getInput()

	for _, command := range input {
		x := parseCmdInt(command)
		fmt.Println(x)
		if strings.Contains(command, "forward") {
			horizontal += x
			depth += (x * aim)
		}

		if strings.Contains(command, "up") {
			aim -= x
		}

		if strings.Contains(command, "down") {
			aim += x
		}
	}
	fmt.Println(horizontal, depth, aim)
	fmt.Println(horizontal * depth)
}

func getInput() []string {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s\n", err)
	}

	str := strings.Trim(string(content[:]), "\n")
	return strings.Split(str, "\n")
}

func parseCmdInt(command string) int {
	byteNum := command[len(command)-1]
	byteInt, err := strconv.Atoi(string(byteNum))
	if err != nil {
		fmt.Printf("error converting byte to int: %s\n", err)
	}
	return byteInt
}
