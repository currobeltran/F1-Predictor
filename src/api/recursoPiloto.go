package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoPiloto : Struct que representa el recurso relacionado con un piloto
type RecursoPiloto struct {
}

//Get : Método para obtener los datos de un piloto
func (api RecursoPiloto) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotoEscogido := pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	c.JSON(200, pilotoEscogido)
}
