package day05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	lower int64
	upper int64
}

func Part1(filename string) int {
	ranges, ids, err := parseFile(filename)
	if err != nil {
		log.Fatalf("failed to parse fresh range and ids: %s", err)
	}

	freshId := func(ranges []Range, id int64) bool {
		for _, v := range ranges {
			if id >= v.lower && id <= v.upper {
				return true
			}
		}
		return false
	}

	freshIds := 0
	for _, id := range ids {
		if freshId(ranges, id) {
			freshIds++
		}
	}
	return freshIds
}

func Part2(filename string) int64 {
	ranges, _, err := parseFile(filename)
	if err != nil {
		log.Fatalf("failed to parse fresh range and ids: %s", err)
	}
	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i].lower < ranges[j].lower
	})

	var freshIds int64 = 0
	var currentHighestupper int64 = 0
	for i := 0; i < len(ranges); i++ {
		if ranges[i].lower > currentHighestupper {
			freshIds += ranges[i].upper - ranges[i].lower + 1
			currentHighestupper = ranges[i].upper
		} else if ranges[i].upper > currentHighestupper {
			freshIds += ranges[i].upper - currentHighestupper
			currentHighestupper = ranges[i].upper
		}
	}
	return freshIds
}

func parseFile(filename string) (ranges []Range, ids []int64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			break
		}

		currentRange := strings.Split(line, "-")
		if len(currentRange) != 2 {
			return nil, nil, fmt.Errorf("incorrect range input")
		}

		lower, err := strconv.ParseInt(currentRange[0], 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing lower range %d", lower)
		}

		upper, err := strconv.ParseInt(currentRange[1], 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing upper range %d", upper)
		}

		ranges = append(ranges, Range{lower, upper})
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			continue
		}

		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing id %d", id)
		}

		ids = append(ids, id)
	}
	return ranges, ids, nil
}
