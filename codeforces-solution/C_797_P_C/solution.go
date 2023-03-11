package main

import (
	"fmt"
)

type Stack []byte

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str byte) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Len() int {
	return len(*s)
}

func New() Stack {
	return Stack{}
}

type AlphabetArr [26]int

// func InitArr(arr AlphabetArr, strInput string) {

// }

func getPosition(strChar byte) int {
	var aChar byte = 'a'
	return int(strChar - aChar)
}

func InitArr(strInput string) AlphabetArr {
	var arr AlphabetArr
	for _, c := range strInput {
		pos := getPosition(byte(c))
		arr[pos] += 1
	}
	return arr
}

func IsCompletelyExplored(arr AlphabetArr) int {
	for ind, c := range arr {
		if c > 0 {
			return ind
		}
	}
	return -1
}

func getMinimalString(strInput string) string {
	t := Stack{}
	u := Stack{}
	trackerArr := InitArr(strInput)
	for i := 0; i < len(strInput); i++ {
		t.Push(strInput[i])
		trackerArr[getPosition(strInput[i])] -= 1
		// t is not empty, the last value in t is minimum than value further and the string is not yet completely traversed.
		for !t.IsEmpty() && IsCompletelyExplored(trackerArr) != -1 && getPosition(t[t.Len()-1]) <= IsCompletelyExplored(trackerArr) {
			e, _ := t.Pop()
			u.Push(e)
		}
	}
	for !t.IsEmpty() {
		e, _ := t.Pop()
		u.Push(e)
	}
	return string(u)
}

func main() {
	var strInput string

	// take the input
	fmt.Scan(&strInput)

	fmt.Print(getMinimalString(strInput))
}
