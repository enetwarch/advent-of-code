package soln

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
	Numbers   []int
	Operation rune
}

func Y2025D6P1(filename string) int {
	homeworks, err := parseMathHomeworkNormally(filename)
	if err != nil {
		log.Fatalf("failed to parse math homework: %s", err)
	}

	total := 0
	for _, homework := range homeworks {
		answer := answerHomework(homework)
		total += answer
	}
	return total
}

func Y2025D6P2(filename string) int {
	homeworks, err := parseMathHomeworkVertically(filename)
	if err != nil {
		log.Fatalf("failed to parse math homework: %s", err)
	}

	total := 0
	for i := len(homeworks) - 1; i >= 0; i-- {
		answer := answerHomework(homeworks[i])
		total += answer
	}
	return total
}

func answerHomework(homework Homework) int {
	answer := homework.Numbers[0]
	for i := 1; i < len(homework.Numbers); i++ {
		switch homework.Operation {
		case '+':
			answer += homework.Numbers[i]
		case '*':
			answer *= homework.Numbers[i]
		}
	}
	return answer
}

func parseMathHomeworkNormally(filename string) (homeworks []Homework, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

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
				homeworks[i].Numbers = append(homeworks[i].Numbers, number)
			}
		} else {
			for i, v := range strings.Fields(line) {
				homeworks[i].Operation = rune(strings.TrimSpace(v)[0])
			}
		}
	}

	return homeworks, nil
}

func parseMathHomeworkVertically(filename string) (homeworks []Homework, err error) {
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
				homeworks[hIndex].Operation = char
			}
		}

		if number == 0 {
			hIndex++
		} else {
			homeworks[hIndex].Numbers = append(homeworks[hIndex].Numbers, number)
		}
	}

	return homeworks, nil
}
