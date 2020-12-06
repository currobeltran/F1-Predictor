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

	esc := r.Group("/api/escuderia/:nombre")
	{
		esc.GET("", func(c *gin.Context) {
			escuderia.Get(c)
		})
	}

	circ := r.Group("/api/circuito/:nombre")
	{
		circ.GET("", func(c *gin.Context) {
			circuito.Get(c)
		})
		circ.PUT("", func(c *gin.Context) {
			circuito.Put(c)
		})
	}

	pil := r.Group("/api/piloto/:nombre")
	{
		pil.GET("", func(c *gin.Context) {
			piloto.Get(c)
		})
	}

	gp := r.Group("api/gp/:nombreCircuito/:temporada")
	{
		gp.GET("/estadisticas", func(c *gin.Context) {
			estadistica.Get(c)
		})
		gp.GET("/sesion/:nombreSesion", func(c *gin.Context) {
			sesion.Get(c)
		})
	}

	r.Run()
}
