package f1predictor

/* Clase donde se define una escudería (o equipo) de Fórmula 1. En el se encuentran
 * los siguientes atributos:
 *
 * 		- Nombre de la escudería
 *		- Pilotos que la componen
 * 		- Títulos mundiales conseguidos
 * 		- Victorias conseguidas
 *		- Pole positions conseguidas
 * 		- Vueltas rápidas conseguidas
 */

type Escuderia struct {
	nombre           string
	pilotos          [2]Piloto
	titulosMundiales int
	victorias        int
	poles            int
	vueltasRapidas   int
}

type MetodosEscuderia interface {
	Constructor(nombre string, pilotos [2]Piloto, titulos int, victorias int, poles int, vr int)

	GetNombre() string
	SetNombre(nombre string)

	GetPilotos() [2]Piloto
	SetPilotos(pilotos [2]Piloto)

	GetTitulosMundiales() int
	SetTitulosMundiales(titulos int)

	GetVictorias() int
	SetVictorias(victorias int)

	GetPoles() int
	SetPoles(poles int)

	GetVueltasRapidas() int
	SetVueltasRapidas(vr int)
}

func (e *Escuderia) Constructor(nombre string, pilotos [2]Piloto, titulos int,
	victorias int, poles int, vr int) {

	e.nombre = nombre
	e.pilotos = pilotos
	e.titulosMundiales = titulos
	e.victorias = victorias
	e.poles = poles
	e.vueltasRapidas = vr

}

func (e *Escuderia) GetNombre() string {
	return e.nombre
}

func (e *Escuderia) SetNombre(nombre string) {
	e.nombre = nombre
}

func (e *Escuderia) GetPilotos() [2]Piloto {
	return e.pilotos
}

func (e *Escuderia) SetPilotos(pilotos [2]Piloto) {
	e.pilotos = pilotos
}

func (e *Escuderia) GetTitulosMundiales() int {
	return e.titulosMundiales
}

func (e *Escuderia) SetTitulosMundiales(titulosMundiales int) {
	e.titulosMundiales = titulosMundiales
}

func (e *Escuderia) GetVictorias() int {
	return e.victorias
}

func (e *Escuderia) SetVictorias(victorias int) {
	e.victorias = victorias
}

func (e *Escuderia) GetPoles() int {
	return e.poles
}

func (e *Escuderia) SetPoles(poles int) {
	e.poles = poles
}

func (e *Escuderia) GetVueltasRapidas() int {
	return e.vueltasRapidas
}

func (e *Escuderia) SetVueltasRapidas(vr int) {
	e.vueltasRapidas = vr
}
