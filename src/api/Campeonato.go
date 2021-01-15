package api

import (
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//Campeonato : Struct que representará un campeonato
//Contará con dos recursos asociados: Escuderias y Pilotos
type Campeonato struct {
	escuderias map[string]f1predictor.Escuderia
	pilotos    map[string]f1predictor.Piloto
}

//RequestPiloto : Tipo de dato creado para realizar peticiones PUT y POST sobre un piloto
type RequestPiloto struct {
	Nombre         string `json:"Nombre" binding:"required"`
	Victorias      int    `json:"Victorias" binding:"required"`
	Poles          int    `json:"Poles" binding:"required"`
	VueltasRapidas int    `json:"Vueltas Rapidas" binding:"required"`
	Mundiales      int    `json:"Mundiales" binding:"required"`
}

//RequestEscuderia : Tipo de dato creado para realizar peticiones PUT y POST sobre una escuderia
type RequestEscuderia struct {
	Nombre           string `json:"Nombre" binding:"required"`
	Piloto1          string `json:"Piloto1" binding:"required"`
	Piloto2          string `json:"Piloto2" binding:"required"`
	TitulosMundiales int    `json:"Titulos" binding:"required"`
	Victorias        int    `json:"Victorias" binding:"required"`
	Poles            int    `json:"Poles" binding:"required"`
	VueltasRapidas   int    `json:"Vueltas Rapidas" binding:"required"`
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

	var peticion RequestEscuderia
	err := c.ShouldBindJSON(&peticion)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var pilotos []f1predictor.Piloto
	pilotos = append(pilotos, Camp.pilotos[peticion.Piloto1])
	pilotos = append(pilotos, Camp.pilotos[peticion.Piloto2])

	escuderiaEscogida.SetNombre(peticion.Nombre)
	escuderiaEscogida.SetPilotos(pilotos)
	escuderiaEscogida.SetPoles(peticion.Poles)
	escuderiaEscogida.SetTitulosMundiales(peticion.TitulosMundiales)
	escuderiaEscogida.SetVictorias(peticion.Victorias)
	escuderiaEscogida.SetVueltasRapidas(peticion.VueltasRapidas)

	Camp.escuderias[peticion.Nombre] = escuderiaEscogida

	c.JSON(200, escuderiaEscogida)
}

//PutPiloto : Método para modificar la información de un piloto
func (Camp *Campeonato) PutPiloto(c *gin.Context) {
	pilotoEscogido := Camp.pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var peticion RequestPiloto
	err := c.ShouldBindJSON(&peticion)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	pilotoEscogido.SetNombre(peticion.Nombre)
	pilotoEscogido.SetMundiales(peticion.Mundiales)
	pilotoEscogido.SetPoles(peticion.Poles)
	pilotoEscogido.SetVictorias(peticion.Victorias)
	pilotoEscogido.SetVueltasRapidas(peticion.VueltasRapidas)

	Camp.pilotos[peticion.Nombre] = pilotoEscogido

	c.JSON(200, pilotoEscogido)
}

//PostEscuderia : Método con el que se podrá crear un nuevo recurso Escuderia
func (Camp *Campeonato) PostEscuderia(c *gin.Context) {
	var peticion RequestEscuderia
	err := c.ShouldBindJSON(&peticion)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	existeEscuderia := Camp.escuderias[peticion.Nombre]

	if existeEscuderia.GetNombre() != "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var nuevaEsc f1predictor.Escuderia

	var pilotos []f1predictor.Piloto
	pilotos = append(pilotos, Camp.pilotos[peticion.Piloto1])
	pilotos = append(pilotos, Camp.pilotos[peticion.Piloto2])

	nuevaEsc.Constructor(peticion.Nombre, pilotos, peticion.TitulosMundiales, peticion.Victorias,
		peticion.Poles, peticion.VueltasRapidas)
	Camp.escuderias[peticion.Nombre] = nuevaEsc

	URI := "/escuderia/" + peticion.Nombre
	c.JSON(201, gin.H{"Location": URI, "datos": nuevaEsc})
}

//PostPiloto : Método para crear un nuevo piloto
func (Camp *Campeonato) PostPiloto(c *gin.Context) {
	var peticion RequestPiloto
	err := c.ShouldBindJSON(&peticion)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	existePiloto := Camp.pilotos[c.PostForm("Nombre")]

	if existePiloto.GetNombre() != "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var nuevoPil f1predictor.Piloto

	nuevoPil.Constructor(peticion.Nombre, peticion.Victorias, peticion.Poles, peticion.VueltasRapidas, peticion.Mundiales)

	Camp.pilotos[peticion.Nombre] = nuevoPil

	URI := "/piloto/" + peticion.Nombre
	c.JSON(201, gin.H{"Location": URI, "datos": nuevoPil})
}
