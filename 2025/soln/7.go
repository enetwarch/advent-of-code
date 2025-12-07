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

	splits, err := calculateBeamSplits(manifold)
	if err != nil {
		log.Fatalf("failed to calculate beam splits: %s", err)
	}

	return splits
}

func calculateBeamSplits(manifold []string) (splits int, err error) {
	splits = 0
	beamIndices := map[int]struct{}{} // Hashset
	for i, v := range manifold[0] {
		if v == 'S' {
			beamIndices[i] = struct{}{}
			break
		}
	}

	if len(beamIndices) != 1 {
		return 0, fmt.Errorf("no beam starting point")
	}

	for i := 1; i < len(manifold); i++ {
		for j := 0; j < len(manifold[i]); j++ {
			if manifold[i][j] != '^' {
				continue
			}

			if _, beamExists := beamIndices[j]; beamExists {
				delete(beamIndices, j)

				if j-1 >= 0 {
					beamIndices[j-1] = struct{}{}
				}

				if j+1 < len(manifold[i]) {
					beamIndices[j+1] = struct{}{}
				}

				splits++
			}
		}
	}

	return splits, nil
}

func parseTachyonManifold(filename string) (manifold []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			manifold = append(manifold, line)
		}
	}

	for _, v := range manifold {
		if len(v) != len(manifold[0]) {
			return nil, fmt.Errorf("unbalanced manifold")
		}
	}

	return manifold, nil
}
