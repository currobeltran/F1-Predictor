package handler

import (
	"encoding/json"
	"fmt"
	f1predictor "m/src/f1predictor" /*Import para traer paquete f1predictor*/
	"net/http"
)

/*Response : Struct que contendrá la respuesta en formato JSON*/
type Response struct {
	Pole string `json:"Pole"`
}

var hamilton f1predictor.Piloto
var melbourne f1predictor.Circuito
var resultados []f1predictor.ResultadoGP
var resultado f1predictor.ResultadoGP

/*Handler : Controla la llegada de una petición HTTP y la creación de una respuesta*/
func Handler(w http.ResponseWriter, r *http.Request) {
	hamilton.Constructor("Lewis Hamilton", 94, 92, 55, 7)
	// bottas.Constructor("Valtteri Bottas", 10, 12, 5, 0)
	// vettel.Constructor("Sebastian Vettel", 55, 32, 30, 4)

	melbourne.Constructor("Melbourne", "Australia")
	resultado.SetPoleman(hamilton)
	resultado.SetTemporada(2020)
	resultados = append(resultados, resultado)

	melbourne.SetResultados(resultados)

	result := melbourne.GetResultadoByTemporada(2020)
	pole := result.GetPoleman().GetNombre()
	data := Response{Pole: pole}

	msg, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(msg))
}
