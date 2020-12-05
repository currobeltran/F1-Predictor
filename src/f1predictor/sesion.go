package f1predictor

import "strconv"

/* Clase donde se relacionan los registros de un Piloto en una determinada sesión. Los atributos
 * son los siguientes:
 *
 *		- Tiempos que ha realizado el Piloto en cada vuelta
 * 		- Piloto con el que se relacionan los Tiempos
 *		- Mejor tiempo de la sesión realizado por el Piloto
 */

type SesionPiloto struct {
	Tiempos     []TiempoVuelta
	Piloto      Piloto
	MejorTiempo TiempoVuelta
}

type MetodosSesionPiloto interface {
	Constructor(Tiempos []TiempoVuelta, Piloto Piloto, MejorTiempo TiempoVuelta)

	GetTiempos() []TiempoVuelta
	SetTiempos(Tiempos []TiempoVuelta)

	GetPiloto() Piloto
	SetPiloto(Piloto Piloto)

	GetMejorTiempo() TiempoVuelta
	SetMejorTiempo(MejorTiempo TiempoVuelta)

	VerClasificacionSesion(Pilotos [20]SesionPiloto) string
}

func (sp *SesionPiloto) Constructor(Tiempos []TiempoVuelta, Piloto Piloto,
	MejorTiempo TiempoVuelta) {
	sp.MejorTiempo = MejorTiempo
	sp.Piloto = Piloto
	sp.Tiempos = Tiempos
}

func (sp SesionPiloto) GetTiempos() []TiempoVuelta {
	return sp.Tiempos
}

func (sp *SesionPiloto) SetTiempos(Tiempos []TiempoVuelta) {
	sp.Tiempos = Tiempos
}

func (sp SesionPiloto) GetPiloto() Piloto {
	return sp.Piloto
}

func (sp *SesionPiloto) SetPiloto(Piloto Piloto) {
	sp.Piloto = Piloto
}

func (sp SesionPiloto) GetMejorTiempo() TiempoVuelta {
	return sp.MejorTiempo
}

func (sp *SesionPiloto) SetMejorTiempo(MejorTiempo TiempoVuelta) {
	sp.MejorTiempo = MejorTiempo
}

/* Método con el que se podrá visualizar de manera ordenada la clasificación de una
 * sesión del Gran Premio
 */
func VerClasificacionSesion(Pilotos [20]SesionPiloto) string {
	var ret string

	for i := 0; i < len(Pilotos); i++ {
		var pil SesionPiloto = Pilotos[i]

		var mejor string = (strconv.Itoa(pil.GetMejorTiempo().GetMinuto()) + ":" +
			strconv.Itoa(pil.GetMejorTiempo().GetSegundo()) + "." + strconv.Itoa(pil.GetMejorTiempo().GetMilesima()))

		var linea string = (strconv.Itoa(i+1) + "º posición: " + pil.GetPiloto().GetNombre() +
			" Mejor Tiempo: " + mejor + "\n")

		ret += linea
	}

	return ret
}
