package main

import (
	"2025/soln"
	"fmt"
	"time"
)

func main() {
	printAnswer(1, 1, soln.Y2025D1P1, "./input/1.txt")
	printAnswer(1, 2, soln.Y2025D1P2, "./input/1.txt")

	printAnswer(2, 1, soln.Y2025D2P1, "./input/2.txt")
	printAnswer(2, 2, soln.Y2025D2P2, "./input/2.txt")

	printAnswer(3, 1, soln.Y2025D3P1, "./input/3.txt")
	printAnswer(3, 2, soln.Y2025D3P2, "./input/3.txt")

	printAnswer(4, 1, soln.Y2025D4P1, "./input/4.txt")
	printAnswer(4, 2, soln.Y2025D4P2, "./input/4.txt")

	printAnswer(5, 1, soln.Y2025D5P1, "./input/5.txt")
	printAnswer(5, 2, soln.Y2025D5P2, "./input/5.txt")

	printAnswer(6, 1, soln.Y2025D6P1, "./input/6.txt")
	printAnswer(6, 2, soln.Y2025D6P2, "./input/6.txt")

	printAnswer(7, 1, soln.Y2025D7P1, "./input/7.txt")
	printAnswer(7, 2, soln.Y2025D7P2, "./input/7.txt")
}

func printAnswer[T any](day, part int, solver func(string) T, filename string) {
	start := time.Now()
	answer := solver(filename)
	elapsed := time.Since(start)

	fmt.Printf("Year 2025 Day %d Part %d: %v (took %s)\n", day, part, answer, elapsed)
}
