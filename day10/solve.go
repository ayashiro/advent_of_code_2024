package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Location struct {
	x, y int
}
func createvisited(H, W int) [][] bool{
	ret := make([][]bool, 0)
	for i := 0; i < H ; i++ {
		ret = append(ret, make([]bool, W))
	}
	return ret
}
func DFS(x, y , H , W int, table,count [][]int, visited [][]bool) {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	if visited[x][y] {
		return
	}
	visited[x][y] = true
	if table[x][y] == 0 {
		count[x][y] += 1
		return
	}
	for i := 0 ; i< 4 ;i ++ {
		nx ,ny := x + dx[i] , y + dy[i]
		if nx < 0 || nx >= H || ny < 0 || ny >= W {
			continue
		}
		if table[nx][ny] == table[x][y] - 1 {
			DFS(nx, ny, H, W, table, count, visited)
		}
	}
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	table := make([][]int, 0)
	count := make([][]int, 0)
	dp := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	valueToLocation := make([][]Location, 10)
	for i := 0 ; i < 10; i++ {
		valueToLocation[i] = make([]Location, i)
	}
	for i:=0; scanner.Scan(); i++{
		line := scanner.Text()
		row := make([]int, 0)
		for j, x:= range line {
			value, _ := strconv.Atoi(string(x))
			row = append(row, value)
			valueToLocation[value] = append(valueToLocation[value], Location{x:i, y:j})
		}
		table = append(table, row)
		count = append(count, make([]int, len(row)))
		dp = append(dp, make([]int, len(row)))
	}
	H := len(table)
	W := len(table[0])
	for i := 0 ; i < H ; i++ {
		for j := 0 ; j < W; j++ {
			if table[i][j] == 9 {
				DFS(i, j, H, W, table, count, createvisited(H, W))
			}
			if table[i][j] == 0 {
				dp[i][j] = 1
			}
		}
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for value := 0 ; value < 9 ; value ++ {

		for _ , location := range valueToLocation[value] {
			x , y := location.x , location.y
			for i := 0 ; i < 4; i ++ {
				nx , ny := x + dx[i], y + dy[i]
				if nx < 0 || nx >= H || ny < 0 || ny >= W {
					continue
				}
				if table[nx][ny] == value + 1  {
					dp[nx][ny] += dp[x][y]
				}
			}
		}
	}
	ans := 0
	ans2 := 0
	for i := 0 ; i < H ; i++ {
		for j := 0 ; j < W; j++ {
			ans += count[i][j]
			if table[i][j] == 9 {
				ans2 += dp[i][j]
			}
		}
	}
	fmt.Println(ans, ans2)


}
