package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func sort(arr []int64, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	//fmt.Println(mid)
	sort(arr, l, mid)
	sort(arr, mid+1, r)
	L := arr[l : mid+1]
	R := arr[mid+1 : r+1]
	//fmt.Println("L:", L)
	//fmt.Println("R:", R)
	i, j, idx := 0, 0, 0
	b := make([]int64, r-l+1)
	for i < len(L) && j < len(R) {
		if L[i] > R[j] {
			b[idx] = L[i]
			i++
		} else {
			b[idx] = R[j]
			j++
		}
		idx++
	}
	for i < len(L) {
		b[idx] = L[i]
		idx++
		i++
	}
	for j < len(R) {
		b[idx] = R[j]
		idx++
		j++
	}
	idx = 0
	for k := l; k <= r; k++ {
		arr[k] = b[idx]
		idx++
	}
	//fmt.Println(arr)
}

func main() {
	var input io.Reader = bufio.NewReader(os.Stdin)
	var test int
	fmt.Fscan(input, &test)
	const N = 200005
	in := make([]int64, N)
	original := make([]int64, N)
	originalSort := make([]int64, N)
	for i := 0; i < test; i++ {
		var n int
		fmt.Fscan(input, &n)
		for j := 0; j < n; j++ {
			fmt.Fscan(input, &in[j])
			original[j] = in[j]
		}
		sort(in[0:n], 0, n-1)
		for j := 0; j < n; j++ {
			originalSort[j] = in[j]
		}
		for j := n - 2; j >= 0; j-- {
			in[j] += in[j+1]
		}
		lst := originalSort[0]
		for j := 1; j < n; j++ {
			if in[j] >= originalSort[j-1] {
				lst = originalSort[j]
			} else {
				break
			}
		}
		cnt := 0
		for j := 0; j < n; j++ {
			if original[j] >= lst {
				cnt++
			}
		}
		fmt.Println(cnt)
		for j := 0; j < n; j++ {
			if original[j] >= lst {
				fmt.Println(j + 1)
			}
		}
		//fmt.Println(in[0:n])
		//fmt.Println(originalSort[0:n])
	}
}
