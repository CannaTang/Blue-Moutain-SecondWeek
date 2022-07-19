package main

import "fmt"

var mp [500][500]int
var vis [500][500]int
var n, m, x, y int
var turn = [][]int{
	{2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	{-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

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
	fmt.Scanf("%d %d %d %d\n", &n, &m, &x, &y)
}

func Init() {
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			mp[i][j] = -1
		}
	}
	q.front = 0
	q.back = 1
	var temp = []int{x, y, 0}
	q.num = append(q.num, temp)
}

func judge(x []int) bool {
	return x[0] >= 1 && x[0] <= n && x[1] >= 1 && x[1] <= m && vis[x[0]][x[1]] == 0
}

func Bfs() {
	for {
		if q.front == q.back {
			break
		}

		now := q.num[q.front]
		q.pop()

		mp[now[0]][now[1]] = now[2]

		for i := 0; i < 8; i++ {
			newArray := make([]int, len(now))
			copy(newArray, now)

			newArray[0] += turn[i][0]
			newArray[1] += turn[i][1]
			newArray[2] += 1

			if judge(newArray) {
				vis[newArray[0]][newArray[1]] = 1
				q.push(newArray)
			}
		}
	}
}

func Output() {
	mp[x][y] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%-5d", mp[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	Read()
	Init()
	Bfs()
	Output()
}
