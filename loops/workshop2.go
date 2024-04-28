package loops

import "fmt"

func Demo4() {
	var girilenSayi int
	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&girilenSayi)

	asalMi := true
	for i := 2; i < girilenSayi; i++ {
		if girilenSayi%i == 0 {
			asalMi = false
		}
	}
	if asalMi == true {
		fmt.Println("Asaldır.")
	} else {
		fmt.Println("değildir.")
	}
}
