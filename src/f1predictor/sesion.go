package f1predictor

/* Clase donde se relacionan los registros de un piloto en una determinada sesión. Los atributos
 * son los siguientes:
 *
 *		- Tiempos que ha realizado el piloto en cada vuelta
 * 		- Piloto con el que se relacionan los tiempos
 *		- Mejor tiempo de la sesión realizado por el piloto
 */

type SesionPiloto struct {
	tiempos     []TiempoVuelta
	piloto      Piloto
	mejorTiempo TiempoVuelta
}

type MetodosSesionPiloto interface {
	Constructor(tiempos []TiempoVuelta, piloto Piloto, mejorTiempo TiempoVuelta)

	GetTiempos() []TiempoVuelta
	SetTiempos(tiempos []TiempoVuelta)

	GetPiloto() Piloto
	SetPiloto(piloto Piloto)

	GetMejorTiempo() TiempoVuelta
	SetMejorTiempo(mejorTiempo TiempoVuelta)
}

func (sp *SesionPiloto) Constructor(tiempos []TiempoVuelta, piloto Piloto,
	mejorTiempo TiempoVuelta) {
	sp.mejorTiempo = mejorTiempo
	sp.piloto = piloto
	sp.tiempos = tiempos
}

func (sp SesionPiloto) GetTiempos() []TiempoVuelta {
	return sp.tiempos
}

func (sp *SesionPiloto) SetTiempos(tiempos []TiempoVuelta) {
	sp.tiempos = tiempos
}

func (sp SesionPiloto) GetPiloto() Piloto {
	return sp.piloto
}

func (sp *SesionPiloto) SetPiloto(piloto Piloto) {
	sp.piloto = piloto
}

func (sp SesionPiloto) GetMejorTiempo() TiempoVuelta {
	return sp.mejorTiempo
}

func (sp *SesionPiloto) SetMejorTiempo(mejorTiempo TiempoVuelta) {
	sp.mejorTiempo = mejorTiempo
}
