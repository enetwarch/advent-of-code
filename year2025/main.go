package main

import (
	"fmt"
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
	printAnswer(1, 1, day01.Part1, "./day01/input.txt")
	printAnswer(1, 2, day01.Part2, "./day01/input.txt")

	printAnswer(2, 1, day02.Part1, "./day02/input.txt")
	printAnswer(2, 2, day02.Part2, "./day02/input.txt")

	printAnswer(3, 1, day03.Part1, "./day03/input.txt")
	printAnswer(3, 2, day03.Part2, "./day03/input.txt")

	printAnswer(4, 1, day04.Part1, "./day04/input.txt")
	printAnswer(4, 2, day04.Part2, "./day04/input.txt")

	printAnswer(5, 1, day05.Part1, "./day05/input.txt")
	printAnswer(5, 2, day05.Part2, "./day05/input.txt")

	printAnswer(6, 1, day06.Part1, "./day06/input.txt")
	printAnswer(6, 2, day06.Part2, "./day06/input.txt")

	printAnswer(7, 1, day07.Part1, "./day07/input.txt")
	printAnswer(7, 2, day07.Part2, "./day07/input.txt")

	printAnswer(8, 1, day08.Part1, "./day08/input.txt")
	printAnswer(8, 2, day08.Part2, "./day08/input.txt")

	printAnswer(9, 1, day09.Part1, "./day09/input.txt")
	printAnswer(9, 2, day09.Part2, "./day09/input.txt")
}

func printAnswer[T any](day, part int, solver func(string) T, filename string) {
	start := time.Now()
	answer := solver(filename)
	elapsed := time.Since(start)

	fmt.Printf("Year 2025 Day %d Part %d: %v (took %s)\n", day, part, answer, elapsed)
}
