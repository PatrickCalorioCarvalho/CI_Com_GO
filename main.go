package main

import (
	"github.com/PatrickCalorioCarvalho/CI_COM_GO/database"
	"github.com/PatrickCalorioCarvalho/CI_COM_GO/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
