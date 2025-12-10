package day08

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coordinates struct {
	x int
	y int
	z int
}

type Distance struct {
	coordinates1 *Coordinates
	coordinates2 *Coordinates
	distance     float64
}

type UnionFind struct {
	parent map[*Coordinates]*Coordinates
	size   map[*Coordinates]int
}

func Part1(filename string) int {
	coordinates, err := parseFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	distances := calculateDistances(coordinates)
	sort.Slice(distances, func(i int, j int) bool {
		// This will sort distances by ascending distance.
		return distances[i].distance < distances[j].distance
	})

	unionFind := newUnionFind(coordinates)
	for i := 0; i < len(coordinates); i++ {
		coord1 := distances[i].coordinates1
		coord2 := distances[i].coordinates2
		if unionFind.find(coord1) != unionFind.find(coord2) {
			unionFind.union(coord1, coord2)
		}
	}

	sizes := []int{}
	for _, size := range unionFind.size {
		sizes = append(sizes, size)
	}
	sort.Slice(sizes, func(i int, j int) bool {
		// Sorts this int slice in descending order.
		return sizes[i] > sizes[j]
	})

	answer := sizes[0]
	for i := 1; i < 3 && i < len(sizes); i++ {
		answer *= sizes[i]
	}
	return answer
}

func Part2(filename string) int {
	coordinates, err := parseFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	distances := calculateDistances(coordinates)
	sort.Slice(distances, func(i int, j int) bool {
		// This will sort distances by ascending distance.
		return distances[i].distance < distances[j].distance
	})

	unionFind := newUnionFind(coordinates)
	connections := 0
	var x1, x2 int
	for _, distance := range distances {
		coord1 := distance.coordinates1
		coord2 := distance.coordinates2
		if unionFind.find(coord1) != unionFind.find(coord2) {
			unionFind.union(coord1, coord2)
			connections++
			if connections == len(coordinates)-1 {
				x1, x2 = coord1.x, coord2.x
				break
			}
		}
	}
	return x1 * x2
}

func newUnionFind(coordinates []Coordinates) UnionFind {
	unionFind := UnionFind{
		parent: make(map[*Coordinates]*Coordinates),
		size:   make(map[*Coordinates]int),
	}
	for _, coordinate := range coordinates {
		unionFind.parent[&coordinate] = &coordinate
		unionFind.size[&coordinate] = 1
	}
	return unionFind
}

func (unionFind *UnionFind) find(child *Coordinates) *Coordinates {
	parent := unionFind.parent[child]
	if parent == nil {
		unionFind.parent[child] = child
		unionFind.size[child] = 1
		return child
	} else if parent != child {
		// Makes every child point to the main parent representative.
		root := unionFind.find(parent)
		unionFind.parent[child] = root
		return root
	}
	return parent
}

func (unionFind *UnionFind) union(coord1, coord2 *Coordinates) {
	parent1 := unionFind.find(coord1)
	parent2 := unionFind.find(coord2)
	if parent1 != parent2 {
		unionFind.parent[parent2] = parent1
		unionFind.size[parent1] += unionFind.size[parent2]
		delete(unionFind.size, parent2)
	}
}

func calculateDistances(coordinates []Coordinates) []Distance {
	// Does not return duplicate coordinates like {A, B}, {B, A}
	distances := []Distance{}
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			coords1 := coordinates[i]
			coords2 := coordinates[j]
			if coords1 != coords2 {
				distances = append(distances, Distance{
					coordinates1: &coords1,
					coordinates2: &coords2,
					distance: math.Sqrt(
						math.Pow(float64(coords1.x-coords2.x), 2) +
							math.Pow(float64(coords1.y-coords2.y), 2) +
							math.Pow(float64(coords1.z-coords2.z), 2),
					),
				})
			}
		}
	}
	return distances
}

func parseFile(filename string) ([]Coordinates, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	coordinates := []Coordinates{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			stringCoords := strings.Split(line, ",")
			if len(stringCoords) != 3 {
				return nil, fmt.Errorf("invalid coordinates: %s", line)
			}

			intCoords := [3]int{}
			for i, stringCoord := range stringCoords {
				intCoord, err := strconv.Atoi(stringCoord)
				if err != nil {
					return nil, err
				}
				intCoords[i] = intCoord
			}

			coordinates = append(coordinates, Coordinates{
				x: intCoords[0], y: intCoords[1], z: intCoords[2],
			})
		}
	}
	return coordinates, nil
}
