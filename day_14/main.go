package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func arr_contains(arr []string, el string) bool {
	for _, e := range arr {
		if e == el {
			return true
		}
	}
	return false
}

func get_max_min_element(pair_map *map[string]int, template string) (int, int) {
	counter := map[string]int{}

	template_split := strings.Split(template, "")

	first_el := template_split[0]
	last_el := template_split[len(template_split)-1]

	for pair, count := range *pair_map {
		for _, el := range strings.Split(pair, "") {
			_, contains := counter[el]
			if !contains {
				counter[el] = count
			} else {
				counter[el] += count
			}
		}

	}

	max := math.MinInt64
	max_key := ""
	min := math.MaxInt64
	min_key := ""

	for key, value := range counter {
		if value > max {
			max = value
			max_key = key
		}
		if value < min {
			min = value
			min_key = key
		}
	}

	if arr_contains(strings.Split(max_key, ""), first_el) || arr_contains(strings.Split(max_key, ""), last_el) {
		max = (max)/2 - 1
	} else {
		max = max / 2
	}

	if arr_contains(strings.Split(min_key, ""), first_el) || arr_contains(strings.Split(min_key, ""), last_el) {
		min = (min)/2 - 1
	} else {
		min = min / 2
	}

	return max, min
}

func create_pair_map(template string) map[string]int {
	template_arr := strings.Split(template, "")
	pair_map := map[string]int{}
	for i := 0; i < len(template_arr)-1; i++ {
		pair := template_arr[i] + template_arr[i+1]
		_, contains := pair_map[pair]
		if contains {
			pair_map[pair]++
		} else {
			pair_map[pair] = 1
		}
	}

	return pair_map
}

func pair_insertion(pair_map *map[string]int, insertion_map map[string]string) {
	new_pair_map := map[string]int{}

	for k, v := range *pair_map {
		new_pair_map[k] = v
	}

	for pair, count := range *pair_map {
		insert_val, contains := insertion_map[pair]
		if !contains {
			continue
		}

		pair_split := strings.Split(pair, "")
		new_pair_1 := pair_split[0] + insert_val
		new_pair_2 := insert_val + pair_split[1]

		_, contains_p_1 := new_pair_map[new_pair_1]
		if !contains_p_1 {
			new_pair_map[new_pair_1] = count
		} else {
			new_pair_map[new_pair_1] += count
		}

		_, contains_p_2 := new_pair_map[new_pair_2]
		if !contains_p_2 {
			new_pair_map[new_pair_2] = count
		} else {
			new_pair_map[new_pair_2] += count

		}

		new_pair_map[pair] -= count

	}

	for key, value := range new_pair_map {
		if value == 0 {
			delete(new_pair_map, key)
		}
	}

	*pair_map = new_pair_map
}

func part_1(template string, insertion_map map[string]string) int {
	pair_map := create_pair_map(template)
	n_steps := 10

	for step := 0; step < n_steps; step++ {
		pair_insertion(&pair_map, insertion_map)
	}

	max, min := get_max_min_element(&pair_map, template)

	return max - min
}

func part_2(template string, insertion_map map[string]string) int {
	pair_map := create_pair_map(template)
	n_steps := 40

	for step := 0; step < n_steps; step++ {
		pair_insertion(&pair_map, insertion_map)
	}

	max, min := get_max_min_element(&pair_map, template)

	return max - min
}

func main() {
	template, rules := read_input("input.txt")
	solution_1 := part_1(template, rules)
	solution_2 := part_2(template, rules)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
