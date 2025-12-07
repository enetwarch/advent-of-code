package soln

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Y2025D7P1(filename string) int {
	manifold, err := parseTachyonManifold(filename)
	if err != nil {
		log.Fatalf("failed to parse tachyon manifold: %s", err)
	}

	splits, _, err := calculateBeamSplitTimelines(manifold)
	if err != nil {
		log.Fatalf("failed to calculate beam splits: %s", err)
	}

	return splits
}

func Y2025D7P2(filename string) int {
	manifold, err := parseTachyonManifold(filename)
	if err != nil {
		log.Fatalf("failed to parse tachyon manifold: %s", err)
	}

	_, timelines, err := calculateBeamSplitTimelines(manifold)
	if err != nil {
		log.Fatalf("failed to calculate beam splits: %s", err)
	}

	return timelines
}

func calculateBeamSplitTimelines(manifold []string) (splits int, timelines int, err error) {
	splits = 0
	timelines = 0

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
		for j := 0; j < len(manifold[i]); j++ {
			if manifold[i][j] != '^' {
				continue
			}

			if count, exists := beamIndices[j]; exists {
				delete(beamIndices, j)
				beamIndices[j-1] += count
				beamIndices[j+1] += count
				timelines += count
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
