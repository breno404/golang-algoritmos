package inversao

func Exec(esquerda *int, direita *int) {
	var temp int = *esquerda
	*esquerda = *direita
	*direita = temp
}
