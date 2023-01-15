package error_handling

import (
	"fmt"
	"os"
)

// type assertion
func Demo1() {
	f, err := os.Open("demo11.txt")
	//dosyayı bulursa err nil olur. Hata yoksa nil, varsa nil'den farklı bir şey hata var demektir.
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok { //hatanın türü pathError ise ok true oluyor hata pErr'e atanıyor, değilse false
			//if true ise
			fmt.Println("Dosya bulunumadi", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadi.") //Genel bir hata oluştu
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
