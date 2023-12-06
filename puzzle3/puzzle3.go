package puzzle3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type part struct {
	partNo string
	x      int
}

type marker struct {
	marker string
	x      int
}

type partMap map[int][]part

type markerMap map[int][]marker

func Part1Solve(input string) {

	parts, markers := getMaps(input)

	// Find valid part numbers

	partSum := 0
	for y, partsList := range parts {

		for _, part := range partsList {
			isValid := false

			partVal, err := strconv.Atoi(part.partNo)

			if err != nil {
				fmt.Println("couldnt convert to number: " + err.Error())
				os.Exit(1)
			}

			// Is there a marker in the same row?
			if !isValid {
				for _, marker := range markers[y] {
					if marker.x == part.x-1 || marker.x == part.x+len(part.partNo) {
						isValid = true
						partSum = partSum + partVal

					}
				}
			}

			// Is there a marker above
			if !isValid {
				for _, marker := range markers[y-1] {

					if marker.x >= part.x-1 && marker.x <= part.x+len(part.partNo) {
						isValid = true
						partSum = partSum + partVal

					}
				}
			}

			// Is there a marker below
			if !isValid {
				for _, marker := range markers[y+1] {

					if marker.x >= part.x-1 && marker.x <= part.x+len(part.partNo) {
						isValid = true
						partSum = partSum + partVal
					}
				}
			}

		}

	}

	fmt.Printf("The sum of all valid parts is: %d\n", partSum)

}

// This is so ugly. I can't believe this is what ended up with...
func Part2Solve(input string) {
	parts, markers := getMaps(input)
	gearTot := 0
	tmpTot := 1

	for y, markerlist := range markers {

		for _, marker := range markerlist {
			if marker.marker == "*" {

				found1 := false

				// Get the relevant parts lines
				partsAbove := parts[y-1]
				partsBelow := parts[y+1]
				partsAdjacent := parts[y]

				// Make new maps out of the parts list. Could probably have done it this way first!

				aboveMap := map[int]part{}

				for _, part := range partsAbove {
					for i := 0; i < len(part.partNo); i++ {
						aboveMap[part.x+i] = part
					}
				}

				belowMap := map[int]part{}

				for _, part := range partsBelow {
					for i := 0; i < len(part.partNo); i++ {
						belowMap[part.x+i] = part
					}

				}

				adjacentMap := map[int]part{}

				for _, part := range partsAdjacent {
					for i := 0; i < len(part.partNo); i++ {
						adjacentMap[part.x+i] = part
					}
				}

				if adjacentMap[marker.x-1].partNo != "" {

					partVal, err := strconv.Atoi(adjacentMap[marker.x-1].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					if found1 {
						gearTot = gearTot + (tmpTot * partVal)
						tmpTot = 1
						found1 = false

						continue
					} else {
						tmpTot = tmpTot * partVal
						found1 = true
					}

				}

				if adjacentMap[marker.x+1].partNo != "" {

					partVal, err := strconv.Atoi(adjacentMap[marker.x+1].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					if found1 {
						gearTot = gearTot + (tmpTot * partVal)
						tmpTot = 1
						found1 = false

						continue

					} else {
						tmpTot = tmpTot * partVal
						found1 = true
					}

				}

				if belowMap[marker.x].partNo != "" {

					partVal, err := strconv.Atoi(belowMap[marker.x].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					found1 = true
					tmpTot = tmpTot * partVal

				}

				if belowMap[marker.x].partNo == "" && belowMap[marker.x-1].partNo != "" {

					partVal, err := strconv.Atoi(belowMap[marker.x-1].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					if found1 {
						gearTot = gearTot + (tmpTot * partVal)
						tmpTot = 1
						found1 = false

						continue
					} else {
						tmpTot = tmpTot * partVal
						found1 = true
					}
				}

				if belowMap[marker.x].partNo == "" && belowMap[marker.x+1].partNo != "" {

					partVal, err := strconv.Atoi(belowMap[marker.x+1].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					if found1 {
						gearTot = gearTot + (tmpTot * partVal)
						tmpTot = 1
						found1 = false

						continue
					} else {
						tmpTot = tmpTot * partVal
						found1 = true
					}

				}

				if aboveMap[marker.x].partNo != "" {

					partVal, err := strconv.Atoi(aboveMap[marker.x].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					if found1 {
						gearTot = gearTot + (tmpTot * partVal)
						tmpTot = 1
						found1 = false

						continue
					} else {
						tmpTot = tmpTot * partVal
						found1 = true
					}

				}

				if aboveMap[marker.x].partNo == "" && aboveMap[marker.x-1].partNo != "" {

					partVal, err := strconv.Atoi(aboveMap[marker.x-1].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					if found1 {
						gearTot = gearTot + (tmpTot * partVal)
						tmpTot = 1
						found1 = false

						continue
					} else {
						tmpTot = tmpTot * partVal
						found1 = true
					}

				}

				if aboveMap[marker.x].partNo == "" && aboveMap[marker.x+1].partNo != "" {

					partVal, err := strconv.Atoi(aboveMap[marker.x+1].partNo)

					if err != nil {
						fmt.Println("couldnt convert to number: " + err.Error())
						os.Exit(1)
					}

					if found1 {
						gearTot = gearTot + (tmpTot * partVal)
						tmpTot = 1
						found1 = false
					}

				}
				tmpTot = 1
			}

		}
	}

	fmt.Printf("The Gear total is: %d\n", gearTot)
}

func getMaps(input string) (partMap, markerMap) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	parts := partMap{}
	markers := markerMap{}

	for scanner.Scan() {

		line := scanner.Text()

		sPartNum := ""
		xParts := []part{}
		foundPart := false
		partX := 0
		xMarkers := []marker{}

		// Get all the parts and markers into maps
		for x, r := range line {
			if unicode.IsDigit(r) {

				if !foundPart {
					partX = x
				}

				sPartNum = strings.TrimSpace(sPartNum + string(r))
				foundPart = true

				// need to cater for the case where the part is at the end of the line
				if x == len(line)-1 {
					part := part{partNo: sPartNum, x: partX}
					xParts = append(xParts, part)
				}

			} else {
				if foundPart {

					part := part{partNo: sPartNum, x: partX}
					xParts = append(xParts, part)
					sPartNum = ""
					foundPart = false

				}

				if string(r) != "." {

					xMarkers = append(xMarkers, marker{marker: string(r), x: x})
				}
			}

		}

		parts[y] = xParts
		markers[y] = xMarkers

		y++
	}

	return parts, markers
}
