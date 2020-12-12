package main

import (
  "bufio"
  "os"
  "fmt"
  "sort"
)

func errorCheck(e error) {
  if e != nil {
    panic(e)
  }
}

func parseSeat(row, col string) (int) {
  seatRows := []int{0,127}
  for _,ele := range row {
    if string(ele) == "F" {
      seatRows[1] -= ((seatRows[1] - seatRows[0]) + 1)  / 2
    } else {
      seatRows[0] += ((seatRows[1] - seatRows[0]) + 1)  / 2
    }
  }

  seatCols := []int{0,7}
  for _,ele := range col {
    if string(ele) == "L" {
      seatCols[1] -= ((seatCols[1] - seatCols[0]) + 1)  / 2
    } else {
      seatCols[0] += ((seatCols[1] - seatCols[0]) + 1) / 2
    }
  }

  return calcSeatID(seatRows[0],seatCols[0])
}


func calcSeatID(row, col int) (int) {
  return row * 8 + col 
}


func main() {
  var lines []string
  var seatIDs []int
  maxSeatID := -1
  var mySeatID int
  
  f, err := os.Open("./input.txt")
  errorCheck(err)

  defer f.Close()
 
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    curr := scanner.Text()
    errorCheck(err)
    lines = append(lines,curr)
  }

  //Part One
  for _,ele := range lines {
    currSeatID := parseSeat(ele[:7],ele[7:])
    if currSeatID > maxSeatID {
      maxSeatID = currSeatID
    }
    seatIDs = append(seatIDs,currSeatID)
  }

  sort.Ints (seatIDs)
  //Part Two
  var prev int = seatIDs[0]
  for _,ele := range seatIDs[1:] {
    if ele == prev + 2 {
      mySeatID = ele - 1
      break
    }
    prev = ele
  }

  fmt.Printf("The highest seat ID is: %d\n", maxSeatID)
  fmt.Printf("My seat ID is: %d", mySeatID)
}
