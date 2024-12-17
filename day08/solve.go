package main

import (
	"bufio"
	"fmt"
	"os"
)

type Location struct {
	x, y int
}

func getAntinode(l1 , l2 Location, power int)  (Location) {
	dx, dy := l1.x - l2.x, l1.y - l2.y
	return Location{
		x: l2.x + dx * power,
		y: l2.y + dy * power,
	}
}


func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	table := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		table = append(table, line)
	}
	H := len(table)
	W := len(table[0])
	visited := make([][]bool, H)
	visited2 := make([][]bool, H)
	for i := 0 ; i < H ; i++ {
		visited[i] = make([]bool, W)
		visited2[i] = make([]bool, W)
		for j:=0 ; j < W ; j++ {
			visited[i][j] = false
			visited2[i][j] = false
		}
	}
	locations := make(map[byte][]Location)
	for x := 0 ; x < H ; x++ {
		for y := 0 ; y < W; y++ {
			c:= table[x][y]
			now := Location{x:x, y:y}
			if c != '.' {
				if _, ok := locations[c] ; !ok {
					locations[c] = make([]Location, 0)
				}
				locations[c] = append(locations[c], now)
			}
		}
	}
	for _, locs := range locations {
		N := len(locs)
		for i := 0 ; i < N ; i++ {
			for j := 0; j < N ; j++ {
				if i == j {
					continue
				}
				antinode := getAntinode(locs[i], locs[j], 2)
				X, Y := antinode.x, antinode.y
				if X>= 0 && X < H && Y >= 0 && Y < W {
					visited[X][Y] = true
				}
				for pow := 1 ; ; pow++ {
					exntendedAntinode := getAntinode(locs[i], locs[j], pow)
					X, Y := exntendedAntinode.x, exntendedAntinode.y
					if X>= 0 && X < H && Y >= 0 && Y < W {
						visited2[X][Y] = true
					} else {
						break
					}

				}
			}
		}
	}
	ans1, ans2  := 0, 0
	for i:= 0; i < H ; i ++ {
		for j := 0 ; j < W ; j++ {
			if visited[i][j] {
				ans1 += 1
			}
			if visited2[i][j] {
				ans2 += 1
			}
		}
	}
	fmt.Println(ans1, ans2)

}
