package handler

import (
	"encoding/json"
	"fmt"
	f1predictor "m/src/f1predictor"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

/*Response : Struct que contendrá la respuesta en formato JSON*/
type Response struct {
	Pole          string `json:"Pole"`
	Clasificacion string `json:"Clasificación"`
}

/*Variables necesarias para realizar la operación*/
var hamilton f1predictor.Piloto
var melbourne f1predictor.Circuito
var resultados []f1predictor.ResultadoGP
var resultado f1predictor.ResultadoGP
var lista [20]f1predictor.SesionPiloto
var hamiltonTiempos f1predictor.SesionPiloto
var tiempos []f1predictor.TiempoVuelta
var record f1predictor.TiempoVuelta

/*Init tendrá como objetivo inicializar todas las variables anteriores
  Estos datos normalmente estarían almacenados en una base de datos, pero aún
  no se ha desarrollado el sistema. Por tanto, para su uso actual, podemos
  utilizar esta información de prueba y comprobar que el método funciona correctamente.
*/
func Init() {
	hamilton.Constructor("Lewis Hamilton", 94, 92, 55, 7)

	melbourne.Constructor("Melbourne", "Australia")
	resultado.SetPoleman(hamilton)
	resultado.SetTemporada(2020)

	record.Constructor(1, 23, 456)
	tiempos = append(tiempos, record)

	hamiltonTiempos.Constructor(tiempos, hamilton, record)

	lista[0] = hamiltonTiempos

	resultado.SetResultadoClasificacion(lista)
	resultados = append(resultados, resultado)

	melbourne.SetResultados(resultados)
}

/*Handler : Controla la llegada de una petición HTTP y la creación de una respuesta*/
func Handler(w http.ResponseWriter, r *http.Request) {
	Init()
	defer r.Body.Close()
	var err string

	if r.URL.String() != "/api/clasificacion" {

		stringBody := strings.Split(r.URL.String(), "&")
		match, _ := regexp.MatchString("year", stringBody[0])

		if !match {
			err = "Error al realizar la petición"
		} else {
			temporada := strings.Split(stringBody[0], "=")
			ntemp, _ := strconv.Atoi(temporada[1])

			result := melbourne.GetResultadoByTemporada(ntemp)
			if result.GetTemporada() == 0 {
				err = "No existe la temporada solicitada"
			} else {
				pole := result.GetPoleman().GetNombre()
				class := f1predictor.ObtenerClasificacion(melbourne, ntemp)

				if len(stringBody) <= 1 {
					w.Header().Add("Content-Type", "text/html")
					fmt.Fprintf(w, `
					<!DOCTYPE html>

					<html lang="es">
						<head>
							<meta charset="utf-8">
							<title>Clasificación</title>
						</head>
					
						<body>
							<h1>Clasificaciones históricas del GP de Australia: Año `+temporada[1]+`</h1>
							<ul>
								<li>Poleman: `+pole+` </li>
								<li>Clasificación: `+class+` </li>
							</ul>
						</body>
					</html>
					`)
				} else {
					data := Response{Pole: pole, Clasificacion: class}

					msg, _ := json.Marshal(data)
					w.Header().Add("Content-Type", "application/json")
					fmt.Fprintf(w, string(msg))
				}
			}
		}
	} else {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, `
		<!DOCTYPE html>

		<html lang="es">
			<head>
				<meta charset="utf-8">
				<title>Clasificación</title>
			</head>
			
			<body>
				<h1>Clasificaciones históricas del GP de Australia</h1>
				<form method="GET">
					<label>
						Introduzca la temporada
						<input type="number" name="year" />
					</label>
					<label>
						¿Devolver como JSON?
						<input type="checkbox" name="json" />
					</label>
					<button type="submit">Aceptar</button>
				</form>
			</body>
		</html>
		`)
	}
	if err != "" {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, `
		<!DOCTYPE html>

		<html lang="es">
			<head>
				<meta charset="utf-8">
				<title>Clasificación</title>
			</head>
		
			<body>
				<h1>`+err+`</h1>
			</body>
		</html>
		`)
	}
}
