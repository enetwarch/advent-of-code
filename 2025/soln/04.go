package soln

import (
	"bufio"
	"log"
	"os"
)

type Coordinates struct {
	I int
	J int
}

func Y2025D04P1(filename string) int {
	grid, err := parseToiletPaperGrid(filename)
	if err != nil {
		log.Fatalf("failed to parse toiler paper grid: %s", err)
	}

	accessibleRolls, _ := accessToiletPaperGrid(grid)
	return accessibleRolls
}

func Y2025D04P2(filename string) int {
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

		cleanToiletPaperGrid(grid, cleaned)
		cleanedRolls += accessed
	}

	return cleanedRolls
}

func accessToiletPaperGrid(grid [][]rune) (accessed int, cleaned []Coordinates) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' && isAccessibleByForklift(grid, i, j) {
				cleaned = append(cleaned, Coordinates{i, j})
				accessed++
			}
		}
	}

	return accessed, cleaned
}

func cleanToiletPaperGrid(grid [][]rune, cleaned []Coordinates) {
	for _, v := range cleaned {
		i, j := v.I, v.J
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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid, nil
}
