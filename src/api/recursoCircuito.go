package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

//RecursoCircuito : Estructura que representará un recurso que sea un circuito
type RecursoCircuito struct {
}

//Get : Método por el cual se definirá el procedimiento a seguir ante una petición GET
func (api RecursoCircuito) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("data/circuitos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var campeonato map[string]f1predictor.Circuito

	err = json.Unmarshal(data, &campeonato)
	if err != nil {
		//TODO Logger
	}

	circuitoEscogido := campeonato[c.Param("nombre")]

	if circuitoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, circuitoEscogido)
}

//Put : Método por el cual se podrá modificar un recurso Circuito de los ya existentes
func (api RecursoCircuito) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("data/circuitos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var campeonato map[string]f1predictor.Circuito

	err = json.Unmarshal(data, &campeonato)
	if err != nil {
		//TODO Logger
	}

	circuitoEscogido := campeonato[c.Param("nombre")]

	if circuitoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	if c.PostForm("nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	circuitoEscogido.SetNombre(c.PostForm("nombre"))

	if c.PostForm("pais") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	circuitoEscogido.SetPais(c.PostForm("pais"))

	campeonato[c.Param("nombre")] = circuitoEscogido

	escribirEnFichero(campeonato, "data/circuitos.json")

	c.JSON(200, circuitoEscogido)
}

// Post : Método para añadir un nuevo recursoCircuito
func (api RecursoCircuito) Post(c *gin.Context) {
	var circuitoNuevo f1predictor.Circuito

	if c.PostForm("nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	circuitoNuevo.SetNombre(c.PostForm("nombre"))

	if c.PostForm("pais") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	circuitoNuevo.SetPais(c.PostForm("pais"))

	data, err := ioutil.ReadFile("data/circuitos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var campeonato map[string]f1predictor.Circuito

	err = json.Unmarshal(data, &campeonato)
	if err != nil {
		//TODO Logger
	}

	campeonato[circuitoNuevo.GetNombre()] = circuitoNuevo

	escribirEnFichero(campeonato, "data/circuitos.json")

	c.JSON(201, circuitoNuevo)
}
