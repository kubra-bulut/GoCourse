package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{20, 30, 50, 10, 2}
	fmt.Println(numbers)

	enBuyuk := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > enBuyuk {
			enBuyuk = numbers[i]
		}
	}
	fmt.Println("En büyük: ", enBuyuk)
}
