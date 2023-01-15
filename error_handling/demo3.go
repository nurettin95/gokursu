package error_handling

import "fmt"

type borderException struct {
	parameter int
	message   string
}

// Error sınıfını implemente ettik
func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{parameter: tahmin, message: "Sinirlarin dişinda kaldi."}
	}
	return "Bildiniz", nil
}
