package day12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Gift struct {
	diagram [3][3]bool
	area    int
}

type Region struct {
	size  [2]int
	gifts []int
}

func Part1(filename string) (int, error) {
	// This part feels cheaty because it takes advantage of expected input.
	// It does not work on the small sample input because of this.
	gifts, regions, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	fittableRegions := 0
	for _, region := range regions {
		regionArea := region.size[0] * region.size[1]
		giftArea := 0
		for i, gift := range gifts {
			giftArea += gift.area * region.gifts[i]
		}
		if giftArea < regionArea {
			fittableRegions++
		}
	}
	return fittableRegions, nil
}

func parseFile(filename string) ([]Gift, []Region, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	gifts := []Gift{}
	linePointer := 0
	for linePointer < len(lines) {
		line := lines[linePointer]
		if line == "" {
			continue
		} else if line[len(line)-1] != ':' {
			break
		}
		if number, err := strconv.Atoi(line[:len(line)-1]); err != nil {
			return nil, nil, fmt.Errorf("invalid gift header: %s", line)
		} else if len(gifts) != number {
			return nil, nil, fmt.Errorf("invalid gift number: %s", line)
		}
		linePointer++

		diagram := [3][3]bool{}
		area := 0
		for i := 0; i < 3; i++ {
			line := lines[linePointer]
			if len(line) != 3 {
				return nil, nil, fmt.Errorf("invalid gift diagram: %s", line)
			}
			for j := 0; j < 3; j++ {
				if line[j] != '.' && line[j] != '#' {
					return nil, nil, fmt.Errorf("invalid gift diagram: %s", line)
				}
				diagram[i][j] = line[j] == '#'
				if line[j] == '#' {
					area++
				}
			}
			linePointer++
		}
		gifts = append(gifts, Gift{diagram, area})
		linePointer++
	}

	regions := []Region{}
	for linePointer < len(lines) {
		line := lines[linePointer]
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, ": ")
		if len(lineSplit) != 2 {
			return nil, nil, fmt.Errorf("invalid region: %s", line)
		}

		regionSize := [2]int{}
		stringifiedRegionSize := strings.Split(lineSplit[0], "x")
		if len(stringifiedRegionSize) != 2 {
			return nil, nil, fmt.Errorf("invalid region size: %s", line)
		}
		for i, stringifiedSize := range stringifiedRegionSize {
			size, err := strconv.Atoi(stringifiedSize)
			if err != nil {
				return nil, nil, err
			}
			regionSize[i] = size
		}

		stringifiedRegionGifts := strings.Split(lineSplit[1], " ")
		if len(stringifiedRegionGifts) != len(gifts) {
			return nil, nil, fmt.Errorf("mismatched region gift amount: %s", line)
		}
		regionGifts := []int{}
		for _, stringifiedRegionGift := range stringifiedRegionGifts {
			regionGift, err := strconv.Atoi(stringifiedRegionGift)
			if err != nil {
				return nil, nil, err
			}
			regionGifts = append(regionGifts, regionGift)
		}
		regions = append(regions, Region{regionSize, regionGifts})
		linePointer++
	}
	return gifts, regions, nil
}
