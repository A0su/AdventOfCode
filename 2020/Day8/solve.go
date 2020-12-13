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

/*	Returns accumulator if finite execution	
 * 	Else -1
*/
func execProg(program []Statement, insLength int) (int) {
	var PC int = 0
	var accumulator int = 0
	var seenIns []int

	for PC < insLength {
		if !in(seenIns, PC) {
			return -1
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
			panic("ERROR: Unknown instruction " + program[PC].ins)
		}
	}

	return accumulator
}


func main() {
	var val int
	var ins string
	var err error
	var properAccumulator int
	var insLength int = 0
	var program []Statement
	var mutate []Statement

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

	//Run through all mutations of program
	for idx,ele := range program {
		if ele.ins == "nop" || ele.ins == "jmp" {
			mutate = make([]Statement, insLength)
			copy(mutate,program)
			if ele.ins == "jmp" {
				mutate[idx].ins = "nop"
			} else {
				mutate[idx].ins = "jmp"
			}
			if retVal := execProg(mutate,insLength); retVal != -1 {
				properAccumulator = retVal
				break
			}
		}
	}



	fmt.Printf("Value of accumulator: %d\n", properAccumulator)
}
