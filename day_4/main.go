package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Number struct {
	value  int
	marked bool
}

type Board struct {
	value [5][5]Number
	won   bool
}

func convert_arr_string_to_int(input []string) []int {
	var ints []int
	for _, el := range input {
		if el == "" {
			continue
		}

		num, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}

		ints = append(ints, num)
	}

	return ints
}

func construct_board(lines []string) Board {
	var board [5][5]Number

	for i, line := range lines {
		str_arr := strings.Split(strings.ReplaceAll(line, "\r", ""), " ")
		numbers := convert_arr_string_to_int(str_arr)
		for j, num := range numbers {
			number := Number{value: num, marked: false}
			board[i][j] = number
		}
	}

	return Board{board, false}
}

func read_input(filename string) ([]int, []Board) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	num_drawn := strings.Split(strings.ReplaceAll(lines[0], "\r", ""), ",")

	numbers_drawn := convert_arr_string_to_int(num_drawn)

	var boards []Board

	for i := 2; i <= len(lines); i += 6 {
		board := construct_board(lines[i : i+5])
		boards = append(boards, board)
	}

	return numbers_drawn, boards
}

func (board *Board) check_win() bool {
	check_arr := func(line [5]Number) bool {
		for _, el := range line {
			if !el.marked {
				return false
			}
		}
		return true
	}

	get_column := func(b Board, col_idx int) [5]Number {
		var col [5]Number
		for i := 0; i < 5; i++ {
			col[i] = board.value[i][col_idx]
		}
		return col
	}

	for i, _ := range board.value {
		if check_arr(board.value[i]) || check_arr(get_column(*board, i)) {
			return true
		}
	}

	return false
}

func (board *Board) mark_number(number int) {
	for i, _ := range board.value {
		for j, _ := range board.value[i] {
			if board.value[i][j].value == number {
				(*board).value[i][j].marked = true
			}
		}
	}
}

func (board *Board) sum_unmarked() int {
	sum := 0
	for i, _ := range board.value {
		for j, _ := range board.value[i] {
			if !board.value[i][j].marked {
				sum += board.value[i][j].value
			}
		}
	}

	return sum
}

func sum(values []int) int {
	s := 0
	for _, el := range values {
		s += el
	}
	return s
}

func part_1(numbers_drawn []int, boards []Board) int {
	for _, number := range numbers_drawn {
		for i := range boards {
			boards[i].mark_number(number)

			if boards[i].check_win() {
				return number * boards[i].sum_unmarked()
			}
		}
	}

	return -1
}

func part_2(numbers_drawn []int, boards []Board) int {
	var winners [100]int
	for _, number := range numbers_drawn {
		for i := range boards {
			boards[i].mark_number(number)

			if boards[i].check_win() && (sum(winners[:])+1) == len(boards) && !boards[i].won {
				return number * boards[i].sum_unmarked()
			}

			if boards[i].check_win() {
				winners[i] = 1
				boards[i].won = true
			}
		}
	}

	return -1
}

func main() {
	numbers_drawn, boards := read_input("input.txt")
	solution_1 := part_1(numbers_drawn, boards)
	solution_2 := part_2(numbers_drawn, boards)

	fmt.Println(solution_1)
	fmt.Println(solution_2)

}
