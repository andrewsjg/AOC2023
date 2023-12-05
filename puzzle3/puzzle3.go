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

		fmt.Println(line)

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

	// Find valid part numbers

	partSum := 0
	for y, partsList := range parts {

		for _, part := range partsList {
			isValid := false
			//fmt.Printf("PartNO: %s, PartX: %d\n", part.partNo, part.x)

			partVal, err := strconv.Atoi(part.partNo)

			if err != nil {
				fmt.Println("couldnt convert to number: " + err.Error())
				os.Exit(1)
			}

			// Is there a marker in the same row?
			if !isValid {
				for _, marker := range markers[y] {
					if marker.x == part.x-1 || marker.x == part.x+len(part.partNo) {
						fmt.Printf("Inline. Part: %s is a valid part. With Marker: %s\n", part.partNo, marker.marker)

						isValid = true
						partSum = partSum + partVal

					}
				}
			}

			// Is there a marker above
			if !isValid {
				for _, marker := range markers[y-1] {

					if marker.x >= part.x-1 && marker.x <= part.x+len(part.partNo) {

						fmt.Printf("Above. Part: %s is a valid part. With Marker: %s\n", part.partNo, marker.marker)
						isValid = true
						partSum = partSum + partVal

					}
				}
			}

			// Is there a marker below
			if !isValid {
				for _, marker := range markers[y+1] {

					if marker.x >= part.x-1 && marker.x <= part.x+len(part.partNo) {

						fmt.Printf("Below. Part: %s is a valid part. With Marker: %s\n", part.partNo, marker.marker)
						isValid = true
						partSum = partSum + partVal
					}
				}
			}

		}

	}
	// 484101 to low
	// 537049 to high
	// 518219 to low
	fmt.Println(markers)
	fmt.Printf("The sum of all valid parts is: %d\n", partSum)

}
