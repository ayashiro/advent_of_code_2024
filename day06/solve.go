package main

import (
	"bufio"
	"fmt"
	"os"
)

type  location struct  {
	x, y, d int
}

func getThreeDimentionVisited(H, W int)[][][]bool {
	ret := make([][][]bool, 0)
	for i := 0 ; i< H ; i++ {
		ret = append(ret, make([][]bool, 0))
		for j := 0;j < W ; j++ {
			ret[i] = append(ret[i], make([]bool, 4))
		}
	}
	return ret
}

func checkLoop(H , W int, table [][]int, sx, sy, ox, oy int)  bool {
	visited := getThreeDimentionVisited(H, W)
	dx , dy := []int{-1,0, 1,0}, []int{0, 1, 0, -1}
	qu := make([]location, 0)
	qu = append(qu, location{x: sx, y:sy, d:0})
	for len(qu) > 0 {
		now := qu[0]
		qu = qu[1:]
		x, y , d := now.x, now.y , now.d
		if visited[x][y][d] {
			return true
		}
		visited[x][y][d] = true
		for i := 0 ; i < 4 ; i++ {
			cd := (d + i) % 4
			nx , ny := x + dx[cd], y + dy[cd]
			if nx < 0 || nx >= H  || ny < 0 || ny >= W {
				break
			}
			if table[nx][ny] == 1 && (nx != ox || ny != oy) {
				qu = append(qu, location{x:nx, y :ny , d:cd})
				break
			}
		}
	}
	return false
}
func main() {
	filename:=os.Args[1]
	file, _ := os.Open(filename)
	defer file.Close()
	table := make([][]int, 0)
	visited := make([][]int, 0)
	visited2 := make([][][]bool, 0)
	scanner := bufio.NewScanner(file)
	var x, y int
	x = 0
	y = 0

	for i:= 0; scanner.Scan(); i++{
		line := scanner.Text()
		row := make([]int, 0)
		row2 := make([]int, 0)
		row3 := make([][]bool, 0)
		for j, v := range line {
			if v == '^' {
				x, y = i, j
			}
			if v == '#' {
				row = append(row, 0)
			} else {
				row = append(row, 1)
			}
			row2 = append(row2, 0)
			row3 = append(row3, make([]bool, 4))
		}
		table = append(table, row)
		visited = append(visited, row2)
		visited2 = append(visited2, row3)
	}
	d := 0
	dx , dy := []int{-1,0, 1,0}, []int{0, 1, 0, -1}
	H := len(table)
	W := len(table[0])
	ans1 := 0
	ans2 := 0
	sx, sy := x, y
	for {
		visited[x][y] = 1
		nx, ny := x + dx[d], y + dy[d]
		if nx < 0 || nx >= H || ny < 0 || ny >= W {
			break
		}
		if table[nx][ny] == 0 {
			d = (d + 1 ) % 4
			nx, ny = x + dx[d], y + dy[d]
		}
		x, y = nx , ny
	}
	for x := 0 ; x < H; x++ {
		for y := 0 ; y < W ; y++ {
			if visited[x][y] > 0{
				ans1 += 1
				if x != sx || y != sy {
					if checkLoop(H,W, table,sx, sy, x, y) {
						ans2 += 1
					}
				}
			}
		}
	}
	fmt.Println("ans1 =", ans1, "ans2 =", ans2)
}
