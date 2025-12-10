package day06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Homework struct {
	numbers   []int
	operation rune
}

func Part1(filename string) int {
	homeworks, err := parseFileHorizontally(filename)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, homework := range homeworks {
		answer := answerHomework(homework)
		total += answer
	}
	return total
}

func Part2(filename string) int {
	homeworks, err := parseFileVertically(filename)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for i := len(homeworks) - 1; i >= 0; i-- {
		answer := answerHomework(homeworks[i])
		total += answer
	}
	return total
}

func answerHomework(homework Homework) int {
	answer := homework.numbers[0]
	for i := 1; i < len(homework.numbers); i++ {
		switch homework.operation {
		case '+':
			answer += homework.numbers[i]
		case '*':
			answer *= homework.numbers[i]
		}
	}
	return answer
}

func parseFileHorizontally(filename string) (homeworks []Homework, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	for _, v := range strings.Fields(line) {
		number, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return nil, err
		}
		homeworks = append(homeworks, Homework{[]int{number}, 0})
	}

	for scanner.Scan() {
		line := scanner.Text()
		if unicode.IsNumber(rune(strings.TrimSpace(line)[0])) {
			for i, v := range strings.Fields(line) {
				number, err := strconv.Atoi(strings.TrimSpace(v))
				if err != nil {
					return nil, err
				}
				homeworks[i].numbers = append(homeworks[i].numbers, number)
			}
		} else {
			for i, v := range strings.Fields(line) {
				homeworks[i].operation = rune(strings.TrimSpace(v)[0])
			}
		}
	}
	return homeworks, nil
}

func parseFileVertically(filename string) (homeworks []Homework, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		if len(line) != len(lines[0]) {
			return nil, fmt.Errorf("unbalanced lines")
		}
	}

	hIndex := 0
	for j := len(lines[0]) - 1; j >= 0; j-- {
		if len(homeworks) == hIndex {
			homeworks = append(homeworks, Homework{})
		}

		number := 0
		for i := 0; i < len(lines); i++ {
			char := rune(lines[i][j])
			if unicode.IsDigit(char) {
				number = (number * 10) + int(char-'0')
			} else if char == '+' || char == '*' {
				homeworks[hIndex].operation = char
			}
		}

		if number == 0 {
			hIndex++
		} else {
			homeworks[hIndex].numbers = append(homeworks[hIndex].numbers, number)
		}
	}
	return homeworks, nil
}
