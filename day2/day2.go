package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strings"
)

func main() {
    f, err := os.Open("2.txt")
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
        // do something with a line
		line := scanner.Text()
		words := strings.Fields(line)

		if (strings.Compare(words[1], "X") == 0) {
			total += 1
			if (strings.Compare(words[0], "A") == 0) {
				total += 3
			} else if (strings.Compare(words[0], "C") == 0) {
				total += 6
			}
		} else if (strings.Compare(words[1], "Y") == 0) {
			total += 2
			if (strings.Compare(words[0], "B") == 0) {
				total += 3
			} else if (strings.Compare(words[0], "A") == 0) {
				total += 6
			}
		} else {
			total += 3
			if (strings.Compare(words[0], "C") == 0) {
				total += 3
			} else if (strings.Compare(words[0], "B") == 0) {
				total += 6
			}
		}

		if (strings.Compare(words[0], "A") == 0) {
			if (strings.Compare(words[1], "Y") == 0) {
				total2 += 4
			} else if (strings.Compare(words[1], "Z") == 0) {
				total2 += 8
			} else {
				total2 += 3
			}
		} else if (strings.Compare(words[0], "B") == 0) {
			if (strings.Compare(words[1], "Y") == 0) {
				total2 += 5
			} else if (strings.Compare(words[1], "Z") == 0) {
				total2 += 9
			} else {
				total2 += 1
			}
		} else {
			if (strings.Compare(words[1], "Y") == 0) {
				total2 += 6
			} else if (strings.Compare(words[1], "Z") == 0) {
				total2 += 7
			} else {
				total2 += 2
			}
		}
    }

	fmt.Printf("Answer 1: %d\n", total)
	fmt.Printf("Answer 2: %d\n", total2)

}