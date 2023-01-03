package main

import "fmt"

/*
[[1,2,3],
 [4,5,6],
 [7,8,9]]
[1,2,3,6,9,8,7,4,5]
*/

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

type pair struct {
	i int
	j int
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	visited := make(map[pair]bool)
	res := make([]int, 0)
	i := 0
	j := 0
	var goRight, goLeft, goDown, goUp bool
	goRight = true
	for true {
		if len(visited) == len(matrix)*len(matrix[0]) {
			break
		}

		fmt.Println(i, j)
		res = append(res, matrix[i][j])
		visited[pair{
			i: i,
			j: j,
		}] = true

		if goRight {
			_, ok := visited[pair{
				i: i,
				j: j + 1,
			}]
			if j+1 >= len(matrix[0]) || ok {
				i++
				goRight = false
				goDown = true
				continue
			}

			j++
		}

		if goDown {
			_, ok := visited[pair{
				i: i + 1,
				j: j,
			}]
			if i+1 >= len(matrix) || ok {
				j--
				goDown = false
				goLeft = true
				continue
			}
			i++
		}

		if goLeft {
			_, ok := visited[pair{
				i: i,
				j: j - 1,
			}]
			if j-1 < 0 || ok {
				i--
				goLeft = false
				goUp = true
				continue
			}
			j--
		}

		if goUp {
			_, ok := visited[pair{
				i: i - 1,
				j: j,
			}]
			if i-1 < 0 || ok {
				goUp = false
				goRight = true
				continue
			}
			i--
		}
	}
	return res
}

/*
if goright {
			if j + 1 >= len(matrix[0]) {
				goright = false
				godown = true
				continue
			}
			if k, ok := visited[pair{
				i: i,
				j: j+1,
			}]; ok && k {
				goright = false
				godown = true
				continue
			}
			res = append(res, matrix[i][j + 1])
			j++
			continue
		}
		if godown {
			if i + 1 >= len(matrix) {
				godown = false
				goleft = true
				continue
			}
			if k, ok := visited[pair{
				i: i+1,
				j: j,
			}]; ok && k {
				godown = false
				goleft = true
				continue
			}
			res = append(res, matrix[i+1][j])
			i++
			continue
		}
		if goleft {
			if j - 1 < 0 {
				goright = false
				godown = true
				continue
			}
			if k, ok := visited[pair{
				i: i,
				j: j+1,
			}]; ok && k {
				goright = false
				godown = true
				continue
			}
			res = append(res, matrix[i][j + 1])
			continue
		}
*/
