package soln

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

		new, _ := rotateDial(dial, rotation, degree)
		dial = new
		if dial == 0 {
			password++
		}
	}

	return password
}

func Y2025D1P2(filename string) int {
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

		new, passes := rotateDial(dial, rotation, degree)
		dial = new
		password += passes
		if dial == 0 {
			password++
		}
		fmt.Printf("Dial: %d, Passes: %d, Password: %d\n", dial, passes, password)
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

	degree, err = strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid degree %w", err)
	}

	return rotation, degree, err
}

func rotateDial(dial int, rotation rune, degree int) (new int, passes int) {
	if rotation == 'L' {
		new = (dial - degree) % 100
		if new < 0 {
			new += 100
		}

		if dial == 0 {
			passes = -(dial - degree) / 100
		} else {
			passes = -(dial - degree - 99) / 100
		}

		return new, passes
	}

	if rotation == 'R' {
		new = (dial + degree) % 100
		passes = (dial + degree - 1) / 100

		return new, passes
	}

	return dial, 0
}
