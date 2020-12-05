package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"net/url"
	"strconv"
)

//RecursoSesion : Struct que representará el recurso relacionado con una sesión de un GP
type RecursoSesion struct {
}

//Get : Metodo para obtener la información de una sesion de un Gran Premio concreto
func (api RecursoSesion) Get(params url.Values) (int, interface{}) {
	data, err := ioutil.ReadFile("data/resultados.json")
	if err != nil {
		return 404, map[string]interface{}{"Error": err.Error()}
	}

	var resultados map[string][]f1predictor.ResultadoGP
	err = json.Unmarshal(data, &resultados)
	if err != nil {
		//TODO Logger
	}

	var temporadaEscogida f1predictor.ResultadoGP
	temporadaNum, _ := strconv.Atoi(params.Get("temporada"))

	for x := 0; x < len(resultados[params.Get("nombreCircuito")]); x++ {
		if resultados[params.Get("nombreCircuito")][x].GetTemporada() == temporadaNum {
			temporadaEscogida = resultados[params.Get("nombreCircuito")][x]
		}
	}

	if temporadaEscogida.GetTemporada() == 0 {
		return 404, map[string]interface{}{"Error": "Not Found"}
	}

	switch params.Get("nombreSesion") {
	case "fp1":
		return 200, temporadaEscogida.GetResultadoFP1()
	case "fp2":
		return 200, temporadaEscogida.GetResultadoFP2()
	case "fp3":
		return 200, temporadaEscogida.GetResultadoFP3()
	case "clasificacion":
		return 200, temporadaEscogida.GetResultadoClasificacion()
	case "carrera":
		return 200, temporadaEscogida.GetResultadoCarrera()
	}

	return 400, map[string]interface{}{"Error": "Bad Request"}
}
