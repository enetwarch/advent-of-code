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
	"year2025/day10"
	"year2025/day11"
	"year2025/day12"
)

type Problem[Type comparable] struct {
	day    int
	part   int
	solver func(string) (Type, error)
}

func main() {
	solve(Problem[int]{day: 1, part: 1, solver: day01.Part1})
	solve(Problem[int]{day: 1, part: 2, solver: day01.Part2})

	solve(Problem[int64]{day: 2, part: 1, solver: day02.Part1})
	solve(Problem[int64]{day: 2, part: 2, solver: day02.Part2})

	solve(Problem[int64]{day: 3, part: 1, solver: day03.Part1})
	solve(Problem[int64]{day: 3, part: 2, solver: day03.Part2})

	solve(Problem[int]{day: 4, part: 1, solver: day04.Part1})
	solve(Problem[int]{day: 4, part: 2, solver: day04.Part2})

	solve(Problem[int]{day: 5, part: 1, solver: day05.Part1})
	solve(Problem[int64]{day: 5, part: 2, solver: day05.Part2})

	solve(Problem[int64]{day: 6, part: 1, solver: day06.Part1})
	solve(Problem[int64]{day: 6, part: 2, solver: day06.Part2})

	solve(Problem[int]{day: 7, part: 1, solver: day07.Part1})
	solve(Problem[int64]{day: 7, part: 2, solver: day07.Part2})

	solve(Problem[int]{day: 8, part: 1, solver: day08.Part1})
	solve(Problem[int64]{day: 8, part: 2, solver: day08.Part2})

	solve(Problem[int64]{day: 9, part: 1, solver: day09.Part1})
	solve(Problem[int64]{day: 9, part: 2, solver: day09.Part2})

	solve(Problem[int]{day: 10, part: 1, solver: day10.Part1})

	solve(Problem[int]{day: 11, part: 1, solver: day11.Part1})
	solve(Problem[int64]{day: 11, part: 2, solver: day11.Part2})

	solve(Problem[int]{day: 12, part: 1, solver: day12.Part1})
}

func solve[Type comparable](problem Problem[Type]) {
	filename := fmt.Sprintf("./day%02d/input.txt", problem.day)
	startingTime := time.Now()
	problemAnswer, err := problem.solver(filename)
	if err != nil {
		log.Fatal(err)
	}
	elapsedTime := time.Since(startingTime)

	label := fmt.Sprintf("Year 2025 Day %d Part %d", problem.day, problem.part)
	value := fmt.Sprintf("%v (took %s)", problemAnswer, elapsedTime)
	fmt.Printf("%s: %s\n", label, value)
}
