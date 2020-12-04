package main

import (
	recursos "m/src/api"

	"github.com/guregodevo/pastis"
)

func main() {
	circuito := new(recursos.RecursoCircuito)
	escuderia := new(recursos.RecursoEscuderia)
	estadistica := new(recursos.RecursoEstadisticas)
	piloto := new(recursos.RecursoPiloto)
	api := pastis.NewAPI()

	api.AddResource("/api/circuito/:nombre", circuito)
	api.AddResource("/api/escuderia/:nombre", escuderia)
	api.AddResource("/api/estadisticas/:nombreCircuito/:temporada", estadistica)
	api.AddResource("/api/piloto/:nombre", piloto)

	api.Start(8080)
}
