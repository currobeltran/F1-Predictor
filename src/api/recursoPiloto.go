package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"net/url"
)

//RecursoPiloto : Struct que representa el recurso relacionado con un piloto
type RecursoPiloto struct {
}

//Get : Método para obtener los datos de un piloto
func (api RecursoPiloto) Get(params url.Values) (int, interface{}) {
	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		return 404, map[string]interface{}{"Error": err.Error()}
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotoEscogido := pilotos[params.Get("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		return 404, map[string]interface{}{"Error": "Not Found"}
	}

	return 200, pilotoEscogido
}
