package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekistenen float64 = 1900

	if cekilmekistenen > hesap {
		fmt.Println("hesaptaki para yetersiz")
	} else if cekilmekistenen == hesap {
		fmt.Println("paranız hazırlanıyor.")
	} else {
		fmt.Println("paranız hazırlanıyor.")
	}

}
