package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    // "strings"
    "strconv"
)

func main() {	

    f, err := os.Open("8.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    ans := 392
    ans2 := 0
    var grid [99][99] int
    row := 0

    for scanner.Scan() {
		line := scanner.Text()

        for i := 0; i < len(line); i++ {
            grid[row][i], err = strconv.Atoi(string(line[i]))
            if err != nil {}
        }

        row++

    }

    for i := 1; i < 98; i++ {
        for j := 1; j < 98; j++ {
            found := false
            l := false
            r := false
            u := false
            d := false
    
            // up
            for k := i; !found && !u && k > 0; k-- {
                if (grid[k - 1][j] >= grid[i][j]) {
                    u = true
                } else if (k == 1) {
                    found = true
                }
            }

            // left
            for k := j; !found && !l && k > 0; k-- {
                if (grid[i][k - 1] >= grid[i][j]) {
                    l = true
                } else if (k == 1) {
                    found = true
                }
            }

            // down
            for k := i; !found && !d && k < 98; k++ {
                if (grid[k + 1][j] >= grid[i][j]) {
                    d = true
                } else if (k == 97) {
                    found = true
                }
            }

            // right
            for k := j; !found && !r && k < 98; k++ {
                if (grid[i][k + 1] >= grid[i][j]) {
                    r = true
                } else if (k == 97) {
                    found = true
                }
            }     
            
            if (found) {
                ans++;
            }

        }    
    }

    for i := 1; i < 98; i++ {
        for j := 1; j < 98; j++ {
            found := false
            l := false
            r := false
            u := false
            d := false
            us := 0
            ls := 0
            ds := 0
            rs := 0
    
            // up
            for k := i; !found && !u && k > 0; k-- {
                if (grid[k - 1][j] >= grid[i][j]) {
                    u = true
                }
                us = i - k + 1
            }

            // left
            for k := j; !found && !l && k > 0; k-- {
                if (grid[i][k - 1] >= grid[i][j]) {
                    l = true
                }
                ls = j - k + 1
            }

            // down
            for k := i; !found && !d && k < 98; k++ {
                if (grid[k + 1][j] >= grid[i][j]) {
                    d = true
                } 
                ds = k + 1 - i
            }

            // right
            for k := j; !found && !r && k < 98; k++ {
                if (grid[i][k + 1] >= grid[i][j]) {
                    r = true
                }
                rs = k + 1 - j
            }     
            
            if (us * ls * ds * rs > ans2) {
                ans2 = us * ls * ds * rs
            }

        }    
    }
    
    fmt.Printf("Answer 1: %d\n", ans)
    fmt.Printf("Answer 2: %d\n", ans2)

}