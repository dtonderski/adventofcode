package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

const (
	file = "input"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	file, err := os.Open(file) 
	check(err)
	defer file.Close()

	dial_pos := 50
	password := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		abs_change, err := strconv.Atoi(line[1:])
		check(err)
		if (line[0] == 'L') {
			dial_pos -= abs_change
		} else {
			dial_pos += abs_change
		}

		if dial_pos%100 == 0 { 
			password ++
		}
	}
	
	fmt.Printf("password = %d\n", password)
}


