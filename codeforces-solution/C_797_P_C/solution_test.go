package main

import (
	"testing"
)

func TestGetPosition(t *testing.T) {
	expectedOuput := 2
	actualOutput := getPosition('c')
	if expectedOuput != actualOutput {
		t.Errorf("Failed getPosition function for, got %d, wanted %d ", actualOutput, expectedOuput)
	}
}

func TestGetMinimalString(t *testing.T) {
	expectedOutputs := [...]string{"abc", "abdc"}
	mockInputs := [...]string{"cab", "acdb"}
	for i := 0; i < len(mockInputs); i++ {
		actualOutput := getMinimalString(mockInputs[i])
		expectedOutput := expectedOutputs[i]
		if actualOutput != expectedOutput {
			t.Errorf("Failed getMinimalString function for %s, got %s, wanted %s", mockInputs[i], actualOutput, expectedOutput)
		}
	}
}
