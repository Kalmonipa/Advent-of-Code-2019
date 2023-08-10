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
	fileName = "../short-input.txt"
)

type Coordinate struct {
	X, Y int
}

// Creates a slice of the first wires coordinates at each step
func mapFirstWire(instrLine string, firstWire []Coordinate, currPos Coordinate) []Coordinate {

	instr := strings.Split(instrLine, ",")

	//fmt.Println("Instructions are: ", instr)

	for _, i := range instr {
		//fmt.Println(i)

		dir := string(i[0])
		steps, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}

		currPos, firstWire = moveFirstWire(steps, dir, firstWire, currPos)

	}

	return firstWire
}

func moveFirstWire(steps int, dir string, firstWire []Coordinate, currPos Coordinate) (Coordinate, []Coordinate) {
	fmt.Printf("Moving %s by %d steps\n", dir, steps)

	dx, dy := 0, 0
	nextPos := Coordinate{X: currPos.X, Y: currPos.Y}

	if dir == "U" {
		dy = 1
	} else if dir == "D" {
		dy = -1
	} else if dir == "L" {
		dx = -1
	} else if dir == "R" {
		dx = 1
	}
	for count := 0; count < steps; count++ {
		nextPos = Coordinate{X: currPos.X + dx, Y: currPos.Y + dy}
		firstWire = append(firstWire, currPos)
		currPos = nextPos
	}

	fmt.Printf("New position is X: %d, Y: %d\n", nextPos.X, nextPos.Y)

	return currPos, firstWire
}

// Gets the coords of the second wire as it makes its movements
// Iterates through the first wires coords to see if there is a crossover point
// and appends that to a slice if it is. Ignores it if not a crossover
func mapSecondWire(instrLine string, currPos Coordinate, crossoverPoints []Coordinate, firstWire []Coordinate) []Coordinate {

	currPos = Coordinate{X: 0, Y: 0}

	instr := strings.Split(instrLine, ",")

	for _, i := range instr {

		dir := string(i[0])
		steps, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch dir {
		case "U":
			fmt.Printf("Moving up by %d steps\n", steps)
			for count := 0; count < steps; count++ {
				newPosY := currPos.Y + 1
				newCoord := Coordinate{X: currPos.X, Y: newPosY}
				if contains(firstWire, newCoord) {
					crossoverPoints = append(crossoverPoints, newCoord)
				}
				currPos = Coordinate{X: currPos.X, Y: newPosY}
			}
			fmt.Printf("New position is X: %d, Y: %d\n", currPos.X, currPos.Y)
		case "D":
			fmt.Printf("Moving down by %d steps\n", steps)
			for count := 0; count < steps; count++ {
				newPosY := currPos.Y - 1
				newCoord := Coordinate{X: currPos.X, Y: newPosY}
				if contains(firstWire, newCoord) {
					crossoverPoints = append(crossoverPoints, newCoord)
				}
				currPos = Coordinate{X: currPos.X, Y: newPosY}
			}
			fmt.Printf("New position is X: %d, Y: %d\n", currPos.X, currPos.Y)
		case "R":
			fmt.Printf("Moving right by %d steps\n", steps)
			for count := 0; count < steps; count++ {
				newPosX := currPos.X + 1
				newCoord := Coordinate{X: newPosX, Y: currPos.Y}
				if contains(firstWire, newCoord) {
					crossoverPoints = append(crossoverPoints, newCoord)
				}
				currPos = Coordinate{X: newPosX, Y: currPos.Y}
			}
			fmt.Printf("New position is X: %d, Y: %d\n", currPos.X, currPos.Y)
		case "L":
			fmt.Printf("Moving left by %d steps\n", steps)
			for count := 0; count < steps; count++ {
				newPosX := currPos.X - 1
				newCoord := Coordinate{X: newPosX, Y: currPos.Y}
				if contains(firstWire, newCoord) {
					crossoverPoints = append(crossoverPoints, newCoord)
				}
				currPos = Coordinate{X: newPosX, Y: currPos.Y}
			}
			fmt.Printf("New position is X: %d, Y: %d\n", currPos.X, currPos.Y)
		}
	}

	return crossoverPoints
}

func getWireInstructions(filename string) (string, string) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the file gets closed after this function finishes
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	var firstLine string
	var secondLine string

	// Loop through each line in the file
	for scanner.Scan() {
		if lineNumber == 1 {
			firstLine = scanner.Text()
			lineNumber += 1
		}
		if lineNumber == 2 {
			secondLine = scanner.Text()
		}
	}

	return firstLine, secondLine
}

func calculateClosestCrossover(crossoverPoints []Coordinate) int {
	closestDistance := 1000000

	for _, c := range crossoverPoints {
		absX := absInt(c.X)
		absY := absInt(c.Y)

		distance := absX + absY

		fmt.Printf("Distance is %d\n", distance)

		if closestDistance > distance {
			closestDistance = distance
		}
	}

	return closestDistance
}

func contains(s []Coordinate, elem Coordinate) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {

	firstWireMapping := []Coordinate{}
	crossoverPoints := []Coordinate{}
	startPos := Coordinate{X: 0, Y: 0}

	firstLine, secondLine := getWireInstructions(fileName)

	// fmt.Println(firstLine)
	// fmt.Println(secondLine)

	fmt.Println("Mapping first line")
	firstWire := mapFirstWire(firstLine, firstWireMapping, startPos)
	fmt.Println("---------------------------\nMapping second line")
	crossoverPoints = mapSecondWire(secondLine, startPos, crossoverPoints, firstWire)

	fmt.Print(crossoverPoints)

	closestDistance := calculateClosestCrossover(crossoverPoints)

	fmt.Print(closestDistance)

}
