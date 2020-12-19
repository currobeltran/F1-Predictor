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

func obtenerPiloto(id string) f1predictor.Piloto {
	var ret f1predictor.Piloto

	data, err := ioutil.ReadFile("../api/data/pilotos.json")
	if err != nil {
		return ret
	}

	var pilotos map[string]f1predictor.Piloto
	err = json.Unmarshal(data, &pilotos)
	if err != nil {
		return ret
	}

	pilotoEscogido := pilotos[id]

	if pilotoEscogido.GetNombre() == "" {
		return ret
	}

	return pilotoEscogido
}

func existeCircuito(id string) bool {
	data, err := ioutil.ReadFile("../api/data/circuitos.json")
	if err != nil {
		return false
	}

	var campeonato map[string]f1predictor.Circuito
	err = json.Unmarshal(data, &campeonato)
	if err != nil {
		return false
	}

	circuitoEscogido := campeonato[id]

	if circuitoEscogido.GetNombre() == "" {
		return false
	}

	return true
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
		aux = obtenerPiloto(piloto)

		ret[i].SetPiloto(aux)

		/*******************Formateamos mejor tiempo************************/
		tiempoFormateado := convertirTiempoString(mejortiempo)
		ret[i].SetMejorTiempo(tiempoFormateado)
	}

	return ret
}

func convertirTiempoString(s string) f1predictor.TiempoVuelta {
	var ret f1predictor.TiempoVuelta

	variables := strings.Split(s, ":")

	minuto := strings.Split(variables[0], ":")[0]
	nmin, _ := strconv.Atoi(minuto)
	ret.SetMinuto(nmin)

	segundo := strings.Split(variables[1], ":")[0]
	ns, _ := strconv.Atoi(segundo)
	ret.SetSegundo(ns)

	milesima := strings.Split(variables[2], ":")[0]
	nmil, _ := strconv.Atoi(milesima)
	ret.SetMilesima(nmil)

	return ret
}

func convertirPodioString(s string) [3]f1predictor.Piloto {
	var ret [3]f1predictor.Piloto

	stringPodio := strings.Split(s, "#")
	for i := 0; i < len(stringPodio); i++ {
		aux := obtenerPiloto(stringPodio[i])
		ret[i] = aux
	}

	return ret
}

func escribirEnFichero(obj interface{}, rutaArchivo string) {
	fichero, _ := json.Marshal(obj)

	f, _ := os.Create(rutaArchivo)

	f.Write(fichero)
	f.Close()
}

//DiseñoRutas : Función para crear las rutas del microservicio
//Se ponen en una función aparte para el correcto testeo
func DiseñoRutas() *gin.Engine {
	circuito := new(RecursoCircuito)
	escuderia := new(RecursoEscuderia)
	estadistica := new(RecursoEstadisticas)
	piloto := new(RecursoPiloto)
	sesion := new(RecursoSesion)

	r := gin.Default()

	var hamilton f1predictor.Piloto
	hamilton.Constructor("Lewis Hamilton", 92, 98, 55, 7)

	var bottas f1predictor.Piloto
	bottas.Constructor("Valtteri Bottas", 10, 12, 14, 0)

	var pilotos []f1predictor.Piloto
	pilotos = append(pilotos, hamilton)
	pilotos = append(pilotos, bottas)

	var melbourne f1predictor.Circuito
	melbourne.Constructor("Albert Park", "Australia")

	var mercedes f1predictor.Escuderia
	mercedes.Constructor("Mercedes", pilotos, 8, 170, 200, 180)

	var tiempo f1predictor.TiempoVuelta
	tiempo.Constructor(1, 23, 456)

	var est f1predictor.EstadisticasGP
	est.Constructor(10, 9, 8, 7, 6, 5, tiempo)

	var lista [20]f1predictor.SesionPiloto
	var sespil f1predictor.SesionPiloto
	var tiempos []f1predictor.TiempoVuelta
	tiempos = append(tiempos, tiempo)
	sespil.Constructor(tiempos, hamilton, tiempo)
	lista[0] = sespil

	podio := [3]f1predictor.Piloto{hamilton, bottas, bottas}

	var res f1predictor.ResultadoGP
	res.Constructor(lista, lista, lista, lista, lista, hamilton, hamilton, podio, 2019, est)

	circuito.AnadirPista(melbourne)
	escuderia.AnadirEscuderia(mercedes)
	piloto.AnadirPiloto(hamilton)
	estadistica.AnadirEstadisticas(est, "Albert Park 2019")
	sesion.AnadirResultado(res, "Albert Park 2019")

	esc := r.Group("/api/escuderia")
	{
		esc.GET("/:nombre", func(c *gin.Context) {
			escuderia.Get(c)
		})
	}

	circ := r.Group("/api/circuito")
	{
		circ.GET("/:nombre", func(c *gin.Context) {
			circuito.Get(c)
		})
	}

	pil := r.Group("/api/piloto")
	{
		pil.GET("/:nombre", func(c *gin.Context) {
			piloto.Get(c)
		})
	}

	gp := r.Group("api/gp")
	{
		gp.GET("/:nombreCircuito/:temporada/estadisticas", func(c *gin.Context) {
			estadistica.Get(c)
		})

		gp.GET("/:nombreCircuito/:temporada/sesion/:nombreSesion", func(c *gin.Context) {
			sesion.Get(c)
		})
	}

	proc := r.Group("api/procesogp/:nombreCircuito")
	{
		proc.GET("media/:parametro", func(c *gin.Context) {
			estadistica.GetMedia(c)
		})
		proc.GET("prediccion/:parametro", func(c *gin.Context) {
			estadistica.GetPrediccion(c)
		})
	}

	return r
}
