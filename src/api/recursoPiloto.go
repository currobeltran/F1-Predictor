package api

import (
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoPiloto : Struct que representa el recurso relacionado con un piloto
type RecursoPiloto struct {
	pilotos map[string]f1predictor.Piloto
}

//AnadirPiloto : Método para llevar a cabo la inyección de dependencias en RecursoPiloto
func (rPil *RecursoPiloto) AnadirPiloto(p f1predictor.Piloto) {
	if rPil.pilotos == nil {
		rPil.pilotos = map[string]f1predictor.Piloto{}
	}
	rPil.pilotos[p.Nombre] = p
}

//Get : Método para obtener los datos de un piloto
func (rPil RecursoPiloto) Get(c *gin.Context) {
	pilotoEscogido := rPil.pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, pilotoEscogido)
}
