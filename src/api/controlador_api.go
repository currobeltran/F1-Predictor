package api

import (
	"m/src/f1predictor"

	"github.com/gin-gonic/gin"
)

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
