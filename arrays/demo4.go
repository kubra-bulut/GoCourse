package arrays

import "fmt"

func Demo4() {
	var numbers [2][3]int

	numbers[0][0] = 5
	numbers[0][1] = 3
	numbers[0][2] = 5
	numbers[1][0] = 1
	numbers[1][1] = 4
	numbers[1][2] = 2
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(numbers[i][j])
		}
	}
}
