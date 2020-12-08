package main

import (
	recursos "m/src/api"

	"github.com/gin-gonic/gin"
)

func main() {
	circuito := new(recursos.RecursoCircuito)
	escuderia := new(recursos.RecursoEscuderia)
	estadistica := new(recursos.RecursoEstadisticas)
	piloto := new(recursos.RecursoPiloto)
	sesion := new(recursos.RecursoSesion)
	granpremio := new(recursos.RecursoGranPremio)

	r := gin.Default()

	esc := r.Group("/api/escuderia")
	{
		esc.GET("/:nombre", func(c *gin.Context) {
			escuderia.Get(c)
		})
		esc.PUT("/:nombre", func(c *gin.Context) {
			escuderia.Put(c)
		})
		esc.POST("", func(c *gin.Context) {
			escuderia.Post(c)
		})
		esc.DELETE("/:nombre", func(c *gin.Context) {
			escuderia.Delete(c)
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
		circ.DELETE("/:nombre", func(c *gin.Context) {
			circuito.Delete(c)
		})
	}

	pil := r.Group("/api/piloto")
	{
		pil.GET("/:nombre", func(c *gin.Context) {
			piloto.Get(c)
		})
		pil.PUT("/:nombre", func(c *gin.Context) {
			piloto.Put(c)
		})
		pil.POST("", func(c *gin.Context) {
			piloto.Post(c)
		})
	}

	gp := r.Group("api/gp")
	{
		gp.GET("/:nombreCircuito/:temporada/estadisticas", func(c *gin.Context) {
			estadistica.Get(c)
		})
		gp.PUT("/:nombreCircuito/:temporada/estadisticas", func(c *gin.Context) {
			estadistica.Put(c)
		})

		gp.GET("/:nombreCircuito/:temporada/sesion/:nombreSesion", func(c *gin.Context) {
			sesion.Get(c)
		})
		gp.PUT("/:nombreCircuito/:temporada/sesion/:nombreSesion", func(c *gin.Context) {
			sesion.Put(c)
		})

		gp.POST("", func(c *gin.Context) {
			granpremio.Post(c)
		})
	}

	r.Run()
}
