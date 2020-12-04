package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "io"
)

func errorCheck(e error) {
  if e != nil {
    panic("An error has occured")
  }
}

func strValidate(freq, rule, pass string) (isValid bool){
  var min, max, count int
  var err error

  split := strings.Split(freq,"-")
  min, err = strconv.Atoi(split[0])
  errorCheck(err)
  max, err = strconv.Atoi(split[1])
  errorCheck(err)
 
  for _, ele := range pass {
    if string(ele) == rule { 
      count += 1
    }
    if count > max {
      return false
    }
  } 

  if count < min {
    return false
  }
  
  return true
}

func strValidatePartTwo(freq, rule, pass string) (isValid bool){
  var idx, idx2 int
  var err error

  split := strings.Split(freq,"-")
  idx, err = strconv.Atoi(split[0])
  errorCheck(err)
  idx2, err = strconv.Atoi(split[1])
  errorCheck(err)

  if (string(pass[idx - 1]) == rule) != (string(pass[idx2 - 1]) == rule) { //no logical xor for bools :(
    return true
  }

  return false
}

func main() {

  f, err := os.Open("./input.txt")
  errorCheck(err)

  defer f.Close()

  count, count2 := 0, 0
  var freq,rule,pass string  
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    curr := scanner.Text()
    split := strings.Split(curr, ": ")
    rule, pass = split[0], split[1]
    split = strings.Split(rule, " ")
    freq, rule = split[0], split[1]
    if strValidate(freq,rule,pass) {
      count += 1
    }
    if strValidatePartTwo(freq,rule,pass) {
      count2 += 1
    }
  }

  fmt.Println("Part 1")
  s := fmt.Sprintf("Count: %d\n", count)
  io.WriteString(os.Stdout, s)
  fmt.Println("Part 2")
  s = fmt.Sprintf("Count: %d\n", count2)
  io.WriteString(os.Stdout, s)
}
