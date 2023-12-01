package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Hello World! Starting to read the input file...")

	Day1Part1("input.txt")

	Day1Part2("input.txt")

	fmt.Println("\nComplete.")
}

func Day1Part1(filename string) {
	// read the input file
	file, err := os.Open(filename)

	// handle potential error with a logger
	if err != nil {
		log.Fatal(err)
	}

	// make sure to clean up resources, defer to after function returns
	defer file.Close()

	// create a scanner to iterate over the lines in the problem
	scanner := bufio.NewScanner(file)

	// sum
	sum := 0

	// iterate over all the lines
	for scanner.Scan() {
		line := scanner.Text()

		value := GetCalibrationCode(line, false)

		val, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		sum += val
	}

	// check for scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// log output
	fmt.Printf("Sum of calibration values for part 1 is: %v\n", sum)
}

func Day1Part2(filename string) {
	// read the input file
	file, err := os.Open(filename)

	// handle potential error with a logger
	if err != nil {
		log.Fatal(err)
	}

	// make sure to clean up resources, defer to after function returns
	defer file.Close()

	// create a scanner to iterate over the lines in the problem
	scanner := bufio.NewScanner(file)

	// sum
	sum := 0

	// iterate over all the lines
	for scanner.Scan() {
		line := scanner.Text()

		value := GetCalibrationCode(line, true)

		val, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		sum += val
	}

	// check for scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// log output
	fmt.Printf("Sum of calibration values for Part 2 is: %v\n", sum)
}

func GetCalibrationCode(line string, words bool) string {
	// we need to extract any numbers AND any spelled out numbers
	var digits []string

	// make the line lower case only
	lower := strings.ToLower(line)

	// iterate over the line, looking for digits
	for loc, char := range lower {
		if unicode.IsDigit(char) {
			digits = append(digits, string(char))
		}

		if words {
			// look for a digit text at this location
			digit := GetDigitByName(lower[loc:])
			if digit != "" {
				// append the digit if we found one
				digits = append(digits, digit)
			}
		}
	}

	// log.Printf("Found digits %v in line %v\n", digits, line)

	// concat and return
	return digits[0] + digits[len(digits)-1]
}

func GetDigitByName(input string) string {
	// could use regex here, but strcmp will suffice
	switch true {
	case strings.HasPrefix(input, "zero"):
		return "0"
	case strings.HasPrefix(input, "one"):
		return "1"
	case strings.HasPrefix(input, "two"):
		return "2"
	case strings.HasPrefix(input, "three"):
		return "3"
	case strings.HasPrefix(input, "four"):
		return "4"
	case strings.HasPrefix(input, "five"):
		return "5"
	case strings.HasPrefix(input, "six"):
		return "6"
	case strings.HasPrefix(input, "seven"):
		return "7"
	case strings.HasPrefix(input, "eight"):
		return "8"
	case strings.HasPrefix(input, "nine"):
		return "9"
	default:
		return ""
	}
}
