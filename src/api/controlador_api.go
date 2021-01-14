package api

import (
	"fmt"
	"m/src/f1predictor"
	"time"

	"github.com/gin-gonic/gin"
)

//DiseñoRutas : Función para crear las rutas del microservicio
//Se ponen en una función aparte para el correcto testeo
func DiseñoRutas() *gin.Engine {
	circuito := new(RecursoCircuito)
	campeonato := new(Campeonato)
	estadistica := new(RecursoEstadisticas)
	sesion := new(RecursoSesion)

	r := gin.New()

	//Creamos un middleware que realice una función de logger
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("Fecha: %s, Método: %s, Recurso: %s, Status: %d \n",
			param.TimeStamp.Format(time.ANSIC),
			param.Method,
			param.Path,
			param.StatusCode,
		)
	}))

	//Establecemos el middleware de recuperación tras errores 5xx
	r.Use(gin.Recovery())

	var hamilton f1predictor.Piloto
	hamilton.Constructor("Lewis Hamilton", 92, 98, 55, 7)

	var bottas f1predictor.Piloto
	bottas.Constructor("Valtteri Bottas", 10, 12, 14, 0)

	var vettel f1predictor.Piloto
	vettel.Constructor("Sebastian Vettel", 55, 30, 20, 4)

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
	campeonato.AnadirEscuderia(mercedes)
	campeonato.AnadirPiloto(hamilton)
	campeonato.AnadirPiloto(bottas)
	campeonato.AnadirPiloto(vettel)
	estadistica.AnadirEstadisticas(est, "Albert Park 2019")
	sesion.AnadirResultado(res, "Albert Park 2019")

	esc := r.Group("/api/escuderia")
	{
		esc.GET("/:nombre", func(c *gin.Context) {
			campeonato.GetEscuderia(c)
		})
		esc.PUT("/:nombre", func(c *gin.Context) {
			campeonato.PutEscuderia(c)
		})
		esc.POST("", func(c *gin.Context) {
			campeonato.PostEscuderia(c)
		})
	}

	circ := r.Group("/api/circuito")
	{
		circ.GET("/:nombre", func(c *gin.Context) {
			circuito.Get(c)
		})

		circ.PUT("/:nombre", func(c *gin.Context) {
			circuito.Put(c)
		})

		circ.POST("", func(c *gin.Context) {
			circuito.Post(c)
		})
	}

	pil := r.Group("/api/piloto")
	{
		pil.GET("/:nombre", func(c *gin.Context) {
			campeonato.GetPiloto(c)
		})
		pil.PUT("/:nombre", func(c *gin.Context) {
			campeonato.PutPiloto(c)
		})
		pil.POST("", func(c *gin.Context) {
			campeonato.PostPiloto(c)
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

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	return r
}
