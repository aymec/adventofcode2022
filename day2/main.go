package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// Read the input file
	data, err := os.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input:=strings.Split(string(data), "\n")
	// init score
	score:=0
	scoreR2:=0

	// Part 1 and 2: Count the score 2 different ways
	for _, line := range input {
		round := strings.Split(string(line), " ")
		op := randomLetterToRPS(round[0])
		me := randomLetterToRPS(round[1])
		score+=shapeScore(me) + outcomeScore(op, me)
		scoreR2+=shapeScoreR2(op, round[1]) + outcomeScoreR2(round[1])
	}

	log.Printf("Part 1 - score: %d\n", score)
	
	log.Printf("Part 2 - score: %d\n", scoreR2)
}

// Transforms A, B, C, X, Y, Z into R, P and S
// for easier manipulation
func randomLetterToRPS (c string) string {
    switch c {
    case "A": return "R"
    case "B": return "P"
    case "C": return "S"
    case "X": return "R"
    case "Y": return "P"
    case "Z": return "S"
    default: return ""
    }
}

func shapeScore (c string) int {
    switch c {
    case "R": return 1
    case "P": return 2
    case "S": return 3
    default: return 0
    }
}

//op: shape of my opponent, me: my shape
//return the score corresponding to the outcome
func outcomeScore (op string, me string) int {
    switch op {
    case "R":
        switch me {
        case "R": return 3
        case "P": return 6
        case "S": return 0
        default: return 0
        }
    case "P":
        switch me {
        case "R": return 0
        case "P": return 3
        case "S": return 6
        default: return 0
        }
    case "S":
        switch me {
        case "R": return 6
        case "P": return 0
        case "S": return 3
        default: return 0
        }
    default: return 0
    }
}

// X: Loss, Y: Draw, Z: Win
func outcomeScoreR2 (c string) int {
    switch c {
    case "X": return 0
    case "Y": return 3
    case "Z": return 6
    default: return 0
    }
}

// Return the score associated to my shape based on outcome
// and my opponent's shape
func shapeScoreR2 (op string, outcome string) int {
    switch op {
    case "R":
        switch outcome {
        case "X": return shapeScore("S")
        case "Y": return shapeScore("R")
        case "Z": return shapeScore("P")
        default: return 0
        }
    case "P":
        switch outcome {
        case "X": return shapeScore("R")
        case "Y": return shapeScore("P")
        case "Z": return shapeScore("S")
        default: return 0
        }
    case "S":
        switch outcome {
        case "X": return shapeScore("P")
        case "Y": return shapeScore("S")
        case "Z": return shapeScore("R")
        default: return 0
        }
    default: return 0
    }
}

