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
	Nombre         string
	Victorias      int
	Poles          int
	VueltasRapidas int
	Mundiales      int
}

type MetodosPiloto interface {
	Constructor(Nombre string, Victorias int, Poles int, VueltasRapidas int, Mundiales int)

	GetNombre() string
	SetNombre(Nombre string)

	GetVictorias() int
	SetVictorias(Victorias int)

	GetPoles() int
	SetPoles(Poles int)

	GetVueltasRapidas() int
	SetVueltasRapidas(VueltasRapidas int)

	GetMundiales() int
	SetMundiales(Mundiales int)

	VerDatosPiloto(n bool, v bool, pol bool, vr bool, m bool) string
}

func (p *Piloto) Constructor(Nombre string, Victorias int,
	Poles int, VueltasRapidas int, Mundiales int) {
	p.Nombre = Nombre
	p.Victorias = Victorias
	p.Poles = Poles
	p.VueltasRapidas = VueltasRapidas
	p.Mundiales = Mundiales
}

func (p Piloto) GetNombre() string {
	return p.Nombre
}

func (p *Piloto) SetNombre(Nombre string) {
	p.Nombre = Nombre
}

func (p Piloto) GetVictorias() int {
	return p.Victorias
}

func (p *Piloto) SetVictorias(Victorias int) {
	p.Victorias = Victorias
}

func (p Piloto) GetPoles() int {
	return p.Poles
}

func (p *Piloto) SetPoles(Poles int) {
	p.Poles = Poles
}

func (p Piloto) GetVueltasRapidas() int {
	return p.VueltasRapidas
}

func (p *Piloto) SetVueltasRapidas(VueltasRapidas int) {
	p.VueltasRapidas = VueltasRapidas
}

func (p Piloto) GetMundiales() int {
	return p.Mundiales
}

func (p *Piloto) SetMundiales(Mundiales int) {
	p.Mundiales = Mundiales
}

/*
 * Método que crea un string para ver de manera ordenada la información correspondiente
 * a un piloto en concreto. Se da la opción de seleccionar cual es la información que se
 * quiere obtener sobre el mismo.
 */

func (p Piloto) VerDatosPiloto(n bool, v bool, pol bool, vr bool, m bool) string {
	var ret string

	if n {
		var Nombre string = ("Nombre del piloto: " + p.Nombre + "\n")
		ret += Nombre
	}

	if v {
		var Victorias string = ("Victorias conseguidas: " + strconv.Itoa(p.Victorias) + "\n")
		ret += Victorias
	}

	if pol {
		var Poles string = ("Poles conseguidas: " + strconv.Itoa(p.Poles) + "\n")
		ret += Poles
	}

	if vr {
		var VueltasRapidas string = ("Vueltas rápidas conseguidas: " + strconv.Itoa(p.VueltasRapidas) + "\n")
		ret += VueltasRapidas
	}

	if m {
		var Mundiales string = ("Mundiales conseguidos: " + strconv.Itoa(p.Mundiales) + "\n")
		ret += Mundiales
	}

	return ret
}
