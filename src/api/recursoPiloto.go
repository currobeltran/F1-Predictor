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
	// data, err := ioutil.ReadFile("../api/data/pilotos.json")
	// if err != nil {
	// 	c.JSON(404, gin.H{"Error": "Not Found"})
	// 	return
	// }

	// var pilotos map[string]f1predictor.Piloto
	// json.Unmarshal(data, &pilotos)

	pilotoEscogido := rPil.pilotos[c.Param("nombre")]

	if pilotoEscogido.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(200, pilotoEscogido)
}

//Put : Método para modificar los datos de un recurso piloto
func (rPil RecursoPiloto) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("../api/data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var pilotos map[string]f1predictor.Piloto
	json.Unmarshal(data, &pilotos)

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

	escribirEnFichero(pilotos, "../api/data/pilotos.json")

	c.JSON(200, pilotoEscogido)
}

//Post : Método para añadir un nuevo recursoPiloto
func (rPil RecursoPiloto) Post(c *gin.Context) {
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

	data, err := ioutil.ReadFile("../api/data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var pilotos map[string]f1predictor.Piloto
	json.Unmarshal(data, &pilotos)

	pilotos[pilotoNuevo.GetNombre()] = pilotoNuevo

	escribirEnFichero(pilotos, "../api/data/pilotos.json")

	c.JSON(201, pilotoNuevo)
}

//Delete : Método para eliminar un piloto
func (rPil RecursoPiloto) Delete(c *gin.Context) {
	/*********************** Eliminamos piloto del archivo ************************/
	data, err := ioutil.ReadFile("../api/data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var pilotos map[string]f1predictor.Piloto
	json.Unmarshal(data, &pilotos)

	if pilotos[c.Param("nombre")].GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	delete(pilotos, c.Param("nombre"))
	escribirEnFichero(pilotos, "../api/data/pilotos.json")

	/*********************** Realizamos comprobación ***************************/
	comprobaciondata, err := ioutil.ReadFile("../api/data/pilotos.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var comprobacion map[string]f1predictor.Piloto
	json.Unmarshal(comprobaciondata, &comprobacion)

	if comprobacion[c.Param("nombre")].GetNombre() == "" {
		c.JSON(200, gin.H{"Registro Eliminado": c.Param("nombre")})
		return
	}
	c.JSON(400, gin.H{"Error": "Bad Request"})
}
