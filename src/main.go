package main

import (
	recursos "m/src/api"

	"github.com/guregodevo/pastis"
)

func main() {
	circuito := new(recursos.RecursoCircuito)
	api := pastis.NewAPI()
	api.AddResource("/api/circuito/:nombre", circuito)
	api.Start(8080)
}
