package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner:= bufio.NewScanner(file)
	data := make([][]int, 0)
	for scanner.Scan() {
		line:= scanner.Text()
		cs := strings.Split(line, "   ")

		tmp := make([]int, 0)
		for i := 0 ; i < 2 ; i++ {
			value, _ := strconv.Atoi(cs[i])
			tmp = append(tmp, value)
		}
		data = append(data, tmp)
	}
	left := make([]int, 0)
	right := make([]int, 0)
	right_occurence := make(map[int]int)
	for i:= 0 ; i < len(data); i++ {
		left = append(left, data[i][0])
		right = append(right, data[i][1])
		right_occurence[data[i][1]] += 1
	}
	sort.Ints(left)
	sort.Ints(right)
	ans := 0
	ans2 := 0
	for i := 0 ; i< len(data); i++ {
		v := left[i] - right[i]
		if v < 0 {
			v = -v
		}
		ans += v
		ans2 += right_occurence[left[i]] * left[i]
	}
	fmt.Printf("ans = %d\n", ans)
	fmt.Printf("ans2 = %d\n", ans2)
}
