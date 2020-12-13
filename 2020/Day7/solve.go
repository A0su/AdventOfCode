package main

import (
	"os"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

}
