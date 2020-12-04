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
	sesion := new(recursos.RecursoSesion)
	api := pastis.NewAPI()

	api.AddResource("/api/circuito/:nombre", circuito)
	api.AddResource("/api/escuderia/:nombre", escuderia)
	api.AddResource("/api/gp/:nombreCircuito/:temporada/estadisticas", estadistica)
	api.AddResource("/api/piloto/:nombre", piloto)
	api.AddResource("/api/gp/:nombreCircuito/:temporada/sesion/:nombreSesion", sesion)

	api.Start(8080)
}
