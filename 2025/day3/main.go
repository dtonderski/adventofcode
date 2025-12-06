package main

import (
	"fmt"
	"log"
	"time"
	"os"
	"bufio"
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

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func getLargestInt(line string) int {
	first, second := 0, 0
	for _, c := range line {
		num, err := strconv.Atoi(string(c))
		check(err)
		if (second > first) {
			first = second
			second = num
		} else if (num > second) {
			second = num
		}
	}
	return 10*first + second
}

func main() {
	log.Printf("starting")
	start_time := time.Now()
	defer timeTrack(start_time, "execution")
	file, err := os.Open(file)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)	
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		largest := getLargestInt(line)
		sum += largest
	}
	
	fmt.Printf("answer = %d\n", sum)
	file.Close()

}
