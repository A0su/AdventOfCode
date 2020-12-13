package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
)


func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}


func findBus(busList []int, earliest int) (int) {
	var minTime int = math.MaxInt32
	var val, busID int
	for _,ele := range busList {
		val = earliest
		for val % ele != 0 {
			val++
		}
		if val < minTime {
			minTime = val
			busID = ele
		}
	}
	return (minTime - earliest) * busID
}


func main() {
	var earliest int
	var atoiErr error
	var waitTime int
	f,err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	earliest,atoiErr = strconv.Atoi(scanner.Text())
	errorCheck(atoiErr)

	scanner.Scan()
	tmp := scanner.Text()
	busList := strings.Split(tmp,",")
	var cleanList []int 
	for _,ele := range busList {
		if string(ele) != "x" {
			tmp,err := strconv.Atoi(ele)
			errorCheck(err)
			cleanList = append(cleanList,tmp)
		}
	}

	waitTime = findBus(cleanList, earliest)


	fmt.Println("The earliest bus ID is: %d", waitTime)
}