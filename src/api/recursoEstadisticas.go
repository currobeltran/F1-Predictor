package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"net/url"
	"strconv"
)

//RecursoEstadisticas : Struct que representará un recurso relacionado con las estadísticas
//de un GP
type RecursoEstadisticas struct {
}

//Get : Método del recurso Estadísticas para obtener los datos de un gran premio determinado
func (api RecursoEstadisticas) Get(params url.Values) (int, interface{}) {
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

	return 200, temporadaEscogida.GetEstadisticas()
}
