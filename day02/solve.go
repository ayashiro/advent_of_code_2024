package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkArray(arr []int) bool {
	N := len(arr)
	for sign := -1; sign <=1 ; sign += 2 {
		flag := true
		for i := 0; (i < N-1) && flag  ; i++ {
			diff := (arr[i+1] - arr[i]) * sign
			if diff < 1 || diff > 3 {
				flag = false
			}
		}
		if flag {
			return true
		}
	}
	return false
}
func checkAbnormal(arr []int, sign int)([]bool, int) {
	diffs := make([]bool, 0)
	cnt := 0
	for i:= 0; i < len(arr) - 1 ; i++ {
		v := (arr[i+1] - arr[i]) * sign
		var t bool
		if v < 1 || v > 3 {
			cnt ++
			t = false
		} else {
			t = true
		}
		diffs = append(diffs, t)
	}
	return diffs, cnt
}
func extractArray(arr []int) bool {
	N := len(arr)
	for _, sign := range []int{1, -1} {
		diffs, errors := checkAbnormal(arr, sign)
		if errors == 0 {
			return true
		}
		if errors > 2 {
			continue
		}
		for i := 0 ; i < N ; i++ {
			if i == 0 {
				// check if removing the first element will be valid.
				if errors == 1 && !diffs[0] {
					return true
				}
			} else if i == N-1 {
				// check if removing the last element will be valid.
				if errors == 1 && !diffs[N-2] {
					return true
				}
			} else {
				d := (arr[i+1] - arr[i-1]) * sign
				if d < 1 || d > 3 {
					// check if the diffs between before and after the specific element are valid
					continue
				}
				if diffs[i-1] && diffs[i] {
					// check if the element is not worth to omit.
					continue
				}
				if (!diffs[i-1]) && (!diffs[i]) {
					return true
				}
				if !(diffs[i-1] && diffs[i]) && errors == 1 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([][]int, 0)
	ans1 := 0
	ans2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		tmp := make([]int, 0)
		for _, v := range strings.Split(line, " ") {
			c, _ := strconv.Atoi(v)
			tmp = append(tmp, c)
		}
		// O(N)
		if checkArray(tmp) {
			ans1 +=1
		}
		//O(N) -- O(N^2) solution is easy to implement
		if extractArray(tmp) {
			ans2 += 1
		}
		data = append(data, tmp)
	}
	fmt.Printf("ans1 = %d ans2 = %d \n", ans1, ans2)
}
