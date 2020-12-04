package main

import (
	recursos "m/src/api"

	"github.com/guregodevo/pastis"
)

func main() {
	circuito := new(recursos.RecursoCircuito)
	escuderia := new(recursos.RecursoEscuderia)
	estadistica := new(recursos.RecursoEstadisticas)
	api := pastis.NewAPI()

	api.AddResource("/api/circuito/:nombre", circuito)
	api.AddResource("/api/escuderia/:nombre", escuderia)
	api.AddResource("/api/estadisticas/:nombreCircuito/:temporada", estadistica)

	api.Start(8080)
}
