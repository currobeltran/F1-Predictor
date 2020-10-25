package f1predictor

import "strconv"

/*
 * Clase que recoge las estadísticas de un GP concreto. Entre sus atributos están:
 *
 * 		- Número de accidentes en el GP
 * 		- Número de coches de seguridad
 * 		- Número de adelantamientos
 *		- Número de banderas amarillas
 * 		- Número de banderas rojas
 *		- Número de sanciones
 * 		- Tiempo de la mejor vuelta del GP
 */

type EstadisticasGP struct {
	accidentes        int
	numeroSafetyCar   int
	adelantamientos   int
	banderasAmarillas int
	banderasRojas     int
	sanciones         int
	mejorVuelta       TiempoVuelta
}

type MetodosEstadistica interface {
	Constructor(ac int, nsc int, ad int, ba int, br int, san int, mv TiempoVuelta)

	GetAccidentes() int
	SetAccidentes(accidentes int)

	GetNumeroSafetyCar() int
	SetNumeroSafetyCar(nsc int)

	GetAdelantamientos() int
	SetAdelantamientos(adelantamientos int)

	GetBanderasAmarillas() int
	SetBanderasAmarillas(banderasAmarillas int)

	GetBanderasRojas() int
	SetBanderasRojas(banderasRojas int)

	GetSanciones() int
	SetSanciones(sanciones int)

	GetMejorVuelta() TiempoVuelta
	SetMejorVuelta(mejorVuelta TiempoVuelta)

	VerEstadisticasGP(ac bool, nsc bool, ad bool, ba bool, br bool, san bool, mv bool) string
}

func (e *EstadisticasGP) Constructor(ac int, nsc int, ad int, ba int,
	br int, san int, mv TiempoVuelta) {

	e.accidentes = ac
	e.adelantamientos = ad
	e.banderasAmarillas = ba
	e.banderasRojas = br
	e.mejorVuelta = mv
	e.numeroSafetyCar = nsc
	e.sanciones = san
}

func (e EstadisticasGP) GetAccidentes() int {
	return e.accidentes
}

func (e *EstadisticasGP) SetAccidentes(accidentes int) {
	e.accidentes = accidentes
}

func (e EstadisticasGP) GetNumeroSafetyCar() int {
	return e.numeroSafetyCar
}

func (e *EstadisticasGP) SetNumeroSafetyCar(nsc int) {
	e.numeroSafetyCar = nsc
}

func (e EstadisticasGP) GetAdelantamientos() int {
	return e.adelantamientos
}

func (e *EstadisticasGP) SetAdelantamientos(adelantamientos int) {
	e.adelantamientos = adelantamientos
}

func (e EstadisticasGP) GetBanderasAmarillas() int {
	return e.banderasAmarillas
}

func (e *EstadisticasGP) SetBanderasAmarillas(banderasAmarillas int) {
	e.banderasAmarillas = banderasAmarillas
}

func (e EstadisticasGP) GetBanderasRojas() int {
	return e.banderasRojas
}

func (e *EstadisticasGP) SetBanderasRojas(banderasRojas int) {
	e.banderasRojas = banderasRojas
}

func (e EstadisticasGP) GetSanciones() int {
	return e.sanciones
}

func (e *EstadisticasGP) SetSanciones(sanciones int) {
	e.sanciones = sanciones
}

func (e EstadisticasGP) GetMejorVuelta() TiempoVuelta {
	return e.mejorVuelta
}

func (e *EstadisticasGP) SetMejorVuelta(mejorVuelta TiempoVuelta) {
	e.mejorVuelta = mejorVuelta
}

/* Método para visualizar las estadísticas de un Gran Premio de manera ordenada
 *
 */

func (e EstadisticasGP) VerEstadisticasGP(ac bool, nsc bool, ad bool,
	ba bool, br bool, san bool, mv bool) string {

	var ret string

	if ad {
		var adelantamientos string = ("Número de adelantamientos: " + strconv.Itoa(e.adelantamientos) + "\n")
		ret += adelantamientos
	}

	if ac {
		var accidentes string = ("Número de accidentes: " + strconv.Itoa(e.accidentes) + "\n")
		ret += accidentes
	}

	if nsc {
		var safetycars string = ("Número de coches de seguridad: " + strconv.Itoa(e.numeroSafetyCar) + "\n")
		ret += safetycars
	}

	if ba {
		var banderasAmarillas string = ("Número de banderas amarillas: " + strconv.Itoa(e.banderasAmarillas) + "\n")
		ret += banderasAmarillas
	}

	if br {
		var banderasRojas string = ("Número de banderas rojas: " + strconv.Itoa(e.banderasRojas) + "\n")
		ret += banderasRojas
	}

	if san {
		var sanciones string = ("Número de sanciones: " + strconv.Itoa(e.sanciones) + "\n")
		ret += sanciones
	}

	if mv {
		var mejor string = (strconv.Itoa(e.mejorVuelta.GetMinuto()) + ":" +
			strconv.Itoa(e.mejorVuelta.GetSegundo()) + "." + strconv.Itoa(e.mejorVuelta.GetMilesima()))

		var mejorVuelta string = ("Tiempo de la mejor vuelta: " + mejor + "\n")

		ret += mejorVuelta
	}

	return ret
}
