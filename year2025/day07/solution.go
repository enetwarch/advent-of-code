package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1(filename string) int {
	manifold, err := parseTachyonManifold(filename)
	if err != nil {
		log.Fatal(err)
	}

	splits, _, err := calculateSplitTimelines(manifold)
	if err != nil {
		log.Fatal(err)
	}
	return splits
}

func Part2(filename string) int {
	manifold, err := parseTachyonManifold(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, timelines, err := calculateSplitTimelines(manifold)
	if err != nil {
		log.Fatal(err)
	}
	return timelines
}

func calculateSplitTimelines(manifold []string) (splits, timelines int, err error) {
	beamIndices := map[int]int{} // Hashmap
	for i, v := range manifold[0] {
		if v == 'S' {
			beamIndices[i] = 1
			timelines++
			break
		}
	}
	if len(beamIndices) != 1 {
		return 0, 0, fmt.Errorf("no beam starting point")
	}

	for i := 1; i < len(manifold); i++ {
		for j, timelinesHere := range beamIndices {
			if manifold[i][j] == '^' {
				delete(beamIndices, j)
				beamIndices[j-1] += timelinesHere
				beamIndices[j+1] += timelinesHere
				timelines += timelinesHere
				splits++
			}
		}
	}
	return splits, timelines, nil
}

func parseTachyonManifold(filename string) (manifold []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			manifold = append(manifold, line)
		}
	}

	for _, v := range manifold {
		if len(v) != len(manifold[0]) || v[0] != '.' || v[len(v)-1] != '.' {
			return nil, fmt.Errorf("invalid manifold")
		}
	}
	return manifold, nil
}
