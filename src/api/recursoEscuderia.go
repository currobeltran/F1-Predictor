package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//RecursoEscuderia : Struct que representará un recurso relacionado con una escuderia
type RecursoEscuderia struct {
}

//Get : Método correspondiente al recurso Escudería para obtener la información de la misma
func (api RecursoEscuderia) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var escuderias map[string]f1predictor.Escuderia
	err = json.Unmarshal(data, &escuderias)
	if err != nil {
		//TODO Logger
	}

	escuderiaEscogida := escuderias[c.Param("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	c.JSON(200, escuderiaEscogida)
}

//Put : Método correspondiente al recurso Escuderia para modificar la información de un elemento
func (api RecursoEscuderia) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var escuderias map[string]f1predictor.Escuderia
	err = json.Unmarshal(data, &escuderias)
	if err != nil {
		//TODO Logger
	}

	escuderiaEscogida := escuderias[c.Param("nombre")]

	if escuderiaEscogida.GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	if c.PostForm("nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request 1"})
		return
	}
	escuderiaEscogida.SetNombre(c.PostForm("nombre"))

	var vectorPilotos []f1predictor.Piloto
	pil1, ex := c.GetPostForm("piloto1")
	if !ex {
		c.JSON(400, gin.H{"Error": "Bad Request 2"})
		return
	}
	piloto1 := modificarStringPiloto(pil1)
	vectorPilotos = append(vectorPilotos, piloto1)

	pil2, ex2 := c.GetPostForm("piloto2")
	if !ex2 {
		c.JSON(400, gin.H{"Error": "Bad Request 3"})
		return
	}
	piloto2 := modificarStringPiloto(pil2)
	vectorPilotos = append(vectorPilotos, piloto2)

	escuderiaEscogida.SetPilotos(vectorPilotos)

	arg, exist := c.GetPostForm("titulos")
	if !exist {
		c.JSON(400, gin.H{"Error": "Bad Request 4"})
		return
	}
	num, _ := strconv.Atoi(arg)
	escuderiaEscogida.SetTitulosMundiales(num)

	arg2, exist2 := c.GetPostForm("victorias")
	if !exist2 {
		c.JSON(400, gin.H{"Error": "Bad Request 5"})
		return
	}
	num2, _ := strconv.Atoi(arg2)
	escuderiaEscogida.SetTitulosMundiales(num2)

	arg3, exist3 := c.GetPostForm("poles")
	if !exist3 {
		c.JSON(400, gin.H{"Error": "Bad Request 6"})
		return
	}
	num3, _ := strconv.Atoi(arg3)
	escuderiaEscogida.SetTitulosMundiales(num3)

	arg4, exist4 := c.GetPostForm("vueltasrapidas")
	if !exist4 {
		c.JSON(400, gin.H{"Error": "Bad Request 7"})
		return
	}
	num4, _ := strconv.Atoi(arg4)
	escuderiaEscogida.SetTitulosMundiales(num4)

	escuderias[c.Param("nombre")] = escuderiaEscogida

	fichero, err := json.Marshal(escuderias)

	if err != nil {
		//TODO Logger
	}

	f, err := os.Create("data/escuderia.json")
	if err != nil {
		//TODO Logger
	}

	f.Write(fichero)
	f.Close()
	c.JSON(200, escuderiaEscogida)
}

func modificarStringPiloto(s string) f1predictor.Piloto {
	var pil f1predictor.Piloto

	variables := strings.Split(s, "#")

	nombre := strings.Split(variables[0], ":")[1]
	pil.SetNombre(nombre)

	victorias := strings.Split(variables[1], ":")[1]
	nvic, _ := strconv.Atoi(victorias)
	pil.SetVictorias(nvic)

	poles := strings.Split(variables[2], ":")[1]
	npol, _ := strconv.Atoi(poles)
	pil.SetPoles(npol)

	vueltasrapidas := strings.Split(variables[3], ":")[1]
	nvr, _ := strconv.Atoi(vueltasrapidas)
	pil.SetVueltasRapidas(nvr)

	mundiales := strings.Split(variables[4], ":")[1]
	nm, _ := strconv.Atoi(mundiales)
	pil.SetMundiales(nm)

	return pil
}
