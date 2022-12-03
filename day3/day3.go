package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {	

    f, err := os.Open("3.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)
	total := 0
	total2 := 0
	counter := 1
	line1 := ""
	line2 := ""

    for scanner.Scan() {
		line := scanner.Text()

		s:=line

		var ss []string
		for i := 1; i < len(s); i++ {
			if i%(len(s)/2) == 0 {
				ss = append(ss, s[:i])
				s = s[i:]
				i = 1
			}
		}
		ss = append(ss, s)

		found:=false

		for i := len(line)/2; i < len(line); i++ {
			if (!found) {
				for j := 0; j < len(ss[0]); j++ {
					if (!found && line[i] == ss[0][j]) {
						if (int(line[i]) > 96) {
							total += (int(line[i]) - 96)
						} else {
							total += (int(line[i]) - 38)
						}
						found = true
					}
				}
			}
		}

		if (counter == 1) {
			line1 = line
 		} else if (counter == 2) {
			 line2 = line
		 } else {
			// find the commonality
			newLine := ""
			ans := ""

			for i := 0; i < len(line1); i++ {
				for j := 0; j < len(line2); j++ {
					if (line1[i] == line2[j]) {
						newLine += string(line1[i])
					}
				}
			}

			found2 := false

			for i := 0; i < len(newLine); i++ {
				for j := 0; j < len(line); j++ {
					if (!found2 && newLine[i] == line[j]) {
						ans += string(newLine[i])
						found2 = true
					}
				}
			}
			
			if (int(ans[0]) > 96) {
				total2 += (int(ans[0]) - 96)
			} else {
				total2 += (int(ans[0]) - 38)
			}


			counter = 0
		 }


		 counter++
	}

	fmt.Printf("Answer 1: %d\n", total)
	fmt.Printf("Answer 2: %d\n", total2)
	

}