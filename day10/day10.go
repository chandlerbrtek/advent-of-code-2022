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

    f, err := os.Open("10.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    reg := 1
    ans := 0
    val := 0
    cycles := 0
    spec := 0
    var crt [240]string

    for scanner.Scan() {
		line := scanner.Text()
        words := strings.Fields(line)
        if (spec == reg || spec == reg - 1 || spec == reg + 1 ) {
            crt[cycles] = "#"
        } else {
            crt[cycles] = "."
        }
        cycles++
        spec++
        if (spec == 40) {
            spec = 0
        }

        if ((cycles - 20) % 40 == 0) {
            ans += (cycles * reg)
        }
      
        if (words[0] != "noop") {
            if (spec == reg || spec == reg - 1 || spec == reg + 1 ) {
                crt[cycles] = "#"
            } else {
                crt[cycles] = "."
            }
            cycles++;
            spec++
            if (spec == 40) {
                spec = 0
            }
            val, err = strconv.Atoi(words[1])
            if err != nil {}
            if ((cycles - 20) % 40 == 0) {
                ans += (cycles * reg)
            }
            reg += val
        }
    }
    fmt.Printf("Answer 1: %d\n", ans)

    for i := 0; i < 6; i++ {
        for j := 0; j < 40; j++ {
            fmt.Printf("%s", crt[(i*40)+j])
        }
        fmt.Printf("\n")    
    }
}