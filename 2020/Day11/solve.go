package main

import (
	"fmt"
	"os"
	"bufio"
)


func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// openSeat
// True: Open
// False: Occupied
func checkAdjacent(seatMap []string, openSeat bool,  x, y int) (bool) {
	var validSeats [][]int
	height := len(seatMap)
	width := len(seatMap[0])

}


//When no change occurs return false
//Else return true
func applySeating(seatMap []string) (int) {


	for outer,row := range seatMap {
		for inner,seat := range row {
			switch string(seat[inner][outer]) {
			case "L":
				if !checkAdjacent(seatMap,true,inner,outer) {
					seat[inner][outer] = "#"
				}
			case "#":
				if checkAdjacent(seatMap,false,inner,outer) >= 4 {
					seat[inner][outer] = "L"
				}
			case ".":
			default: panic("Unknown character in seatMap: " + string(seat[inner][outer]))
			}
		}
	}

}


func findOpenSeats(seatMap []string) (int) {
	var numSeats int = 0
	for _,row := range seatMap {
		for _,seat := range row {
			if string(seat) == "L" {
				numSeats++
			}
		}
	}
	return numSeats
}



func main() {
	var seatMap []string

	f,err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr := scanner.Text()
		seatMap = append(seatMap, curr)
	}


	for findOpenSeats(seatMap) {}
	fmt.Printf("The number of seats unoccupied is: %d", findOpenSeats())
}