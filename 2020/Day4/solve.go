package main

import  (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func validate(passport []string) (bool) {
	m := map[string]bool{
		"byr":false,
		"iyr":false,
		"eyr":false,
		"hgt":false,
		"hcl":false,
		"ecl":false,
		"pid":false,
	}

	for _,ele := range passport {
		strs := strings.Split(ele," ")
		for _,substr := range strs {
			if substr[0:3] != "cid" {
				m[substr[0:3]] = true
			}
		}
	}

	for _,v := range m {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	var passport []string
	var count int = 0

	f,err := os.Open("./input.txt")
	errorCheck(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		curr := scanner.Text()
		if curr == "" {
			if validate(passport) {
				count++
			}
			passport = []string{}
		} else {
			passport = append(passport, curr)
		}
	}

	fmt.Printf("Count: %d\n", count)
}
