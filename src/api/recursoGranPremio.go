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

//RecursoGranPremio : Estructura de datos que representara un recurso de Gran Premio
type RecursoGranPremio struct {
}

//Post : Método para crear un nuevo recurso de Gran Premio
func (api RecursoGranPremio) Post(c *gin.Context) {
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
	nuevoResultado.SetResultadoFP3(sesion)

	if c.PostForm("carrera") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	sesion = convertirStringSesion(c.PostForm("carrera"))
	nuevoResultado.SetResultadoFP3(sesion)

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

	piloto = obtenerPiloto(c.PostForm("poleman"))
	nuevoResultado.SetPoleman(piloto)

	var podio [3]f1predictor.Piloto
	if c.PostForm("podio") == "" {
		c.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}

	stringPodio := strings.Split(c.PostForm("podio"), "#")
	for i := 0; i < len(stringPodio); i++ {
		aux := obtenerPiloto(stringPodio[i])
		podio[i] = aux
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

	data, err := ioutil.ReadFile("data/resultados.json")
	if err != nil {
		c.JSON(404, gin.H{"Error": "Not Found"})
	}

	var resultados map[string][]f1predictor.ResultadoGP
	err = json.Unmarshal(data, &resultados)
	if err != nil {
		//TODO Logger
	}

	resultados[c.PostForm("nombre")] = append(resultados[c.PostForm("nombre")], nuevoResultado)

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

	c.JSON(201, nuevoResultado)
}

func convertirEstadisticasString(s string) f1predictor.EstadisticasGP {
	var ret f1predictor.EstadisticasGP

	atributos := strings.Split(s, "#")

	nac, _ := strconv.Atoi(atributos[0])
	ret.SetAccidentes(nac)

	nsc, _ := strconv.Atoi(atributos[1])
	ret.SetNumeroSafetyCar(nsc)

	nad, _ := strconv.Atoi(atributos[2])
	ret.SetAdelantamientos(nad)

	nba, _ := strconv.Atoi(atributos[3])
	ret.SetBanderasAmarillas(nba)

	nbr, _ := strconv.Atoi(atributos[4])
	ret.SetBanderasRojas(nbr)

	nsan, _ := strconv.Atoi(atributos[5])
	ret.SetSanciones(nsan)

	tiempo := convertirTiempoString(atributos[6])
	ret.SetMejorVuelta(tiempo)

	return ret
}

func obtenerPiloto(id string) f1predictor.Piloto {
	var ret f1predictor.Piloto

	data, err := ioutil.ReadFile("data/pilotos.json")
	if err != nil {
		return ret
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		//TODO Logger
	}

	pilotoEscogido := pilotos[id]

	if pilotoEscogido.GetNombre() == "" {
		return ret
	}

	return pilotoEscogido
}
