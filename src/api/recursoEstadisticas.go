package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"strconv"

	"github.com/gin-gonic/gin"
)

//RecursoEstadisticas : Struct que representará un recurso relacionado con las estadísticas
//de un GP
type RecursoEstadisticas struct {
}

//Get : Método del recurso Estadísticas para obtener los datos de un gran premio determinado
func (api RecursoEstadisticas) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var resultados map[string][]f1predictor.ResultadoGP
	err = json.Unmarshal(data, &resultados)
	if err != nil {
		//TODO Logger
	}

	var temporadaEscogida f1predictor.ResultadoGP
	temporadaNum, _ := strconv.Atoi(c.Param("temporada"))

	for x := 0; x < len(resultados[c.Param("nombreCircuito")]); x++ {
		if resultados[c.Param("nombreCircuito")][x].GetTemporada() == temporadaNum {
			temporadaEscogida = resultados[c.Param("nombreCircuito")][x]
		}
	}

	if temporadaEscogida.GetTemporada() == 0 {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	c.JSON(200, temporadaEscogida.GetEstadisticas())
}
