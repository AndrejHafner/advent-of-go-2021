package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func read_input(filename string) []int {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	str_nums := strings.Split(string(bs), ",")

	var positions []int

	for _, num := range str_nums {
		num_parse, _ := strconv.Atoi(num)
		positions = append(positions, num_parse)
	}

	return positions
}

func sum(in []int) int {
	sum := 0
	for _, el := range in {
		sum += el
	}
	return sum
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func abs_arr(in *[]int) {
	for i := range *in {
		(*in)[i] = abs((*in)[i])
	}
}

func min(in []int) int {
	min_value := 999999999999

	for _, val := range in {
		if val < min_value {
			min_value = val
		}
	}
	return min_value
}

func max(in []int) int {
	max_value := -999999999999

	for _, val := range in {
		if val > max_value {
			max_value = val
		}
	}
	return max_value
}

func increasing_cost(in int) int {
	out := 0
	for i := 1; i <= in; i++ {
		out += i
	}
	return out
}

func add_increasing_cost(in *[]int) {
	for i := range *in {
		(*in)[i] = increasing_cost((*in)[i])
	}
}

func part_1(positions []int) int {
	var sums []int
	max_pos := max(positions)
	for pos := 0; pos <= max_pos; pos++ {
		positions_cp := make([]int, len(positions))
		copy(positions_cp, positions)

		for i := 0; i < len(positions_cp); i++ {
			positions_cp[i] -= pos
		}
		abs_arr(&positions_cp)
		sum := sum(positions_cp)
		sums = append(sums, sum)
	}

	return min(sums)
}

func part_2(positions []int) int {
	var sums []int
	max_pos := max(positions)
	for pos := 0; pos <= max_pos; pos++ {
		positions_cp := make([]int, len(positions))
		copy(positions_cp, positions)

		for i := 0; i < len(positions_cp); i++ {
			positions_cp[i] -= pos
		}
		abs_arr(&positions_cp)
		add_increasing_cost(&positions_cp)
		sum := sum(positions_cp)
		sums = append(sums, sum)
	}

	return min(sums)
}

func main() {
	input := read_input("input.txt")
	solution_1 := part_1(input)
	solution_2 := part_2(input)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
