package day09

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type PointPair struct {
	first  *Point
	second *Point
	area   int64
}

func Part1(filename string) (int64, error) {
	points, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	var largestArea int64 = 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := calculateArea(points[i], points[j])
			if area > largestArea {
				largestArea = area
			}
		}
	}
	return largestArea, nil
}

func Part2(filename string) (int64, error) {
	points, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	rectangles := []*PointPair{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			rectangles = append(rectangles, &PointPair{
				first:  points[i],
				second: points[j],
				area:   calculateArea(points[i], points[j]),
			})
		}
	}
	sort.Slice(rectangles, func(i int, j int) bool {
		// Sort area by descending order.
		return rectangles[i].area > rectangles[j].area
	})

	for _, rectangle := range rectangles {
		rxMin, rxMax := getMinMax(rectangle.first.x, rectangle.second.x)
		ryMin, ryMax := getMinMax(rectangle.first.y, rectangle.second.y)
		if isRectangleInsidePolygon(points, rxMin, rxMax, ryMin, ryMax) {
			return rectangle.area, nil
		}
	}
	return 0, fmt.Errorf("no valid rectangles") // Should never happen
}

func parseFile(filename string) ([]*Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := []*Point{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			stringCoords := strings.Split(line, ",")
			if len(stringCoords) != 2 {
				return nil, fmt.Errorf("invalid points: %s", line)
			}

			x, err := strconv.Atoi(stringCoords[0])
			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(stringCoords[1])
			if err != nil {
				return nil, err
			}

			points = append(points, &Point{x: x, y: y})
		}
	}
	return points, nil
}

func calculateArea(point1, point2 *Point) int64 {
	abs := func(number int) int {
		if number < 0 {
			return -number
		}
		return number
	}

	x1, x2 := point1.x, point2.x
	y1, y2 := point1.y, point2.y
	return int64((abs(x1-x2) + 1) * (abs(y1-y2) + 1))
}

func getMinMax(number1, number2 int) (min, max int) {
	if number1 < number2 {
		return number1, number2
	}
	return number2, number1
}

func isRectangleInsidePolygon(points []*Point, rxMin, rxMax, ryMin, ryMax int) bool {
	for i := 0; i < len(points); i++ {
		ePoint1, ePoint2 := points[i], points[(i+1)%len(points)]
		exMin, exMax := getMinMax(ePoint1.x, ePoint2.x)
		eyMin, eyMax := getMinMax(ePoint1.y, ePoint2.y)
		if (exMin < rxMax && exMax > rxMin) && (eyMin < ryMax && eyMax > ryMin) {
			return false
		}
	}
	return true
}
