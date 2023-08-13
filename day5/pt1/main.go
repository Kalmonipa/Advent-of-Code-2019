package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	filename   = "../input.txt"
	inputValue = 1
)

func getOpCodes(filename string) ([]int, error) {
	var ops []int

	// Read the file content
	data, err := os.ReadFile(filename)
	if err != nil {
		//fmt.Println("Error reading file:", err)
		return ops, err
	}

	// Split the content by commas
	parts := strings.Split(string(data), ",")

	for _, part := range parts {
		// Convert each part to an integer
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			fmt.Printf("Error converting '%s' to integer: %s\n", part, err)
			return ops, err
		}
		ops = append(ops, num)
	}

	return ops, nil
}

// Opcode 1 adds together numbers read from two positions and stores the result in a third position.
// The three integers immediately after the opcode tell you these three positions -
// the first two indicate the positions from which you should read the input values, and
// the third indicates the position at which the output should be stored.

// For example, if your Intcode computer encounters 1,10,20,30, it should read the values at positions
// 10 and 20, add those values, and then overwrite the value at position 30 with their sum.

func OpCodeOne(ops []int, position int) []int {
	//fmt.Println("Ops code one")
	posOne := ops[position+1]
	posTwo := ops[position+2]
	destination := ops[position+3]

	ops[destination] = ops[posOne] + ops[posTwo]

	return ops
}

// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them.
// Again, the three integers after the opcode indicate where the inputs and outputs are, not their values.
func OpCodeTwo(ops []int, position int) []int {
	//fmt.Println("Ops code two")
	posOne := ops[position+1]
	posTwo := ops[position+2]
	destination := ops[position+3]

	ops[destination] = ops[posOne] * ops[posTwo]

	return ops
}

// Takes a single integer as input and saves it to the position given by
// its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
func OpsCodeThree(ops []int, position int) []int {
	target := ops[position+1]

	ops[target] = inputValue
	return ops
}

// outputs the value of its only parameter. For example,
// the instruction 4,50 would output the value at address 50.
func OpsCodeFour(ops []int, position int) int {
	return ops[position]
}

func main() {

	position := 0

	ops, err := getOpCodes(filename)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(ops)

	for {
		if ops[position] == 1 {
			ops = OpCodeOne(ops, position)
			position += 4
		} else if ops[position] == 2 {
			ops = OpCodeTwo(ops, position)
			position += 4
		} else if ops[position] == 3 {
			ops = OpsCodeThree(ops, position)
			position += 2
		} else if ops[position] == 4 {
			fmt.Println(ops[position+1])
		} else if ops[position] == 99 {
			fmt.Println(ops)
			return
		}
	}

}
