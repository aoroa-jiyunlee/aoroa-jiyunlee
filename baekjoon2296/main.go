package main

import (
	"fmt"
	"github.com/aoroa-jiyunlee/util"
	"sort"
)

type Building struct {
	x   int
	y   int
	val int
}

func getMaxProfit(N int, buildings []Building) int {
	dp := make([]int, N)

	for i := 0; i < N; i++ {
		dp[i] = buildings[i].val
		for j := 0; j < i; j++ {
			if buildings[j].y < buildings[i].y {
				dp[i] = util.MaxInt([]int{dp[i], dp[j] + buildings[i].val})
			}
		}
	}

	return util.MaxInt(dp)
}

func inputBuildings(N int) []Building {
	fmt.Println("각 건물의 좌표와 이익을 입력하세요.")
	buildings := make([]Building, N)
	for n := 0; n < N; n++ {
		x := util.PromptInt(fmt.Sprintf("%d 번 건물의 x 좌표를 입력하세요: ", n+1))
		y := util.PromptInt(fmt.Sprintf("%d 번 건물의 y 좌표를 입력하세요: ", n+1))
		val := util.PromptInt(fmt.Sprintf("%d 번 건물의 이익를 입력하세요: ", n+1))
		buildings[n] = Building{x, y, val}
	}
	return buildings
}

func main() {
	/*
		*** 예시 ***
			N := 4
			buildings := []Building{{1, 1, 2}, {2, 5, 4}, {4, 6, 2}, {5, 2, 5}}
			출력기대값: 9
	*/

	N := util.PromptInt("건물의 개수를 입력하세요: ")
	buildings := inputBuildings(N)

	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i].x > buildings[j].x
	})
	desc := getMaxProfit(N, buildings)

	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i].x < buildings[j].x
	})
	asc := getMaxProfit(N, buildings)

	fmt.Printf("건물 수: %d, 각 건물의 좌표: %v\n", N, buildings)
	fmt.Printf("출력값: %d", util.MaxInt([]int{desc, asc}))
}
