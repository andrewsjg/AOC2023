package puzzle10

import (
	"bufio"
	"fmt"
	"os"
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

func followMap(startCoordinates coordinate, pipeMap map[coordinate]location) (stepCount int, loopMap map[coordinate]location) {

	startLocation := pipeMap[startCoordinates]
	currentLocation := startLocation.connected[0]
	previousLocation := startLocation

	//stepCount := 0

	for currentLocation != startCoordinates {

		tmpLocation := currentLocation

		if pipeMap[currentLocation].connected[0] != previousLocation.gridLocation {

			currentLocation = pipeMap[currentLocation].connected[0]
		} else {
			currentLocation = pipeMap[currentLocation].connected[1]
		}

		previousLocation = pipeMap[tmpLocation]

		stepCount++
	}

	return stepCount, loopMap
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
					connections = append(connections, coordinate{x, y + 1})
					connections = append(connections, coordinate{x + 1, y})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}

				case 'J':
					connections := []coordinate{}
					connections = append(connections, coordinate{x, y - 1})
					connections = append(connections, coordinate{x - 1, y})

					pipeMap[coordinate{x, y}] = location{coordinate{x, y}, connections, entity}

				case 'S':
					connections := []coordinate{}
					startCoordinates = coordinate{x, y}
					//fmt.Println(pipeMap[coordinate{x, y}])

					if x > 0 && line[x-1] != '.' {

						connections = append(connections, coordinate{x - 1, y})

					}

					if x < len(line) && line[x+1] != '.' {

						connections = append(connections, coordinate{x + 1, y})

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
