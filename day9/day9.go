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

    f, err := os.Open("9.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    // lets pretend start is 0,0...

    hx := 0;
    hy := 0;
    tx := 0;
    ty := 0;
    var coords []string

    for scanner.Scan() {
		line := scanner.Text()
        words := strings.Fields(line)

        dir := words[0]
        dist, err := strconv.Atoi(words[1])
        if err != nil {}

        if dir == "U" {
            for i := 0; i < dist; i++ { 
                hy++;
                if (outOfBounds(hx, hy, tx, ty)) {
                    tx = hx;
                    ty++;
                }
                ygrids := strconv.Itoa(tx) + "," + strconv.Itoa(ty)
                if (!contains(coords, ygrids)) {
                    coords = append(coords, ygrids)
                }
            }
        } else if dir == "D" {
            for i := 0; i < dist; i++ { 
                hy--;
                if (outOfBounds(hx, hy, tx, ty)) {
                    tx = hx;
                    ty--;
                }
                ygrids := strconv.Itoa(tx) + "," + strconv.Itoa(ty)
                if (!contains(coords, ygrids)) {
                    coords = append(coords, ygrids)
                }
            }
        } else if dir == "L" {
            for i := 0; i < dist; i++ { 
                hx--;
                if (outOfBounds(hx, hy, tx, ty)) {
                    tx--;
                    ty = hy;
                }
                ygrids := strconv.Itoa(tx) + "," + strconv.Itoa(ty)
                if (!contains(coords, ygrids)) {
                    coords = append(coords, ygrids)
                }
            }
        } else { // R
            for i := 0; i < dist; i++ { 
                hx++;
                if (outOfBounds(hx, hy, tx, ty)) {
                    tx++;
                    ty = hy;
                }
                ygrids := strconv.Itoa(tx) + "," + strconv.Itoa(ty)
                if (!contains(coords, ygrids)) {
                    coords = append(coords, ygrids)
                }
            }
        }

    }
    
    fmt.Printf("Answer 1: %d\n", len(coords))

    g, err := os.Open("9.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer g.Close()

    scanner = bufio.NewScanner(g)

    xs := [10]int{0,0,0,0,0,0,0,0,0,0}
    ys := [10]int{0,0,0,0,0,0,0,0,0,0}
    var coords2 []string
    // upward := false

    for scanner.Scan() {
		line := scanner.Text()
        words := strings.Fields(line)

        dir := words[0]
        dist, err := strconv.Atoi(words[1])
        if err != nil {}

        if dir == "U" {
            for i := 0; i < dist; i++ { 
                ys[0]++;
                prevx := xs[0]
                for j := 1; j < 10; j++ {
                    tempx := xs[j]
                    if (outOfBounds(xs[j-1], ys[j-1], xs[j], ys[j])) {
                        xs[j] = prevx;
                        if (ys[j] < ys[j-1]) {
                            ys[j]++
                        } else {
                            ys[j] = ys[j-1]
                        }               
                    }     
                    prevx = tempx
                }
                ygrids := strconv.Itoa(xs[9]) + "," + strconv.Itoa(ys[9])
                if (!contains(coords2, ygrids)) {
                    coords2 = append(coords2, ygrids)
                }
            }
        } else if dir == "D" {
            for i := 0; i < dist; i++ { 
                ys[0]--;
                prevx := xs[0]
                for j := 1; j < 10; j++ {
                    tempx := xs[j]
                    if (outOfBounds(xs[j-1], ys[j-1], xs[j], ys[j])) {
                        xs[j] = prevx;
                        if (ys[j] > ys[j-1]) {
                            ys[j]--
                        } else {
                            ys[j] = ys[j-1]
                        }
                    }
                    prevx = tempx
                }
                ygrids := strconv.Itoa(xs[9]) + "," + strconv.Itoa(ys[9])
                if (!contains(coords2, ygrids)) {
                    coords2 = append(coords2, ygrids)
                }
            }
        } else if dir == "L" {
            for i := 0; i < dist; i++ { 
                xs[0]--;
                prevy := ys[0]
                for j := 1; j < 10; j++ {
                    tempy := ys[j]
                    if (outOfBounds(xs[j-1], ys[j-1], xs[j], ys[j])) {
                        if (xs[j] > xs[j-1]) {
                            xs[j]--
                        } else {
                            xs[j] = xs[j-1]
                        }
                        ys[j] = prevy;    
                    }
                    prevy = tempy
                }
                ygrids := strconv.Itoa(xs[9]) + "," + strconv.Itoa(ys[9])
                if (!contains(coords2, ygrids)) {
                    coords2 = append(coords2, ygrids)
                }
            }
        } else { // R
            for i := 0; i < dist; i++ { 
                xs[0]++;
                prevy := ys[0]
                for j := 1; j < 10; j++ {
                    tempy := ys[j]
                    if (outOfBounds(xs[j-1], ys[j-1], xs[j], ys[j])) {
                        if (xs[j] < xs[j-1]) {
                            xs[j]++
                        } else {
                            xs[j] = xs[j-1]
                        }                        
                        ys[j] = prevy;
                    }
                    prevy = tempy
                }
                ygrids := strconv.Itoa(xs[9]) + "," + strconv.Itoa(ys[9])
                if (!contains(coords2, ygrids)) {
                    coords2 = append(coords2, ygrids)
                }
            }
        }
    }
    
    fmt.Printf("Answer 2: %d\n", len(coords2))

}

func outOfBounds(hx int, hy int, tx int, ty int) bool {
    return (absDiffInt(hx, tx) > 1 || absDiffInt(hy, ty) > 1);
}

func absDiffInt(x, y int) int {
    if x < y {
       return y - x
    }
    return x - y
 } 

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
