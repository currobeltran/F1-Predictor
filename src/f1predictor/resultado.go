package f1predictor

import "strconv"

/* Clase donde se almacenan los resultados correspondientes a un Gran Premio. Los atributos son:
 *
 *		- Resultado de los Entrenamientos Libres 1
 *		- Resultado de los Entrenamientos Libres 2
 * 		- Resultado de los Entrenamientos Libres 3
 * 		- Resultado de la clasificación
 *		- Resultado de la carrera
 * 		- Ganador del Gran Premio
 * 		- Poleman del Gran Premio
 * 		- Podio correspondiente a la carrera
 *		- Número de la temporada
 * 		- Estadísticas del Gran Premio (adelantamientos, accidentes...)
 */
type ResultadoGP struct {
	resultadoFP1           [20]SesionPiloto
	resultadoFP2           [20]SesionPiloto
	resultadoFP3           [20]SesionPiloto
	resultadoClasificacion [20]SesionPiloto
	resultadoCarrera       [20]SesionPiloto
	ganador                Piloto
	poleman                Piloto
	podio                  [3]Piloto
	temporada              int
	estadisticas           EstadisticasGP
}

type MetodosResultado interface {
	Constructor(fp1 [20]SesionPiloto, fp2 [20]SesionPiloto, fp3 [20]SesionPiloto,
		qualy [20]SesionPiloto, carrera [20]SesionPiloto, gan Piloto, pole Piloto,
		podio [3]Piloto, nTemp int, est EstadisticasGP)

	GetResultadoFP1() [20]SesionPiloto
	SetResultadoFP1(fp1 [20]SesionPiloto)

	GetResultadoFP2() [20]SesionPiloto
	SetResultadoFP2(fp2 [20]SesionPiloto)

	GetResultadoFP3() [20]SesionPiloto
	SetResultadoFP3(fp3 [20]SesionPiloto)

	GetResultadoClasificacion() [20]SesionPiloto
	SetResultadoClasificacion(clasificacion [20]SesionPiloto)

	GetResultadoCarrera() [20]SesionPiloto
	SetResultadoCarrera(carrera [20]SesionPiloto)

	GetGanador() Piloto
	SetGanador(ganador Piloto)

	GetPoleman() Piloto
	SetPoleman(pole Piloto)

	GetPodio() [3]Piloto
	SetPodio(podio [3]Piloto)

	GetTemporada() int
	SetTemporada(nTemp int)

	GetEstadisticas() EstadisticasGP
	SetEstadisticas(est EstadisticasGP)

	VerClasificacionSesion(pilotos []SesionPiloto) string

	VerDatosGranPremio(fp1 bool, fp2 bool, fp3 bool, q bool, c bool, g bool,
		pol bool, p bool, t bool, e bool) string
}

func (r *ResultadoGP) Constructor(fp1 [20]SesionPiloto, fp2 [20]SesionPiloto, fp3 [20]SesionPiloto,
	qualy [20]SesionPiloto, carrera [20]SesionPiloto, gan Piloto, pole Piloto,
	podio [3]Piloto, nTemp int, est EstadisticasGP) {

	r.estadisticas = est
	r.ganador = gan
	r.podio = podio
	r.poleman = pole
	r.resultadoCarrera = carrera
	r.resultadoClasificacion = qualy
	r.resultadoFP1 = fp1
	r.resultadoFP2 = fp2
	r.resultadoFP3 = fp3
	r.temporada = nTemp
}

func (r *ResultadoGP) GetResultadoFP1() [20]SesionPiloto {
	return r.resultadoFP1
}

func (r *ResultadoGP) SetResultadoFP1(fp1 [20]SesionPiloto) {
	r.resultadoFP1 = fp1
}

func (r *ResultadoGP) GetResultadoFP2() [20]SesionPiloto {
	return r.resultadoFP2
}

func (r *ResultadoGP) SetResultadoFP2(fp2 [20]SesionPiloto) {
	r.resultadoFP2 = fp2
}
func (r *ResultadoGP) GetResultadoFP3() [20]SesionPiloto {
	return r.resultadoFP3
}

func (r *ResultadoGP) SetResultadoFP3(fp3 [20]SesionPiloto) {
	r.resultadoFP3 = fp3
}

func (r *ResultadoGP) GetResultadoClasificacion() [20]SesionPiloto {
	return r.resultadoClasificacion
}

func (r *ResultadoGP) SetResultadoClasificacion(clasificacion [20]SesionPiloto) {
	r.resultadoClasificacion = clasificacion
}

func (r *ResultadoGP) GetResultadoCarrera() [20]SesionPiloto {
	return r.resultadoCarrera
}

func (r *ResultadoGP) SetResultadoCarrera(carrera [20]SesionPiloto) {
	r.resultadoCarrera = carrera
}

func (r *ResultadoGP) GetGanador() Piloto {
	return r.ganador
}

func (r *ResultadoGP) SetGanador(ganador Piloto) {
	r.ganador = ganador
}

func (r *ResultadoGP) GetPoleman() Piloto {
	return r.poleman
}

func (r *ResultadoGP) SetPoleman(pole Piloto) {
	r.poleman = pole
}

func (r *ResultadoGP) GetPodio() [3]Piloto {
	return r.podio
}

func (r *ResultadoGP) SetPodio(podio [3]Piloto) {
	r.podio = podio
}

func (r *ResultadoGP) GetTemporada() int {
	return r.temporada
}

func (r *ResultadoGP) SetTemporada(nTemp int) {
	r.temporada = nTemp
}

func (r ResultadoGP) GetEstadisticas() EstadisticasGP {
	return r.estadisticas
}

func (r ResultadoGP) SetEstadisticas(est EstadisticasGP) {
	r.estadisticas = est
}

/* Método con el que se podrá visualizar de manera ordenada la clasificación de una
 * sesión del Gran Premio
 */
func (r ResultadoGP) VerClasificacionSesion(pilotos []SesionPiloto) string {
	var ret string

	for i := 0; i < len(pilotos); i++ {
		var pil SesionPiloto = pilotos[i]

		var mejor string = (strconv.Itoa(pil.GetMejorTiempo().GetMinuto()) + ":" +
			strconv.Itoa(pil.GetMejorTiempo().GetSegundo()) + "." + strconv.Itoa(pil.GetMejorTiempo().GetMilesima()))

		var linea string = (strconv.Itoa(i+1) + "º posición: " + pil.GetPiloto().GetNombre() +
			" Mejor Tiempo: " + mejor + "\n")

		ret += linea
	}

	return ret
}
