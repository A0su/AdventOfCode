package main

import  (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func validate(passport []string) (bool) {
	accepted := []string{"blu","brn","gry","grn","hzl","oth"}
	hairValidate := []string{"0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f"}
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
				switch substr[0:3] {
					case "byr":

					case "iyr":

					case "eyr":

					case "hgt":

					case "hcl":
						param := substr[4:]
						if len(param) == 7 {
							if param[0] == "#" {
								for _,ele := range param[1:] {
									if 
								}
							}
						}
					case "ecl":
						param := substr[4:]
						for _,ele := range accepted {
							if param == ele {
								m["ecl"] = true
								break
							}
						}
					case "pid":
						if len(substr[4:]) == 9 {
							m["pid"] = true
						}
				}
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

	if validate(passport) {
		count++
	}


	fmt.Printf("Count: %d\n", count)
}
