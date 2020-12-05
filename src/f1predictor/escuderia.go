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
	Nombre           string
	Pilotos          []Piloto
	TitulosMundiales int
	Victorias        int
	Poles            int
	VueltasRapidas   int
}

type MetodosEscuderia interface {
	Constructor(Nombre string, Pilotos [2]Piloto, titulos int, Victorias int, Poles int, vr int)

	GetNombre() string
	SetNombre(Nombre string)

	GetPilotos() []Piloto
	SetPilotos(Pilotos []Piloto)

	GetTitulosMundiales() int
	SetTitulosMundiales(titulos int)

	GetVictorias() int
	SetVictorias(Victorias int)

	GetPoles() int
	SetPoles(Poles int)

	GetVueltasRapidas() int
	SetVueltasRapidas(vr int)

	VerDatosEscuderia(n bool, p bool, m bool, v bool, pol bool, vr bool) string
}

func (e *Escuderia) Constructor(Nombre string, Pilotos []Piloto, titulos int,
	Victorias int, Poles int, vr int) {

	e.Nombre = Nombre
	e.Pilotos = Pilotos
	e.TitulosMundiales = titulos
	e.Victorias = Victorias
	e.Poles = Poles
	e.VueltasRapidas = vr

}

func (e Escuderia) GetNombre() string {
	return e.Nombre
}

func (e *Escuderia) SetNombre(Nombre string) {
	e.Nombre = Nombre
}

func (e Escuderia) GetPilotos() []Piloto {
	return e.Pilotos
}

func (e *Escuderia) SetPilotos(Pilotos []Piloto) {
	e.Pilotos = Pilotos
}

func (e Escuderia) GetTitulosMundiales() int {
	return e.TitulosMundiales
}

func (e *Escuderia) SetTitulosMundiales(TitulosMundiales int) {
	e.TitulosMundiales = TitulosMundiales
}

func (e Escuderia) GetVictorias() int {
	return e.Victorias
}

func (e *Escuderia) SetVictorias(Victorias int) {
	e.Victorias = Victorias
}

func (e Escuderia) GetPoles() int {
	return e.Poles
}

func (e *Escuderia) SetPoles(Poles int) {
	e.Poles = Poles
}

func (e Escuderia) GetVueltasRapidas() int {
	return e.VueltasRapidas
}

func (e *Escuderia) SetVueltasRapidas(vr int) {
	e.VueltasRapidas = vr
}

/* Método para visualizar ordenadamente la información sobre una escuderia. Se pueden seleccionar
 * los campos que se desean visualizar y los que no.
 *
 */

func (e Escuderia) VerDatosEscuderia(n bool, p bool, m bool, v bool, pol bool, vr bool) string {
	var ret string

	if n {
		var Nombre string = ("Nombre de la escuderia: " + e.Nombre + "\n")
		ret += Nombre
	}

	if p {
		var Pilotos string = ("Pilotos de la escuderia: " + e.Pilotos[0].GetNombre() +
			" y " + e.Pilotos[1].GetNombre() + "\n")

		ret += Pilotos
	}

	if m {
		var mundiales string = ("Mundiales obtenidos: " + strconv.Itoa(e.TitulosMundiales) + "\n")
		ret += mundiales
	}

	if v {
		var Victorias string = ("Victorias conseguidas: " + strconv.Itoa(e.Victorias) + "\n")
		ret += Victorias
	}

	if pol {
		var Poles string = ("Pole positions conseguidas: " + strconv.Itoa(e.Poles) + "\n")
		ret += Poles
	}

	if vr {
		var VueltasRapidas string = ("Vueltas rápidas conseguidas: " + strconv.Itoa(e.VueltasRapidas) + "\n")
		ret += VueltasRapidas
	}

	return ret
}
