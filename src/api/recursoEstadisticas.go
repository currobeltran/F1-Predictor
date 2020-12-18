package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoEstadisticas : Struct que representará un recurso relacionado con las estadísticas
//de un GP
type RecursoEstadisticas struct {
	estadisticas map[string]f1predictor.EstadisticasGP
}

//AnadirEstadisticas : Método que nos ayudará a realizar la inyección de dependencias en este recurso
func (rEst *RecursoEstadisticas) AnadirEstadisticas(e f1predictor.EstadisticasGP, n string) {
	if rEst.estadisticas == nil {
		rEst.estadisticas = map[string]f1predictor.EstadisticasGP{}
	}
	rEst.estadisticas[n] = e
}

//Get : Método del recurso Estadísticas para obtener los datos de un gran premio determinado
func (rEst RecursoEstadisticas) Get(c *gin.Context) {
	str := c.Param("nombreCircuito") + " " + c.Param("temporada")

	estadisticas, exist := rEst.estadisticas[str]

	if !exist {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, estadisticas)
}

//GetMedia : Método para devolver la media de una estadistica concreta
func (rEst RecursoEstadisticas) GetMedia(c *gin.Context) {
	data, err := ioutil.ReadFile("../api/data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var resultados map[string][]f1predictor.ResultadoGP
	json.Unmarshal(data, &resultados)

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
func (rEst RecursoEstadisticas) GetPrediccion(c *gin.Context) {
	data, err := ioutil.ReadFile("../api/data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var resultados map[string][]f1predictor.ResultadoGP
	json.Unmarshal(data, &resultados)

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
