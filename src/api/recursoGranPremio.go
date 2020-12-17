package api

import (
	"encoding/json"
	"io/ioutil"
	"m/src/f1predictor"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

//RecursoGranPremio : Estructura de datos que representara un recurso de Gran Premio
type RecursoGranPremio struct {
}

//Post : Método para crear un nuevo recurso de Gran Premio
func (rGP RecursoGranPremio) Post(c *gin.Context) {
	if !existeCircuito(c.PostForm("nombre")) {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	var nuevoResultado f1predictor.ResultadoGP

	var sesion [20]f1predictor.SesionPiloto

	if c.PostForm("fp1") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesion = convertirStringSesion(c.PostForm("fp1"))
	nuevoResultado.SetResultadoFP1(sesion)

	if c.PostForm("fp2") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesion = convertirStringSesion(c.PostForm("fp2"))
	nuevoResultado.SetResultadoFP2(sesion)

	if c.PostForm("fp3") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesion = convertirStringSesion(c.PostForm("fp3"))
	nuevoResultado.SetResultadoFP3(sesion)

	if c.PostForm("clasificacion") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesion = convertirStringSesion(c.PostForm("clasificacion"))
	nuevoResultado.SetResultadoClasificacion(sesion)

	if c.PostForm("carrera") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesion = convertirStringSesion(c.PostForm("carrera"))
	nuevoResultado.SetResultadoCarrera(sesion)

	var piloto f1predictor.Piloto

	if c.PostForm("ganador") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	piloto = obtenerPiloto(c.PostForm("ganador"))
	nuevoResultado.SetGanador(piloto)

	if c.PostForm("poleman") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	poleman := obtenerPiloto(c.PostForm("poleman"))
	if poleman.GetNombre() == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nuevoResultado.SetPoleman(poleman)

	if c.PostForm("podio") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	podio := convertirPodioString(c.PostForm("podio"))
	if !reflect.DeepEqual(piloto, podio[0]) {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	nuevoResultado.SetPodio(podio)

	if c.PostForm("temporada") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ntemp, _ := strconv.Atoi(c.PostForm("temporada"))
	nuevoResultado.SetTemporada(ntemp)

	if c.PostForm("estadisticas") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	est := convertirEstadisticasString(c.PostForm("estadisticas"))
	nuevoResultado.SetEstadisticas(est)

	data, err := ioutil.ReadFile("../api/data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
		return
	}

	var resultados map[string][]f1predictor.ResultadoGP
	json.Unmarshal(data, &resultados)

	resultados[c.PostForm("nombre")] = append(resultados[c.PostForm("nombre")], nuevoResultado)

	escribirEnFichero(resultados, "../api/data/resultados.json")

	c.JSON(201, nuevoResultado)
}
