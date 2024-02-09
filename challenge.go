package main

import "fmt"

const ebulicaoK = 373.15

func main() {
	tempK := ebulicaoK
	tempC := (tempK) - 273.15

	fmt.Println("O valor da conversão do ponto de ebulição da água de graus Kelvin para graus Celsius é:", tempC, "graus")

}
