package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	lower int
	upper int
}

func Part1(filename string) (int, error) {
	ranges, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	invalid := func(id int) bool {
		stringifiedId := strconv.Itoa(id)
		midIndex := len(stringifiedId) / 2
		return (len(stringifiedId)%2 == 0 &&
			stringifiedId[midIndex:] == stringifiedId[:midIndex])
	}

	var answer int = 0
	for _, v := range ranges {
		for id := v.lower; id <= v.upper; id++ {
			if invalid(id) {
				answer += id
			}
		}
	}
	return answer, nil
}

func Part2(filename string) (int, error) {
	ranges, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	invalid := func(id int) bool {
		stringifiedId := strconv.Itoa(id)
		for length := 1; length <= len(stringifiedId)/2; length++ {
			if len(stringifiedId)%length != 0 {
				continue
			}

			toRepeat := stringifiedId[:length]
			isInvalid := true
			for i := length; i < len(stringifiedId); i += length {
				if stringifiedId[i:(i+length)] != toRepeat {
					isInvalid = false
					break
				}
			}
			if isInvalid {
				return true
			}
		}
		return false
	}

	var answer int = 0
	for _, v := range ranges {
		for id := v.lower; id <= v.upper; id++ {
			if invalid(id) {
				answer += id
			}
		}
	}
	return answer, nil
}

func parseFile(filename string) (ranges []Range, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges = []Range{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineRanges := strings.Split(line, ",")
		for _, v := range lineRanges {
			if v == "" {
				continue
			}

			ids := strings.Split(v, "-")
			if len(ids) != 2 {
				return nil, fmt.Errorf("failed to process ids")
			}

			lower, err := strconv.Atoi(ids[0])
			if err != nil {
				return nil, fmt.Errorf("failed to process lower id: %s", err)
			}

			upper, err := strconv.Atoi(ids[1])
			if err != nil {
				return nil, fmt.Errorf("failed to process upper id: %s", err)
			}

			ranges = append(ranges, Range{lower, upper})
		}
	}
	return ranges, nil
}
