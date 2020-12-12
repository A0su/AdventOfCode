package main

import ( 
	"fmt"
	"os"
	"bufio"
	"strings"
)


func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	var groups [][]string
	var totalCount, groupCount int
	
	f,err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var buff []string
	var curr string
	for scanner.Scan() {
		curr = scanner.Text()
		if string(curr) == "" {
			groups = append(groups,buff)
			buff = []string{}
		} else {
			buff = append(buff,curr)
		}
	}
	groups = append(groups,buff)
	
	var qMap map[string]int
	for _,ele := range groups {
		qMap = make(map[string]int)
		groupCount = 0
		var numPeople int = len(ele)
		for _,person := range ele {
			tmp := strings.Split(person,"")
			for _,char := range tmp {
				qMap[string(char)]++
			} 
		}
		for _,value := range qMap {
			if value == numPeople {
				groupCount++
			}
		}
		totalCount += groupCount
	}

	fmt.Printf("The sum of the counts is: %d", totalCount)
}