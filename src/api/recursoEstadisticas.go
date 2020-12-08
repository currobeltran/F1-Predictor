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

	c.JSON(200, temporadaEscogida.GetEstadisticas())
}

//Put : Método del recurso Estadisticas para modificar la información estadística de un GP
func (api RecursoEstadisticas) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("data/resultados.json")
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

	est := temporadaEscogida.GetEstadisticas()

	if c.PostForm("accidentes") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nac, err := strconv.Atoi(c.PostForm("accidentes"))
	est.SetAccidentes(nac)

	if c.PostForm("numerosafety") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nsc, err := strconv.Atoi(c.PostForm("numerosafety"))
	est.SetNumeroSafetyCar(nsc)

	if c.PostForm("adelantamientos") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nad, err := strconv.Atoi(c.PostForm("adelantamientos"))
	est.SetAdelantamientos(nad)

	if c.PostForm("banderasamarillas") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nba, err := strconv.Atoi(c.PostForm("banderasamarillas"))
	est.SetBanderasAmarillas(nba)

	if c.PostForm("banderasrojas") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nbr, err := strconv.Atoi(c.PostForm("banderasrojas"))
	est.SetBanderasRojas(nbr)

	if c.PostForm("sanciones") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nsan, err := strconv.Atoi(c.PostForm("sanciones"))
	est.SetSanciones(nsan)

	if c.PostForm("mejortiempo") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	tvuelta := convertirTiempoString(c.PostForm("mejortiempo"))
	est.SetMejorVuelta(tvuelta)

	temporadaEscogida.SetEstadisticas(est)

	for x := 0; x < len(resultados[c.Param("nombreCircuito")]); x++ {
		if resultados[c.Param("nombreCircuito")][x].GetTemporada() == temporadaNum {
			resultados[c.Param("nombreCircuito")][x] = temporadaEscogida
		}
	}

	escribirEnFichero(resultados, "data/resultados.json")

	c.JSON(200, temporadaEscogida.GetEstadisticas())
}
