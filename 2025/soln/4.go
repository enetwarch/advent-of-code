package soln

import (
	"bufio"
	"log"
	"os"
)

func Y2025D4P1(filename string) int {
	grid, err := parseToiletPaperGrid(filename)
	if err != nil {
		log.Fatalf("failed to parse toiler paper grid: %s", err)
	}

	accessibleRolls, _ := accessToiletPaperGrid(grid)
	return accessibleRolls
}

func Y2025D4P2(filename string) int {
	grid, err := parseToiletPaperGrid(filename)
	if err != nil {
		log.Fatalf("failed to parse toiler paper grid: %s", err)
	}

	cleanedRolls := 0
	for {
		accessed, cleaned := accessToiletPaperGrid(grid)
		if accessed <= 0 {
			break
		}

		cleanedRolls += accessed
		cleanToiletPaperGrid(grid, cleaned)
	}

	return cleanedRolls
}

func accessToiletPaperGrid(grid [][]rune) (accessed int, cleaned [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' && isAccessibleByForklift(grid, i, j) {
				cleaned = append(cleaned, []int{i, j})
				accessed++
			}
		}
	}

	return accessed, cleaned
}

func cleanToiletPaperGrid(grid [][]rune, cleaned [][]int) {
	for _, v := range cleaned {
		i, j := v[0], v[1]
		grid[i][j] = '.'
	}
}

func isAccessibleByForklift(grid [][]rune, i int, j int) bool {
	adjacentRunes := [8]rune{
		accessGridElement(grid, i, j+1),   // NORTH
		accessGridElement(grid, i+1, j+1), // NORTHEAST
		accessGridElement(grid, i+1, j),   // EAST
		accessGridElement(grid, i+1, j-1), // SOUTHEAST
		accessGridElement(grid, i, j-1),   // SOUTH
		accessGridElement(grid, i-1, j-1), // SOUTHWEST
		accessGridElement(grid, i-1, j),   // WEST
		accessGridElement(grid, i-1, j+1), // NORTHWEST
	}

	count := 0
	for _, v := range adjacentRunes {
		if v == '@' {
			count++
		}
	}

	return count < 4
}

func accessGridElement(grid [][]rune, i int, j int) rune {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return 0
	}

	return grid[i][j]
}

func parseToiletPaperGrid(filename string) (grid [][]rune, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid, nil
}
