package main

import (
	"log"
	"os"
	"strings"
	"strconv"
)

type FileType int
const (
    FILE FileType = iota
    DIR  FileType = iota
)

// Tree type
type Node struct {
    name     string
    size     int
    typ      FileType
    parent   *Node
    children []*Node
}

func main() {
	// Read the input file
	data, err := os.ReadFile("day7/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input:=strings.Split(string(data), "\n")
	// init score
	var scoreR1 int64 = 0
	scoreR2 := 0
	
    top := Node{
        name: "/",
        typ: DIR,
        size: 0,
    }
    currentDir := &top

	for _, line := range input {
	    processLine(line, &top, &currentDir)
	}
	
	//Part 1
	stack := []*Node{&top}
    calculateScore1(stack, &scoreR1)
    
    //Part 2
    stack = []*Node{&top}
    //Total size of disk is 70.000.000
    //Min free space needed is 30.000.000
    //=> Max used space should be 40.000.000
    //=> Min Space to delete is space occupied - max used space
    var minSpaceToDel int = top.size - 40000000
    calculateScore2(stack, &scoreR2, minSpaceToDel)
	
	log.Printf("Part 1 - answer: %d\n", scoreR1)
	log.Printf("Part 2 - answer: %d\n", scoreR2)
}

func processLine(line string, top *Node, currentDir **Node) {
    input:=strings.Split(line, " ")
    switch input[0] {
    case "$": // command line
        processCmd(input[1], input[2:], top, currentDir)
    case "dir": // new dir
        // do nothing
    default: // this is a file
        if len(input) > 1 {
            processFile(input[1], input[0], top, currentDir)
        }
    }
}

//Process commands such as cd and ls
func processCmd(cmd string, args []string, top *Node, currentDir **Node) {
    switch cmd {
    case "cd":
        switch args[0] {
        case "/":
            // move to top dir
            *currentDir = top
        case "..":
            // move to parent's directory, and if it's nil, stay in current dir
            if (*currentDir).parent != nil {
                *currentDir = (*currentDir).parent
            }
        default:
            // As we move to a new directy, add it to currentDir's list of children
            addChild(currentDir, DIR, args[0], 0)
        }
    //No need to do anything when `ls`
    }
}

//Add file to current directory's list of children
func processFile(fileName string, size string, top *Node, currentDir **Node) {
    if v, err := strconv.Atoi(size); err == nil {
        addChild(currentDir, FILE, fileName, v)
    }
}

//Adds a child to a directory. If the child is a directory, returns a pointer
//to that child. If the child is a file, returns a pointer to the currentDir
//If the child is a file, it adds the file's size to the size of all its ancestors
func addChild(currentDir **Node, typ FileType, name string, size int) {
    // Only add the child if it's not already in
    for _, node := range (*currentDir).children {
        if name == node.name {
            // this dir is already a child of the current dir
            // No need to add it and we can return the existing node
            if(node.typ == DIR){
                (*currentDir) = node
            }
            return
        }
    }
    // if we get to this point, we need to create a node and
    // add it to the children of the current directory node
    child := Node{
        name: name,
        typ: typ,
        size: size,
        parent: *currentDir,
    }
    if (*currentDir).children == nil {
        (*currentDir).children = []*Node{&child}
    } else {
        (*currentDir).children = append((*currentDir).children, &child)
    }
    if typ == DIR {
        (*currentDir) = &child
    } else {
        //Add file's size to all its ancestors' size
        ancestor := &child
        for ancestor.parent != nil {
            ancestor.parent.size += child.size
            ancestor = ancestor.parent
        }
    }
}

//Pass the top node, and it will return the sum of the sizes
//of all directories whose individual size is 100000 at most 
func calculateScore1(stack []*Node, score *int64) {
    if(stack == nil || len(stack) == 0) {
        return
    }
        
    n := stack[0]
    stack = stack[1:]
    if n.typ == DIR {
        if n.size <= 100000 {
            *score += int64(n.size)
        }
        for _, c := range n.children {
            if c.typ == DIR {
                stack = append(stack, c)
            }
        }
    }
    calculateScore1(stack, score)
}

//Pass the top node, and how much space needs to be deleted
//It will return the size of the smallest directory that needs to be
//delete to free the necessary space 
func calculateScore2(stack []*Node, score *int, minSpaceToDel int) {
    if(stack == nil || len(stack) == 0) {
        return
    }
        
    n := stack[0]
    stack = stack[1:]
    if n.typ == DIR {
        if n.size > minSpaceToDel && (*score == 0 || n.size < *score) {
            *score = n.size
        }
        for _, c := range n.children {
            if c.typ == DIR && c.size > minSpaceToDel {
                stack = append(stack, c)
            }
        }
    }
    calculateScore2(stack, score, minSpaceToDel)
}
