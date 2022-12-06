package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    // "strings"
    // "strconv"
)

func main() {	

    f, err := os.Open("6.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    ans := 0
    ans2 := 0

    for scanner.Scan() {
		line := scanner.Text()

        for i := 0; i < len(line); i++ {
            if (i >= 3) {
                found := true

                for j := i - 3; j <= i; j++ {
                    for k := i - 3; k < i; k++ {
                        if (j != k) {
                            if (line[j] == line[k]) {
                                found = false
                            }
                        }
                    }
                }

                if (ans == 0 && found) {
                    ans = i + 1;
                }

            }

            if (i >= 13) {
                found := true

                for j := i - 13; j <= i; j++ {
                    for k := i - 13; k < i; k++ {
                        if (j != k) {
                            if (line[j] == line[k]) {
                                found = false
                            }
                        }
                    }
                }

                if (ans2 == 0 && found) {
                    ans2 = i + 1;
                }

            }
        }


    }

    fmt.Printf("Answer 1: %d\n", ans)
    fmt.Printf("Answer 2: %d\n", ans2)

}