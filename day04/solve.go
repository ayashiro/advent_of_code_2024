package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)
	defer file.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	base := "XMAS"
	base2 := "MAS"
	dx := []int {-1, -1,  1,  1, 1, -1,  0,  0}
	dy := []int {-1,  1, -1,  1, 0,  0,  1, -1}
	H := len(lines)
	W := len(lines[0])
	count := make([][]int, H)
	for i := 0 ; i < H ; i++ {
		count[i] = make([]int, W)
	}
	ans1 := 0
	for i := 0 ; i< H; i ++ {
		for j := 0 ; j < W ; j ++ {
			for d:= 0; d < 8; d++ {
				flag := true

				for ind := 0 ; ind < 4 && flag; ind++ {
					x, y := i + dx[d] * ind , j + dy[d] * ind
					if x < 0 || x >= H || y < 0 || y >= W {
						flag = false
						continue
					} else if base[ind] != lines[x][y] {
						flag = false
					}
				}
				if flag {
					ans1 += 1
				}
				if d >= 4 {
					continue
				}
				flag = true
				for ind := 0 ; ind < 3; ind ++ {
					x, y := i + dx[d] * ind , j + dy[d] * ind
					if x < 0 || x >= H || y < 0 || y >= W {
						flag = false
						continue
					} else if base2[ind] != lines[x][y] {
						flag = false
					}
				}
				if flag {
					count[i + dx[d]][j+dy[d]] += 1
				}

			}
		}
	}
	ans2 :=0
	for i := 0 ; i < H ; i++ {
		for j:= 0 ;j < W ; j++ {
			if count[i][j] == 2 {
				ans2 += 1
			}
		}
	}
	fmt.Println("ans1 =", ans1, "ans2 =", ans2)

}
