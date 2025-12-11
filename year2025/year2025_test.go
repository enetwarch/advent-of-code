package main_test

import (
	"fmt"
	"os"
	"testing"
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
)

type Test[Type comparable] struct {
	day      int
	solver   func(string) (Type, error)
	expected Type
}

func TestDay01Part1(t *testing.T) {
	test(t, Test[int]{day: 1, solver: day01.Part1, expected: 1078})
}

func TestDay01Part2(t *testing.T) {
	test(t, Test[int]{day: 1, solver: day01.Part2, expected: 6412})
}

func TestDay02Part1(t *testing.T) {
	test(t, Test[int64]{day: 2, solver: day02.Part1, expected: 19128774598})
}

func TestDay02Part2(t *testing.T) {
	test(t, Test[int64]{day: 2, solver: day02.Part2, expected: 21932258645})
}

func TestDay03Part1(t *testing.T) {
	test(t, Test[int64]{day: 3, solver: day03.Part1, expected: 17301})
}

func TestDay03Part2(t *testing.T) {
	test(t, Test[int64]{day: 3, solver: day03.Part2, expected: 172162399742349})
}

func TestDay04Part1(t *testing.T) {
	test(t, Test[int]{day: 4, solver: day04.Part1, expected: 1393})
}

func TestDay04Part2(t *testing.T) {
	test(t, Test[int]{day: 4, solver: day04.Part2, expected: 8643})
}

func TestDay05Part1(t *testing.T) {
	test(t, Test[int]{day: 5, solver: day05.Part1, expected: 517})
}

func TestDay05Part2(t *testing.T) {
	test(t, Test[int64]{day: 5, solver: day05.Part2, expected: 336173027056994})
}

func TestDay06Part1(t *testing.T) {
	test(t, Test[int64]{day: 6, solver: day06.Part1, expected: 4076006202939})
}

func TestDay06Part2(t *testing.T) {
	test(t, Test[int64]{day: 6, solver: day06.Part2, expected: 7903168391557})
}

func TestDay07Part1(t *testing.T) {
	test(t, Test[int]{day: 7, solver: day07.Part1, expected: 1649})
}

func TestDay07Part2(t *testing.T) {
	test(t, Test[int64]{day: 7, solver: day07.Part2, expected: 16937871060075})
}

func TestDay08Part1(t *testing.T) {
	test(t, Test[int]{day: 8, solver: day08.Part1, expected: 46398})
}

func TestDay08Part2(t *testing.T) {
	test(t, Test[int64]{day: 8, solver: day08.Part2, expected: 8141888143})
}

func TestDay09Part1(t *testing.T) {
	test(t, Test[int64]{day: 9, solver: day09.Part1, expected: 4759531084})
}

func TestDay09Part2(t *testing.T) {
	test(t, Test[int64]{day: 9, solver: day09.Part2, expected: 1539238860})
}

func TestDay10Part1(t *testing.T) {
	test(t, Test[int]{day: 10, solver: day10.Part1, expected: 436})
}

func test[Type comparable](t *testing.T, test Test[Type]) {
	filename := fmt.Sprintf("./day%02d/input.txt", test.day)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("file does not exist: %s", filename)
	}

	if answer, err := test.solver(filename); err != nil {
		t.Errorf("solver failed due to err (%s)", err)
	} else if answer != test.expected {
		t.Errorf("expected answer was %v but got: %v", test.expected, answer)
	}
}
