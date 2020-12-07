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

	pil := r.Group("/api/piloto/:nombre")
	{
		pil.GET("", func(c *gin.Context) {
			piloto.Get(c)
		})
		pil.PUT("", func(c *gin.Context) {
			piloto.Put(c)
		})
	}

	gp := r.Group("api/gp/:nombreCircuito/:temporada")
	{
		gp.GET("/estadisticas", func(c *gin.Context) {
			estadistica.Get(c)
		})
		gp.PUT("/estadisticas", func(c *gin.Context) {
			estadistica.Put(c)
		})

		gp.GET("/sesion/:nombreSesion", func(c *gin.Context) {
			sesion.Get(c)
		})
		gp.PUT("/sesion/:nombreSesion", func(c *gin.Context) {
			sesion.Put(c)
		})
	}

	r.Run()
}
