package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

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
			return 0, err
		}

		newDial, _ := rotateDial(rotation, dial, degree)
		dial = newDial
		if dial == 0 {
			password++
		}
	}
	return password, nil
}

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

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
			return 0, err
		}

		new, passes := rotateDial(rotation, dial, degree)
		dial = new
		password += passes
		if dial == 0 {
			password++
		}
	}
	return password, nil
}

func processLine(line string) (rotation rune, degree int, err error) {
	if len(line) <= 1 {
		return 0, 0, fmt.Errorf("line is too short: %s", line)
	}

	rotation = rune(line[0])
	if !(rotation == 'L' || rotation == 'R') {
		return 0, 0, fmt.Errorf("unknown rotation rune: %c", rotation)
	}

	degree, err = strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid degree: %s", line)
	}

	return rotation, degree, err
}

func rotateDial(rotation rune, dial, degree int) (newDial int, zeroPasses int) {
	if rotation == 'L' {
		newDial = (dial - degree) % 100
		if newDial < 0 {
			newDial += 100
		}
		if dial == 0 {
			zeroPasses = -(dial - degree) / 100
		} else {
			zeroPasses = -(dial - degree - 99) / 100
		}
		return newDial, zeroPasses
	} else if rotation == 'R' {
		newDial = (dial + degree) % 100
		zeroPasses = (dial + degree - 1) / 100
		return newDial, zeroPasses
	}
	return dial, 0 // Should never land in this line
}
