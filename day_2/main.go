package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Command struct {
	command string
	dist    int
}

func read_commands(filename string) []Command {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	sonar_sweep := strings.Split(string(bs), "\n")

	var commands []Command

	for _, el := range sonar_sweep {
		tuple := strings.Split(el, " ")
		dist, err := strconv.Atoi(tuple[1])
		if err != nil {
			panic(err)
		}

		command := Command{command: tuple[0], dist: dist}

		commands = append(commands, command)
	}
	return commands
}

func part_1(commands []Command) int {
	x := 0
	depth := 0

	for _, cmd := range commands {
		switch cmd.command {
		case "forward":
			x += cmd.dist
		case "up":
			depth -= cmd.dist
		case "down":
			depth += cmd.dist

		}
	}

	return x * depth
}

func part_2(commands []Command) int {
	x := 0
	depth := 0
	aim := 0

	for _, cmd := range commands {
		switch cmd.command {
		case "forward":
			x += cmd.dist
			depth += aim * cmd.dist
		case "up":
			aim -= cmd.dist
		case "down":
			aim += cmd.dist

		}
	}

	return x * depth
}

func main() {
	commands := read_commands("input.txt")

	solution_1 := part_1(commands)
	solution_2 := part_2(commands)

	fmt.Println(solution_1)
	fmt.Println(solution_2)
}
