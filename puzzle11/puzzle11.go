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

	galaxyMap := expandGalaxyMap(input, false)

	_, galaxyLocations := readMap(galaxyMap)

	runningTotal := 0
	total := totalDistances(galaxyLocations, &runningTotal)

	fmt.Printf("Total distances between all galaxies: %d\n", total)

}

func Part2Solve(input string) {
	galaxyMap := expandGalaxyMap(input, true)
	_, galaxyLocations := readMap(galaxyMap)

	fmt.Println(galaxyMap)

	// Test Path

	start := 0
	end := 6

	startGal := galaxyLocations[start]
	endGal := galaxyLocations[end]

	fmt.Printf("TEST PATH Start Gal X: %d, Y: %d\n", startGal.x, startGal.y)
	fmt.Printf("TEST PATH End Gal X: %d, Y: %d\n", endGal.x, endGal.y)
	fmt.Println()

	for galNo, gal := range galaxyLocations {
		fmt.Printf("Galaxy: %d has X: %d, Y: %d\n", galNo, gal.x, gal.y)
	}

	foundPath, _, found := path(startGal, endGal)
	if found {

		fmt.Printf("TEST PATH - Distance between galaxy %d and galaxy %d: %d\n", start, end, len(foundPath)-1)
		// 0 to 6 = 5000021 vs 15 for Part 1
	}

	/*
		runningTotal := 0
		total := totalDistances(galaxyLocations, &runningTotal)

		fmt.Printf("Total distances between all galaxies in the massively expanded universe: %d\n", total)
	*/

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

func expandGalaxyMap(input string, part2 bool) string {
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

	//colsMarked := false

	for scanner.Scan() {
		line := scanner.Text()

		for _, col := range cols {

			if part2 {

				line = line[:col] + "&" + line[col:]
				/*
					if !colsMarked {
						line = line[:col] + "&" + line[col:]
					} else {
						line = line[:col] + "." + line[col:]
					} */

			} else {
				line = line[:col] + "." + line[col:]
			}
		}
		//colsMarked = true
		if slices.Contains(rows, rowCount) {
			if part2 {
				line = "%" + line[1:]

			} else {
				line = line + "\n" + line
			}
		}

		rowCount++

		expandedMap += line + "\n"
	}
	return expandedMap
}

func readMap(input string) (world galacticMap, galaxyLocations []*Location) {

	Map := galacticMap{}

	mapReader := bufio.NewScanner(strings.NewReader(input))

	xMultiplier := 0
	yMultiplier := 0

	y := 0

	for mapReader.Scan() {
		line := mapReader.Text()

		mapLine := string(line)

		for x, chr := range mapLine {

			//fmt.Println(xMultiplier, yMultiplier)
			switch chr {
			case '.':
				Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*1000000), y+(yMultiplier*1000000))

			case '#':
				// a Galaxy has the same movement cost as normal space
				Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*1000000), y+(yMultiplier*1000000))
				galaxyLocation := Map.getLocation(x+(xMultiplier*1000000), y+(yMultiplier*1000000))
				galaxyLocations = append(galaxyLocations, galaxyLocation)

			case '&':
				// the '&' character represents a million time increase
				// in the expansion along the x axis.

				for i := x + (xMultiplier * 1000000); i <= x+(xMultiplier*1000000)+1000000; i++ {
					Map.setLocation(&Location{topoType: '.'}, i, y+(yMultiplier*1000000))
				}

				xMultiplier++

			case '%':
				// the '%' character represents a million time increase
				// in the expansion along the y axis.

				for i := y + (yMultiplier * 1000000); i <= y+(yMultiplier*1000000)+1000000; i++ {
					Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*1000000), i)
				}

				yMultiplier++

			default:
				Map.setLocation(&Location{topoType: int(chr)}, x+(xMultiplier*1000000), y+(yMultiplier*1000000))

			}
		}
		xMultiplier = 0
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
