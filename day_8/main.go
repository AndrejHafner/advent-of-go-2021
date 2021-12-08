package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var SEGMENT_MAP = map[int][]string{
	0: {"a", "b", "c", "e", "f", "g"},
	1: {"c", "f"},
	2: {"a", "c", "d", "e", "g"},
	3: {"a", "c", "d", "f", "g"},
	4: {"b", "d", "c", "f"},
	5: {"a", "b", "d", "f", "g"},
	6: {"a", "b", "d", "e", "f", "g"},
	7: {"a", "c", "f"},
	8: {"a", "b", "c", "d", "e", "f", "g"},
	9: {"a", "b", "c", "d", "f", "g"},
}

var INDEX_CHAR_MAP = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
	5: "f",
	6: "g",
}

type Entry struct {
	patterns []string
	digits   []string
}

type Map struct {
	from string
	to   string
}

type SegmentMap struct {
	maps []Map
}

func read_input(filename string) []Entry {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	var entries []Entry

	for _, line := range lines {
		line_split := strings.Split(line, "|")
		patterns := strings.Split(strings.Trim(strings.ReplaceAll(line_split[0], "\r", ""), " "), " ")
		digits := strings.Split(strings.Trim(strings.ReplaceAll(line_split[1], "\r", ""), " "), " ")

		entries = append(entries, Entry{patterns, digits})
	}

	return entries
}

func create_segment_maps(patterns []string) [7]SegmentMap {
	var segment_map [7]SegmentMap

	var maps []Map

	for _, pattern := range patterns {
		switch len(pattern) {
		case 2: // 1
			pattern_split := strings.Split(pattern, "")
			for _, from := range pattern_split {
				for _, to := range SEGMENT_MAP[1] {
					maps = append(maps, Map{from, to})
				}
			}

		case 3: // 7
			pattern_split := strings.Split(pattern, "")
			for _, from := range pattern_split {
				for _, to := range SEGMENT_MAP[7] {
					maps = append(maps, Map{from, to})
				}
			}
		case 4: // 4
			pattern_split := strings.Split(pattern, "")
			for _, from := range pattern_split {
				for _, to := range SEGMENT_MAP[4] {
					maps = append(maps, Map{from, to})
				}
			}
		case 7: // 8
			// pattern_split := strings.Split(pattern, "")
			// for _, from := range pattern_split {
			// 	for _, to := range SEGMENT_MAP[8] {
			// 		maps = append(maps, Map{from, to})
			// 	}
			// }
		case 5: // 2, 3, 5
			// for _, i := range []int{2, 3, 5} {
			// 	pattern_split := strings.Split(pattern, "")
			// 	for _, from := range pattern_split {
			// 		for _, to := range SEGMENT_MAP[i] {
			// 			maps = append(maps, Map{from, to})
			// 		}
			// 	}
			// }
		case 6: // 0, 6, 9
			// for _, i := range []int{0, 6, 9} {
			// 	pattern_split := strings.Split(pattern, "")
			// 	for _, from := range pattern_split {
			// 		for _, to := range SEGMENT_MAP[i] {
			// 			maps = append(maps, Map{from, to})
			// 		}
			// 	}
			// }
		}
	}
	for i := range segment_map {
		var curr_maps []Map
		for _, m := range maps {
			if m.from == INDEX_CHAR_MAP[i] {
				curr_maps = append(curr_maps, m)
			}
		}

		segment_map[i] = SegmentMap{curr_maps}
	}

	return segment_map
}

func decode_segment_map(segment_map SegmentMap, from string) Map {
	counter := make(map[string]int)

	// Populate the map
	for i := 0; i < 7; i++ {
		counter[INDEX_CHAR_MAP[i]] = 0
	}

	for _, m := range segment_map.maps {
		counter[m.to]++
	}

	return Map{}
}

func part_1(entries []Entry) int {
	count := 0
	for _, entry := range entries {
		for _, digit := range entry.digits {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				count++
			}
		}
	}
	return count
}

func part_2(entries []Entry) int {

	var segment_maps [][7]SegmentMap
	for _, entry := range entries {
		segment_map := create_segment_maps(entry.patterns)
		segment_maps = append(segment_maps, segment_map)
		for i := 0; i < 7; i++ {
			r := decode_segment_map(segment_map[i], INDEX_CHAR_MAP[i])
			fmt.Println(r)
		}
	}

	return 4
}

func main() {
	input := read_input("input_2.txt")
	solution_1 := part_1(input)
	solution_2 := part_2(input)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
