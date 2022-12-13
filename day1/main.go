package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"container/heap"
)

// An IntHeap is a max-heap of ints. (in this case, by reversing the Less function)
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Read the input file. It contains a list of integers representing calories carried
	// by elves. Consecutive values are carried by the same elf. Elves' load are 
	// separated by an empty line
	data, err := os.ReadFile("day1/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input:=strings.Split(string(data), "\n")
	
	//Init a max heap that will contain the ordered list of calories 
	h := &IntHeap{}
	heap.Init(h)
	currentCount := 0

	// Part 1: Count calories per elf and keep the max
	for _, line := range input {
		if line != "" {
		    data, err := strconv.Atoi(line)
		    if err != nil {
		        log.Fatal(err)
		    }
		    currentCount+=data
		} else {
		    heap.Push(h, currentCount)
		    currentCount=0
		}
	}
	//In case the last elf carries the max load
	heap.Push(h, currentCount)

	log.Printf("Part 1 - Calories carried by the elf with most calories: %d\n", (*h)[0])

    top3Count := heap.Pop(h).(int)
    top3Count+= heap.Pop(h).(int)
    top3Count+= heap.Pop(h).(int)
	
	log.Printf("Part 2 - Calories carried by the top 3 elves: %d\n", top3Count)
}


