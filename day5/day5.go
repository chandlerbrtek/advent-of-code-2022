package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func main() {	

    f, err := os.Open("5.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    s1 := new(Stack)
    s1.Push("D")
    s1.Push("L")
    s1.Push("V")
    s1.Push("T")
    s1.Push("M")
    s1.Push("H")
    s1.Push("F")

    s2 := new(Stack)
    s2.Push("H")
    s2.Push("Q")
    s2.Push("G")
    s2.Push("J")
    s2.Push("C")
    s2.Push("T")
    s2.Push("N")
    s2.Push("P")

    s3 := new(Stack)
    s3.Push("R")
    s3.Push("S")
    s3.Push("D")
    s3.Push("M")
    s3.Push("P")
    s3.Push("H")

    s4 := new(Stack)
    s4.Push("L")
    s4.Push("B")
    s4.Push("V")
    s4.Push("F")

    s5 := new(Stack)
    s5.Push("N")
    s5.Push("H")
    s5.Push("G")
    s5.Push("L")
    s5.Push("Q")

    s6 := new(Stack)
    s6.Push("W")
    s6.Push("B")
    s6.Push("D")
    s6.Push("G")
    s6.Push("R")
    s6.Push("M")
    s6.Push("P")

    s7 := new(Stack)
    s7.Push("G")
    s7.Push("M")
    s7.Push("N")
    s7.Push("R")
    s7.Push("C")
    s7.Push("H")
    s7.Push("L")
    s7.Push("Q")

    s8 := new(Stack)
    s8.Push("C")
    s8.Push("L")
    s8.Push("W")

    s9 := new(Stack)
    s9.Push("R")
    s9.Push("D")
    s9.Push("L")
    s9.Push("Q")
    s9.Push("J")
    s9.Push("Z")
    s9.Push("M")
    s9.Push("T")

    tempStack := new(Stack)

    stacks := [9]Stack{*s1, *s2, *s3, *s4, *s5, *s6, *s7, *s8, *s9}

    for scanner.Scan() {
		line := scanner.Text()
        words := strings.Fields(line)

        numToMove, err := strconv.Atoi(words[1])
        fromCol, err := strconv.Atoi(words[3])
        toCol, err := strconv.Atoi(words[5])

        if err != nil {
            fmt.Println(err)
        }

        for i := 0; i < numToMove; i++ {
            tempStack.Push(stacks[fromCol - 1].Pop())
        }

        for i := 0; i < numToMove; i++ {
            stacks[toCol - 1].Push(tempStack.Pop())
        }
        
	}

    for i := 0; i < 9; i++ {
        fmt.Printf("Answer: %s\n", stacks[i].Peek())
    }
	
}

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}	
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}