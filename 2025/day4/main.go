package main

import (
	"log"
	"time"
	"os"
	"bufio"
	"fmt"
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

type offset struct {
	x,y int
}

func main() {
	log.Printf("starting")
	start_time := time.Now()
	defer timeTrack(start_time, "execution")
	file, err := os.Open(file)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)	

	rune_arr := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		rune_arr = append(rune_arr,  []rune(line))
	}

	offsets := []offset{ offset{-1, -1},
		offset{-1, 0},
		offset{-1, 1},
		offset{0,-1},
		offset{0,1},
		offset{1,-1},
		offset{1,0},
		offset{1,1},
	}

	totSum := 0

	for true {
		neighbor_arr := make([][]int, len(rune_arr))	
		for i, _ := range neighbor_arr {
			neighbor_arr[i] = make([]int, len(rune_arr[i]))
		}



		for i, row := range rune_arr {
			for j, cell := range row { 
				update_neighbors(i, j, cell, neighbor_arr, offsets)			
			}
		}

		sum := 0

		for i, row := range neighbor_arr {
			for j, cell := range row {
				if cell < 4 && rune_arr[i][j] == '@' {
					rune_arr[i][j] = 'x'
					sum += 1
				}

			}
		}
		totSum += sum
		// remove || true for day4.2
		if sum == 0 || true {
			break
		}
	}
	fmt.Printf("TotSum: %d\n", totSum)

}

func update_neighbors(i int, j int, char rune, neighbor_arr [][]int, offsets []offset) {
	max_row := len(neighbor_arr)
	max_col := len(neighbor_arr[0])

	if char != '@' {
		return
	}

	for _, offset := range offsets {
		new_x := i + offset.x
		new_y := j + offset.y

		if (new_x >= max_row || new_x < 0) {
			continue
		}

		if new_y >= max_col || new_y < 0 {
			continue
		}
		neighbor_arr[new_x][new_y] += 1
	}
}
