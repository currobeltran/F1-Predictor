package f1predictor

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

func (e *EstadisticasGP) GetAccidentes() int {
	return e.accidentes
}

func (e *EstadisticasGP) SetAccidentes(accidentes int) {
	e.accidentes = accidentes
}

func (e *EstadisticasGP) GetNumeroSafetyCar() int {
	return e.numeroSafetyCar
}

func (e *EstadisticasGP) SetNumeroSafetyCar(nsc int) {
	e.numeroSafetyCar = nsc
}

func (e *EstadisticasGP) GetAdelantamientos() int {
	return e.adelantamientos
}

func (e *EstadisticasGP) SetAdelantamientos(adelantamientos int) {
	e.adelantamientos = adelantamientos
}

func (e *EstadisticasGP) GetBanderasAmarillas() int {
	return e.banderasAmarillas
}

func (e *EstadisticasGP) SetBanderasAmarillas(banderasAmarillas int) {
	e.banderasAmarillas = banderasAmarillas
}

func (e *EstadisticasGP) GetBanderasRojas() int {
	return e.banderasRojas
}

func (e *EstadisticasGP) SetBanderasRojas(banderasRojas int) {
	e.banderasRojas = banderasRojas
}

func (e *EstadisticasGP) GetSanciones() int {
	return e.sanciones
}

func (e *EstadisticasGP) SetSanciones(sanciones int) {
	e.sanciones = sanciones
}

func (e *EstadisticasGP) GetMejorVuelta() TiempoVuelta {
	return e.mejorVuelta
}

func (e *EstadisticasGP) SetMejorVuelta(mejorVuelta TiempoVuelta) {
	e.mejorVuelta = mejorVuelta
}
