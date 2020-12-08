package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"strconv"

	"github.com/gin-gonic/gin"
)

//RecursoPiloto : Struct que representa el recurso relacionado con un piloto
type RecursoPiloto struct {
}

//Get : Método para obtener los datos de un piloto
func (api RecursoPiloto) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotoEscogido := pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, pilotoEscogido)
}

//Put : Método para modificar los datos de un recurso piloto
func (api RecursoPiloto) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotoEscogido := pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	if c.PostForm("nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	pilotoEscogido.SetNombre(c.PostForm("nombre"))

	if c.PostForm("victorias") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nvic, err := strconv.Atoi(c.PostForm("victorias"))
	pilotoEscogido.SetVictorias(nvic)

	if c.PostForm("poles") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	npol, err := strconv.Atoi(c.PostForm("poles"))
	pilotoEscogido.SetPoles(npol)

	if c.PostForm("vueltasrapidas") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nvr, err := strconv.Atoi(c.PostForm("vueltasrapidas"))
	pilotoEscogido.SetVueltasRapidas(nvr)

	if c.PostForm("mundiales") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nm, err := strconv.Atoi(c.PostForm("mundiales"))
	pilotoEscogido.SetMundiales(nm)

	pilotos[c.Param("nombre")] = pilotoEscogido

	escribirEnFichero(pilotos, "data/pilotos.json")

	c.JSON(200, pilotoEscogido)
}

//Post : Método para añadir un nuevo recursoPiloto
func (api RecursoPiloto) Post(c *gin.Context) {
	var pilotoNuevo f1predictor.Piloto

	if c.PostForm("nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	pilotoNuevo.SetNombre(c.PostForm("nombre"))

	if c.PostForm("victorias") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nvic, err := strconv.Atoi(c.PostForm("victorias"))
	pilotoNuevo.SetVictorias(nvic)

	if c.PostForm("poles") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	npol, err := strconv.Atoi(c.PostForm("poles"))
	pilotoNuevo.SetPoles(npol)

	if c.PostForm("vueltasrapidas") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nvr, err := strconv.Atoi(c.PostForm("vueltasrapidas"))
	pilotoNuevo.SetVueltasRapidas(nvr)

	if c.PostForm("mundiales") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nm, err := strconv.Atoi(c.PostForm("mundiales"))
	pilotoNuevo.SetMundiales(nm)

	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotos[pilotoNuevo.GetNombre()] = pilotoNuevo

	escribirEnFichero(pilotos, "data/pilotos.json")

	c.JSON(201, pilotoNuevo)
}

//Delete : Método para eliminar un piloto
func (api RecursoPiloto) Delete(c *gin.Context) {
	/*********************** Eliminamos piloto del archivo ************************/
	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	if pilotos[c.Param("nombre")].GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	delete(pilotos, c.Param("nombre"))
	escribirEnFichero(pilotos, "data/pilotos.json")

	/*********************** Realizamos comprobación ***************************/
	comprobaciondata, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var comprobacion map[string]f1predictor.Piloto
	err = json.Unmarshal(comprobaciondata, &comprobacion)
	if err != nil {
		//TODO Logger
	}

	if comprobacion[c.Param("nombre")].GetNombre() == "" {
		c.JSON(200, gin.H{"Registro Eliminado": c.Param("nombre")})
		return
	}
	c.JSON(400, gin.H{"Error": "Bad Request"})
}
