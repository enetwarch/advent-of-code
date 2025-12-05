package soln

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Y2025D1P1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	password := 0
	dial := 50

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 1 {
			break
		}

		rotation, degree, err := processLine(line)
		if err != nil {
			log.Fatalf("failed to process line: %s", err)
		}

		dial = rotateDial(dial, rotation, degree)
		if dial == 0 {
			password++
		}
	}

	return password
}

func processLine(line string) (rotation rune, degree int, err error) {
	if len(line) <= 1 {
		return 0, 0, errors.New("line is too short")
	}

	rotation = rune(line[0])
	if !(rotation == 'L' || rotation == 'R') {
		return 0, 0, fmt.Errorf("unknown rotation rune %c", rotation)
	}

	degree, err = strconv.Atoi(line[1:])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid degree %w", err)
	}

	return rotation, degree, err
}

func rotateDial(dial int, rotation rune, degree int) int {
	switch rotation {
	case 'L':
		dial -= degree
	case 'R':
		dial += degree
	}

	dial %= 100
	if dial < 0 {
		dial += 100
	}

	return dial
}
