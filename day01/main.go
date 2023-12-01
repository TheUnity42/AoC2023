package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("Hello World! Starting to read the input file...")

	Day1Part1("input.txt")

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
		fmt.Printf("Read line from input file: %v\n", line)

		var digits []rune

		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}

		if len(digits) < 1 {
			log.Fatalf("No digits were found in line! Line was: %v", line)
		}

		value := string(digits[0]) + string(digits[len(digits)-1])

		fmt.Printf("len: %v, val: %v\n", len(digits), value)

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
	fmt.Printf("Sum of calibration values is: %v\n", sum)
}
