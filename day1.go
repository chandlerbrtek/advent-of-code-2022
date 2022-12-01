package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
	"reflect"
	"math"
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
	max := 0.0

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
			max = math.Max(max, total)
			total = 0
		}
    }

	fmt.Printf("answer part 1: %d\n", int(max))

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}