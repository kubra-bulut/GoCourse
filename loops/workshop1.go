package loops

import "fmt"

//bir değişkende sayı tut.
//fmt.Scanln ile kullanıcıdan sayı alınabilir.
func Demo3() {
	var aklimdakiSayi int = 80
	tahminEdilenSayi := 0
	tahminSayisi := 0
	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi++
	for aklimdakiSayi != tahminEdilenSayi {

		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha büyük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi++
		} else if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi++
		}
	}

	fmt.Println("doğru sayı")
	fmt.Printf("Tahmin sayınız %v", tahminSayisi)
}
