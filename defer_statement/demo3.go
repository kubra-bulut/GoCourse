package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "laptop", unitPrice: 5000}
	defer p.save()
	fmt.Println("İşlem başarılı")
}
