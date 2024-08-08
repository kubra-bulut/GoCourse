package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayisi: ", len(sozluk))
	fmt.Println("Sözlük: ", sozluk)
	//eleman silme
	delete(sozluk, "book")
	fmt.Println("Eleman sayisi: ", len(sozluk))

	fmt.Println("Sözlük: ", sozluk)

	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu: ", varMi)
}
