package variables

import "fmt" //fmt paketi kullanıldı.

// kullanılmayan paketler silinmeli.

func Demo1() {
	fmt.Print("Merhaba ")
	fmt.Print("Dünya")
	fmt.Println("!")

	//Değişken Tanımlama
	var metin string = "Merhaba Dünya!"
	fmt.Println(metin)

	var kdv int = 10 //sadece tam sayıları tutar
	fmt.Println(kdv)

	var kdv2 float32 = 0.1 //ondalıklı sayıları tutabilir
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	//Değişken Tanımlama Yöntemi
	kdv3 := 25
	fmt.Println(kdv3)
	fmt.Printf("veri tipi: %T\n", kdv3)

	//Mantıksal Veri Tipi
	var durum bool

	var metin1 string = "Kübra"
	var metin2 string = "Ahmet"

	durum = metin1 == metin2 //tek eşittir değer atar. çift eşittir değerlerin eşitliğini sorgular.
	fmt.Println(durum)
}
