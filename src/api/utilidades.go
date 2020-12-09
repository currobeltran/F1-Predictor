package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"m/src/f1predictor"
	"os"
	"strconv"
	"strings"
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
	fichero, err := json.Marshal(obj)

	if err != nil {
		//TODO Logger
	}

	f, err := os.Create(rutaArchivo)
	if err != nil {
		//TODO Logger
	}

	f.Write(fichero)
	f.Close()
}
