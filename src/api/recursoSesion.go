package api

import (
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoSesion : Struct que representará el recurso relacionado con una sesión de un GP
type RecursoSesion struct {
	resultados map[string]f1predictor.ResultadoGP
}

//AnadirResultado : Método para realizar la inyección de dependencias en recursoSesion
func (rSes *RecursoSesion) AnadirResultado(r f1predictor.ResultadoGP, n string) {
	if rSes.resultados == nil {
		rSes.resultados = map[string]f1predictor.ResultadoGP{}
	}
	rSes.resultados[n] = r
}

//Get : Metodo para obtener la información de una sesion de un Gran Premio concreto
func (rSes RecursoSesion) Get(c *gin.Context) {
	str := c.Param("nombreCircuito") + " " + c.Param("temporada")

	temporadaEscogida := rSes.resultados[str]

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
