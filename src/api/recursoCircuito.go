package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"net/url"
)

//RecursoCircuito : Estructura que representará un recurso que sea un circuito
type RecursoCircuito struct {
}

//Get : Método por el cual se definirá el procedimiento a seguir ante una petición GET
func (api RecursoCircuito) Get(params url.Values) (int, interface{}) {
	data, err := ioutil.ReadFile("data/circuitos.json")
	if err != nil {
		return 404, map[string]interface{}{"Error": err.Error()}
	}

	var campeonato map[string]f1predictor.Circuito

	err = json.Unmarshal(data, &campeonato)
	if err != nil {
		//TODO Logger
	}

	circuitoEscogido := campeonato[params.Get("nombre")]

	if circuitoEscogido.GetNombre() == "" {
		return 404, map[string]interface{}{"Error": "Not Found"}
	}

	return 200, circuitoEscogido
}
