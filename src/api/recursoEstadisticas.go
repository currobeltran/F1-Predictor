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

	c.JSON(200, temporadaEscogida.GetEstadisticas())
}

//Put : Método del recurso Estadisticas para modificar la información estadística de un GP
func (api RecursoEstadisticas) Put(c *gin.Context) {
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

	escribirEnFichero(resultados, "../api/data/resultados.json")

	c.JSON(200, temporadaEscogida.GetEstadisticas())
}

//GetMedia : Método para devolver la media de una estadistica concreta
func (api RecursoEstadisticas) GetMedia(c *gin.Context) {
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

	resultadosCircuito := resultados[c.Param("nombreCircuito")]

	if resultadosCircuito == nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	circuito := new(f1predictor.Circuito)
	circuito.SetResultados(resultadosCircuito)

	valor := float32(-1)
	valor = circuito.Media(c.Param("parametro"))

	if valor == float32(-1) {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	c.JSON(200, gin.H{"Media": valor})
}

//GetPrediccion : Método para devolver la prediccion de una estadistica concreta
func (api RecursoEstadisticas) GetPrediccion(c *gin.Context) {
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

	resultadosCircuito := resultados[c.Param("nombreCircuito")]

	if resultadosCircuito == nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	circuito := new(f1predictor.Circuito)
	circuito.SetResultados(resultadosCircuito)

	switch c.Param("parametro") {
	case "accidentes":
		valor := circuito.PosibilidadAccidentes()
		c.JSON(200, gin.H{"Posibles Accidentes": valor})
		return
	case "safetycar":
		valor := circuito.PosibilidadSafetyCar()
		c.JSON(200, gin.H{"Posibles Safety Car": valor})
		return
	case "adelantamientos":
		valor := circuito.PosibilidadAdelantamiento()
		c.JSON(200, gin.H{"Posibles Adelantamientos": valor})
		return
	case "banderaamarilla":
		valor := circuito.PosibilidadBanderaAmarilla()
		c.JSON(200, gin.H{"Posibles Banderas Amarillas": valor})
		return
	case "banderaroja":
		valor := circuito.PosibilidadBanderaRoja()
		c.JSON(200, gin.H{"Posibles Banderas Rojas": valor})
		return
	case "sancion":
		valor := circuito.PosibilidadSancion()
		c.JSON(200, gin.H{"Posibles Sanciones": valor})
		return
	}

	c.JSON(400, gin.H{"Error": "Bad Request"})
}
