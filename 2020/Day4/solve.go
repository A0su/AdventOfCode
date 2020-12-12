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

func strVal(inputStr string, accepted []string) (bool) {
  var in bool
  for _,ele := range inputStr {
    for _,comp := range accepted {
      if string(ele) == comp {
        in = true
        break
      } 
    }
    if !in { 
      return false
    }
    in = false
  }
  return true 
}

func heightVal(inputStr string) (bool) {
  var magn, unit string
  var sb, sb2 strings.Builder
  sb.WriteString(magn)
  sb2.WriteString(unit)
  for _,ele := range inputStr {
    if ele < 58 {
      sb.WriteRune(ele)
    } else {
      sb2.WriteRune(ele)
    }
  }
  if  sb2.String() == "cm" {
    val, err := strconv.Atoi(sb.String())
    errorCheck(err)
    if val > 193 || val < 150 {
      return false
    }
  } else {
    val, err := strconv.Atoi(sb.String())
    errorCheck(err)
    if val > 76 || val < 59 {
      return false
    }
  }
  return true
}

func validate(passport []string) (bool) {
  accepted := []string{"amb","blu","brn","gry","grn","hzl","oth"}
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
						val, err := strconv.Atoi(substr[4:])
						errorCheck(err)
						if val >= 1920 && val <= 2002 { 
							m["byr"] = true
						}
					case "iyr":
						val, err := strconv.Atoi(substr[4:])
						errorCheck(err)
						if val >= 2010 && val <= 2020 {
							m["iyr"] = true
						}	
					case "eyr":
						val, err := strconv.Atoi(substr[4:])
						errorCheck(err)
						if val >= 2020 && val <= 2030 {
							m["eyr"] = true
						}	
					case "hgt":
						if heightVal(substr[4:]) {
							m["hgt"] =  true
						}
					case "hcl":
						param := substr[4:]
						if len(param) == 7 {
							if string(param[0]) == "#" {
								if strVal(substr[5:],hairValidate) {
									m["hcl"] = true
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
