package f1predictor

import "strconv"

/*
 * Clase que contiene información sobre un piloto de Fórmula 1. Los atributos son los siguientes:
 *
 *		- Nombre del piloto
 *		- Victorias conseguidas
 * 		- Poles conseguidas
 * 		- Vueltas rápidas conseguidas
 *		- Mundiales conseguidos
 */

type Piloto struct {
	nombre         string
	victorias      int
	poles          int
	vueltasRapidas int
	mundiales      int
}

type MetodosPiloto interface {
	Constructor(nombre string, victorias int, poles int, vueltasRapidas int, mundiales int)

	GetNombre() string
	SetNombre(nombre string)

	GetVictorias() int
	SetVictorias(victorias int)

	GetPoles() int
	SetPoles(poles int)

	GetVueltasRapidas() int
	SetVueltasRapidas(vueltasRapidas int)

	GetMundiales() int
	SetMundiales(mundiales int)

	VerDatosPiloto(n bool, v bool, pol bool, vr bool, m bool) string
}

func (p *Piloto) Constructor(nombre string, victorias int,
	poles int, vueltasRapidas int, mundiales int) {
	p.nombre = nombre
	p.victorias = victorias
	p.poles = poles
	p.vueltasRapidas = vueltasRapidas
	p.mundiales = mundiales
}

func (p Piloto) GetNombre() string {
	return p.nombre
}

func (p Piloto) SetNombre(nombre string) {
	p.nombre = nombre
}

func (p Piloto) GetVictorias() int {
	return p.victorias
}

func (p Piloto) SetVictorias(victorias int) {
	p.victorias = victorias
}

func (p Piloto) GetPoles() int {
	return p.poles
}

func (p Piloto) SetPoles(poles int) {
	p.poles = poles
}

func (p Piloto) GetVueltasRapidas() int {
	return p.vueltasRapidas
}

func (p Piloto) SetVueltasRapidas(vueltasRapidas int) {
	p.vueltasRapidas = vueltasRapidas
}

func (p Piloto) GetMundiales() int {
	return p.mundiales
}

func (p Piloto) SetMundiales(mundiales int) {
	p.mundiales = mundiales
}

/*
 * Método que crea un string para ver de manera ordenada la información correspondiente
 * a un piloto en concreto. Se da la opción de seleccionar cual es la información que se
 * quiere obtener sobre el mismo.
 */

func (p Piloto) VerDatosPiloto(n bool, v bool, pol bool, vr bool, m bool) string {
	var ret string

	if n {
		var nombre string = ("Nombre del piloto: " + p.nombre + "\n")
		ret += nombre
	}

	if v {
		var victorias string = ("Victorias conseguidas: " + strconv.Itoa(p.victorias) + "\n")
		ret += victorias
	}

	if pol {
		var poles string = ("Poles conseguidas: " + strconv.Itoa(p.poles) + "\n")
		ret += poles
	}

	if vr {
		var vueltasRapidas string = ("Vueltas rápidas conseguidas: " + strconv.Itoa(p.vueltasRapidas) + "\n")
		ret += vueltasRapidas
	}

	if m {
		var mundiales string = ("Mundiales conseguidos: " + strconv.Itoa(p.mundiales) + "\n")
		ret += mundiales
	}

	return ret
}
