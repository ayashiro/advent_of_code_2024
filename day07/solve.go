package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func canReach(target, current, ind int, elements []int) bool{
	if ind == len(elements)  {
		return target == current
	}
	return canReach(target, current + elements[ind], ind + 1, elements )||
	 canReach(target, current * elements[ind], ind +1 , elements)
}
func canReach2(target, current, ind int, elements []int) bool{
	if ind == len(elements)  {
		return target == current
	}
	 concated, _ := strconv.Atoi(strconv.Itoa(current) +strconv.Itoa(elements[ind]))

	if (canReach2(target, current + elements[ind], ind + 1, elements )||
	 canReach2(target, current * elements[ind], ind +1 , elements)) {
		return true
	 }
	 return canReach2(target, concated, ind + 1, elements)
}
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ans1 := 0
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, ": ")
		target, _ := strconv.Atoi(tmp[0])
		values := strings.Split(line, " ")
		arr := make([]int, 0)
		for _, value := range(values[1:]) {
			numericValue, _ := strconv.Atoi(value)
			arr = append(arr, numericValue)
		}
			if canReach(target, arr[0], 1, arr) {
				ans1 += target
			}
			if canReach2(target, arr[0], 1 , arr) {
				ans2 += target
			}

	}
	fmt.Println("ans1 =", ans1, "ans2 =", ans2)
}
