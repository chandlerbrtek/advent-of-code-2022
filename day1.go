package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
	"reflect"
	"sort"
)

func main() {
    f, err := os.Open("inputs/1.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)
	total := 0.0
	var s []int

    for scanner.Scan() {
        // do something with a line
		line := scanner.Text()
		if line != "" {
			val, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(val, err, reflect.TypeOf(val))
			}
			total += float64(val)
		} else {
			s = append(s, int(total))
			total = 0
		}
    }

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	fmt.Printf("Answer 1: %d\n", s[0])
	fmt.Printf("Answer 2: %d\n", s[0] + s[1] + s[2])
	
	
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}