package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLastTwoDigits(line string) int {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == '.' || r == ',' || r == ' '
	})
	times, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		panic(err)
	}
	if times == 0 {
		return 0
	}

	coverage, err := strconv.Atoi(fields[len(fields)-2])
	if err != nil {
		panic(err)
	}
	return coverage
}

func main() {
	filePath := "golang-gin-realworld-example-app/coverage.out"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scannar := bufio.NewScanner(file)
	scannar.Scan()
	count := 0
	for scannar.Scan() {
		line := scannar.Text()
		nums := getLastTwoDigits(line)
		count += nums
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(count)
}
