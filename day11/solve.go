package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNext(n int) []int {
	ret := make([]int, 0)
	s  := strconv.Itoa(n)
	if n == 0 {
		ret = append(ret, 1)
	} else if  len(s) % 2  == 0 {
		N := len(s) / 2
		left, right := s[:N], s[N:]
		l , _  := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)
		ret = append(ret, l)
		ret = append(ret, r)

	} else {
		ret = append(ret, n * 2024)
	}
	return ret
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	current := make(map[int]int)
	for _, v:= range strings.Split(line, " ") {
		value, _ := strconv.Atoi(v)
		current[value] ++
	}
	ans1 := 0
	ans2 := 0
	for i := 0; i< 75 ; i ++ {
		nxt := make(map[int]int)
		for k, v:= range current {
			for _, x := range getNext(k) {
				nxt[x] += v
			}
		}
		current = nxt
		if i+1 == 25 || i+1 == 75 {
			for _, v := range current {
				if i + 1 == 25 {
					ans1 += v
				}
				if i + 1 == 75 {
					ans2 += v
				}
			}
		}
	}
	fmt.Println(ans1, ans2)
}
