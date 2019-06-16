package matematica

// AULA82 - GERANDO RELATORIOS COM COBERTURA DE TESTES

import (
	"fmt"
	"strconv"
)

// Media calcula a media dos valores
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
