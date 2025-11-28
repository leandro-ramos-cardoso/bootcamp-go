package main

import "fmt"

func main() {

	defer fmt.Println("Treino finalizado. Goku saiu da Sala do Tempo")

	for treino := 1; treino <= 10; treino++ {

		if treino < 5 {
			fmt.Println("Goku está aquecendo");
		} else if treino >= 5 && treino <= 8 {
			fmt.Println("Goku está aumentando o Ki!")
		} else {
			fmt.Println("Goku está no limite!")
		}

		switch treino {
		case 7:
			fmt.Println("Transformação Super Saiyajin!")
		case 9:
			fmt.Println("Transformação Super Saiyajin 2!")
		default:
			fmt.Println("Treinando normalmente nos outros casos")
		}		
	}
}
