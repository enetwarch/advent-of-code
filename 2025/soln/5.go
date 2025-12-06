package soln

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Y2025D5P1(filename string) int {
	ranges, ids, err := parseFreshRangeAndIDs(filename)
	if err != nil {
		log.Fatalf("failed to parse fresh range and ids: %s", err)
	}

	freshIds := 0
	for _, id := range ids {
		if isFreshId(ranges, id) {
			freshIds++
		}
	}

	return freshIds
}

func Y2025D5P2(filename string) int64 {
	ranges, _, err := parseFreshRangeAndIDs(filename)
	if err != nil {
		log.Fatalf("failed to parse fresh range and ids: %s", err)
	}

	return countFreshIds(ranges)
}

func isFreshId(ranges []Range, id int64) bool {
	for _, v := range ranges {
		if id >= v.Lower && id <= v.Upper {
			return true
		}
	}

	return false
}

func countFreshIds(ranges []Range) int64 {
	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i].Lower < ranges[j].Lower
	})

	var freshIds int64 = 0
	var currentHighestUpper int64 = 0

	for i := 0; i < len(ranges); i++ {
		if ranges[i].Lower > currentHighestUpper {
			freshIds += ranges[i].Upper - ranges[i].Lower + 1
			currentHighestUpper = ranges[i].Upper
		} else if ranges[i].Upper > currentHighestUpper {
			freshIds += ranges[i].Upper - currentHighestUpper
			currentHighestUpper = ranges[i].Upper
		}
	}

	return freshIds
}

func parseFreshRangeAndIDs(filename string) (ranges []Range, ids []int64, err error) {
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
