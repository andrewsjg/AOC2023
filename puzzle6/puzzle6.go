package puzzle6

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

func Part1Solve(input string) {

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	races := getRaces(stringInput)
	totalWinners := 1

	for _, race := range races {
		raceWinners := 0

		for chargeTime := 1; chargeTime <= race.time; chargeTime++ {
			distanceTraveled := computeDistance(race, chargeTime)
			if distanceTraveled > race.distance {
				raceWinners++
			}
		}

		totalWinners = totalWinners * raceWinners
	}

	fmt.Printf("Winner Totals: %d\n", totalWinners)
}
func computeDistance(race race, chargeTime int) (distance int) {

	movementTime := race.time - chargeTime
	distance = movementTime * chargeTime

	return
}

func Part2Solve(input string) {

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	totalWinners := 1
	raceWinners := 0

	races := getRaces(stringInput)

	raceTimeString := ""
	distanceString := ""
	for race := range races {
		raceTimeString += strconv.Itoa(races[race].time)
		distanceString += strconv.Itoa(races[race].distance)
	}

	raceTime, _ := strconv.Atoi(raceTimeString)
	dist, _ := strconv.Atoi(distanceString)

	race := race{raceTime, dist}

	for chargeTime := 1; chargeTime <= race.time; chargeTime++ {
		distanceTraveled := computeDistance(race, chargeTime)
		if distanceTraveled > race.distance {

			raceWinners++
		}
	}

	totalWinners = totalWinners * raceWinners

	fmt.Printf("Winner Totals: %d\n", totalWinners)

}

// There MUST be a better way to do this. Regular expressions suck.
func getRaces(input string) (races []race) {

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "Time: ") {

			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllStringSubmatch(line, -1)

			for _, match := range matches {
				race := race{}
				race.time, _ = strconv.Atoi(match[0])

				races = append(races, race)
			}

		}

		if strings.Contains(line, "Distance: ") {

			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllStringSubmatch(line, -1)

			for idx, match := range matches {
				race := races[idx]
				race.distance, _ = strconv.Atoi(match[0])
				races[idx] = race
			}
		}

	}
	return
}
