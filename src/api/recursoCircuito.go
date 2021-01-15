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

//Put : Método por el cual se modificará un circuito ya existente
func (rCirc *RecursoCircuito) Put(c *gin.Context) {
	circuitoEscogido := rCirc.circuitos[c.Param("nombre")]

	if circuitoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	if (c.PostForm("Nombre") == "") || (c.PostForm("Pais") == "") {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	circuitoEscogido.SetNombre(c.PostForm("Nombre"))
	circuitoEscogido.SetPais(c.PostForm("Pais"))

	rCirc.circuitos[c.Param("nombre")] = circuitoEscogido

	c.JSON(200, circuitoEscogido)
}

//Post : Método por el cual se podrá crear un nuevo circuito
func (rCirc *RecursoCircuito) Post(c *gin.Context) {
	if (c.PostForm("Nombre") == "") || (c.PostForm("Pais") == "") {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	existeCircuito := rCirc.circuitos[c.PostForm("Nombre")]

	if existeCircuito.GetNombre() != "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var auxCirc f1predictor.Circuito
	auxCirc.Constructor(c.PostForm("Nombre"), c.PostForm("Pais"))
	rCirc.circuitos[c.PostForm("Nombre")] = auxCirc

	URI := "/circuito/" + c.PostForm("Nombre")
	c.JSON(201, gin.H{"Location": URI, "datos": auxCirc})
}
