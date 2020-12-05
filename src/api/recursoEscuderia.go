package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"net/url"
)

//RecursoEscuderia : Struct que representará un recurso relacionado con una escuderia
type RecursoEscuderia struct {
}

//Get : Método correspondiente al recurso Escudería para obtener la información de la misma
func (api RecursoEscuderia) Get(params url.Values) (int, interface{}) {
	data, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		return 404, map[string]interface{}{"Error": err.Error()}
	}

	var escuderias map[string]f1predictor.Escuderia
	err = json.Unmarshal(data, &escuderias)
	if err != nil {
		//TODO Logger
	}

	escuderiaEscogida := escuderias[params.Get("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		return 404, map[string]interface{}{"Error": "Not Found"}
	}

	return 200, escuderiaEscogida
}
