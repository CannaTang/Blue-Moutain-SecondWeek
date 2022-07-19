package main

import (
	"fmt"
	"sort"
)

var n int
var ans [][]int

type queue struct {
	front int
	back  int
	num   [][]int
}

func (q *queue) push(x []int) {
	q.num = append(q.num, x)
	q.back++
}
func (q *queue) pop() {
	q.front++
}

var q queue

func Read() {
	_, _ = fmt.Scanln(&n)
}
func Init() {
	q.front = 0
	q.back = n / 2
	for i := 1; i <= n/2; i++ {
		var temp = []int{i, i}
		q.num = append(q.num, temp)
	}
}
func Bfs() {
	for {
		if q.front == q.back {
			break
		}
		now := q.num[q.front]
		q.pop()
		if now[0] == n {
			ans = append(ans, now)
			continue
		}

		for i := now[len(now)-1]; now[0]+i <= n; i++ {
			newArray := make([]int, len(now))
			copy(newArray, now)
			newArray = append(newArray, i)
			newArray[0] += i
			q.push(newArray)
		}
	}
}
func Sort() {
	sort.Slice(ans, func(i int, j int) bool {
		len1 := len(ans[i])
		for m := 0; m < len1; m++ {
			if ans[i][m] != ans[j][m] {
				return ans[i][m] < ans[j][m]
			}
		}
		return i < j
	})
}
func output() {
	for i := range ans {
		for j := 1; j < len(ans[i]); j++ {
			if j == 1 {
				fmt.Printf("%d", ans[i][j])
			} else {
				fmt.Printf("+%d", ans[i][j])
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	Read()
	Init()
	Bfs()
	Sort()
	output()
}
