package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type Vertex struct {
	name        string
	visit_count int
	neighbours  []*Vertex
}

func (v *Vertex) IsSmall() bool {
	for _, r := range (*v).name {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (v *Vertex) IsStart() bool {
	return (*v).name == "start"
}

func (v *Vertex) IsEnd() bool {
	return (*v).name == "end"
}

func get_start_vert(vertices []*Vertex) *Vertex {
	for i := range vertices {
		if vertices[i].IsStart() {
			return vertices[i]
		}
	}

	return nil
}

type Stack struct {
	s []*Vertex
}

func (stack *Stack) Push(vert *Vertex) {
	(*stack).s = append((*stack).s, vert)
}

func (stack *Stack) Pop() *Vertex {
	item := (*stack).s[len((*stack).s)-1]
	(*stack).s = (*stack).s[:len((*stack).s)-1]
	return item
}

func (stack *Stack) Size() int {
	return len((*stack).s)
}

func read_input(filename string) []*Vertex {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	var vertices []*Vertex
	vert_map := make(map[string]*Vertex)

	for _, line := range lines {
		line_split := strings.Split(strings.ReplaceAll(line, "\r", ""), "-")
		vert1 := line_split[0]
		vert2 := line_split[1]

		_, vert1_in := vert_map[vert1]

		if !vert1_in {
			v := Vertex{name: vert1}
			vert_map[vert1] = &v
			vertices = append(vertices, &v)
		}

		_, vert2_in := vert_map[vert2]

		if !vert2_in {
			v := Vertex{name: vert2}
			vert_map[vert2] = &v
			vertices = append(vertices, &v)
		}

		v1, _ := vert_map[vert1]
		v2, _ := vert_map[vert2]

		v1.neighbours = append(v1.neighbours, v2)
		v2.neighbours = append(v2.neighbours, v1)

	}

	return vertices
}

func find_paths_1(vert *Vertex, path_cnt *int) {
	if vert.IsEnd() {
		(*path_cnt)++
		return
	}

	(*vert).visit_count++

	for _, neigh := range (*vert).neighbours {
		if !neigh.IsStart() && ((neigh.IsSmall() && (*neigh).visit_count == 0) || !neigh.IsSmall()) {
			find_paths_1(neigh, path_cnt)
		}
	}

	(*vert).visit_count--

}

func find_paths_2(vert *Vertex, path_cnt *int) {
	if vert.IsEnd() {
		(*path_cnt)++
		return
	}

	(*vert).visit_count++

	for _, neigh := range (*vert).neighbours {
		if neigh.IsStart() {
			continue
		}

		if neigh.IsSmall() && (*neigh).visit_count == 1 {
			find_paths_1(neigh, path_cnt)
		}

		if !neigh.IsSmall() || (neigh.IsSmall() && (*neigh).visit_count == 0) {
			find_paths_2(neigh, path_cnt)
		}

	}

	(*vert).visit_count--
}

func part_1(vertices []*Vertex) int {
	path_cnt := 0
	find_paths_1(get_start_vert(vertices), &path_cnt)
	return path_cnt
}

func part_2(vertices []*Vertex) int {
	path_cnt := 0
	find_paths_2(get_start_vert(vertices), &path_cnt)
	return path_cnt
}

func main() {
	input := read_input("input.txt")
	solution_1 := part_1(input)
	solution_2 := part_2(input)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
