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

    f, err := os.Open("4.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)
	total := 0
	total2 := 0

    for scanner.Scan() {
		line := scanner.Text()
        stringSlice := strings.Split(line, ",")
        pairOne := strings.Split(stringSlice[0], "-")
        pairTwo := strings.Split(stringSlice[1], "-")

        p1s, err := strconv.Atoi(pairOne[0])
        p1e, err := strconv.Atoi(pairOne[1])
        p2s, err := strconv.Atoi(pairTwo[0])
        p2e, err := strconv.Atoi(pairTwo[1])

        if err != nil {
            fmt.Println(err)
        }

        if (p1s <= p2s && p1e >= p2e ||
            p2s <= p1s && p2e >= p1e) {
            total++;
        }

        for i := p1s; i <= p1e; i++ {
            if (i >= p2s && i <= p2e) {
                total2++;
                break;
            }
        }

	}

	fmt.Printf("Answer 1: %d\n", total)
	fmt.Printf("Answer 2: %d\n", total2)
	
}