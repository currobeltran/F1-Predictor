package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"strconv"

	"github.com/gin-gonic/gin"
)

//RecursoSesion : Struct que representará el recurso relacionado con una sesión de un GP
type RecursoSesion struct {
}

//Get : Metodo para obtener la información de una sesion de un Gran Premio concreto
func (api RecursoSesion) Get(c *gin.Context) {
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

	switch c.Param("nombreSesion") {
	case "fp1":
		c.JSON(200, temporadaEscogida.GetResultadoFP1())
		break
	case "fp2":
		c.JSON(200, temporadaEscogida.GetResultadoFP2())
		break
	case "fp3":
		c.JSON(200, temporadaEscogida.GetResultadoFP3())
		break
	case "clasificacion":
		c.JSON(200, temporadaEscogida.GetResultadoClasificacion())
		break
	case "carrera":
		c.JSON(200, temporadaEscogida.GetResultadoCarrera())
		break
	default:
		c.JSON(400, gin.H{"Error": "Bad Request"})
	}

}
