package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)


func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}


func validatePreamble(preamble []int, target int) (bool) {
	diffMap := make(map[int]int)

	for _,ele := range preamble {
		if _,pres := diffMap[ele]; pres {
			return true
		} else {
			diffMap[target-ele] = ele
		}
	}
	return false
}


func MinMax(nums []int) (int,int) {
	var min,max int = nums[0],nums[0]
	for _,ele := range nums {
		if min > ele {
			min = ele
		}
		if max < ele {
			max = ele
		}
	}
	return min,max
}


func findWeakness(seen []int, target int) (int) {
	var sum int = 0
	var contArray []int
	for idx,ele := range seen {
		contArray = append(contArray,ele)
		for _,num := range seen[idx:] {
			if sum == target { 
				val1, val2 := MinMax(contArray)
				return val1 + val2
			} else if sum < target {
				sum += num
				contArray = append(contArray,num)
			} else {
				sum = 0
				contArray = make([]int,0)
				break
			}
		}
	}
	return -1 //No match found
}


func main() {
	var preamble = make([]int,25)
	var incorrectNum int
	var seen []int
	f,err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

	var num int
	var atoiErr error
	scanner := bufio.NewScanner(f)
	for idx := 0; idx < 25; idx++ {
		scanner.Scan()
		num, err = strconv.Atoi(scanner.Text())
		preamble[idx] = num
		seen = append(seen,num)
	}
	
	for scanner.Scan() {
		num, atoiErr = strconv.Atoi(scanner.Text())
		errorCheck(atoiErr)
		if !validatePreamble(preamble, num) {
			incorrectNum = num 
			break
		}
		preamble = append(preamble[1:], num)
		seen = append(seen,num)
	}
	fmt.Printf("The first number without the property is %d\n", incorrectNum)
	fmt.Printf("The encryption weekness is: %d\n", findWeakness(seen, incorrectNum))
}