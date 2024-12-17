package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func checkOrder(v1, v2 int , order map[int]map[int]bool) bool{
	if _, ok:= order[v1]; ok {
		return order[v1][v2]
	}
	return false
}
func main() {
	pattern, _ := regexp.Compile(`^(\d+)\|(\d+)$`)
	order := make(map[int]map[int]bool)
	filename := os.Args[1]
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matcher := pattern.FindAllStringSubmatch(line, -1)
		if len(matcher) == 0 {
			break
		}
		v1, _ := strconv.Atoi(matcher[0][1])
		v2, _ := strconv.Atoi(matcher[0][2])
		if order[v1] == nil {
			order[v1] = make(map[int]bool)
		}
		order[v1][v2] = true
	}
	ans := 0
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		arr := make([]int, 0)
		for _, elem :=  range strings.Split(line, ",") {
				value, _ := strconv.Atoi(elem)
				arr = append(arr, value)
		}
		flag := true
		for i := 0 ; i < len(arr) - 1  && flag ; i ++  {
			if !checkOrder(arr[i], arr[i+1], order){
				flag = false
			}
		}
		if flag {
			ans += arr[len(arr)/2]
		} else {
			sort.Slice(arr, func(i, j int) bool {
				return checkOrder(arr[i], arr[j], order)
			})
			ans2 += arr[len(arr)/2]
		}
	}
	fmt.Println("ans1 =", ans, "ans2 =", ans2)
}
