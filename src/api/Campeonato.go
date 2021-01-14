package api

import (
	"m/src/f1predictor"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Campeonato : Struct que representará un campeonato
//Contará con dos recursos asociados: Escuderias y Pilotos
type Campeonato struct {
	escuderias map[string]f1predictor.Escuderia
	pilotos    map[string]f1predictor.Piloto
}

// AnadirEscuderia : Método que nos permitirá realizar la inyección de dependencias en RecursoEscuderia
func (Camp *Campeonato) AnadirEscuderia(e f1predictor.Escuderia) {
	if Camp.escuderias == nil {
		Camp.escuderias = map[string]f1predictor.Escuderia{}
	}
	Camp.escuderias[e.Nombre] = e
}

//AnadirPiloto : Método para llevar a cabo la inyección de dependencias en RecursoPiloto
func (Camp *Campeonato) AnadirPiloto(p f1predictor.Piloto) {
	if Camp.pilotos == nil {
		Camp.pilotos = map[string]f1predictor.Piloto{}
	}
	Camp.pilotos[p.Nombre] = p
}

//GetEscuderia : Método correspondiente al recurso Escudería para obtener la información de la misma
func (Camp Campeonato) GetEscuderia(c *gin.Context) {
	escuderiaEscogida := Camp.escuderias[c.Param("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, escuderiaEscogida)
}

//GetPiloto : Método para obtener los datos de un piloto
func (Camp Campeonato) GetPiloto(c *gin.Context) {
	pilotoEscogido := Camp.pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, pilotoEscogido)
}

//PutEscuderia : Método para cambiar las características de una escudería
func (Camp *Campeonato) PutEscuderia(c *gin.Context) {
	escuderiaEscogida := Camp.escuderias[c.Param("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	if c.PostForm("Nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	if (c.PostForm("Piloto1") == "") || (c.PostForm("Piloto2") == "") {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	poles, _ := strconv.Atoi(c.PostForm("Poles"))
	victorias, _ := strconv.Atoi(c.PostForm("Victorias"))
	titulos, _ := strconv.Atoi(c.PostForm("Titulos"))
	vr, _ := strconv.Atoi(c.PostForm("Vueltas Rápidas"))

	var pilotos []f1predictor.Piloto
	pilotos = append(pilotos, Camp.pilotos[c.PostForm("Piloto1")])
	pilotos = append(pilotos, Camp.pilotos[c.PostForm("Piloto2")])

	escuderiaEscogida.SetNombre(c.PostForm("Nombre"))
	escuderiaEscogida.SetPilotos(pilotos)
	escuderiaEscogida.SetPoles(poles)
	escuderiaEscogida.SetTitulosMundiales(titulos)
	escuderiaEscogida.SetVictorias(victorias)
	escuderiaEscogida.SetVueltasRapidas(vr)

	Camp.escuderias[c.Param("nombre")] = escuderiaEscogida

	c.JSON(200, escuderiaEscogida)
}

//PutPiloto : Método para modificar la información de un piloto
func (Camp *Campeonato) PutPiloto(c *gin.Context) {
	pilotoEscogido := Camp.pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	if c.PostForm("Nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	poles, _ := strconv.Atoi(c.PostForm("Poles"))
	victorias, _ := strconv.Atoi(c.PostForm("Victorias"))
	mundiales, _ := strconv.Atoi(c.PostForm("Mundiales"))
	vr, _ := strconv.Atoi(c.PostForm("Vueltas Rapidas"))

	pilotoEscogido.SetNombre(c.PostForm("Nombre"))
	pilotoEscogido.SetMundiales(mundiales)
	pilotoEscogido.SetPoles(poles)
	pilotoEscogido.SetVictorias(victorias)
	pilotoEscogido.SetVueltasRapidas(vr)

	Camp.pilotos[c.Param("nombre")] = pilotoEscogido

	c.JSON(200, pilotoEscogido)
}

//PostEscuderia : Método con el que se podrá crear un nuevo recurso Escuderia
func (Camp *Campeonato) PostEscuderia(c *gin.Context) {
	if c.PostForm("Nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	if (c.PostForm("Piloto1") == "") || (c.PostForm("Piloto2") == "") {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	existeEscuderia := Camp.escuderias[c.PostForm("Nombre")]

	if existeEscuderia.GetNombre() != "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var nuevaEsc f1predictor.Escuderia

	poles, _ := strconv.Atoi(c.PostForm("Poles"))
	victorias, _ := strconv.Atoi(c.PostForm("Victorias"))
	titulos, _ := strconv.Atoi(c.PostForm("Titulos"))
	vr, _ := strconv.Atoi(c.PostForm("Vueltas Rapidas"))

	var pilotos []f1predictor.Piloto
	pilotos = append(pilotos, Camp.pilotos[c.PostForm("Piloto1")])
	pilotos = append(pilotos, Camp.pilotos[c.PostForm("Piloto2")])

	nuevaEsc.Constructor(c.PostForm("Nombre"), pilotos, titulos, victorias, poles, vr)
	Camp.escuderias[c.PostForm("Nombre")] = nuevaEsc

	URI := "/api/escuderia/" + c.PostForm("Nombre")
	c.JSON(200, gin.H{"Location": URI, "datos": nuevaEsc})
}

//PostPiloto : Método para crear un nuevo piloto
func (Camp *Campeonato) PostPiloto(c *gin.Context) {
	if c.PostForm("Nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	existePiloto := Camp.pilotos[c.PostForm("Nombre")]

	if existePiloto.GetNombre() != "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var nuevoPil f1predictor.Piloto

	poles, _ := strconv.Atoi(c.PostForm("Poles"))
	victorias, _ := strconv.Atoi(c.PostForm("Victorias"))
	mundiales, _ := strconv.Atoi(c.PostForm("Mundiales"))
	vr, _ := strconv.Atoi(c.PostForm("Vueltas Rapidas"))

	nuevoPil.Constructor(c.PostForm("Nombre"), victorias, poles, vr, mundiales)

	Camp.pilotos[c.PostForm("Nombre")] = nuevoPil

	URI := "/api/piloto/" + c.PostForm("Nombre")
	c.JSON(200, gin.H{"Location": URI, "datos": nuevoPil})
}
