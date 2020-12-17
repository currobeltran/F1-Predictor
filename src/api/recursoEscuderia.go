package api

import (
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoEscuderia : Struct que representará un recurso relacionado con una escuderia
type RecursoEscuderia struct {
	escuderias map[string]f1predictor.Escuderia
}

// AnadirEscuderia : Método que nos permitirá realizar la inyección de dependencias en RecursoEscuderia
func (rEsc *RecursoEscuderia) AnadirEscuderia(e f1predictor.Escuderia) {
	if rEsc.escuderias == nil {
		rEsc.escuderias = map[string]f1predictor.Escuderia{}
	}
	rEsc.escuderias[e.Nombre] = e
}

//Get : Método correspondiente al recurso Escudería para obtener la información de la misma
func (rEsc RecursoEscuderia) Get(c *gin.Context) {
	escuderiaEscogida := rEsc.escuderias[c.Param("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, escuderiaEscogida)
}
