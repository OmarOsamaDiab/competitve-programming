package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var input io.Reader = bufio.NewReader(os.Stdin)
	var m int
	fmt.Fscan(input, &m)
	arr := make([][]int, 6)
	for i := 1; i <= 5; i++ {
		arr[i] = make([]int, 6)
	}
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(input, &x, &y)
		arr[x][y] = 1
		arr[y][x] = 1
	}
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == i {
				continue
			}
			for k := 1; k <= 5; k++ {
				if k == i || k == j {
					continue
				}
				sum := arr[i][j] + arr[j][k] + arr[i][k]
				if sum == 3 || sum == 0 {
					fmt.Println("WIN")
					os.Exit(0)
				}
			}
		}
	}
	fmt.Println("FAIL")
}
