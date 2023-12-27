package puzzle11

import (
	"bufio"
	"fmt"
	"math"
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

	start := 7
	end := 8

	startGal := galaxyLocations[start]
	endGal := galaxyLocations[end]

	fmt.Printf("TEST PATH - Start Gal X: %d, Y: %d\n", startGal.x, startGal.y)
	fmt.Printf("TEST PATH - End Gal X: %d, Y: %d\n", endGal.x, endGal.y)

	distance := math.Abs(float64(endGal.x)-float64(startGal.x)) + math.Abs((float64(endGal.y) - float64(startGal.y)))
	fmt.Printf("TEST PATH - Distance between galaxy %d and galaxy %d: %d\n", start, end, int(distance))

	fmt.Println()
	for galNo, gal := range galaxyLocations {
		fmt.Printf("Galaxy: %d has X: %d, Y: %d\n", galNo, gal.x, gal.y)
	}

	runningTotal := 0
	total := totalDistances(galaxyLocations, &runningTotal)

	// Using sample input and 100 times expansion this result should be 8410 but I get 8538
	fmt.Printf("Total distances between all galaxies in the massively expanded universe: %d\n", total)
	// 678627324165 - TO HIGH

}

func totalDistances(galaxyLocations []*Location, runningTotal *int) (distanceTotal int) {
	if len(galaxyLocations) > 1 {

		dist := 0.0

		startGal := galaxyLocations[0]

		for i := 1; i < len(galaxyLocations); i++ {

			endGal := galaxyLocations[i]

			dist = dist + math.Abs(float64(endGal.x)-float64(startGal.x)) + math.Abs((float64(endGal.y) - float64(startGal.y)))

		}

		*runningTotal = int(dist) + totalDistances(galaxyLocations[1:], runningTotal)
	}

	return *runningTotal
}

func totalDistancesAStar(galaxyLocations []*Location, runningTotal *int) (distanceTotal int) {
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

	for scanner.Scan() {
		line := scanner.Text()

		for _, col := range cols {

			if part2 {

				line = line[:col] + "&" + line[col:]

			} else {
				line = line[:col] + "." + line[col:]
			}
		}

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
			// 1000000
			//fmt.Println(xMultiplier, yMultiplier)

			// ..&.#.&...&.
			// 0,1,1-100,102,103(gal),104,106+100,207, ...
			switch chr {
			case '.':
				Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*100), y+(yMultiplier*100))

			case '#':
				// a Galaxy has the same movement cost as normal space

				galaxyLocation := &Location{}

				if xMultiplier == 0 && yMultiplier == 0 {
					Map.setLocation(&Location{topoType: '.'}, x, y)
					galaxyLocation = Map.getLocation(x, y)

				} else if xMultiplier != 0 && yMultiplier != 0 {
					Map.setLocation(&Location{topoType: '.'}, x-1+(xMultiplier*100+x), y-1+(yMultiplier*100))
					galaxyLocation = Map.getLocation(x-1+(xMultiplier*100+x), y-1+(yMultiplier*100))
				} else if yMultiplier != 0 && xMultiplier == 0 {
					Map.setLocation(&Location{topoType: '.'}, x, y-1+(yMultiplier*100))
					galaxyLocation = Map.getLocation(x, y-1+(yMultiplier*100))
				} else if xMultiplier != 0 && yMultiplier == 0 {
					Map.setLocation(&Location{topoType: '.'}, x, y-1+(yMultiplier*100))
					galaxyLocation = Map.getLocation(x-1+(xMultiplier*100), y)
				}
				/*
					if xMultiplier == 0 && yMultiplier == 0 {
						Map.setLocation(&Location{topoType: '.'}, x, y)
						galaxyLocation = Map.getLocation(x, y)

					} else if xMultiplier != 0 && yMultiplier != 0 {
						Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*100), y+(yMultiplier*100))
						galaxyLocation = Map.getLocation(x+(xMultiplier*100), y+(yMultiplier*100))
					} else if yMultiplier != 0 && xMultiplier == 0 {
						Map.setLocation(&Location{topoType: '.'}, x, y+(yMultiplier*100))
						galaxyLocation = Map.getLocation(x, y+(yMultiplier*100))
					} else if xMultiplier != 0 && yMultiplier == 0 {
						Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*100), y)
						galaxyLocation = Map.getLocation(x+(xMultiplier*100), y)
					} */

				// for part2 we only really care about the x,y values for the galaxies
				// if we arent using pathfinding then we dont care about the neighbour x,y values
				//galaxyLocation := Map.getLocation(x+(xMultiplier*100), y+(yMultiplier*100))
				galaxyLocations = append(galaxyLocations, galaxyLocation)

			case '&':
				// the '&' character represents a million time increase
				// in the expansion along the x axis.

				xMultiplier++
				//Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*100), y+(yMultiplier*100))

			case '%':
				// the '%' character represents a million time increase
				// in the expansion along the y axis.

				yMultiplier++
				//Map.setLocation(&Location{topoType: '.'}, x+(xMultiplier*100), y+(yMultiplier*100))

			default:
				Map.setLocation(&Location{topoType: int(chr)}, x+(xMultiplier*100), y+(yMultiplier*100))

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
