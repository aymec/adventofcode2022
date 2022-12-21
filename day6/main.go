package main

import (
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// Read the input file
	data, err := os.ReadFile("day6/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input:=strings.Split(string(data), "\n")
	// init score
	scoreR1 := 0
	scoreR2 := 0
	//A map to contain the latest 4 characters and their count
	//If there are 4 different chars, len will equal 4 
	char4Map := make(map[rune]int)
	//A map to contain the latest 14 characters and their count
    //If there are 14 different chars, len will equal 14
	char14Map := make(map[rune]int)
 
	for idx, c := range input[0] {
	    var char4ago,char14ago rune
	    //Part 1
        //Save char that is over 4 chars ago
        if idx > 3 {
            char4ago, _ = utf8.DecodeRuneInString(input[0][idx-4:])
        }
        processCharMap(c, char4ago, &char4Map)
        if scoreR1 == 0 && len(char4Map) == 4 {
            scoreR1=idx+1
        }
        //Part 2
        //Save char that is over 4 chars ago
        if idx > 13 {
            char14ago, _ = utf8.DecodeRuneInString(input[0][idx-14:])
        }
        processCharMap(c, char14ago, &char14Map)
        if scoreR2 == 0 && len(char14Map) == 14 {
            scoreR2=idx+1
            break //We just found the 14 consecutive chars, stop now
        }
	}
	
	log.Printf("Part 1 - answer: %d\n", scoreR1)
	log.Printf("Part 2 - answer: %d\n", scoreR2)
}

func processCharMap(addRune rune, remRune rune, m *map[rune]int) {
    // Add latest char to map
    if v, ok := (*m)[addRune]; ok != false {
        (*m)[addRune]=v+1
    } else {
        (*m)[addRune]=1
    }
    //Remove char that is over size chars ago
    //It will not remove anything if remRune is nil
    if v, ok := (*m)[remRune]; ok != false {
        if v>1 {
            (*m)[remRune]=v-1
        } else {
            delete(*m, remRune)
        }
    }
}
