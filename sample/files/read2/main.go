package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	lines := make(map[int]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		lines[number] = struct{}{}
	}

	result := make([]int, 0, len(lines))
	for i := range lines {
		result = append(result, i)
	}
	sort.Ints(result)
	fmt.Println(result)
}
