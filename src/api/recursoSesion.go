package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"m/src/f1predictor"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//RecursoSesion : Struct que representará el recurso relacionado con una sesión de un GP
type RecursoSesion struct {
}

//Get : Metodo para obtener la información de una sesion de un Gran Premio concreto
func (api RecursoSesion) Get(c *gin.Context) {
	data, err := ioutil.ReadFile("data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var resultados map[string][]f1predictor.ResultadoGP
	err = json.Unmarshal(data, &resultados)
	if err != nil {
		//TODO Logger
	}

	var temporadaEscogida f1predictor.ResultadoGP
	temporadaNum, _ := strconv.Atoi(c.Param("temporada"))

	for x := 0; x < len(resultados[c.Param("nombreCircuito")]); x++ {
		if resultados[c.Param("nombreCircuito")][x].GetTemporada() == temporadaNum {
			temporadaEscogida = resultados[c.Param("nombreCircuito")][x]
		}
	}

	if temporadaEscogida.GetTemporada() == 0 {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	switch c.Param("nombreSesion") {
	case "fp1":
		c.JSON(200, temporadaEscogida.GetResultadoFP1())
		break
	case "fp2":
		c.JSON(200, temporadaEscogida.GetResultadoFP2())
		break
	case "fp3":
		c.JSON(200, temporadaEscogida.GetResultadoFP3())
		break
	case "clasificacion":
		c.JSON(200, temporadaEscogida.GetResultadoClasificacion())
		break
	case "carrera":
		c.JSON(200, temporadaEscogida.GetResultadoCarrera())
		break
	default:
		c.JSON(400, gin.H{"Error": "Bad Request"})
	}

}

//Put : Método para modificar la información de una sesión
func (api RecursoSesion) Put(c *gin.Context) {
	data, err := ioutil.ReadFile("data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var resultados map[string][]f1predictor.ResultadoGP
	err = json.Unmarshal(data, &resultados)
	if err != nil {
		//TODO Logger
	}

	var temporadaEscogida f1predictor.ResultadoGP
	temporadaNum, _ := strconv.Atoi(c.Param("temporada"))

	for x := 0; x < len(resultados[c.Param("nombreCircuito")]); x++ {
		if resultados[c.Param("nombreCircuito")][x].GetTemporada() == temporadaNum {
			temporadaEscogida = resultados[c.Param("nombreCircuito")][x]
		}
	}

	if temporadaEscogida.GetTemporada() == 0 {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var sesionEscogida [20]f1predictor.SesionPiloto

	if c.PostForm("sesion") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesionEscogida = convertirStringSesion(c.PostForm("sesion"))

	switch c.Param("nombreSesion") {
	case "fp1":
		temporadaEscogida.SetResultadoFP1(sesionEscogida)
		break
	case "fp2":
		temporadaEscogida.SetResultadoFP1(sesionEscogida)
		break
	case "fp3":
		temporadaEscogida.SetResultadoFP1(sesionEscogida)
		break
	case "clasificacion":
		temporadaEscogida.SetResultadoFP1(sesionEscogida)
		break
	case "carrera":
		temporadaEscogida.SetResultadoFP1(sesionEscogida)
		break
	}

	for x := 0; x < len(resultados[c.Param("nombreCircuito")]); x++ {
		if resultados[c.Param("nombreCircuito")][x].GetTemporada() == temporadaNum {
			resultados[c.Param("nombreCircuito")][x] = temporadaEscogida
		}
	}

	fichero, err := json.Marshal(resultados)

	if err != nil {
		//TODO Logger
	}

	f, err := os.Create("data/resultados.json")
	if err != nil {
		//TODO Logger
	}

	f.Write(fichero)
	f.Close()

	c.JSON(200, sesionEscogida)
}

func convertirStringSesion(s string) [20]f1predictor.SesionPiloto {
	var ret [20]f1predictor.SesionPiloto

	clasif := strings.Split(s, "/")

	for i := 0; i < len(clasif); i++ {
		sesion := strings.Split(clasif[i], "#")
		fmt.Println(sesion)

		tiempos := sesion[0]
		piloto := sesion[1]
		mejortiempo := sesion[2]

		/*******************Formateamos tiempo************************/
		vectorTiempos := strings.Split(tiempos, "|")
		var tiemp []f1predictor.TiempoVuelta
		for j := 0; j < len(vectorTiempos); j++ {
			auxt := convertirTiempoString(vectorTiempos[j])
			tiemp = append(tiemp, auxt)
		}
		ret[i].SetTiempos(tiemp)

		/*******************Formateamos piloto************************/
		var aux f1predictor.Piloto
		datosPiloto := strings.Split(piloto, "|")

		nombre := datosPiloto[0]
		aux.SetNombre(nombre)

		nvic, _ := strconv.Atoi(datosPiloto[1])
		aux.SetVictorias(nvic)

		npol, _ := strconv.Atoi(datosPiloto[2])
		aux.SetPoles(npol)

		nvr, _ := strconv.Atoi(datosPiloto[3])
		aux.SetVueltasRapidas(nvr)

		nm, _ := strconv.Atoi(datosPiloto[4])
		aux.SetMundiales(nm)

		ret[i].SetPiloto(aux)

		/*******************Formateamos mejor tiempo************************/
		tiempoFormateado := convertirTiempoString(mejortiempo)
		ret[i].SetMejorTiempo(tiempoFormateado)
	}

	return ret
}
