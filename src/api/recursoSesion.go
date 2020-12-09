package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

//RecursoSesion : Struct que representará el recurso relacionado con una sesión de un GP
type RecursoSesion struct {
}

//Get : Metodo para obtener la información de una sesion de un Gran Premio concreto
func (api RecursoSesion) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("../api/data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
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
		return
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

//Put : Método para modificar la información de una sesión
func (api RecursoSesion) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("../api/data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
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

	var sesionEscogida [20]f1predictor.SesionPiloto

	if c.PostForm("sesion") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesionEscogida = convertirStringSesion(c.PostForm("sesion"))

	switch c.Param("nombreSesion") {
	case "fp1":
		temporadaEscogida.SetResultadoFP1(sesionEscogida)
		break
	case "fp2":
		temporadaEscogida.SetResultadoFP2(sesionEscogida)
		break
	case "fp3":
		temporadaEscogida.SetResultadoFP3(sesionEscogida)
		break
	case "clasificacion":
		if c.PostForm("poleman") == "" {
			c.JSON(400, gin.H{"Error": "Bad Request"})
			return
		}
		poleman := obtenerPiloto(c.PostForm("poleman"))
		temporadaEscogida.SetPoleman(poleman)
		temporadaEscogida.SetResultadoClasificacion(sesionEscogida)
		break
	case "carrera":
		if c.PostForm("ganador") == "" {
			c.JSON(400, gin.H{"Error": "Bad Request"})
			return
		}
		ganador := obtenerPiloto(c.PostForm("ganador"))
		temporadaEscogida.SetGanador(ganador)

		if c.PostForm("podio") == "" {
			c.JSON(400, gin.H{"Error": "Bad Request"})
			return
		}
		podio := convertirPodioString(c.PostForm("podio"))
		if !reflect.DeepEqual(ganador, podio[0]) {
			c.JSON(400, gin.H{"Error": "Bad Request"})
			return
		}
		temporadaEscogida.SetPodio(podio)

		temporadaEscogida.SetResultadoCarrera(sesionEscogida)
		break
	}

	for x := 0; x < len(resultados[c.Param("nombreCircuito")]); x++ {
		if resultados[c.Param("nombreCircuito")][x].GetTemporada() == temporadaNum {
			resultados[c.Param("nombreCircuito")][x] = temporadaEscogida
		}
	}

	escribirEnFichero(resultados, "../api/data/resultados.json")

	c.JSON(200, sesionEscogida)
}
