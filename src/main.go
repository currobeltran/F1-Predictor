package main

import (
	recursos "m/src/api"

	"github.com/guregodevo/pastis"
)

func main() {
	circuito := new(recursos.RecursoCircuito)
	escuderia := new(recursos.RecursoEscuderia)
	api := pastis.NewAPI()

	api.AddResource("/api/circuito/:nombre", circuito)
	api.AddResource("/api/escuderia/:nombre", escuderia)

	api.Start(8080)
}
