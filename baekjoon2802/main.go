package main

import (
	"fmt"
	"github.com/aoroa-jiyunlee/util"
)

func findCuttingHeight(N, M int, trees []int) int {
	var min, max = 0, 0

	max = util.MaxInt(trees)

	for min < max {
		var mid = (min + max) / 2
		var sum = 0

		for _, tree := range trees {
			if tree > mid {
				sum += tree - mid
			}
		}

		if sum < M {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min - 1
}

func main() {
	/*
		*** 예시 ***
			N := 4
			M := 7
			trees := []int{10, 15, 20, 17}
			출력기대값: 15
	*/

	N := util.PromptInt("나무의 수를 입력하세요: ")
	M := util.PromptInt("가져가려는 나무의 길이를 입력하세요: ")
	fmt.Println("각 나무의 길이를 입력하세요.")

	trees := make([]int, N)
	for n := 0; n < N; n++ {
		trees[n] = util.PromptInt(fmt.Sprintf("%d 번 나무의 길이를 입력하세요: ", n+1))
	}

	fmt.Printf("나무의수: %d, 가져가려는 나우믜 길이: %d, 각 나무의 길이: %v\n", N, M, trees)
	fmt.Printf("출력값: %d", findCuttingHeight(N, M, trees)) // 15
}
