package main

import (
	"fmt"
	"bufio"
	"os"
)


func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	var floor int = 0
	var pos, set int
	f,err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := scanner.Text()

	for idx,ele := range input {
		if string(ele) == "(" {
			floor++
		} else {
			floor--
		}
		if floor == -1 && set != 1 {
			set = 1
			pos = idx + 1 //1 indexed
		}
	}
	fmt.Printf("The floor is: %d\n", floor)
	fmt.Printf("The position of the first character to enter the basement is: %d\n", pos)
}