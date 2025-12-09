package soln

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type CoordinatesD09 struct {
	x int
	y int
}

func Y2025D09P1(filename string) int {
	coordinates, err := parseD09(filename)
	if err != nil {
		log.Fatal(err)
	}

	largestArea := 0
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			x1, x2 := coordinates[i].x, coordinates[j].x
			y1, y2 := coordinates[i].y, coordinates[j].y
			area := int((math.Abs(float64(x1-x2)) + 1) * (math.Abs(float64(y1-y2)) + 1))
			if area > largestArea {
				largestArea = area
			}
		}
	}
	return largestArea
}

func parseD09(filename string) ([]*CoordinatesD09, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coordinates := []*CoordinatesD09{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			stringCoords := strings.Split(line, ",")
			if len(stringCoords) != 2 {
				return nil, fmt.Errorf("invalid coordinates: %s", line)
			}

			intCoords := [2]int{}
			for i, stringCoord := range stringCoords {
				intCoord, err := strconv.Atoi(stringCoord)
				if err != nil {
					return nil, err
				}
				intCoords[i] = intCoord
			}

			coordinates = append(coordinates, &CoordinatesD09{
				x: intCoords[0],
				y: intCoords[1],
			})
		}
	}
	return coordinates, nil
}
