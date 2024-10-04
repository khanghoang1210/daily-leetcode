package main

import "fmt"

// Time: O(N^2) Space: O(N^2)
func generate(numRows int) [][]int {
	res := [][]int{{1}}
	for i := range numRows - 1 {
		i = i
		tmp := append(res[len(res)-1], 0)
		tmp = append([]int{0}, tmp...)
		row := []int{}
		for j := range len(res[len(res)-1]) + 1 {
			row = append(row, tmp[j]+tmp[j+1])
		}
		res = append(res, row)
	}
	return res
}

func main() {
	numRows := 5
	pascalTriangle := generate(numRows)

	for _, row := range pascalTriangle {
		fmt.Println(row)
	}
}