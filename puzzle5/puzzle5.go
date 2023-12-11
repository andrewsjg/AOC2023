package puzzle5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rangeMap struct {
	destMin int
	destMax int
	srcMin  int
	srcMax  int
}

func Part1Solve(input string) {
	seeds := getSeeds(input)

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

	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
			//fmt.Println(tmpLine)
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
				destinationRangeStart + rangeLength,
				sourceRangeStart,
				sourceRangeStart + rangeLength,
			}

			retMap = append(retMap, rangeMap)

		}
	}

	return retMap
}
