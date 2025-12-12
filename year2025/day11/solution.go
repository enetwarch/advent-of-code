package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1(filename string) (int, error) {
	deviceMap, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	outPaths := 0
	dfsStack := []string{}
	dfsStack = append(dfsStack, deviceMap["you"]...)
	for len(dfsStack) > 0 {
		device := dfsStack[len(dfsStack)-1]
		dfsStack = dfsStack[:len(dfsStack)-1]
		if device == "out" {
			outPaths++
			continue
		}
		dfsStack = append(dfsStack, deviceMap[device]...)
	}
	return outPaths, nil
}

func parseFile(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	deviceMap := map[string][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			devices := strings.Split(line, ": ")
			if len(devices) != 2 {
				return nil, fmt.Errorf("invalid device: %s", line)
			}
			deviceMap[devices[0]] = strings.Split(devices[1], " ")
		}
	}
	return deviceMap, nil
}
