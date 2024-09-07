package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	parkingLot  []string
	carParkSlot = make(map[string]int)
)

func main() {
	// open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// reading file line by line
	scanner := bufio.NewScanner(file)

	// loop through lines of file
	for scanner.Scan() {
		line := scanner.Text()
		readCommand(line)
	}

	// check errors in file
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}

func readCommand(str string) {
	input := strings.Split(str, " ")

	switch input[0] {
	// create
	case "create":
		if len(input) < 2 || input[1] == "" {
			fmt.Println("Please input a size!")
			return
		}

		create(input[1])

	// park
	case "park":
		if len(input) < 2 || input[1] == "" {
			fmt.Println("Please input a car number!")
			return
		}

		park(input[1])

	// leave
	case "leave":
		if len(input) < 3 || input[1] == "" || input[2] == "" {
			fmt.Println("Please input a car number and hours on how long the car has parked!")
			return
		}

		leave(input[1], input[2])
	// status
	case "status":
		status()

	// default
	default:
		fmt.Println("Command unrecognized")
	}
}

func create(input string) {
	size, _ := strconv.Atoi(input)

	parkingLot = make([]string, size)
	fmt.Println("Created parking lot with " + strconv.Itoa(size) + " slots")
}

func park(input string) {
	for i, v := range parkingLot {
		if v == "" {
			parkingLot[i] = input
			carParkSlot[input] = i

			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}

	fmt.Println("Sorry, parking lot is full")
}

func status() {
	fmt.Printf("%-6s %s\n", "Slot No.", "Registration No.")
	for i, v := range parkingLot {
		if v != "" {
			fmt.Printf("%-8d %s\n", i+1, v)
		}
	}
}

func leave(input1, input2 string) {
	slot, isExist := carParkSlot[input1]
	if !isExist {
		fmt.Printf("Registration Number %s not found\n", input1)
		return
	}

	parkingLot[slot] = ""
	delete(carParkSlot, input1)

	hours, _ := strconv.Atoi(input2)
	charge := 0
	if hours <= 2 {
		charge = 10
	} else {
		charge = (hours - 1) * 10
	}

	fmt.Printf("Registration Number %s from Slot %d has left with Charge %d\n", input1, slot+1, charge)
}
