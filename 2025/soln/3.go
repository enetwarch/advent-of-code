package soln

import (
	"bufio"
	"log"
	"os"
)

func Y2025D3P1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	totalJoltage := 0

	for scanner.Scan() {
		line := scanner.Text()
		totalJoltage += getLargestJoltage(line)
	}

	return totalJoltage
}

func getLargestJoltage(bank string) int {
	leftNumber := '0'
	rightNumber := '0'

	for i := 0; i < len(bank); i++ {
		currentNumber := rune(bank[i])
		if currentNumber > leftNumber && i < len(bank)-1 {
			nextNumber := rune(bank[i+1])

			leftNumber = currentNumber
			rightNumber = nextNumber
		} else if currentNumber > rightNumber {
			rightNumber = currentNumber
		}
	}

	return (int(leftNumber-'0') * 10) + int(rightNumber-'0')
}
