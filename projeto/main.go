package main

import (
	"fmt"
	"projeto/utils"
)

func statusSistema() string {
	return "Sistema online"
}

func main() {
	status := statusSistema()

	fmt.Println("Status da plataforma: ", status)
	utils.Mensagem("Pedro")
}
