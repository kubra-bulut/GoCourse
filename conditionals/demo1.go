package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekistenen float64 = 900

	if cekilmekistenen > hesap {
		fmt.Println("hesaptaki para yetersiz")
	}
	if cekilmekistenen <= hesap {
		fmt.Println("paranız hazırlanıyor.")
		hesap = hesap - cekilmekistenen
	}
	fmt.Println("bitti. Hesaptaki para: " + fmt.Sprintf("%v", hesap))
}
