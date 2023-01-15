package loops

import "fmt"

//1. Kullanıcıdan bir sayı girmesini isteyiniz.
//23 : Asaldır

func Demo4() {

	sayi := 0
	fmt.Println("Bir sayi giriniz...")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi == true {
		fmt.Println("Asaldir.")
	} else {
		fmt.Println("Asal değildir.")
	}

}
