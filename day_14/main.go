package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func read_input(filename string) (string, map[string]string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	template := strings.ReplaceAll(lines[0], "\r", "")

	insertion_map := map[string]string{}

	for i := 2; i < len(lines); i++ {
		line := strings.ReplaceAll(strings.ReplaceAll(lines[i], "\r", ""), " ", "")

		line_split := strings.Split(line, "->")
		insertion_map[line_split[0]] = line_split[1]
	}

	return template, insertion_map
}

func pair_insertion(template *[]string, insertion_map map[string]string) {
	var new_template []string

	new_template = append(new_template, (*template)[0])

	for i := 0; i < len(*template)-1; i++ {
		pair := (*template)[i] + (*template)[i+1]
		insert_value, contains := insertion_map[pair]

		if contains {
			new_template = append(new_template, insert_value)
		}
		new_template = append(new_template, (*template)[i+1])
	}

	*template = new_template
}

func get_max_min_element(template_arr *[]string) (int, int) {
	counter := map[string]int{}

	for _, el := range *template_arr {
		_, contains := counter[el]
		if !contains {
			counter[el] = 1
		} else {
			counter[el]++
		}
	}

	return 0, 0
}

func part_1(template string, insertion_map map[string]string) int {

	template_arr := strings.Split(template, "")
	n_steps := 40

	for step := 0; step < n_steps; step++ {
		fmt.Println(step)
		pair_insertion(&template_arr, insertion_map)
	}

	max, min := get_max_min_element(&template_arr)

	return max - min
}

func part_2(template string, insertion_map map[string]string) int {
	return 0
}

func main() {
	template, rules := read_input("input.txt")
	solution_1 := part_1(template, rules)
	solution_2 := part_2(template, rules)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
