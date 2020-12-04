package main

import (
  "bufio"
  "os"
  "fmt"
  "io"
)

func errorCheck(e error) {
  if e != nil {
    panic("An error has occured")
  }
}

/*
  input(s):
    gameMap []string : ASCII map describing the forest
    pathing []int    : slice of ints describing the pathing 
  returns: number of trees hit for given pathing
*/
func traverseMap(gameMap []string, pathing []int) (int) {
  height := len(gameMap)
  length := len(gameMap[0])
  var treesFound int = 0
  coord := []int{0, 0} 
  
  for coord[0] < height {
    if coord[1] >= length {
      coord[1] = coord[1] % length
    }

    if string(gameMap[coord[0]][coord[1]]) == "#" {
      treesFound += 1
    }

    coord[0] += pathing[0]
    coord[1] += pathing[1]
  }
  return treesFound
}


func main() {
  f, err := os.Open("./input.txt")
  errorCheck(err)

  defer f.Close()

  var gameMap []string
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    curr := scanner.Text()
    gameMap = append(gameMap,curr)
  }

  //Part 1
  pathing := []int{1,3}
  count := traverseMap(gameMap, pathing)

  //Part 2
  var product int = 1
  paths := [][]int{
    {1, 1},
    {1, 3},
    {1, 5},
    {1, 7},
    {2, 1},
  }
  
  for _,ele := range paths {
    product *= traverseMap(gameMap, ele)
  }

  fmt.Println("Part 1")
  s := fmt.Sprintf("Count: %d\n", count)
  io.WriteString(os.Stdout, s)
  fmt.Println("Part 2")
  s = fmt.Sprintf("Count: %d\n", product)
  io.WriteString(os.Stdout, s)

}
