package main

import (
	"net/http"

	"github.com/usuário/site/gofirstweb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
