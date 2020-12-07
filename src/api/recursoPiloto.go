package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"os"
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
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotoEscogido := pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	c.JSON(200, pilotoEscogido)
}

//Put : Método para modificar los datos de un recurso piloto
func (api RecursoPiloto) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotoEscogido := pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
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

	fichero, err := json.Marshal(pilotos)

	if err != nil {
		//TODO Logger
	}

	f, err := os.Create("data/pilotos.json")
	if err != nil {
		//TODO Logger
	}

	f.Write(fichero)
	f.Close()

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
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotos[pilotoNuevo.GetNombre()] = pilotoNuevo

	fichero, err := json.Marshal(pilotos)

	if err != nil {
		//TODO Logger
	}

	f, err := os.Create("data/pilotos.json")
	if err != nil {
		//TODO Logger
	}

	f.Write(fichero)
	f.Close()

	c.JSON(201, pilotoNuevo)
}
