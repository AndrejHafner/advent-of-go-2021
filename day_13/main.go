package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Fold struct {
	axis  string
	value int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func make_matrix(size_x, size_y int) [][]bool {
	paper := make([][]bool, size_y)
	for i := range paper {
		paper[i] = make([]bool, size_x)
	}

	return paper
}

func read_input(filename string) ([]Point, []Fold) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	var points []Point
	var folds []Fold

	empty_line_idx := 0

	for i, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		if len(line) == 0 {
			empty_line_idx = i
			break
		}
		line_split := strings.Split(line, ",")

		x, _ := strconv.Atoi(line_split[0])
		y, _ := strconv.Atoi(line_split[1])

		points = append(points, Point{x: x, y: y})
	}

	for line_idx := empty_line_idx + 1; line_idx < len(lines); line_idx++ {
		line := strings.ReplaceAll(lines[line_idx], "\r", "")

		line_split := strings.Split(line, " ")

		entry_split := strings.Split(line_split[2], "=")

		axis := entry_split[0]
		value, _ := strconv.Atoi(entry_split[1])

		folds = append(folds, Fold{axis: axis, value: value})
	}

	return points, folds
}

func max_x_y(points []Point) (int, int) {
	max_x, max_y := math.MinInt32, math.MinInt32
	for _, pt := range points {
		if pt.x > max_x {
			max_x = pt.x
		}
		if pt.y > max_y {
			max_y = pt.y
		}
	}

	return max_x, max_y
}

func count_dots(in [][]bool) int {
	count := 0
	for i := range in {
		for j := range in[i] {
			if in[i][j] {
				count++
			}
		}
	}

	return count
}

func flip_y_axis(part [][]bool) [][]bool {
	size_x := len(part[0])
	size_y := len(part)
	out := make_matrix(size_x, size_y)

	for y := 0; y < size_y; y++ {
		out[size_y-y-1] = part[y]
	}

	return out
}

func flip_x_axis(part [][]bool) [][]bool {
	size_x := len(part[0])
	size_y := len(part)
	out := make_matrix(size_x, size_y)

	for y := 0; y < size_y; y++ {
		for x := 0; x < size_x; x++ {
			out[y][size_x-x-1] = part[y][x]
		}
	}

	return out
}

func combine_parts_vertical(upper, bottom [][]bool) [][]bool {
	var out [][]bool

	bottom = flip_y_axis(bottom)

	if len(upper) >= len(bottom) {
		out = upper
		y_diff := len(upper) - len(bottom)
		for y := range bottom {
			for x := range bottom[y] {
				out[y+y_diff][x] = out[y+y_diff][x] || bottom[y][x]
			}
		}
	} else {
		out = bottom
		y_diff := len(bottom) - len(upper)
		for y := range upper {
			for x := range upper[y] {
				out[y+y_diff][x] = upper[y][x] || out[y+y_diff][x]
			}
		}
		out = flip_y_axis(out)
	}

	return out
}

func combine_parts_horizontal(left, right [][]bool) [][]bool {
	var out [][]bool

	right = flip_x_axis(right)

	if len(left) >= len(right) {
		out = left
		x_diff := len(left[0]) - len(right[0])
		for y := range right {
			for x := range right[y] {
				out[y][x+x_diff] = out[y][x+x_diff] || right[y][x]
			}
		}
	} else {
		out = right
		x_diff := len(right[0]) - len(left[0])
		for y := range left {
			for x := range left[y] {
				out[y][x+x_diff] = left[y][x] || out[y][x+x_diff]
			}
		}
	}

	return out
}

func get_columns(in [][]bool, from, to int) [][]bool {
	out := make_matrix(to-from, len(in))
	for y := range in {
		for x := 0; x < (to - from); x++ {
			out[y][x] = in[y][x+from]
		}
	}

	return out
}

func fold_paper(paper *[][]bool, folds []Fold) {

	for _, fold := range folds {
		if fold.axis == "y" { // Vertical fold --> bottom part up
			upper_part := (*paper)[:fold.value]
			bottom_part := (*paper)[fold.value+1:]
			(*paper) = combine_parts_vertical(upper_part, bottom_part)
		} else { // Horizontal fold --> right to left
			left_part := get_columns(*paper, 0, fold.value)
			right_part := get_columns(*paper, fold.value+1, len((*paper)[0]))
			(*paper) = combine_parts_horizontal(left_part, right_part)
		}
	}
}

func print_paper(paper [][]bool) {
	for y := range paper {
		for x := range paper[y] {
			if paper[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part_1(points []Point, folds []Fold) int {
	max_x, max_y := max_x_y(points)

	paper := make_matrix(max_x+1, max_y+1)

	for _, pt := range points {
		paper[pt.y][pt.x] = true
	}

	fold_paper(&paper, folds[:1])
	return count_dots(paper)
}

func part_2(points []Point, folds []Fold) int {
	max_x, max_y := max_x_y(points)

	paper := make_matrix(max_x+1, max_y+1)

	for _, pt := range points {
		paper[pt.y][pt.x] = true
	}

	fold_paper(&paper, folds)
	print_paper(paper)
	return 0
}

func main() {
	points, folds := read_input("input.txt")
	solution_1 := part_1(points, folds)
	solution_2 := part_2(points, folds)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
