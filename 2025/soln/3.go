package soln

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Y2025D3P1(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	var totalJoltage int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		joltage, err := getLargestJoltage(line, 2)
		if err != nil {
			log.Fatalf("failed to get largest joltage: %s", err)
		}

		totalJoltage += joltage
	}

	return totalJoltage
}

func Y2025D3P2(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	var totalJoltage int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		joltage, err := getLargestJoltage(line, 12)
		if err != nil {
			log.Fatalf("failed to get largest joltage: %s", err)
		}

		totalJoltage += joltage
	}

	return totalJoltage
}

func getLargestJoltage(bank string, digits int) (joltage int64, err error) {
	joltageBank := []rune(strings.Repeat("0", digits))

	for _, v := range bank {
		index, _ := findReplaceableJoltageRune(joltageBank, v)
		if index == -1 {
			continue
		}

		for i := index; i < len(joltageBank)-1; i++ {
			joltageBank[i] = joltageBank[i+1]
		}
		joltageBank[len(joltageBank)-1] = v
	}

	joltage, err = strconv.ParseInt(string(joltageBank), 10, 64)
	if err != nil {
		return 0, err
	}

	return joltage, nil
}

func findReplaceableJoltageRune(bank []rune, insert rune) (index int, jolt rune) {
	for i := 0; i < len(bank)-1; i++ {
		if bank[i] < bank[i+1] {
			return i, bank[i]
		}
	}

	for i := len(bank) - 1; i >= 0; i-- {
		if bank[i] < insert {
			return i, bank[i]
		}
	}

	return -1, '0'
}
