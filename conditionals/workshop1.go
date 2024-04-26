package conditionals

import "fmt"

func Demo3() {
	//Senaryo: 3 adet int değişken tanımlanacak ve en büyük olan ekrana yazdırılacak.
	var num1 = 1
	var num2 = 2
	var num3 = 3

	var enBuyuk = num1
	if num2 > enBuyuk {
		enBuyuk = num2
	}
	if num3 > enBuyuk {
		enBuyuk = num3
	}

	fmt.Printf("En büyük sayı: %v", enBuyuk)
}
