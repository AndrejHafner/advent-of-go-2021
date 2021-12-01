package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func read_measurements(filename string) []int {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file!")
		panic(err)
	}

	sonar_sweep := strings.Split(string(bs), "\n")

	var measurements []int

	for _, el := range sonar_sweep {
		n, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}

		measurements = append(measurements, n)
	}
	return measurements
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func part_1(measurements []int) int {
	depth_increases := 0

	for i := 1; i < len(measurements); i++ {

		if measurements[i] > measurements[i-1] {
			depth_increases++
		}
	}

	return depth_increases
}

func part_2(measurements []int) int {
	depth_increases := 0

	for i := 1; i < len(measurements)-2; i++ {

		if sum(measurements[i:i+3]) > sum(measurements[i-1:i+2]) {
			depth_increases++
		}
	}

	return depth_increases
}

func main() {

	measurements := read_measurements("input.txt")

	solution_1 := part_1(measurements)
	solution_2 := part_2(measurements)

	fmt.Println(solution_1)
	fmt.Println(solution_2)

}
