package api

import (
	"m/src/f1predictor"
	"strconv"

	"github.com/gin-gonic/gin"
)

//RecursoEscuderia : Struct que representará un recurso relacionado con una escuderia
type RecursoEscuderia struct {
	escuderias map[string]f1predictor.Escuderia
	pilotos    map[string]f1predictor.Piloto
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

//Put : Método para cambiar las características de una escudería
func (rEsc RecursoEscuderia) Put(c *gin.Context) {
	escuderiaEscogida := rEsc.escuderias[c.Param("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	if c.PostForm("Nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	if c.PostForm("Pilotos") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	poles, _ := strconv.Atoi(c.PostForm("Poles"))
	victorias, _ := strconv.Atoi(c.PostForm("Titulos"))
	titulos, _ := strconv.Atoi(c.PostForm("Victorias"))
	vr, _ := strconv.Atoi(c.PostForm("Vueltas Rápidas"))

	var pilotos []f1predictor.Piloto
	// pilotos = append(pilotos, RecursoPiloto.obtenerPiloto(c.PostForm("Piloto1")))
	// pilotos = append(pilotos, RecursoPiloto.obtenerPiloto(c.PostForm("Piloto2")))

	escuderiaEscogida.SetNombre(c.PostForm("Nombre"))
	escuderiaEscogida.SetPilotos(pilotos)
	escuderiaEscogida.SetPoles(poles)
	escuderiaEscogida.SetTitulosMundiales(titulos)
	escuderiaEscogida.SetVictorias(victorias)
	escuderiaEscogida.SetVueltasRapidas(vr)

	c.JSON(200, escuderiaEscogida)
}
