package iteracao

// Repetir recebe um caractere e um número de repetições e retorna o carectere repetido
func Repetir(caractere string, numeroRepeticoes int) string {
	var repeticoes string
	for i := 0; i < numeroRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
