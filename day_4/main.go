package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MathewKostiuk/adventhelpers"
)

type Boards struct {
	boards []*Board
}

type Board struct {
	winner bool
	lines  []*Line
	score  int
}

type Line struct {
	winner    bool
	positions []*Position
}

type Position struct {
	winner bool
	number int
}

var winners []*Board

func main() {
	input := adventhelpers.GetInput()
	ni := input[0]
	bi := input[1:]
	nums := strings.Split(ni, ",")

	boards := buildBoards(bi)
	for i := 0; i < len(nums); i++ {
		num, err := strconv.Atoi(string(nums[i]))
		if err != nil {
			fmt.Printf("error converting string to int: %s", err)
		}
		checkBoards(boards, num)
	}
	fmt.Println(winners[0].score)
	fmt.Println(winners[len(winners)-1].score)
}

func buildBoards(rows []string) *Boards {
	var boards []*Board

	for i := 1; i < len(rows); i += 6 {
		var r []string = rows[i : i+5]
		var lines []*Line

		// Build rows
		for j := 0; j < 5; j++ {
			row := strings.Split(strings.TrimSpace(strings.ReplaceAll(r[j], "  ", " ")), " ")
			var positions []*Position
			for _, p := range row {
				number, err := strconv.Atoi(p)
				if err != nil {
					fmt.Printf("error converting string to int: %s", err)
				}
				positions = append(positions, &Position{false, number})
			}

			lines = append(lines, &Line{false, positions})
		}

		// Build columns
		for l := 0; l < 5; l++ {
			var positions []*Position
			positions = append(positions, lines[0].positions[l])
			positions = append(positions, lines[1].positions[l])
			positions = append(positions, lines[2].positions[l])
			positions = append(positions, lines[3].positions[l])
			positions = append(positions, lines[4].positions[l])
			lines = append(lines, &Line{false, positions})
		}
		boards = append(boards, &Board{false, lines, 0})
	}

	return &Boards{boards}
}

func checkBoards(boards *Boards, num int) {
	for _, board := range boards.boards {
		if board.winner {
			continue
		}
		checkBoard(board, num)
		if board.winner {
			calculateScore(board, num)
			winners = append(winners, board)
		}
	}
}

func checkBoard(board *Board, num int) {
	for _, line := range board.lines {
		checkLine(line, num)
		if line.winner {
			board.winner = true
		}
	}
}

func checkLine(line *Line, num int) {
	isWinner := true
	for _, position := range line.positions {
		if position.winner {
			continue
		}

		if num == position.number {
			position.winner = true
		} else {
			isWinner = false
		}
	}

	if isWinner {
		line.winner = true
	}
}

func calculateScore(board *Board, num int) {
	var score int
	for _, line := range board.lines {
		for _, position := range line.positions {
			if !position.winner {
				score += position.number
			}
		}
	}
	score = (score / 2) * num
	board.score = score
}
