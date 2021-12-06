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

	var lanternfish []int

	for _, num := range str_nums {
		num_parse, _ := strconv.Atoi(num)
		lanternfish = append(lanternfish, num_parse)
	}

	return lanternfish
}

func simulate_day(day_fish_map *[9]int) {
	prev_day := (*day_fish_map)[8]
	new_borns := (*day_fish_map)[0]
	for i := 7; i >= 0; i-- {
		curr_day := (*day_fish_map)[i]
		(*day_fish_map)[i] = prev_day
		prev_day = curr_day
	}
	(*day_fish_map)[8] = new_borns
	(*day_fish_map)[6] += new_borns
}

func simulate(lanternfish []int, days int) int {

	var day_fish_map [9]int

	for _, fish := range lanternfish {
		day_fish_map[fish]++
	}

	for day := 1; day <= days; day++ {
		simulate_day(&day_fish_map)
	}

	sum := 0
	for _, el := range day_fish_map {
		sum += el
	}

	return sum
}

func part_1(lanternfish []int) int {
	return simulate(lanternfish, 80)
}

func part_2(lanternfish []int) int {
	return simulate(lanternfish, 256)
}

func main() {
	input := read_input("input.txt")
	solution_1 := part_1(input)
	solution_2 := part_2(input)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
