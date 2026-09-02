package main

import "fmt"

func main() {
	weight := 115.0
	height := 1.78

	imt := weight / (height * height)

	fmt.Print("Ваш ИМТ = ", imt)
}