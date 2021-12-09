package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const MAP_SIZE_X = 102
const MAP_SIZE_Y = 102

var visited [MAP_SIZE_X][MAP_SIZE_Y]bool

func read_input(filename string) [MAP_SIZE_X][MAP_SIZE_Y]int {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	var heightmap [MAP_SIZE_X][MAP_SIZE_Y]int

	for i := 0; i < MAP_SIZE_X; i++ {
		heightmap[i][0] = 9
		heightmap[i][MAP_SIZE_Y-1] = 9
	}

	for i := 0; i < MAP_SIZE_Y; i++ {
		heightmap[0][i] = 9
		heightmap[MAP_SIZE_X-1][i] = 9
	}

	for i, line := range lines {
		line_split := strings.Split(strings.ReplaceAll(line, "\r", ""), "")
		for j, height := range line_split {
			h, _ := strconv.Atoi(height)
			heightmap[i+1][j+1] = h
		}
	}

	return heightmap
}

func measure_basin(heightmap [MAP_SIZE_X][MAP_SIZE_Y]int, i, j int) int {
	i_neigh := []int{i - 1, i, i + 1, i}
	j_neigh := []int{j, j - 1, j, j + 1}

	size := 0
	for k := range j_neigh {
		if heightmap[i_neigh[k]][j_neigh[k]] < 9 && !visited[i_neigh[k]][j_neigh[k]] {
			visited[i_neigh[k]][j_neigh[k]] = true
			size += measure_basin(heightmap, i_neigh[k], j_neigh[k])
		}
	}

	return size + 1
}

func part_1(heightmap [MAP_SIZE_X][MAP_SIZE_Y]int) int {
	risk_sum := 0

	for i := 1; i < MAP_SIZE_X-1; i++ {
		for j := 1; j < MAP_SIZE_Y-1; j++ {
			height := heightmap[i][j]
			if heightmap[i-1][j] > height && heightmap[i][j-1] > height && heightmap[i+1][j] > height && heightmap[i][j+1] > height {
				risk_sum += height + 1
			}
		}
	}

	return risk_sum
}

func part_2(heightmap [MAP_SIZE_X][MAP_SIZE_Y]int) int {
	var basin_sizes []int

	for i := 1; i < MAP_SIZE_X-1; i++ {
		for j := 1; j < MAP_SIZE_Y-1; j++ {
			height := heightmap[i][j]
			if heightmap[i-1][j] > height && heightmap[i][j-1] > height && heightmap[i+1][j] > height && heightmap[i][j+1] > height {
				// Found a bottom of the basin
				basin_size := measure_basin(heightmap, i, j) - 1
				basin_sizes = append(basin_sizes, basin_size)
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basin_sizes)))

	return basin_sizes[0] * basin_sizes[1] * basin_sizes[2]
}

func main() {
	input := read_input("input.txt")
	solution_1 := part_1(input)
	solution_2 := part_2(input)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
