package puzzle11

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
)

type Location struct {
	x, y     int
	topoType int

	Map galacticMap
}

type galacticMap map[int]map[int]*Location

func Part1Solve(input string) {

	galaxyMap := expandGalaxyMap(input)

	_, galaxyLocations := readMap(galaxyMap)

	runningTotal := 0
	total := totalDistances(galaxyLocations, &runningTotal)

	fmt.Printf("Total distances between all galaxies: %d\n", total)

}

func totalDistances(galaxyLocations []*Location, runningTotal *int) (distanceTotal int) {
	// Using go routines naively speeds this up, but not enough.
	// There will be way better way to speed this up, but its not slow enough to worry about for part 1

	if len(galaxyLocations) > 1 {

		dist := 0

		startGalaxy := galaxyLocations[0]
		resultsChan := make(chan int, len(galaxyLocations))
		var wg sync.WaitGroup

		for i := 1; i < len(galaxyLocations); i++ {
			wg.Add(1)
			endGalaxy := galaxyLocations[i]
			go distanceCalc(startGalaxy, endGalaxy, &wg, resultsChan)

		}

		wg.Wait()
		close(resultsChan)

		for resultDist := range resultsChan {
			dist = resultDist + dist
		}

		*runningTotal = dist + totalDistances(galaxyLocations[1:], runningTotal)
	}

	return *runningTotal
}

func distanceCalc(start *Location, end *Location, wg *sync.WaitGroup, resultsChan chan int) {
	defer wg.Done()

	foundPath, _, found := path(start, end)

	if found {
		resultsChan <- len(foundPath) - 1
	}

}

func expandGalaxyMap(input string) string {
	expandedMap := ""

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	scanner := bufio.NewScanner(strings.NewReader(stringInput))

	rows := []int{}
	cols := []int{}

	rowCount := 0
	nonEmptyCols := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "#") {
			rows = append(rows, rowCount)
		}

		for col, char := range line {

			if char == '.' {

				if !slices.Contains(cols, col) && !slices.Contains(nonEmptyCols, col) {
					cols = append(cols, col)
				}

			}

			if char == '#' {

				if !slices.Contains(nonEmptyCols, col) {
					nonEmptyCols = append(nonEmptyCols, col)
				}

				if slices.Contains(cols, col) {
					cols = removeItem(cols, col)

				}
			}
		}

		rowCount++
	}

	scanner = bufio.NewScanner(strings.NewReader(stringInput))

	for idx := range cols {

		cols[idx] = cols[idx] + idx
	}

	rowCount = 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, col := range cols {

			line = line[:col] + "." + line[col:]
		}

		if slices.Contains(rows, rowCount) {
			line = line + "\n" + line
		}

		rowCount++

		expandedMap += line + "\n"
	}
	return expandedMap
}

func readMap(input string) (world galacticMap, galaxyLocations []*Location) {

	Map := galacticMap{}

	mapReader := bufio.NewScanner(strings.NewReader(input))

	y := 0
	for mapReader.Scan() {
		line := mapReader.Text()

		mapLine := string(line)

		for x, chr := range mapLine {
			switch chr {
			case '.':
				Map.setLocation(&Location{topoType: '.'}, x, y)

			case '#':
				//fmt.Printf("Found Galaxy at: %d,%d\n", x, y)
				//Map.setLocation(&Location{topoType: '#'}, x, y)
				Map.setLocation(&Location{topoType: '.'}, x, y)
				galaxyLocation := Map.getLocation(x, y)
				galaxyLocations = append(galaxyLocations, galaxyLocation)

			default:
				Map.setLocation(&Location{topoType: int(chr)}, x, y)

			}
		}

		y++
	}

	return Map, galaxyLocations
}

func (t galacticMap) setLocation(l *Location, x, y int) {
	if t[x] == nil {
		t[x] = map[int]*Location{}
	}

	l.x = x
	l.y = y

	t[x][y] = l
	l.Map = t
}

func (l Location) pathToNeighBourCost(to Pather) float64 {
	toLocation := to.(*Location)

	return float64(toLocation.topoType)
}

func (l Location) pathCostEstimate(to Pather) float64 {

	toLocation := to.(*Location)

	deltaX := toLocation.x - l.x
	deltaY := toLocation.y - l.y

	// Make absolutes
	if deltaX < 0 {
		deltaX = -deltaX
	}

	if deltaY < 0 {
		deltaY = -deltaY
	}

	if deltaX > deltaY {

		return float64(14*deltaY + 10*(deltaX-deltaY))
	}

	return float64(14*deltaX + 10*(deltaY-deltaX))
}

func (l Location) neighbours() []Pather {
	neighbours := []Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {

		if neighbour := l.Map.getLocation(l.x+offset[0], l.y+offset[1]); neighbour != nil {

			neighbours = append(neighbours, neighbour)
			/*
				diff := neighbour.topoType - l.topoType


					// Exclude neighbours that are to high
					if diff < 2 {

						neighbours = append(neighbours, neighbour)
					} */
		}
	}
	return neighbours
}

/*
func (l Location) neighbours() []Pather {
	neighbours := []Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {

		if neighbour := l.Map.getLocation(l.x+offset[0], l.y+offset[1]); neighbour != nil {

			diff := neighbour.topoType - l.topoType

			// Exclude neighbours that are to high
			if diff < 2 {
				neighbours = append(neighbours, neighbour)
			}
		}
	}
	return neighbours
} */

func (t galacticMap) getLocation(x, y int) *Location {
	if t[x] == nil {
		return nil
	}
	return t[x][y]
}

func removeItem[I comparable](input []I, item I) []I {
	output := []I{}

	for _, inputItem := range input {
		if item != inputItem {
			output = append(output, inputItem)
		}
	}

	return output
}
