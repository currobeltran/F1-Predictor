package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"strconv"

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
		return
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

	c.JSON(200, escuderiaEscogida)
}

//Put : Método correspondiente al recurso Escuderia para modificar la información de un elemento
func (api RecursoEscuderia) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
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
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	escuderiaEscogida.SetNombre(c.PostForm("nombre"))

	var vectorPilotos []f1predictor.Piloto
	pil1, ex := c.GetPostForm("piloto1")
	if !ex {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	piloto1 := obtenerPiloto(pil1)
	vectorPilotos = append(vectorPilotos, piloto1)

	pil2, ex2 := c.GetPostForm("piloto2")
	if !ex2 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	piloto2 := obtenerPiloto(pil2)
	vectorPilotos = append(vectorPilotos, piloto2)

	escuderiaEscogida.SetPilotos(vectorPilotos)

	arg, exist := c.GetPostForm("titulos")
	if !exist {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num, _ := strconv.Atoi(arg)
	escuderiaEscogida.SetTitulosMundiales(num)

	arg2, exist2 := c.GetPostForm("victorias")
	if !exist2 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num2, _ := strconv.Atoi(arg2)
	escuderiaEscogida.SetVictorias(num2)

	arg3, exist3 := c.GetPostForm("poles")
	if !exist3 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num3, _ := strconv.Atoi(arg3)
	escuderiaEscogida.SetPoles(num3)

	arg4, exist4 := c.GetPostForm("vueltasrapidas")
	if !exist4 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num4, _ := strconv.Atoi(arg4)
	escuderiaEscogida.SetVueltasRapidas(num4)

	escuderias[c.Param("nombre")] = escuderiaEscogida

	escribirEnFichero(escuderias, "data/escuderia.json")

	c.JSON(200, escuderiaEscogida)
}

//Post : Crea un nuevo elemento de recursoEscuderia
func (api RecursoEscuderia) Post(c *gin.Context) {
	/************ Creamos nueva escuderia **************/
	var escuderiaNueva f1predictor.Escuderia

	if c.PostForm("nombre") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	escuderiaNueva.SetNombre(c.PostForm("nombre"))

	var vectorPilotos []f1predictor.Piloto
	pil1, ex := c.GetPostForm("piloto1")
	if !ex {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	piloto1 := obtenerPiloto(pil1)
	vectorPilotos = append(vectorPilotos, piloto1)

	pil2, ex2 := c.GetPostForm("piloto2")
	if !ex2 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	piloto2 := obtenerPiloto(pil2)
	vectorPilotos = append(vectorPilotos, piloto2)

	escuderiaNueva.SetPilotos(vectorPilotos)

	arg, exist := c.GetPostForm("titulos")
	if !exist {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num, _ := strconv.Atoi(arg)
	escuderiaNueva.SetTitulosMundiales(num)

	arg2, exist2 := c.GetPostForm("victorias")
	if !exist2 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num2, _ := strconv.Atoi(arg2)
	escuderiaNueva.SetVictorias(num2)

	arg3, exist3 := c.GetPostForm("poles")
	if !exist3 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num3, _ := strconv.Atoi(arg3)
	escuderiaNueva.SetPoles(num3)

	arg4, exist4 := c.GetPostForm("vueltasrapidas")
	if !exist4 {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	num4, _ := strconv.Atoi(arg4)
	escuderiaNueva.SetVueltasRapidas(num4)

	/************ Añadimos nueva escuderia a archivo **************/

	data, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}
	var escuderias map[string]f1predictor.Escuderia
	err = json.Unmarshal(data, &escuderias)
	if err != nil {
		//TODO Logger
	}

	escuderias[escuderiaNueva.GetNombre()] = escuderiaNueva
	escribirEnFichero(escuderias, "data/escuderia.json")

	c.JSON(201, escuderiaNueva)
}

//Delete : Método para eliminar un recurso de Escuderia
func (api RecursoEscuderia) Delete(c *gin.Context) {
	data, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var escuderias map[string]f1predictor.Escuderia
	err = json.Unmarshal(data, &escuderias)
	if err != nil {
		//TODO Logger
	}

	if escuderias[c.Param("nombre")].GetNombre() == "" {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	delete(escuderias, c.Param("nombre"))

	escribirEnFichero(escuderias, "data/escuderia.json")

	/*************** Recuperamos de nuevo el archivo para comprobar borrado *******************/

	comprobaciondata, err := ioutil.ReadFile("data/escuderia.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var comprobacion map[string]f1predictor.Escuderia
	err = json.Unmarshal(comprobaciondata, &comprobacion)
	if err != nil {
		//TODO Logger
	}

	if comprobacion[c.Param("nombre")].GetNombre() == "" {
		c.JSON(200, gin.H{"ID Registro Eliminado": c.Param("nombre")})
		return
	}

	c.JSON(400, gin.H{"Error": "Bad Request"})
	return
}
