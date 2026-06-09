package main

import "fmt"

func main() {
	var capitais map[string]string = make(map[string]string)

	capitais["Brasil"] = "Brasilia"
	capitais["Franca"] = "Paris"
	capitais["Italia"] = "Roma"
	capitais["Japao"] = "Tokio"

	fmt.Println(capitais)

	// Mudando o nome de um valor
	capitais["Japao"] = "Teste"
	fmt.Println(capitais)
}
