package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lukpank/go-glpk/glpk"
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

func Part2(filename string) (int, error) {
	machines, err := parseFile(filename)
	if err != nil {
		return 0, err
	}

	totalPresses := 0
	/*
		[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
			[0, 0, 0, 0, 1, 1 | 3]
			[0, 1, 0, 0, 0, 1 | 5]
			[0, 0, 1, 1, 1, 0 | 4]
			[1, 1, 0, 1, 0, 0 | 7]
	*/
	for _, machine := range machines {
		lp := glpk.New()
		lp.SetObjDir(glpk.ObjDir(glpk.MIN))
		lp.AddRows(len(machine.joltageRequirements))
		for i := 0; i < len(machine.joltageRequirements); i++ {
			requirement := float64(machine.joltageRequirements[i])
			lp.SetRowBnds(i+1, glpk.BndsType(glpk.FX), requirement, requirement)
		}
		lp.AddCols(len(machine.wiringSchematics))
		for i := 0; i < len(machine.wiringSchematics); i++ {
			lp.SetColBnds(i+1, glpk.BndsType(glpk.LO), 0.0, 0.0)
			lp.SetColKind(i+1, glpk.VarType(glpk.IV))
			lp.SetObjCoef(i+1, 1)
		}

		contains := func(array []int, value int) bool {
			for i := 0; i < len(array); i++ {
				if array[i] == value {
					return true
				}
			}
			return false
		}

		iArray, jArray := []int32{0}, []int32{0}
		array := []float64{0.0}
		for i := 0; i < len(machine.joltageRequirements); i++ {
			for j := 0; j < len(machine.wiringSchematics); j++ {
				iArray = append(iArray, int32(i+1))
				jArray = append(jArray, int32(j+1))
				if contains(machine.wiringSchematics[j], i) {
					array = append(array, 1.0)
				} else {
					array = append(array, 0.0)
				}
			}
		}
		lp.LoadMatrix(iArray, jArray, array)
		iocp := glpk.NewIocp()
		iocp.SetMsgLev(glpk.MsgLev(glpk.MSG_OFF))
		iocp.SetPresolve(true)
		if err := lp.Intopt(iocp); err != nil {
			return 0, nil
		}
		for i := 0; i < len(machine.wiringSchematics); i++ {
			totalPresses += int(lp.MipColVal(i + 1))
		}
		lp.Delete()
	}
	return totalPresses, nil
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
