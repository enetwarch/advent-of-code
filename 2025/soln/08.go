package soln

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

type CoordinatesD08 struct {
	x int
	y int
	z int
}

type DistanceD08 struct {
	coordinates1 *CoordinatesD08
	coordinates2 *CoordinatesD08
	distance     float64
}

type UnionFindD08 struct {
	parent map[*CoordinatesD08]*CoordinatesD08
	size   map[*CoordinatesD08]int
}

func Y2025D08P1(filename string) int {
	coordinates, err := parseD08(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Coordinate size: %d\n", len(coordinates))

	distances := distancesD08(coordinates)
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

func Y2025D08P2(filename string) int {
	coordinates, err := parseD08(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Coordinate size: %d\n", len(coordinates))

	distances := distancesD08(coordinates)
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
	fmt.Println(x1, x2)
	return x1 * x2
}

func newUnionFind(coordinates []*CoordinatesD08) UnionFindD08 {
	unionFind := UnionFindD08{
		parent: make(map[*CoordinatesD08]*CoordinatesD08),
		size:   make(map[*CoordinatesD08]int),
	}
	for _, coordinate := range coordinates {
		unionFind.parent[coordinate] = coordinate
		unionFind.size[coordinate] = 1
	}
	return unionFind
}

func (unionFind *UnionFindD08) find(child *CoordinatesD08) *CoordinatesD08 {
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

func (unionFind *UnionFindD08) union(coord1 *CoordinatesD08, coord2 *CoordinatesD08) {
	parent1 := unionFind.find(coord1)
	parent2 := unionFind.find(coord2)
	if parent1 != parent2 {
		unionFind.parent[parent2] = parent1
		unionFind.size[parent1] += unionFind.size[parent2]
		delete(unionFind.size, parent2)
	}
}

func distancesD08(coordinates []*CoordinatesD08) []*DistanceD08 {
	// Does not return duplicate coordinates like {A, B}, {B, A}
	distances := []*DistanceD08{}
	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			coords1 := coordinates[i]
			coords2 := coordinates[j]
			if coords1 != coords2 {
				distances = append(distances, &DistanceD08{
					coordinates1: coords1,
					coordinates2: coords2,
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

func parseD08(filename string) ([]*CoordinatesD08, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	coordinates := []*CoordinatesD08{}
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

			coordinates = append(coordinates, &CoordinatesD08{
				x: intCoords[0],
				y: intCoords[1],
				z: intCoords[2],
			})
		}
	}
	return coordinates, nil
}
