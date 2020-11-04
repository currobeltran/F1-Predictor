package f1predictor

import "strconv"

/*
 * Clase circuito con la podremos acceder a distintos datos correspondientes
 * a un circuito del campeonato de Formula 1. Los atributos son:
 * 		- Nombre del circuito.
 * 		- País donde se encuentra.
 *		- Record de tiempo sobre la pista.
 *		- Acumulación histórica de resultados de los GPs disputados en dicha pista.
 */

type Circuito struct {
	nombre         string
	pais           string
	recordCircuito TiempoVuelta
	resultados     []ResultadoGP
}

type MetodosCircuito interface {
	Constructor(nombre string, pais string)

	GetNombre() string
	SetNombre(nombre string)

	GetPais() string
	SetPais(pais string)

	GetRecordCircuito() TiempoVuelta
	SetRecordCircuito(nuevoRecord TiempoVuelta)

	GetResultados() []ResultadoGP
	SetResultados(resultados []ResultadoGP)
	SetResultadoByIndex(i int, resultado ResultadoGP)

	Media(estadistica string) float32

	VerDatosCircuito(n bool, p bool, t bool) string
}

func (c *Circuito) Constructor(nombre string, pais string) {
	c.nombre = nombre
	c.pais = pais

	var tmp TiempoVuelta
	tmp.Constructor(0, 0, 0)
	c.recordCircuito = tmp

	c.resultados = make([]ResultadoGP, 1)
}

func (c Circuito) GetNombre() string {
	return c.nombre
}

func (c *Circuito) SetNombre(nombre string) {
	c.nombre = nombre
}

func (c Circuito) GetPais() string {
	return c.pais
}

func (c *Circuito) SetPais(pais string) {
	c.pais = pais
}

func (c Circuito) GetRecordCircuito() TiempoVuelta {
	return c.recordCircuito
}

func (c *Circuito) SetRecordCircuito(nuevoRecord TiempoVuelta) {
	c.recordCircuito = nuevoRecord
}

func (c Circuito) GetResultados() []ResultadoGP {
	return c.resultados
}

func (c *Circuito) SetResultados(resultados []ResultadoGP) {
	c.resultados = resultados
}

func (c *Circuito) SetResultadosByIndex(i int, resultado ResultadoGP) {
	c.resultados[i] = resultado
}

/*
 * Método para calcular la media de una estadística concreta de un circuito
 * concreto
 */

func (c Circuito) Media(estadistica string) float32 {
	var acumulado int = 0
	var x float32 = 0

	for i := 0; i < len(c.resultados); i++ {
		var est EstadisticasGP = c.resultados[i].GetEstadisticas()

		switch estadistica {
		case "accidentes":
			acumulado += est.GetAccidentes()
			break
		case "banderasAmarillas":
			acumulado += est.GetBanderasAmarillas()
			break
		case "banderasRojas":
			acumulado += est.GetBanderasRojas()
			break
		case "adelantamientos":
			acumulado += est.GetAdelantamientos()
			break
		case "sanciones":
			acumulado += est.GetSanciones()
			break
		case "numeroSafetyCar":
			acumulado += est.GetNumeroSafetyCar()
			break
		case "default":
			break
		}

	}

	x = float32(acumulado) / float32(len(c.resultados))

	return x
}

/*
 * Método para ver de manera ordenada la información correspondiente a un
 * circuito. Se puede seleccionar que datos se quieren ver y cuales no.
 */

func (c Circuito) VerDatosCircuito(n bool, p bool, t bool) string {
	var ret string

	if n {
		var nombre string = ("Nombre del circuito: " + c.nombre + "\n")
		ret += nombre
	}

	if p {
		var pais string = ("País: " + c.pais + "\n")
		ret += pais
	}

	if t {
		var record string = (strconv.Itoa(c.recordCircuito.GetMinuto()) + ":" +
			strconv.Itoa(c.recordCircuito.GetSegundo()) + "." +
			strconv.Itoa(c.recordCircuito.GetMilesima()) + "\n")

		var tiempo string = ("Record de la pista: " + record)

		ret += tiempo
	}

	return ret
}

/* Método para calcular un las posibilidades de que salga un SafetyCar en una carrera disputada
 * en el circuito.
 */
func (c Circuito) PosibilidadSafetyCar() float64 {
	var ret float64 = 0.0
	var suma float64 = 0.0
	var div float64 = 0.0
	var ponderacion float64 = 1.0

	for i := len(c.resultados) - 1; i >= 0; i-- {
		if ponderacion > 0 {
			var est EstadisticasGP = c.resultados[i].GetEstadisticas()
			suma += (float64(est.GetNumeroSafetyCar()) * ponderacion)

			div += ponderacion
			ponderacion -= 0.05
		}
	}

	ret = suma / div

	return ret
}

/* Método para calcular cuantos adelantamientos habrá en una carrera disputada en el circuito*/
func (c Circuito) PosibilidadAdelantamiento() float64 {
	var ret float64 = 0.0
	var suma float64 = 0.0
	var div float64 = 0.0
	var ponderacion float64 = 1.0

	for i := len(c.resultados) - 1; i >= 0; i-- {
		if ponderacion > 0 {
			var est EstadisticasGP = c.resultados[i].GetEstadisticas()
			suma += (float64(est.GetAdelantamientos()) * ponderacion)
			div += ponderacion
			ponderacion -= 0.05
		}
	}

	ret = suma / div

	return ret
}

/* Método para calcular la probable cantidad de banderas amarillas que habrá en un circuito durante
 * un Gran Premio
 */
func (c Circuito) PosibilidadBanderaAmarilla() float64 {
	var ret float64 = 0.0
	var suma float64 = 0.0
	var div float64 = 0.0
	var ponderacion float64 = 1.0

	for i := len(c.resultados) - 1; i >= 0; i-- {
		if ponderacion > 0 {
			var est EstadisticasGP = c.resultados[i].GetEstadisticas()
			suma += (float64(est.GetBanderasAmarillas()) * ponderacion)
			div += ponderacion
			ponderacion -= 0.05
		}
	}

	ret = suma / div

	return ret
}

//PosibilidadAccidentes
//Método para calcular la posible cantidad de accidentes que puedan acontecer en un Gran Premio
//disputado en un determinado circuito
func (c Circuito) PosibilidadAccidentes() float64 {
	var ret float64 = 0.0

	return ret
}
