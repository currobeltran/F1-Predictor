package f1predictor

import "strconv"

/*
 * Clase que recoge las estadísticas de un GP concreto. Entre sus atributos están:
 *
 * 		- Número de Accidentes en el GP
 * 		- Número de coches de seguridad
 * 		- Número de Adelantamientos
 *		- Número de banderas amarillas
 * 		- Número de banderas rojas
 *		- Número de Sanciones
 * 		- Tiempo de la mejor vuelta del GP
 */

type EstadisticasGP struct {
	Accidentes        int
	NumeroSafetyCar   int
	Adelantamientos   int
	BanderasAmarillas int
	BanderasRojas     int
	Sanciones         int
	MejorVuelta       TiempoVuelta
}

type MetodosEstadistica interface {
	Constructor(ac int, nsc int, ad int, ba int, br int, san int, mv TiempoVuelta)

	GetAccidentes() int
	SetAccidentes(Accidentes int)

	GetNumeroSafetyCar() int
	SetNumeroSafetyCar(nsc int)

	GetAdelantamientos() int
	SetAdelantamientos(Adelantamientos int)

	GetBanderasAmarillas() int
	SetBanderasAmarillas(BanderasAmarillas int)

	GetBanderasRojas() int
	SetBanderasRojas(BanderasRojas int)

	GetSanciones() int
	SetSanciones(Sanciones int)

	GetMejorVuelta() TiempoVuelta
	SetMejorVuelta(MejorVuelta TiempoVuelta)

	VerEstadisticasGP(ac bool, nsc bool, ad bool, ba bool, br bool, san bool, mv bool) string
}

func (e *EstadisticasGP) Constructor(ac int, nsc int, ad int, ba int,
	br int, san int, mv TiempoVuelta) {

	e.Accidentes = ac
	e.Adelantamientos = ad
	e.BanderasAmarillas = ba
	e.BanderasRojas = br
	e.MejorVuelta = mv
	e.NumeroSafetyCar = nsc
	e.Sanciones = san
}

func (e EstadisticasGP) GetAccidentes() int {
	return e.Accidentes
}

func (e *EstadisticasGP) SetAccidentes(Accidentes int) {
	e.Accidentes = Accidentes
}

func (e EstadisticasGP) GetNumeroSafetyCar() int {
	return e.NumeroSafetyCar
}

func (e *EstadisticasGP) SetNumeroSafetyCar(nsc int) {
	e.NumeroSafetyCar = nsc
}

func (e EstadisticasGP) GetAdelantamientos() int {
	return e.Adelantamientos
}

func (e *EstadisticasGP) SetAdelantamientos(Adelantamientos int) {
	e.Adelantamientos = Adelantamientos
}

func (e EstadisticasGP) GetBanderasAmarillas() int {
	return e.BanderasAmarillas
}

func (e *EstadisticasGP) SetBanderasAmarillas(BanderasAmarillas int) {
	e.BanderasAmarillas = BanderasAmarillas
}

func (e EstadisticasGP) GetBanderasRojas() int {
	return e.BanderasRojas
}

func (e *EstadisticasGP) SetBanderasRojas(BanderasRojas int) {
	e.BanderasRojas = BanderasRojas
}

func (e EstadisticasGP) GetSanciones() int {
	return e.Sanciones
}

func (e *EstadisticasGP) SetSanciones(Sanciones int) {
	e.Sanciones = Sanciones
}

func (e EstadisticasGP) GetMejorVuelta() TiempoVuelta {
	return e.MejorVuelta
}

func (e *EstadisticasGP) SetMejorVuelta(MejorVuelta TiempoVuelta) {
	e.MejorVuelta = MejorVuelta
}

/* Método para visualizar las estadísticas de un Gran Premio de manera ordenada
 *
 */

func (e EstadisticasGP) VerEstadisticasGP(ac bool, nsc bool, ad bool,
	ba bool, br bool, san bool, mv bool) string {

	var ret string

	if ad {
		var Adelantamientos string = ("Número de Adelantamientos: " + strconv.Itoa(e.Adelantamientos) + "\n")
		ret += Adelantamientos
	}

	if ac {
		var Accidentes string = ("Número de Accidentes: " + strconv.Itoa(e.Accidentes) + "\n")
		ret += Accidentes
	}

	if nsc {
		var safetycars string = ("Número de coches de seguridad: " + strconv.Itoa(e.NumeroSafetyCar) + "\n")
		ret += safetycars
	}

	if ba {
		var BanderasAmarillas string = ("Número de banderas amarillas: " + strconv.Itoa(e.BanderasAmarillas) + "\n")
		ret += BanderasAmarillas
	}

	if br {
		var BanderasRojas string = ("Número de banderas rojas: " + strconv.Itoa(e.BanderasRojas) + "\n")
		ret += BanderasRojas
	}

	if san {
		var Sanciones string = ("Número de Sanciones: " + strconv.Itoa(e.Sanciones) + "\n")
		ret += Sanciones
	}

	if mv {
		var mejor string = (strconv.Itoa(e.MejorVuelta.GetMinuto()) + ":" +
			strconv.Itoa(e.MejorVuelta.GetSegundo()) + "." + strconv.Itoa(e.MejorVuelta.GetMilesima()))

		var MejorVuelta string = ("Tiempo de la mejor vuelta: " + mejor + "\n")

		ret += MejorVuelta
	}

	return ret
}
