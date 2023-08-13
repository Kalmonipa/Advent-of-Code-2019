package main

import "testing"

func indexOf(element int, data []int) int {
	for index, value := range data {
		if value == element {
			return index
		}
	}
	return -1 // Return -1 if the element is not found
}

func TestOpCodeOne(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	position := indexOf(1, input)

	expectedRes := []int{2, 0, 0, 0, 99}

	res := OpCodeOne(input, position)

	for ind, _ := range expectedRes {
		if res[ind] != expectedRes[ind] {
			t.Fatalf("Result %v is not expected", res)
		}
	}
}

func TestOpCodeTwo(t *testing.T) {
	input := []int{2, 3, 0, 3, 99}
	position := indexOf(2, input)

	expectedRes := []int{2, 3, 0, 6, 99}

	res := OpCodeTwo(input, position)

	for ind, _ := range expectedRes {
		if res[ind] != expectedRes[ind] {
			t.Fatalf("Result %v is not expected", res)
		}
	}
}
