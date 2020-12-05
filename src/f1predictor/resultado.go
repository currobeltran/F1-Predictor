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
 *		- Número de la Temporada
 * 		- Estadísticas del Gran Premio (adelantamientos, accidentes...)
 */
type ResultadoGP struct {
	ResultadoFP1           [20]SesionPiloto
	ResultadoFP2           [20]SesionPiloto
	ResultadoFP3           [20]SesionPiloto
	ResultadoClasificacion [20]SesionPiloto
	ResultadoCarrera       [20]SesionPiloto
	Ganador                Piloto
	Poleman                Piloto
	Podio                  [3]Piloto
	Temporada              int
	Estadisticas           EstadisticasGP
}

type MetodosResultado interface {
	Constructor(fp1 [20]SesionPiloto, fp2 [20]SesionPiloto, fp3 [20]SesionPiloto,
		qualy [20]SesionPiloto, carrera [20]SesionPiloto, gan Piloto, pole Piloto,
		Podio [3]Piloto, nTemp int, est EstadisticasGP)

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
	SetGanador(Ganador Piloto)

	GetPoleman() Piloto
	SetPoleman(pole Piloto)

	GetPodio() [3]Piloto
	SetPodio(Podio [3]Piloto)

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
	Podio [3]Piloto, nTemp int, est EstadisticasGP) {

	r.Estadisticas = est
	r.Ganador = gan
	r.Podio = Podio
	r.Poleman = pole
	r.ResultadoCarrera = carrera
	r.ResultadoClasificacion = qualy
	r.ResultadoFP1 = fp1
	r.ResultadoFP2 = fp2
	r.ResultadoFP3 = fp3
	r.Temporada = nTemp
}

func (r *ResultadoGP) GetResultadoFP1() [20]SesionPiloto {
	return r.ResultadoFP1
}

func (r *ResultadoGP) SetResultadoFP1(fp1 [20]SesionPiloto) {
	r.ResultadoFP1 = fp1
}

func (r *ResultadoGP) GetResultadoFP2() [20]SesionPiloto {
	return r.ResultadoFP2
}

func (r *ResultadoGP) SetResultadoFP2(fp2 [20]SesionPiloto) {
	r.ResultadoFP2 = fp2
}
func (r *ResultadoGP) GetResultadoFP3() [20]SesionPiloto {
	return r.ResultadoFP3
}

func (r *ResultadoGP) SetResultadoFP3(fp3 [20]SesionPiloto) {
	r.ResultadoFP3 = fp3
}

func (r *ResultadoGP) GetResultadoClasificacion() [20]SesionPiloto {
	return r.ResultadoClasificacion
}

func (r *ResultadoGP) SetResultadoClasificacion(clasificacion [20]SesionPiloto) {
	r.ResultadoClasificacion = clasificacion
}

func (r *ResultadoGP) GetResultadoCarrera() [20]SesionPiloto {
	return r.ResultadoCarrera
}

func (r *ResultadoGP) SetResultadoCarrera(carrera [20]SesionPiloto) {
	r.ResultadoCarrera = carrera
}

func (r *ResultadoGP) GetGanador() Piloto {
	return r.Ganador
}

func (r *ResultadoGP) SetGanador(Ganador Piloto) {
	r.Ganador = Ganador
}

func (r *ResultadoGP) GetPoleman() Piloto {
	return r.Poleman
}

func (r *ResultadoGP) SetPoleman(pole Piloto) {
	r.Poleman = pole
}

func (r *ResultadoGP) GetPodio() [3]Piloto {
	return r.Podio
}

func (r *ResultadoGP) SetPodio(Podio [3]Piloto) {
	r.Podio = Podio
}

func (r *ResultadoGP) GetTemporada() int {
	return r.Temporada
}

func (r *ResultadoGP) SetTemporada(nTemp int) {
	r.Temporada = nTemp
}

func (r ResultadoGP) GetEstadisticas() EstadisticasGP {
	return r.Estadisticas
}

func (r *ResultadoGP) SetEstadisticas(est EstadisticasGP) {
	r.Estadisticas = est
}

/* Método para visualizar los resultados de un Gran Premio de manera ordenada. Se puede seleccionar
 * que información se desea que aparezca.
 */
func (r ResultadoGP) VerDatosGranPremio(fp1 bool, fp2 bool, fp3 bool, q bool, c bool, g bool,
	pol bool, p bool, t bool, e bool) string {

	var ret string

	if fp1 {
		var clas string = VerClasificacionSesion(r.ResultadoFP1)
		var libres1 string = ("Resultado Libres 1: \n" + clas)

		ret += libres1
	}

	if fp2 {
		var clas string = VerClasificacionSesion(r.ResultadoFP2)
		var libres2 string = ("Resultado Libres 2: \n" + clas)

		ret += libres2
	}

	if fp3 {
		var clas string = VerClasificacionSesion(r.ResultadoFP3)
		var libres3 string = ("Resultado Libres 3: \n" + clas)

		ret += libres3
	}

	if q {
		var clas string = VerClasificacionSesion(r.ResultadoClasificacion)
		var qualy string = ("Resultado Clasificación: \n" + clas)

		ret += qualy
	}

	if c {
		var clas string = VerClasificacionSesion(r.ResultadoCarrera)
		var carrera string = ("Resultado Carrera: \n" + clas)

		ret += carrera
	}

	if g {
		var Ganador string = ("Ganador del Gran Premio: " + r.Ganador.GetNombre() + "\n")

		ret += Ganador
	}

	if pol {
		var Poleman string = ("Poleman del Gran Premio: " + r.Poleman.GetNombre() + "\n")

		ret += Poleman
	}

	if p {
		var Podio string = ("Podio del Gran Premio: \n1º posición: " + r.Podio[0].GetNombre() +
			"\n2º posición: " + r.Podio[1].GetNombre() + "\n3º posición: " + r.Podio[2].GetNombre() + "\n")

		ret += Podio
	}

	if t {
		var Temporada string = ("Gran Premio correspondiente a la Temporada: " + strconv.Itoa(r.Temporada) + "\n")

		ret += Temporada
	}

	if e {
		var Estadisticas string = r.Estadisticas.VerEstadisticasGP(true, true, true, true, true, true, true)

		ret += Estadisticas
	}

	return ret
}
