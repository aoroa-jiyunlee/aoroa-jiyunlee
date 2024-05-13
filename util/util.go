package util

import (
	"bufio"
	"fmt"
	"os"
)

func PromptInt(req string) int {
	reader := bufio.NewReader(os.Stdin)

	var input int
	for {
		fmt.Print(req)
		if _, err := fmt.Fscanln(reader, &input); err != nil {
			fmt.Print("입력값을 확인해주세요. ")
			reader.ReadString('\n')
		}
		if input != 0 {
			break
		}
	}
	return input
}

func MaxInt(ints []int) int {
	var max int
	for _, i := range ints {
		if max < i {
			max = i
		}
	}
	return max
}
