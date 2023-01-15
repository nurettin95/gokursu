package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {

	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasinda bir sayi giriniz.")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayi giriniz.", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayi giriniz.", nil
	}

	return "Bildiniz..", nil
}

func Demo2() {
	mesaj, hata := TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(hata)
	// mesaj2, _ := TahminEt(50) nil yazdırmak istemezsek
	// fmt.Println(mesaj2)
}
