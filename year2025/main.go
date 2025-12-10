package main

import (
	"fmt"
	"log"
	"time"
	"year2025/day01"
	"year2025/day02"
	"year2025/day03"
	"year2025/day04"
	"year2025/day05"
	"year2025/day06"
	"year2025/day07"
	"year2025/day08"
	"year2025/day09"
)

func main() {
	solve(1, 1, day01.Part1)
	solve(1, 2, day01.Part2)

	solve(2, 1, day02.Part1)
	solve(2, 2, day02.Part2)

	solve(3, 1, day03.Part1)
	solve(3, 2, day03.Part2)

	solve(4, 1, day04.Part1)
	solve(4, 2, day04.Part2)

	solve(5, 1, day05.Part1)
	solve(5, 2, day05.Part2)

	solve(6, 1, day06.Part1)
	solve(6, 2, day06.Part2)

	solve(7, 1, day07.Part1)
	solve(7, 2, day07.Part2)

	solve(8, 1, day08.Part1)
	solve(8, 2, day08.Part2)

	solve(9, 1, day09.Part1)
	solve(9, 2, day09.Part2)
}

func solve[T any](day, part int, solver func(string) (T, error)) {
	filename := fmt.Sprintf("./day%02d/input.txt", day)
	startingTime := time.Now()
	problemAnswer, err := solver(filename)
	if err != nil {
		log.Fatal(err)
	}
	elapsedTime := time.Since(startingTime)

	label := fmt.Sprintf("Year 2025 Day %d Part %d", day, part)
	value := fmt.Sprintf("%v (took %s)", problemAnswer, elapsedTime)
	fmt.Printf("%s: %s\n", label, value)
}
