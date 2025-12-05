package main

import (
	"time"
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
	"log"
)

const (
	file = "input"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func consists_of(str string, part string) bool {
	if len(part) == 0 {
		return false
	}

	if len(str) % len(part) != 0 {
		return false
	}

	for i := 0; i < len(str); i++ {
		if str[i] != part[i%len(part)] {
			return false
		}
	}
	return true
}

func is_illegal_string(s string, exactly_twice bool) bool {
	if (exactly_twice) {
		if len(s) % 2 != 0 {
			return false
		}
		if consists_of(s, s[0:len(s)/2]) {
			fmt.Println(s, s[0:len(s)/2])
			return true
		}
		return false
	}

	for i, _ := range s {
		if consists_of(s, s[0:i]) {
			return true
		}
	}
	return false
}

type interval struct {
	start, stop int
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

func main() {
	log.Printf("starting")
	start_time := time.Now()
	defer timeTrack(start_time, "execution")
	file, err := os.Open(file)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string

	for scanner.Scan() {
		line = scanner.Text()
	}

	file.Close()

	range_strings := strings.Split(line, ",")
	intervals := make([]interval, 0)
	for _, range_string := range range_strings {
		split_range := strings.Split(range_string, "-") 

		start, err := strconv.Atoi(split_range[0])
		check(err)

		stop, stop_err := strconv.Atoi(split_range[1])
		check(stop_err)

		intervals = append(intervals, interval {start, stop})

	}

	numbers := make([]string, 0)
	for _, interval := range intervals {
		cur := interval.start
		for cur <= interval.stop {
			numbers = append(numbers, strconv.Itoa(cur))
			cur ++
		}
	}

	sum := 0
	for _, number := range numbers {
		if (is_illegal_string(number, false))	{
			num, err := strconv.Atoi(number)

			check(err)
			sum += num
		}
	}
	fmt.Printf("Answer: %d\n", sum)
}

