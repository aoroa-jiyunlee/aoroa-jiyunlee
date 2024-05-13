package util

import (
	"bufio"
	"fmt"
	"os"
)

func PromptInt(req string) int {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var input int
	for {
		fmt.Print(req)
		fmt.Fscanln(reader, &input)

		if input != 0 {
			break
		} else {
			fmt.Print("입력값을 확인해주세요. ")
			reader.ReadString('\n')
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
