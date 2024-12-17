package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type block struct {
	start, length, index int
}
func printArr(arr []int) {
	for _, v:= range arr {
		if v < 0 {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", v)
		}
	}
	fmt.Println()
}
func printArrFromFile(files []block, length int) {
	arr := make([]int, length)
	for i := 0 ; i< length; i++ {
		arr[i] = -1
	}
	for _, f := range files {
		for i:= 0; i< f.length; i++ {
			arr[f.start + i] = f.index
		}
	}
	for _, v := range arr {
		if v < 0 {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", v)
		}
	}
	fmt.Println()
}
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
					line = scanner.Text()
	}
	arr := make([]int, 0)
	files := make([]block, 0)
	free := make([]block, 0)
	position := 0
	for i := 0 ; i < len(line) ; i++ {
		x, _ := strconv.Atoi(string(line[i]))
		var t int
		if i % 2 == 0 {
			t = i/ 2
			files = append(files, block{start:position, length: x, index:i / 2})
		} else {
			t = -1
			free  = append(free, block{start:position, length: x, index:-1})
		}
		for cnt := 0 ; cnt < x; cnt ++ {
			arr = append(arr, t)
		}

		position += x
	}
	left := 0
	right := len(arr) - 1
	for arr[left]  >= 0 {
		left ++
	}
	for arr[right] < 0 {
		right --
	}
	for left < right {
		arr[left] , arr[right] = arr[right], arr[left]
		for ;left < len(arr) ; left ++{
			if arr[left] < 0 {
				break
			}
		}
		for ; right >= 0 ; right-- {
			if arr[right] >= 0 {
				break
			}
		}
	}
	for right := len(files) -1 ;right  >= 0 ; right--{
		rightPosition := files[right].start
		rightLength := files[right].length
		for left := 0 ; left < len(free) && free[left].start < rightPosition ; left++ {
			if rightLength <= free[left].length {
				files[right].start = free[left].start
				free[left].start += rightLength
				free[left].length -= rightLength
				break
			}
		}
	}
	ans1 := 0
	for i:= 0 ; i < len(arr); i++ {
		if arr[i] >= 0 {
			ans1 += i * arr[i]
		}
	}
	ans2 := 0
	for _, f := range files  {
		for pos := f.start; pos < f.start + f.length ; pos++ {
			ans2 += pos * f.index
		}
	}
	fmt.Println("ans1 =",ans1, "ans2 =",ans2)
}
