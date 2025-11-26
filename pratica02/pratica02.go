package main

import (
	"fmt"
	"math"
)

const SaiyanMultiplier = 1.5

func comparePower(poderGoku int, poderVegeta int) (string, int) {
	
	if poderGoku > poderVegeta {
		return "Goku vence!", poderGoku - poderVegeta
	} else if poderVegeta > poderGoku {
		return  "Vegeta vence!", poderVegeta - poderGoku
	}

	return "Empate perfeito!", 0
}

func main() {
	var poderGokuBase int = 8000
	poderVegetaBase := 5000

	gokuFinal := int(float64(poderGokuBase) * SaiyanMultiplier)
	vegetaFinal := int(float64(poderVegetaBase) * SaiyanMultiplier)

	fmt.Println("Poder final de Goku: ", gokuFinal)
	fmt.Println("Poder final de Vegeta: ", vegetaFinal)

	result, diff := comparePower(gokuFinal, vegetaFinal)

	fmt.Println("Resultado da batalha: ", result)
	fmt.Println("Diferenca do poder: ", diff)

	var kiExtra float64
	fmt.Println("Ki extra não carregado: ", kiExtra)

	fmt.Println("Raiz da diferença: ", math.Sqrt(float64(diff)))
}
