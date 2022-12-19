package main

import (
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	// Read the input file
	data, err := os.ReadFile("day5/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input:=strings.Split(string(data), "\n")
	// init score
	var answerR1 string = ""
	var answerR2 string = ""
	var stacks [][]string
	var stacks2 [][]string
	isStacksInitiated := false
 
	for _, line := range input {
	    //Part 1
	    if line == "" {
	        isStacksInitiated = true
	        //copyStacks(&stacks, &stacks2)
	    } else {
	        if !isStacksInitiated {
	            buildStacks(line, &stacks)
	            buildStacks(line, &stacks2)
            } else {
                processCommands1(line, &stacks)
                processCommands2(line, &stacks2)
            }
	    }
	}

    // Build answer for part 1
    for _, charSlice := range stacks {
        answerR1 += charSlice[0]
    }
    
    // Build answer for part 1
    for _, charSlice := range stacks2 {
        answerR2 += charSlice[0]
    }
	
	log.Printf("Part 1 - answer: %s\n", answerR1)
	log.Printf("Part 2 - answer: %s\n", answerR2)
}

func processCommands1(line string, stacks *[][]string) {
    words := strings.Split(string(line), " ")
    var cmd [3]int
    idx := 0 
    
    //Extract values for # of crates, stack from and stack to
    for _, w := range words {
        if v, err := strconv.Atoi(w); err == nil {
            cmd[idx] = v
            idx++
        }
    }

    //Moving the crates, one by one
    for i:=0; i<cmd[0]; i++ {
        tmp := (*stacks)[cmd[1]-1][0]
        (*stacks)[cmd[1]-1] = (*stacks)[cmd[1]-1][1:]
        (*stacks)[cmd[2]-1] = append([]string{tmp}, (*stacks)[cmd[2]-1]...)
    }
}

func processCommands2(line string, stacks *[][]string) {
    words := strings.Split(string(line), " ")
    var cmd [3]int
    idx := 0 
    
    //Extract values for # of crates, stack from and stack to
    for _, w := range words {
        if v, err := strconv.Atoi(w); err == nil {
            cmd[idx] = v
            idx++
        }
    }

    //Moving the crates multiple at once
    tmp := append([]string{}, (*stacks)[cmd[1]-1][0:cmd[0]]...)
    (*stacks)[cmd[2]-1] = append(tmp, (*stacks)[cmd[2]-1]...)
    (*stacks)[cmd[1]-1] = (*stacks)[cmd[1]-1][cmd[0]:]
}

func buildStacks(line string, stacks *[][]string) {
    for j, c := range line {
        if (j%4)==1 {
            if len(*stacks) < (j/4)+1 {
                *stacks = append(*stacks, nil)
            }
            if c >= 'A' && c<= 'Z' {
                (*stacks)[j/4] = append((*stacks)[j/4], string(c))
            }
        }
    }
}

func copyStacks(stack1 *[][]string, stacks2 *[][]string) {
    for i, sub1 := range (*stack1) {
        (*stacks2) = append((*stacks2), nil)
        (*stacks2)[i] = append((*stacks2)[i], sub1...) 
    }
}
