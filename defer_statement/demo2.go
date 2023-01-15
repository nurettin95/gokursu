package deferstatement

import "fmt"

func Demo2(sayi int) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		fmt.Println("Çift sayi çalisti.")
		return "Çift sayidir."
	}

	if sayi%2 != 0 {
		fmt.Println("Tek sayi çalisti.")
		return "Tek sayidir."
	}

	return "Belli değil"
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
