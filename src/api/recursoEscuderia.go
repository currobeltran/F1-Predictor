package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoEscuderia : Struct que representará un recurso relacionado con una escuderia
type RecursoEscuderia struct {
}

//Get : Método correspondiente al recurso Escudería para obtener la información de la misma
func (api RecursoEscuderia) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var escuderias map[string]f1predictor.Escuderia
	err = json.Unmarshal(data, &escuderias)
	if err != nil {
		//TODO Logger
	}

	escuderiaEscogida := escuderias[c.Param("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	c.JSON(200, escuderiaEscogida)
}
