package conditionals

import "fmt"

func Demo3() {
	//üç adet int değişken tanımla
	//ekrana en büyük olanı yazdır.

	var sayi1, sayi2, sayi3 int = 10, 5, 18

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En Büyük sayi : %v", enBuyuk)

}
