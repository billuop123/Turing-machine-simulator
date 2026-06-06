package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ReadState int

const (
	Zero ReadState = iota
	One
)

type (
	Move   string
	States string
)

const (
	R Move = "R"
	L Move = "L"
)

const (
	B States = "B"
	C States = "C"
	A States = "A"
)

type rule struct {
	current States
	read    ReadState
	write   ReadState
	move    Move
	next    States
}
type Index int

const (
	current Index = iota
	read
	write
	move
	next
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		val := fmt.Errorf("format is <statefile> <inputfile>")
		fmt.Println(val)
		return
	}
	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Print(err)
	}

	allStates := string(data)
	lines := strings.Split(allStates, "\n")
	filteredLines := []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filteredLines = append(filteredLines, strings.TrimSpace(line))
		}
	}
	allRules := make([]rule, len(filteredLines))
	for i, line := range filteredLines {
		arr := strings.Split(line, " ")
		if len(arr) < 5 {
			continue
		}
		populateRules(allRules, i, arr)
	}
	inputs, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Print(err)
		return
	}
	inputArr := strings.Split(string(inputs), " ")
	fmt.Println("Initial State:")
	for j := range len(inputArr) {
		fmt.Printf("%s ", inputArr[j])
	}
	fmt.Println()
	fmt.Println("^")
	compute(string(allRules[0].current), inputArr, 0, allRules)
}

func compute(next string, inputArr []string, i int, allRules []rule) {
	if next == "HALT" {
		fmt.Println("The machine is halted")
		return
	}
	r := rule{}
	readVal, _ := strconv.Atoi(inputArr[i])
	for _, rule := range allRules {
		if rule.current == States(next) && rule.read == ReadState(readVal) {
			r = rule
		}
	}
	if r.current == "" {
		fmt.Println("halting...")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if strings.TrimSpace(input) != "" {
		fmt.Println("Enter space or enter only")
	}
	fmt.Println("Current State:" + r.current)
	inputArr[i] = strconv.Itoa(int(r.write))
	switch r.move {
	case L:
		i--
		if i < 0 {
			i = len(inputArr) - 1
		}
	case R:
		i++
		if i == len(inputArr) {
			i = 0
		}
	}
	for j := range len(inputArr) {
		fmt.Printf("%s ", inputArr[j])
	}
	fmt.Println()
	for j := range len(inputArr) {
		if j == i {
			fmt.Print("^")
		} else {
			fmt.Print("  ")
		}
	}
	compute(string(r.next), inputArr, i, allRules)
}

func populateRules(allRules []rule, i int, arr []string) {
	allRules[i].current = States(arr[current])
	readVal, _ := strconv.Atoi(arr[read])
	allRules[i].read = ReadState(readVal)
	writeVal, _ := strconv.Atoi(arr[write])
	allRules[i].write = ReadState(writeVal)
	allRules[i].move = Move(arr[move])
	allRules[i].next = States(arr[next])
}
