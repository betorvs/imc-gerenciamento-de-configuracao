package main

import (
	"fmt"
	"strconv"
)

func main() {
	var peso, altura string

	fmt.Print("Entre com seu peso (kg): ")
	fmt.Scanf("%s", &peso)
	fmt.Print("Entre com sua altura (metros): ")
	fmt.Scanf("%s", &altura)
	p, _ := strconv.ParseFloat(peso, 64)
	a, _ := strconv.ParseFloat(altura, 64)
	a2 := a * a
	imc := p / a2
	var indice string
	switch {
	case imc < 18.5:
		indice = "magreza"
	case imc > 18.5 && imc < 24.9:
		indice = "normal"
	case imc > 25.0 && imc < 29.9:
		indice = "sobrepeso"
	case imc > 30.0 && imc < 34.9:
		indice = "obesidade grau I"
	case imc > 35.0 && imc < 39.9:
		indice = "obesidade grau II"
	case imc > 40.0:
		indice = "obesidade grau III"
	}
	fmt.Printf("IMC %.4f\nClassificado como %s\n", imc, indice)
}
