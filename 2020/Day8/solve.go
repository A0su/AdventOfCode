package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Statement struct {
	ins    string
	offset int
}

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func in(vals []int, PC int) bool {
	for _, v := range vals {
		if v == PC {
			return false
		}
	}
	return true
}

func main() {
	var seenIns []int
	var val int
	var ins string
	var err error
	var PC int = 0
	var insLength int = 0
	var accumulator int = 0
	var program []Statement

	f, err := os.Open("./input.txt")
	errorCheck(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr := scanner.Text()
		strSplit := strings.Split(curr, " ")
		val, err = strconv.Atoi(strSplit[1])
		errorCheck(err)
		ins = strSplit[0]
		program = append(program, Statement{ins, val})
		insLength++
	}

	for PC < insLength {
		if !in(seenIns, PC) {
			fmt.Println("Preventing running ins twice")
			break
		}
		seenIns = append(seenIns, PC)
		switch program[PC].ins {
		case "acc":
			accumulator += program[PC].offset
			PC++
		case "jmp":
			PC += program[PC].offset
		case "nop":
			PC++
		default:
			panic("ERROR: Unknown instruction " + ins)
		}
	}

	fmt.Printf("Value of accumulator: %d\n", accumulator)
}
