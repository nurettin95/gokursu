package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz.")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paraniz hazirlaniyor")
		fmt.Println("Dikkat, hesapta para kalmadi")
	} else {
		fmt.Println("Paraniz hazirlaniyor")
	}

}
