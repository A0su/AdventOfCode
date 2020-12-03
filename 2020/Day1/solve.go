package main

import (
  "os"
  "bufio"
  "strconv"
  "fmt"
)

const target = 2020

func errorCheck(e error) {
  if e != nil {
    panic("Error occured")
  }
}

func main() {
  m := make(map[int]int)
  var lines []int

  f, err := os.Open("./input.txt")
  errorCheck(err)

  defer f.Close()
  
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    curr, err := strconv.Atoi(scanner.Text())
    errorCheck(err)
    lines = append(lines,curr)
  }

  fmt.Println("Part 1")
  for _, ele := range lines {
    _, prs := m[ele]

    if prs {
      fmt.Println("MATCH FOUND")
      fmt.Println(m[ele] * ele)
    }
    m[target - ele] = ele 
  }

  //Part 2
  fmt.Println("\nPart 2")
  for _, ele := range lines{ 
    for _, inner := range lines {
      _, prs := m[ele+inner]

      if prs {
        fmt.Println("MATCH FOUND")
        fmt.Println(ele*inner*m[ele+inner])
        return  
      }
    }
  }  

}
