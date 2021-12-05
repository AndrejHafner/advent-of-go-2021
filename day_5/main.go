package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	point_1 Point
	point_2 Point
}

const FLOOR_SIZE int = 1000

func read_input(filename string) []Line {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	var fmt_lines []Line

	for _, line := range lines {
		line_split := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, "\r", ""), " ", ""), "->")
		pt1 := strings.Split(line_split[0], ",")
		pt2 := strings.Split(line_split[1], ",")

		x_1, _ := strconv.Atoi(pt1[0])
		y_1, _ := strconv.Atoi(pt1[1])
		x_2, _ := strconv.Atoi(pt2[0])
		y_2, _ := strconv.Atoi(pt2[1])

		line_s := Line{point_1: Point{x_1, y_1}, point_2: Point{x_2, y_2}}

		fmt_lines = append(fmt_lines, line_s)
	}

	return fmt_lines
}

func make_range(min, max int, reverse bool) []int {
	a := make([]int, max-min+1)
	if !reverse {
		for i := range a {
			a[i] = min + i
		}
	} else {
		for i := range a {
			a[i] = max - i
		}
	}

	return a
}

func cover_floor(floor *[FLOOR_SIZE][FLOOR_SIZE]int, line Line, horizontal bool) {
	if line.point_1.x == line.point_2.x {
		if line.point_1.y > line.point_2.y {
			for y := line.point_2.y; y <= line.point_1.y; y++ {
				(*floor)[line.point_1.x][y]++
			}
		} else {
			for y := line.point_1.y; y <= line.point_2.y; y++ {
				(*floor)[line.point_1.x][y]++
			}
		}
	} else if line.point_1.y == line.point_2.y {
		if line.point_1.x > line.point_2.x {
			for x := line.point_2.x; x <= line.point_1.x; x++ {
				(*floor)[x][line.point_1.y]++
			}
		} else {
			for x := line.point_1.x; x <= line.point_2.x; x++ {
				(*floor)[x][line.point_1.y]++
			}
		}
	} else if horizontal {
		var x_range, y_range []int
		if line.point_1.x > line.point_2.x {
			x_range = make_range(line.point_2.x, line.point_1.x, true)
		} else {
			x_range = make_range(line.point_1.x, line.point_2.x, false)
		}

		if line.point_1.y > line.point_2.y {
			y_range = make_range(line.point_2.y, line.point_1.y, true)
		} else {
			y_range = make_range(line.point_1.y, line.point_2.y, false)
		}

		for i := range x_range {
			(*floor)[x_range[i]][y_range[i]]++
		}
	}
}

func count_overlaps(floor *[FLOOR_SIZE][FLOOR_SIZE]int, threshold int) int {
	count := 0
	for _, row := range *floor {
		for _, entry := range row {
			if entry >= threshold {
				count++
			}
		}
	}

	return count
}

func part_1(lines []Line) int {
	var floor [FLOOR_SIZE][FLOOR_SIZE]int

	for _, line := range lines {
		cover_floor(&floor, line, false)
	}
	return count_overlaps(&floor, 2)
}

func part_2(lines []Line) int {
	var floor [FLOOR_SIZE][FLOOR_SIZE]int

	for _, line := range lines {
		cover_floor(&floor, line, true)
	}

	return count_overlaps(&floor, 2)
}

func main() {
	input := read_input("input.txt")

	solution_1 := part_1(input)
	solution_2 := part_2(input)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
