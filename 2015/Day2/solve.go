package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	f,err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr := scanner.Text()
		tmp := strings.Split(curr,"x")
	}
}