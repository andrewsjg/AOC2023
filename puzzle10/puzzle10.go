package puzzle10

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type location struct {
	gridLocation coordinate
	connected    []coordinate
	locationChar rune
}

func Part1Solve(input string) {
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	pipeMap, startCoords := makeMap(stringInput)

	stepCount, _ := followMap(startCoords, pipeMap)
	furtherest := float32(stepCount) / float32(2)

	fmt.Printf("Furherest point: %d\n", int(furtherest+0.5))
}

func Part2Solve(input string) {
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	pipeMap, startCoords := makeMap(stringInput)

	_, loopMap := followMap(startCoords, pipeMap)

	scanner := bufio.NewScanner(strings.NewReader(stringInput))

	y := 0
	insideCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := range line {

			testCoordinates := coordinate{x, y}
			if insideLoop(loopMap, testCoordinates) {
				insideCount++
			}
		}
		y++
	}

	fmt.Printf("Total locations inside the loop: %d\n", insideCount)

}
func followMap(startCoordinates coordinate, pipeMap map[coordinate]location) (stepCount int, loopMap map[coordinate]location) {

	startLocation := pipeMap[startCoordinates]
	currentLocation := startLocation.connected[0]
	previousLocation := startLocation

	loopMap = make(map[coordinate]location)
	loopMap[startCoordinates] = startLocation

	for currentLocation != startCoordinates {

		tmpLocation := currentLocation
		loopMap[currentLocation] = pipeMap[currentLocation]

		//currentLocation = pipeMap[currentLocation].connected[1]

		if pipeMap[currentLocation].connected[1] != previousLocation.gridLocation {

			currentLocation = pipeMap[currentLocation].connected[1]

		} else {
			currentLocation = pipeMap[currentLocation].connected[0]
		}

		previousLocation = pipeMap[tmpLocation]

		stepCount++
	}

	return stepCount, loopMap
}

// Is a passed coordinate inside a loop?
func insideLoop(loopMap map[coordinate]location, checkCoord coordinate) bool {
	/*
		The key to this solution is 2 things:

		1.) If any coordinate is inside the loop, traversing the map will require crossing the pipe an odd number of times
		2.) The L7 and FJ junctions can be considered a vertical pipe and treated as one when counting pipe crossings. All other pipe configuriations can be ignored

	*/

	edges := []rune{'|'}
	corners := []rune{'7', 'J', 'F', 'L'}

	edgeCount := 0
	cornerCount := 0

	// coordinate is not part of the pipe
	if _, ok := loopMap[checkCoord]; ok {

		return false
	}

	// Look to the left from coordinate and count the number of times we cross the loop edge
	lastCorner := ' '
	for i := checkCoord.x; i >= 0; i-- {

		testCoord := coordinate{x: i, y: checkCoord.y}
		loopPosition := loopMap[testCoord]

		if slices.Contains(edges, loopPosition.locationChar) {
			lastCorner = ' '
			edgeCount++

		}

		if slices.Contains(corners, loopPosition.locationChar) {
			cornerCount++

			if loopPosition.locationChar == 'L' && lastCorner == '7' {
				edgeCount++
			}

			if loopPosition.locationChar == 'F' && lastCorner == 'J' {
				edgeCount++
			}

			lastCorner = loopPosition.locationChar

		}

	}

	return edgeCount%2 == 1

}

func makeMap(input string) (pipeMap map[coordinate]location, startCoordinates coordinate) {

	scanner := bufio.NewScanner(strings.NewReader(input))

	pipeMap = make(map[coordinate]location)
	y := 0

	lineAbove := ""
	startCoordinates = coordinate{-2, -2}
	startFound := false

	for scanner.Scan() {
		line := scanner.Text()

		if startFound {

			connections := pipeMap[startCoordinates].connected

			if lineAbove[startCoordinates.x] != '.' {

				connections = append(connections, coordinate{startCoordinates.x, y - 1})

			}

			if line[startCoordinates.x] != '.' {

				connections = append(connections, coordinate{startCoordinates.x, y})
			}

			pipeMap[startCoordinates] = location{startCoordinates, connections, 'S'}
			startFound = false
		}

		for x, entity := range line {

			if entity != '.' {
				switch entity {

				case '|':
					connections := []coordinate{}

					connections = append(connections, coordinate{x, y - 1})
					connections = append(connections, coordinate{x, y + 1})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}

				case '-':
					connections := []coordinate{}
					connections = append(connections, coordinate{x + 1, y})
					connections = append(connections, coordinate{x - 1, y})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}
				case 'L':
					connections := []coordinate{}
					connections = append(connections, coordinate{x, y - 1})
					connections = append(connections, coordinate{x + 1, y})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}
				case '7':
					connections := []coordinate{}
					connections = append(connections, coordinate{x, y + 1})
					connections = append(connections, coordinate{x - 1, y})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}

				case 'F':
					connections := []coordinate{}

					connections = append(connections, coordinate{x + 1, y})
					connections = append(connections, coordinate{x, y + 1})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}

				case 'J':
					connections := []coordinate{}
					connections = append(connections, coordinate{x, y - 1})
					connections = append(connections, coordinate{x - 1, y})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}

				case 'S':
					connections := []coordinate{}
					startCoordinates = coordinate{x, y}

					if x < len(line) && line[x+1] != '.' {

						connections = append(connections, coordinate{x + 1, y})

					}

					if x > 0 && line[x-1] != '.' {

						connections = append(connections, coordinate{x - 1, y})

					}

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}
					startFound = true

				}
			}
		}

		y++
		if !startFound {
			lineAbove = line
		}
	}

	return
}
