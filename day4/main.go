package main

import (
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	// Read the input file
	data, err := os.ReadFile("day4/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input:=strings.Split(string(data), "\n")
	// init score
	var scoreR1 int64 = 0
	var scoreR2 int64 = 0
 
	for _, line := range input {
	    //Part 1
	    scoreR1 += score1(line)
	    //Part 2
	    scoreR2 += score2(line)
	}
	
	log.Printf("Part 1 - score: %d\n", scoreR1)
	log.Printf("Part 2 - score: %d\n", scoreR2)
}

func score1(line string) int64 {
    areas := strings.Split(line, ",")
    if len(areas) != 2 {
        return 0
    }
    area1 := strings.Split(areas[0], "-")
    area2 := strings.Split(areas[1], "-")
    area11, _ := strconv.Atoi(area1[0])
    area12, _ := strconv.Atoi(area1[1])
    area21, _ := strconv.Atoi(area2[0])
    area22, _ := strconv.Atoi(area2[1]) 
    if (area11 <= area21 && area12 >= area22) ||
       (area11 >= area21 && area12 <= area22) {
        return 1
    }
    return 0
}

func score2(line string) int64 {
    areas := strings.Split(line, ",")
    if len(areas) != 2 {
        return 0
    }
    area1 := strings.Split(areas[0], "-")
    area2 := strings.Split(areas[1], "-")
    area11, _ := strconv.Atoi(area1[0])
    area12, _ := strconv.Atoi(area1[1])
    area21, _ := strconv.Atoi(area2[0])
    area22, _ := strconv.Atoi(area2[1]) 
    if (area11 <= area21 && area12 >= area21) ||
       (area21 <= area11 && area22 >= area11) ||
       (area11 <= area21 && area12 >= area22) ||
       (area11 >= area21 && area12 <= area22) {
        return 1
    }
    return 0
}
