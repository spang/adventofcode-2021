package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func valid(window []int) bool {
	return len(window) == 3
}

func sum(window []int) int {
	s := 0
	for _, v := range window {
		s += v
	}
	return s
}

func main() {
	dat, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(dat)

	counter := 0
	window := make([]int, 0, 3)
	fmt.Println(window)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		i, err := strconv.Atoi(line)
		check(err)
		if valid(window) {
			last_sum := sum(window)
			window = append(window, i)
			if len(window) > 3 {
				window = window[1:]
			}
			current_sum := sum(window)
			if current_sum > last_sum {
				counter += 1
				fmt.Printf("%v increased\n", window)
			} else {
				fmt.Printf("%v decreased\n", window)
			}
		} else {
			fmt.Printf("Window %v not yet valid\n", window)
			window = append(window, i)
		}
	}
	fmt.Printf("%v\n", counter)
}
