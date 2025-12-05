package soln

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Lower int64
	Upper int64
}

func Y2025D2P1(filename string) int64 {
	ranges, err := parseFile(filename)
	if err != nil {
		log.Fatalf("failed to parse file: %s", err)
	}

	var answer int64 = 0
	for _, v := range ranges {
		for id := v.Lower; id <= v.Upper; id++ {
			if isInvalidID(id) {
				answer += id
			}
		}
	}

	return answer
}

func isInvalidID(id int64) bool {
	stringifiedId := strconv.FormatInt(id, 10)
	midIndex := len(stringifiedId) / 2

	return (len(stringifiedId)%2 == 0 &&
		stringifiedId[midIndex:] == stringifiedId[:midIndex])
}

func parseFile(filename string) (ranges []Range, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", err)
	}

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

			lower, err := strconv.ParseInt(ids[0], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to process lower id: %s", err)
			}

			upper, err := strconv.ParseInt(ids[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to process upper id: %s", err)
			}

			ranges = append(ranges, Range{lower, upper})
		}
	}

	return ranges, nil
}
