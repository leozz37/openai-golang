package main

import (
	"fmt"
	"openai-golang/openai"
)

func main() {

	prompt := `Monte uma dieta para mim, com detalhes de calorias, com 4 refeicões:

	Animal: 4 refeicões, diabético, gosto de queijo
	Names: Café da manhã:
			- 2 ovos mexidos com 1 fatia de queijo cheddar
			- 1 fatia de pão de queijo sem glúten
			- 1 xícara de chá verde
			* Aproximadamente 400 calorias
	
		Lanche da manhã:
			- 1 pêra com 2 colheres de sopa de queijo branco
			* Aproximadamente 200 calorias
		
		Almoço:
			- Salada de frutas com mix de folhas verdes, tomate, abacate e nozes
			- Filé de frango grelhado
			- Arroz integral sem glúten
			- 1 xícara de chá de ervas
			* Aproximadamente 400-450 calorias
	
		Jantar:
			- Sopa de legumes com queijo ralado
			- 2 colheres de sopa de quinoa sem glúten
			- Salmão grelhado
			- 1 xícara de chá de camomila
			* Aproximadamente 350-400 calorias
	Animal: 6 refeicões, intolerante a lactose, gosto de frutas
	Names: Café da manhã:
			- Panqueca de aveia e banana
			- 1 xícara de chá verde
			* Aproximadamente 300 calorias
			
		Lanche da manhã:
			- 1 mamão cortado com 1 colher de sopa de granola sem lactose
			* Aproximadamente 200 calorias
			
		Almoço:
			- Salada de frutas com mix de folhas verdes, tomate, abacate e nozes
			- Peixe grelhado
			- Arroz integral
			- 1 xícara de chá de ervas
			* Aproximadamente 400-450 calorias
			
		Lanche da tarde:
			- 1 maçã com 2 colheres de sopa de nozes
			* Aproximadamente 200 calorias
	
		Jantar:
			- Sopa de legumes
			- 2 colheres de sopa de quinoa
			- Frango grelhado
			- 1 xícara de chá de camomila
			* Aproximadamente 350-400 calorias
	
		Ceia:
			- 2 fatias de pão de aveia sem lactose com 1 colher de sopa de pasta de amendoim
			- 1 xícara de chá de ervas
			* Aproximadamente 200-250 calorias
	Animal: 5 refeicões, sou alérgico a chá, gosto de batata, com detalhes de carboidratos
	Names:`

	x, err := openai.Create(prompt, "text-davinci-003", 5000, 0.1)
	if err != nil {
		panic(err)
	}

	fmt.Println(x)
}
