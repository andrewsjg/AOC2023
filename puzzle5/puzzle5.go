package puzzle5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type rangeMap struct {
	destMin int
	destMax int
	srcMin  int
	srcMax  int
}

func Part2Solve(input string) {

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	seeds := getSeeds(stringInput)

	seedcount := 0

	for idx := 0; idx < len(seeds); idx += 2 {

		for seed := seeds[idx]; seed <= (seeds[idx] + seeds[idx+1]); seed++ {
			seedcount++

		}
	}

	allMaps := map[string][]rangeMap{}

	tokens := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	for _, token := range tokens {
		allMaps[token] = getMap(stringInput, token)
	}

	lowestLocation := -1

	locChannel := make(chan int, seedcount)

	var wg sync.WaitGroup
	for idx := 0; idx < len(seeds); idx += 2 {

		for seed := seeds[idx]; seed <= (seeds[idx] + seeds[idx+1]); seed++ {
			wg.Add(1)
			go getLocation(seed, allMaps, locChannel, &wg)
		}

	}

	wg.Wait()
	close(locChannel)

	for location := range locChannel {
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}

	}

	fmt.Printf("Lowest Location: %d\n", lowestLocation)
}

func getLocation(seed int, allMaps map[string][]rangeMap, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	tokens := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	myResult := seed

	for _, token := range tokens {
		myRange := allMaps[token]
		myVal := myResult

		for _, rnge := range myRange {

			if myVal >= rnge.srcMin && myVal <= rnge.srcMax {
				myResult = rnge.destMax - (rnge.srcMax - myVal)

			}

		}

	}

	ch <- myResult
}

func Part1Solve(input string) {

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	seeds := getSeeds(stringInput)

	seedToSoil := getMap(input, "seed-to-soil")
	soilToFertilizer := getMap(input, "soil-to-fertilizer")
	fertilizerToWater := getMap(input, "fertilizer-to-water")
	waterToLight := getMap(input, "water-to-light")
	lightToTemparature := getMap(input, "light-to-temperature")
	temparatureToHumidity := getMap(input, "temperature-to-humidity")
	humidityToLocation := getMap(input, "humidity-to-location")

	soil := -1
	fertilizer := -1
	water := -1
	light := -1
	temparature := -1
	humidity := -1
	location := -1
	lowestLocation := -1

	for _, seed := range seeds {

		soil = getMapVal(seedToSoil, seed)

		fertilizer = getMapVal(soilToFertilizer, soil)
		water = getMapVal(fertilizerToWater, fertilizer)
		light = getMapVal(waterToLight, water)
		temparature = getMapVal(lightToTemparature, light)
		humidity = getMapVal(temparatureToHumidity, temparature)
		location = getMapVal(humidityToLocation, humidity)

		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}

	}

	fmt.Printf("Lowest Location: %d\n", lowestLocation)
}

func getMapVal(rangeMap []rangeMap, myVal int) (myResult int) {

	myResult = myVal
	for _, rnge := range rangeMap {
		if myVal >= rnge.srcMin && myVal <= rnge.srcMax {
			myResult = rnge.destMax - (rnge.srcMax - myVal)
		}
	}

	return myResult
}

func getSeeds(input string) (seeds []int) {

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		if strings.Contains(line, "seeds: ") {

			tmpArry := strings.Split(line[strings.Index(line, "seeds: ")+7:], " ")

			for _, item := range tmpArry {
				seed, err := strconv.Atoi(item)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				seeds = append(seeds, seed)
			}
		}
	}

	return seeds
}

/*
Tokens:

	seed-to-soil
	soil-to-fertilizer
	fertilizer-to-water
	water-to-light
	light-to-temperature
	temperature-to-humidity
	humidity-to-location
*/

func getMap(input string, token string) []rangeMap {
	retMap := []rangeMap{}

	tokenFound := false

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		if strings.Contains(line, strings.ToLower(token)) {
			tokenFound = true
			continue
		}

		if line == "" && tokenFound {
			tokenFound = false
		}

		if tokenFound && line != "" {
			tmpLine := line

			destinationRangeStart, err := strconv.Atoi(line[:strings.Index(tmpLine, " ")])

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tmpLine = tmpLine[strings.Index(tmpLine, " ")+1:]
			sourceRangeStart, err := strconv.Atoi(tmpLine[:strings.Index(tmpLine, " ")])

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tmpLine = tmpLine[strings.Index(tmpLine, " ")+1:]
			rangeLength, err := strconv.Atoi(tmpLine)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			rangeMap := rangeMap{
				destinationRangeStart,
				destinationRangeStart + rangeLength - 1,
				sourceRangeStart,
				sourceRangeStart + rangeLength - 1,
			}

			retMap = append(retMap, rangeMap)

		}
	}

	return retMap
}
