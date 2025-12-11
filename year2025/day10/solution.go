package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	lightDiagram        string
	wiringSchematics    [][]int
	joltageRequirements []int
}

func Part1(filename string) (int, error) {
	machines, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	type State struct {
		diagram uint16
		presses int
	}

	fewestButtonPresses := 0
	for _, machine := range machines {
		targetDiagram := machine.bitmaskDiagram()
		schematics := machine.bitmaskSchematics()
		encounteredDiagrams := map[uint16]bool{}
		bfsQueue := []State{{diagram: 0, presses: 0}}
		for len(bfsQueue) > 0 {
			state := bfsQueue[0]
			bfsQueue = bfsQueue[1:]
			if state.diagram == targetDiagram {
				fewestButtonPresses += state.presses
				break
			}

			encounteredDiagrams[state.diagram] = true
			for _, schematic := range schematics {
				newDiagram := state.diagram ^ schematic
				if !encounteredDiagrams[newDiagram] {
					bfsQueue = append(bfsQueue, State{
						diagram: newDiagram,
						presses: state.presses + 1,
					})
				}
			}
		}
	}
	return fewestButtonPresses, nil
}

func parseFile(filename string) ([]Machine, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	machines := []Machine{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if line[0] != '[' || line[len(line)-1] != '}' {
			return nil, fmt.Errorf("invalid machine: %s", line)
		}
		runes := []rune(line)

		parseNumbers := func(runes []rune, left, right int) ([]int, error) {
			numbers := []int{}
			stringifiedNumbers := strings.Split(string(runes[left+1:right]), ",")
			for _, stringifiedNumber := range stringifiedNumbers {
				number, err := strconv.Atoi(stringifiedNumber)
				if err != nil {
					return nil, err
				}
				numbers = append(numbers, number)
			}
			return numbers, nil
		}

		machine := Machine{}
		leftPointer, rightPointer := 0, 0
		for rightPointer < len(runes) {
			switch runes[rightPointer] {
			case '[', '(', '{':
				leftPointer = rightPointer
			case ']':
				if runes[leftPointer] != '[' {
					return nil, fmt.Errorf("unclosed bracket: %s", line)
				} else if machine.lightDiagram != "" {
					return nil, fmt.Errorf("duplicate indicator light diagram: %s", line)
				}

				machine.lightDiagram = string(line[leftPointer+1 : rightPointer])
				leftPointer = rightPointer
			case ')':
				if runes[leftPointer] != '(' {
					return nil, fmt.Errorf("unclosed parenthesis: %s", line)
				}

				schematic, err := parseNumbers(runes, leftPointer, rightPointer)
				if err != nil {
					return nil, err
				}
				machine.wiringSchematics = append(machine.wiringSchematics, schematic)
				leftPointer = rightPointer
			case '}':
				if runes[leftPointer] != '{' {
					return nil, fmt.Errorf("unclosed braces: %s", line)
				} else if machine.joltageRequirements != nil {
					return nil, fmt.Errorf("duplicate joltage requirements: %s", line)
				}

				requirements, err := parseNumbers(runes, leftPointer, rightPointer)
				if err != nil {
					return nil, err
				}
				machine.joltageRequirements = requirements
				leftPointer = rightPointer
			}
			rightPointer++
		}
		machines = append(machines, machine)
	}
	return machines, nil
}

func (machine *Machine) bitmaskDiagram() uint16 {
	var maskedDiagram uint16 = 0
	for i, diagramNode := range machine.lightDiagram {
		if diagramNode == '#' {
			maskedDiagram |= 1 << i
		}
	}
	return maskedDiagram
}

func (machine *Machine) bitmaskSchematics() []uint16 {
	maskedSchematics := []uint16{}
	for _, wiringSchematic := range machine.wiringSchematics {
		var schematic uint16 = 0
		for _, wireNumber := range wiringSchematic {
			schematic |= 1 << wireNumber
		}
		maskedSchematics = append(maskedSchematics, schematic)
	}
	return maskedSchematics
}
