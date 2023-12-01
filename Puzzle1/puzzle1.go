package puzzle1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func P1Solve(input string) {

	calibrationTotal := 0

	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := 0
		lastDigit := 0

		for _, r := range line {

			firstDigit, err = getDigit(r)
			if err == nil {
				break
			}

		}

		for i := len(line) - 1; i >= 0; i-- {
			lastDigit, err = getDigit(rune(line[i]))

			if err == nil {
				break
			}

		}

		// This is a dirty way to make a 2 digit number
		calVal, err := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))

		if err != nil {
			fmt.Println("couldnt create a value: " + err.Error())
			os.Exit(1)
		}

		calibrationTotal = calibrationTotal + calVal

	}

	fmt.Println(calibrationTotal)
}

// Check if the passed rune is a digit then return the digit as an int
func getDigit(r rune) (int, error) {

	var err error
	digit := -9

	if unicode.IsDigit(r) {
		digit, err = strconv.Atoi(string(r))
	} else {
		err = errors.New("not a digit")
	}

	return digit, err
}
