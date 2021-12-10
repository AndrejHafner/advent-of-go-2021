package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Syntax struct {
	sequence []string
}

type Stack struct {
	s []string
}

var ILLEGAL_CHAR_SCORE = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var INCOMPLETE_CHAR_SCORE = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func (stack *Stack) Push(str string) {
	(*stack).s = append((*stack).s, str)
}

func (stack *Stack) Pop() string {
	item := (*stack).s[len((*stack).s)-1]
	(*stack).s = (*stack).s[:len((*stack).s)-1]
	return item
}

func (stack *Stack) Size() int {
	return len((*stack).s)
}

func read_input(filename string) []Syntax {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	lines := strings.Split(string(bs), "\n")

	var sequences []Syntax

	for _, line := range lines {
		line_split := strings.Split(strings.ReplaceAll(line, "\r", ""), "")
		sequences = append(sequences, Syntax{line_split})
	}

	return sequences
}

func part_1(sequences []Syntax) int {
	score := 0

	for _, seq := range sequences {
		stack := &Stack{}
		for _, char := range seq.sequence {
			if char == "(" || char == "{" || char == "[" || char == "<" {
				// Opening chars
				stack.Push(char)
			} else {
				item := stack.Pop()
				if (item == "{" && char == "}") || (item == "[" && char == "]") || (item == "(" && char == ")") || (item == "<" && char == ">") {
					continue
				} else {
					score += ILLEGAL_CHAR_SCORE[char]
				}
			}
		}
	}

	return score
}

func part_2(sequences []Syntax) int {
	score := 0
	var scores []int
	corrupted_line := false
	for _, seq := range sequences {
		stack := &Stack{}
		score = 0
		corrupted_line = false
		for _, char := range seq.sequence {
			if char == "(" || char == "{" || char == "[" || char == "<" {
				// Opening chars
				stack.Push(char)
			} else {
				item := stack.Pop()
				if (item == "{" && char == "}") || (item == "[" && char == "]") || (item == "(" && char == ")") || (item == "<" && char == ">") {
					continue
				} else {
					corrupted_line = true
					break
				}
			}
		}

		if corrupted_line {
			continue
		}

		for {
			if stack.Size() == 0 {
				break
			}

			item := stack.Pop()
			score *= 5
			score += INCOMPLETE_CHAR_SCORE[item]
		}
		scores = append(scores, score)
	}
	sort.Sort(sort.IntSlice(scores))

	idx := (len(scores) / 2)

	return scores[idx]
}

func main() {
	input := read_input("input.txt")
	solution_1 := part_1(input)
	solution_2 := part_2(input)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
