package api

import (
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoCircuito : Estructura que representará un recurso que sea un circuito
type RecursoCircuito struct {
	circuitos map[string]f1predictor.Circuito
}

//AnadirPista : Método que nos permitirá llevar a cabo la inyección de dependencias
func (rCirc *RecursoCircuito) AnadirPista(c f1predictor.Circuito) {
	if rCirc.circuitos == nil {
		rCirc.circuitos = map[string]f1predictor.Circuito{}
	}
	rCirc.circuitos[c.Nombre] = c
}

//Get : Método por el cual se definirá el procedimiento a seguir ante una petición GET
func (rCirc RecursoCircuito) Get(c *gin.Context) {
	circuitoEscogido := rCirc.circuitos[c.Param("nombre")]

	if circuitoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, circuitoEscogido)
}
