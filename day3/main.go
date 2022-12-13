package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// Read the input file
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input:=strings.Split(string(data), "\n")
	// init score
	var scoreR1 int64 = 0
	var scoreR2 int64 = 0
	var groupOf3 [3]string
 
	for i, line := range input {
	    //Part 1
		compartment1 := line[0:len(line)/2]
		compartment2 := line[len(line)/2:len(line)]
		scoreR1 += countScore(diffCompartments(compartment1, compartment2))
		
		//Part 2
		groupOf3[i%3] = line
		if i%3 == 2 {
		    scoreR2 += countScore2(groupOf3)
		}
	}
	
	log.Printf("Part 1 - score: %d\n", scoreR1)
	log.Printf("Part 2 - score: %d\n", scoreR2)
}

func countScore2(arr [3]string) int64 {
    c1StrCounts := charCount(arr[0])
    c2StrCounts := charCount(arr[1])
    c3StrCounts := charCount(arr[2])
    
    comp1 := findCommonLetters(c1StrCounts, c2StrCounts)
    comp2 := findCommonLetters(c1StrCounts, c3StrCounts)
    comp := findCommonLetters(comp1, comp2)
    
    return countScore(comp)
}

func countScore(m map[byte]int) int64 {
    var score int64 = 0
    for k, v := range m {
        if strings.ToUpper(string(k)) == string(k) {
            score += int64(v*(27 + int(k-'A')))
        } else {
            score += int64(v*(1 + int(k-'a')))
        }
    }
    return score
}

func charCount(s string) map[byte]int {
    strCounts := make(map[byte]int)
    for i:=0; i<len(s); i++ {
        v, contains := strCounts[s[i]]
        if !contains {
            strCounts[s[i]] = 1
        } else {
            strCounts[s[i]] = v+1
        }
    }
    return strCounts
}

func findCommonLetters(m1 map[byte]int, m2 map[byte]int) map[byte]int {
    ret := make(map[byte]int)
    for k, _ := range m1 {
        _, contains := m2[k]
        if contains {
            ret[k] = 1
        }
    }
    return ret
}

// Return a map with the char present only in string or the other
// and the corresponding number of occurrences
func diffCompartments(c1 string, c2 string) map[byte]int {
    ret := make(map[byte]int)
    if len(c1)!=len(c2) {
        return ret
    }
    c1StrCounts := charCount(c1)
    c2StrCounts := charCount(c2)
    
    //compare maps and build the map to return
    ret = findCommonLetters(c1StrCounts, c2StrCounts)
    
    return ret
}

