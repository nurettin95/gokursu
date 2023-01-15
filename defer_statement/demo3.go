package deferstatement

import "fmt"

type Product struct {
	productName string
	unitPrice   int
}

func (p Product) save() {
	fmt.Println("Ürün kaydedildi : ", p.productName)
	defer Log() //Log'un her zaman çalışmasını istiyorum
}

func Demo3() {

	p := Product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = Product{productName: "Mouse", unitPrice: 45}

	fmt.Println("İşlem başarili")
	fmt.Println(p.productName)
}

func Log() {
	fmt.Println("Loglandi.")
}
