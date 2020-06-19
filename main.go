package main

import (
	"net/http"

	"github.com/usuário/site/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
