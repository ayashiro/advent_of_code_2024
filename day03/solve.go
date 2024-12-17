package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	pattern, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	filename := os.Args[1]
	file, _ := os.Open(filename)
	defer file.Close()
	reader := bufio.NewScanner(file)
	statements := make([]string, 0)
	for reader.Scan() {
		statements = append(statements, reader.Text())
	}
	val := strings.Join(statements, "")
	lines := pattern.FindAllStringSubmatch(val, -1)
	enable := true
	ans1 := 0
	ans2 := 0
	for _,line  := range lines {
		if line[0] == "do()" {
			enable = true
		} else if line[0] == "don't()" {
			enable = false
		} else {
			v1, _ := strconv.Atoi(line[1])
			v2, _ := strconv.Atoi(line[2])
			ans1 += v1 * v2
			if enable {
				ans2 += v1 * v2
			}
		}
	}
	fmt.Println("ans1 =",ans1, "ans2 =",ans2)

}
