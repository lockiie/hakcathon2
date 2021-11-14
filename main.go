package main

import (
	"fmt"

	//iniciar o servidor e o banco de dados
	_ "github.com/lucas/hackathon/src/db"
	//Inicia a conexão com as rotas
	_ "github.com/lucas/hackathon/src/routers"
)

func main() {
	fmt.Println("Iniciando app")
}
