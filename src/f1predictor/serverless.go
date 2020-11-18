package f1predictor

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

}

// ObtenerClasificacion: Función para obtener la información de una sesión de clasificación
// determinada para desplegarla utilizando un servicio serverless
func ObtenerClasificacion(c Circuito, temp int) string {
	var s string = ""

	result := c.GetResultadoByTemporada(temp)

	list := result.GetResultadoClasificacion()

	for i := 0; i < len(list); i++ {
		s += fmt.Sprint(i+1) + "ª posición: " + list[i].GetPiloto().GetNombre() + ", tiempo: " + list[i].GetMejorTiempo().ImprimirTiempo() + "\n"
	}

	return s
}
