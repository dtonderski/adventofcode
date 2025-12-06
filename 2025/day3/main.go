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

func getLargestIntOld(line string) int {
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

func getLargestInt(line string, size int) int {
	if (len(line) < size) {
		panic("wtf")
	}

	
	line_ints := make([]int, len(line), len(line))
	arr := make([]int, size, size)

	for i := 0; i < len(line); i++ {
		num, err := strconv.Atoi(string(line[i])) 
		check(err)
		line_ints[i] = num
		if (i < size) {
			arr[i] = num
		}
	}

	for i := size; i < len(line); i++ {
		num := line_ints[i]

		shifting := false
		for j := 0; j < size-1; j++ {
			if arr[j] < arr[j+1] {
				shifting = true
			}

			if shifting {
				arr[j] = arr[j+1]
			}
		}
		if shifting || num > arr[size-1] {
			arr[size-1] = num
		}
	}
	
	s := ""
	for _, num := range arr {
		s += strconv.Itoa(num)
	}
		
	res, err := strconv.Atoi(s)
	check(err)
	return res

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
		// largest := getLargestIntOld(line)
		largest := getLargestInt(line, 12)
		sum += largest
	}
	
	fmt.Printf("answer = %d\n", sum)
	file.Close()

}
