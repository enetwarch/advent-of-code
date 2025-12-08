package main

import (
	"2025/soln"
	"fmt"
	"time"
)

func main() {
	printAnswer(1, 1, soln.Y2025D01P1, "./input/01.txt")
	printAnswer(1, 2, soln.Y2025D01P2, "./input/01.txt")

	printAnswer(2, 1, soln.Y2025D02P1, "./input/02.txt")
	printAnswer(2, 2, soln.Y2025D02P2, "./input/02.txt")

	printAnswer(3, 1, soln.Y2025D03P1, "./input/03.txt")
	printAnswer(3, 2, soln.Y2025D03P2, "./input/03.txt")

	printAnswer(4, 1, soln.Y2025D04P1, "./input/04.txt")
	printAnswer(4, 2, soln.Y2025D04P2, "./input/04.txt")

	printAnswer(5, 1, soln.Y2025D05P1, "./input/05.txt")
	printAnswer(5, 2, soln.Y2025D05P2, "./input/05.txt")

	printAnswer(6, 1, soln.Y2025D06P1, "./input/06.txt")
	printAnswer(6, 2, soln.Y2025D06P2, "./input/06.txt")

	printAnswer(7, 1, soln.Y2025D07P1, "./input/07.txt")
	printAnswer(7, 2, soln.Y2025D07P2, "./input/07.txt")
}

func printAnswer[T any](day, part int, solver func(string) T, filename string) {
	start := time.Now()
	answer := solver(filename)
	elapsed := time.Since(start)

	fmt.Printf("Year 2025 Day %d Part %d: %v (took %s)\n", day, part, answer, elapsed)
}
