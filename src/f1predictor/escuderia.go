package f1predictor

import "strconv"

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
	pilotos          []Piloto
	titulosMundiales int
	victorias        int
	poles            int
	vueltasRapidas   int
}

type MetodosEscuderia interface {
	Constructor(nombre string, pilotos [2]Piloto, titulos int, victorias int, poles int, vr int)

	GetNombre() string
	SetNombre(nombre string)

	GetPilotos() []Piloto
	SetPilotos(pilotos []Piloto)

	GetTitulosMundiales() int
	SetTitulosMundiales(titulos int)

	GetVictorias() int
	SetVictorias(victorias int)

	GetPoles() int
	SetPoles(poles int)

	GetVueltasRapidas() int
	SetVueltasRapidas(vr int)

	VerDatosEscuderia(n bool, p bool, m bool, v bool, pol bool, vr bool) string
}

func (e *Escuderia) Constructor(nombre string, pilotos []Piloto, titulos int,
	victorias int, poles int, vr int) {

	e.nombre = nombre
	e.pilotos = pilotos
	e.titulosMundiales = titulos
	e.victorias = victorias
	e.poles = poles
	e.vueltasRapidas = vr

}

func (e Escuderia) GetNombre() string {
	return e.nombre
}

func (e *Escuderia) SetNombre(nombre string) {
	e.nombre = nombre
}

func (e Escuderia) GetPilotos() []Piloto {
	return e.pilotos
}

func (e *Escuderia) SetPilotos(pilotos []Piloto) {
	e.pilotos = pilotos
}

func (e Escuderia) GetTitulosMundiales() int {
	return e.titulosMundiales
}

func (e *Escuderia) SetTitulosMundiales(titulosMundiales int) {
	e.titulosMundiales = titulosMundiales
}

func (e Escuderia) GetVictorias() int {
	return e.victorias
}

func (e *Escuderia) SetVictorias(victorias int) {
	e.victorias = victorias
}

func (e Escuderia) GetPoles() int {
	return e.poles
}

func (e *Escuderia) SetPoles(poles int) {
	e.poles = poles
}

func (e Escuderia) GetVueltasRapidas() int {
	return e.vueltasRapidas
}

func (e *Escuderia) SetVueltasRapidas(vr int) {
	e.vueltasRapidas = vr
}

/* Método para visualizar ordenadamente la información sobre una escuderia. Se pueden seleccionar
 * los campos que se desean visualizar y los que no.
 *
 */

func (e Escuderia) VerDatosEscuderia(n bool, p bool, m bool, v bool, pol bool, vr bool) string {
	var ret string

	if n {
		var nombre string = ("Nombre de la escuderia: " + e.nombre + "\n")
		ret += nombre
	}

	if p {
		var pilotos string = ("Pilotos de la escuderia: " + e.pilotos[0].GetNombre() +
			" y " + e.pilotos[1].GetNombre() + "\n")

		ret += pilotos
	}

	if m {
		var mundiales string = ("Mundiales obtenidos: " + strconv.Itoa(e.titulosMundiales) + "\n")
		ret += mundiales
	}

	if v {
		var victorias string = ("Victorias conseguidas: " + strconv.Itoa(e.victorias) + "\n")
		ret += victorias
	}

	if pol {
		var poles string = ("Pole positions conseguidas: " + strconv.Itoa(e.poles) + "\n")
		ret += poles
	}

	if vr {
		var vueltasRapidas string = ("Vueltas rápidas conseguidas: " + strconv.Itoa(e.vueltasRapidas) + "\n")
		ret += vueltasRapidas
	}

	return ret
}
