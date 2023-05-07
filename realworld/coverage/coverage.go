package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getLastTwoDigits(line string) int {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == '.' || r == ',' || r == ' ' || r == ':'
	})
	times, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		panic(err)
	}
	if times == 0 {
		return 0
	}
	// fmt.Println(fields)
	// fmt.Println(fields[3], fields[5])

	// coverage, err := strconv.Atoi(fields[len(fields)-2])
	coverage1, err := strconv.Atoi(fields[3])
	if err != nil {
		panic(err)
	}
	coverage2, err := strconv.Atoi(fields[5])
	if err != nil {
		panic(err)
	}
	return coverage2 - coverage1
}

// func getLastTwoDigits(line string) int {
// 	fields := strings.FieldsFunc(line, func(r rune) bool {
// 		return r == '.' || r == ',' || r == ' '
// 	})
// 	times, err := strconv.Atoi(fields[len(fields)-1])
// 	if err != nil {
// 		panic(err)
// 	}
// 	if times == 0 {
// 		return 0
// 	}

// 	coverage, err := strconv.Atoi(fields[len(fields)-2])
// 	if err != nil {
// 		panic(err)
// 	}
// 	return coverage
// }
func getRunLine(coverage string) (int, int, int, int) {
	coverage = strings.ReplaceAll(coverage, "\r", "")
	lines := strings.Split(coverage, "\n")
	allCount := 0
	articlesCount := 0
	usersCount := 0
	concatCount := 0
	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == '.' || r == ',' || r == ' ' || r == ':'
		})
		if len(fields) < 3 {
			continue
		}

		times, err := strconv.Atoi(fields[len(fields)-1])
		if err != nil {
			panic(err)
		}
		if times == 0 {
			continue
		}

		names := strings.Split(fields[1], "/")
		articlesLine := 0
		usersLine := 0

		coverage1, err := strconv.Atoi(fields[3])
		if err != nil {
			panic(err)
		}
		coverage2, err := strconv.Atoi(fields[5])
		if err != nil {
			panic(err)
		}

		allCoverage := (coverage2 - coverage1) + 1
		if names[3] == "articles" {
			articlesLine = allCoverage
		} else if names[3] == "users" {
			usersLine = allCoverage
		}

		allCount += allCoverage
		concatCount += (articlesLine + usersLine)
		articlesCount += articlesLine
		usersCount += usersLine
	}
	return allCount, concatCount, articlesCount, usersCount
}

// func getRunLine(coverage string) int {
// 	coverage = strings.ReplaceAll(coverage, "\r", "")
// 	lines := strings.Split(coverage, "\n")
// 	count := 0
// 	for _, line := range lines {
// 		cov := strings.Split(line, " ")
// 		if len(cov) < 3 {
// 			continue
// 		}

// 		times, err := strconv.Atoi(cov[2])
// 		if err != nil {
// 			panic(err)
// 		}
// 		if times == 0 {
// 			continue
// 		}
// 		runLine, err := strconv.Atoi(cov[1])
// 		if err != nil {
// 			panic(err)
// 		}
// 		count += runLine
// 	}
// 	return count
// }

func main() {
	cmd := exec.Command("/bin/sh", "-c", "lsof -i -P -n | grep goc | grep LISTEN")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(out))
	// matches any string that contains an IP address and port number in the format x.x.x.x:y
	re := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d+)`)
	matches := re.FindStringSubmatch(string(out))
	// if len(matches) > 1 {
	// 	fmt.Println(matches[0]) // e.g.: 127.0.0.1:44785
	// }

	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Set the log output to the file
	log.SetOutput(f)
	log.SetFlags(0)
	log.SetPrefix("")
	round := 1

	// Record coverage per second
	for {
		cmd = exec.Command("/bin/sh", "-c", "goc profile --center=\"http://"+matches[0]+"\"")
		out, err = cmd.Output()
		if err != nil {
			panic(err)
		}
		allCount, concatCount, articlesCount, usersCount := getRunLine(string(out))

		log.Printf("%d %d %d %d %d", round, allCount, concatCount, articlesCount, usersCount)
		round++
		time.Sleep(1 * time.Second)
	}

	// filePath := "golang-gin-realworld-example-app/coverage.out"
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// scannar := bufio.NewScanner(file)
	// scannar.Scan()
	// count := 0
	// for scannar.Scan() {
	// 	line := scannar.Text()
	// 	nums := getLastTwoDigits(line)
	// 	count += nums
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println(count)
}
