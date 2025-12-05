package main

import (
	"bufio"
	"fmt"
	"os"
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

func floorDiv(a, b int) int {
	if a < 0 && a%b != 0 {
		return a/b - 1
	}
	return a / b
}

func move_dial_and_check_password_increment(dial *int, password *int, input string, new_algo bool) {
	abs_change, _ := strconv.Atoi(input[1:])

	target := *dial
	if input[0] == 'L' {
		target -= abs_change
	} else {
		target += abs_change
	}

	if !new_algo {
		*dial = target
		if *dial%100 == 0 {
			*password += 1
		}
		return
	}

	if target > *dial {
		next_boundary := (floorDiv(*dial, 100) + 1) * 100
		for next_boundary <= target {
			*password++
			next_boundary += 100
		}
	} else {
		next_boundary := floorDiv(*dial-1, 100) * 100
		for next_boundary >= target {
			*password++
			next_boundary -= 100
		}
	}

	*dial = target
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
		move_dial_and_check_password_increment(&dial_pos, &password, line, true)
	}

	fmt.Printf("password = %d\n", password)
}
