package soln

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Y2025D6P1(filename string) int {
	calculations, err := parseMathHomework(filename)
	if err != nil {
		log.Fatalf("failed to parse math homework: %s", err)
	}

	total := 0
	for _, calculation := range calculations {
		answer := calculateCalculations(calculation)
		total += answer
	}
	return total
}

func calculateCalculations(calculation Calculations) int {
	answer := calculation.Numbers[0]
	for i := 1; i < len(calculation.Numbers); i++ {
		switch calculation.Operation {
		case '+':
			answer += calculation.Numbers[i]
		case '*':
			answer *= calculation.Numbers[i]
		}
	}
	return answer
}

func parseMathHomework(filename string) (calculations []Calculations, err error) {
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

		calculations = append(calculations, Calculations{[]int{number}, 0})
	}

	for scanner.Scan() {
		line := scanner.Text()
		if unicode.IsNumber(rune(strings.TrimSpace(line)[0])) {
			for i, v := range strings.Fields(line) {
				number, err := strconv.Atoi(strings.TrimSpace(v))
				if err != nil {
					return nil, err
				}
				calculations[i].Numbers = append(calculations[i].Numbers, number)
			}
		} else {
			for i, v := range strings.Fields(line) {
				calculations[i].Operation = rune(strings.TrimSpace(v)[0])
			}
		}
	}

	return calculations, nil
}
