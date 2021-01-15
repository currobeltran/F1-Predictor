package api

import (
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoCircuito : Estructura que representará un recurso que sea un circuito
type RecursoCircuito struct {
	circuitos map[string]f1predictor.Circuito
}

//RequestCircuito : Tipo de dato creado para realizar peticiones PUT y POST sobre un recursoCircuito
type RequestCircuito struct {
	Nombre string `json:"Nombre" binding:"required"`
	Pais   string `json:"Pais" binding:"required"`
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

//Put : Método por el cual se modificará un circuito ya existente
func (rCirc *RecursoCircuito) Put(c *gin.Context) {
	circuitoEscogido := rCirc.circuitos[c.Param("nombre")]

	if circuitoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var peticion RequestCircuito
	err := c.ShouldBindJSON(&peticion)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	circuitoEscogido.SetNombre(peticion.Nombre)
	circuitoEscogido.SetPais(peticion.Pais)

	rCirc.circuitos[c.Param("nombre")] = circuitoEscogido

	c.JSON(200, circuitoEscogido)
}

//Post : Método por el cual se podrá crear un nuevo circuito
func (rCirc *RecursoCircuito) Post(c *gin.Context) {
	var peticion RequestCircuito
	err := c.ShouldBindJSON(&peticion)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	existeCircuito := rCirc.circuitos[c.PostForm("Nombre")]

	if existeCircuito.GetNombre() != "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var auxCirc f1predictor.Circuito
	auxCirc.Constructor(peticion.Nombre, peticion.Pais)
	rCirc.circuitos[peticion.Nombre] = auxCirc

	URI := "/circuito/" + peticion.Nombre
	c.JSON(201, gin.H{"Location": URI, "datos": auxCirc})
}
