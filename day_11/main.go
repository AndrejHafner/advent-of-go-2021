package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Octopus struct {
	energy  int
	flashed bool
}

type Grid struct {
	value [10][10]Octopus
}

func (grid *Grid) Step() {
	for i, _ := range (*grid).value {
		for j, _ := range (*grid).value[i] {
			(*grid).value[i][j].energy++
		}
	}
}

func (grid *Grid) FlashAll() {

	var octopus_flashes [][2]int

	for i, _ := range (*grid).value {
		for j, _ := range (*grid).value[i] {
			if (*grid).value[i][j].energy > 9 {
				octopus_flashes = append(octopus_flashes, [2]int{i, j})
			}
		}
	}

	for _, coords := range octopus_flashes {
		if !(*grid).value[coords[0]][coords[1]].flashed {
			octo_flash(grid, coords[0], coords[1])
		}
	}
}

func (grid *Grid) DidAllFlash() bool {
	count := 0

	for i, _ := range (*grid).value {
		for j, _ := range (*grid).value[i] {
			if (*grid).value[i][j].flashed {
				count++
			}
		}
	}

	return count == 100
}

func (grid *Grid) CountAndReset() int {
	count := 0

	for i, _ := range (*grid).value {
		for j, _ := range (*grid).value[i] {
			if (*grid).value[i][j].flashed {
				count++
				(*grid).value[i][j].energy = 0
				(*grid).value[i][j].flashed = false
			}
		}
	}

	return count
}

func octo_flash(grid *Grid, i, j int) {
	neigh_i := [8]int{i - 1, i - 1, i - 1, i, i, i + 1, i + 1, i + 1}
	neigh_j := [8]int{j - 1, j, j + 1, j - 1, j + 1, j - 1, j, j + 1}

	(*grid).value[i][j].flashed = true

	for k, _ := range neigh_i {
		if neigh_i[k] >= 0 && neigh_i[k] < 10 && neigh_j[k] >= 0 && neigh_j[k] < 10 {
			(*grid).value[neigh_i[k]][neigh_j[k]].energy++

		}
	}

	for k, _ := range neigh_i {
		if (neigh_i[k] >= 0 && neigh_i[k] < 10 && neigh_j[k] >= 0 && neigh_j[k] < 10) &&
			(*grid).value[neigh_i[k]][neigh_j[k]].energy > 9 &&
			!(*grid).value[neigh_i[k]][neigh_j[k]].flashed {
			octo_flash(grid, neigh_i[k], neigh_j[k])
		}
	}
}

func read_input(filename string) *Grid {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	var octopuses [10][10]Octopus

	for i, line := range lines {
		line_split := strings.Split(strings.ReplaceAll(line, "\r", ""), "")
		for j, octo := range line_split {
			num, _ := strconv.Atoi(octo)
			octopuses[i][j] = Octopus{energy: num}
		}
	}

	return &Grid{octopuses}
}

func part_1(grid *Grid) int {
	flashes := 0

	for i := 0; i < 100; i++ {
		grid.Step()
		grid.FlashAll()
		flashes += grid.CountAndReset()
	}

	return flashes
}

func part_2(grid *Grid) int {
	for i := 0; i < 1000; i++ {
		grid.Step()
		grid.FlashAll()
		if grid.DidAllFlash() {
			return i + 1
		}
		grid.CountAndReset()
	}

	return 0
}

func main() {
	input := read_input("input.txt")
	input_2 := Grid{}
	copy(input_2.value[:], input.value[:])

	solution_1 := part_1(input)
	solution_2 := part_2(&input_2)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
