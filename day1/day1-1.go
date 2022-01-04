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

func main() {
	dat, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(dat)

	counter := 0
	last_val := 9999
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		i, err := strconv.Atoi(line)
		if err == nil && i > last_val {
			fmt.Printf("%v increased\n", i)
			counter += 1
		} else {
			fmt.Printf("%v decreased\n", i)
		}
		last_val = i
	}
	fmt.Printf("%v\n", counter)
}
