package f1predictor

import (
	"reflect"
	"testing"
)

func TestGetNombreEscuderia(t *testing.T) {
	t.Log("Test GetNombre")

	var nombre string = ferrari.GetNombre()

	if nombre != "Ferrari" {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s", "Ferrari", nombre)
	}
}

func TestSetNombreEscuderia(t *testing.T) {
	t.Log("Test SetNombre")

	ferrari.SetNombre("Mercedes")

	if ferrari.GetNombre() != "Mercedes" {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s", "Mercedes", ferrari.GetNombre())
	}
}

func TestGetPilotos(t *testing.T) {
	t.Log("Test GetPilotos")

	var pilotosAux []Piloto = ferrari.GetPilotos()

	if !reflect.DeepEqual(pilotosAux, pilotos) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s y %s, obtenido %s y %s",
			pilotosAux[0].GetNombre(), pilotosAux[1].GetNombre(),
			pilotos[0].GetNombre(), pilotos[1].GetNombre())
	}
}

func TestSetPilotos(t *testing.T) {
	t.Log("Test SetPilotos")

	ferrari.SetPilotos(pilotos2)

	if !reflect.DeepEqual(ferrari.GetPilotos(), pilotos2) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s y %s, obtenido %s y %s",
			ferrari.GetPilotos()[0].GetNombre(), ferrari.GetPilotos()[1].GetNombre(),
			pilotos2[0].GetNombre(), pilotos2[1].GetNombre())
	}
}

func TestGetTitulosMundialesEscuderia(t *testing.T) {
	t.Log("Test GetTitulosMundiales")

	var titulos int = ferrari.GetTitulosMundiales()

	if titulos != 20 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 20, ferrari.GetTitulosMundiales())
	}
}

func TestSetTitulosMundialesEscuderia(t *testing.T) {
	t.Log("Test SetTitulosMundiales")

	ferrari.SetTitulosMundiales(19)

	if ferrari.GetTitulosMundiales() != 19 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 19, ferrari.GetTitulosMundiales())
	}
}

func TestGetVictoriasEscuderia(t *testing.T) {
	t.Log("Test GetVictorias")

	var victorias int = ferrari.GetVictorias()

	if victorias != 150 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 150, ferrari.GetVictorias())
	}
}

func TestSetVictoriasEscuderia(t *testing.T) {
	t.Log("Test SetVictorias")

	ferrari.SetVictorias(190)

	if ferrari.GetVictorias() != 190 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 190, ferrari.GetVictorias())
	}
}

func TestGetPolesEscuderia(t *testing.T) {
	t.Log("Test GetPoles")

	var poles int = ferrari.GetPoles()

	if poles != 120 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 120, ferrari.GetPoles())
	}
}

func TestSetPolesEscuderia(t *testing.T) {
	t.Log("Test SetPoles")

	ferrari.SetPoles(190)

	if ferrari.GetPoles() != 190 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 190, ferrari.GetPoles())
	}
}

func TestGetVueltasRapidasEscuderia(t *testing.T) {
	t.Log("Test GetVueltasRapidas")

	var vueltasrapidas int = ferrari.GetVueltasRapidas()

	if vueltasrapidas != 110 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 110, ferrari.GetVueltasRapidas())
	}
}

func TestSetVueltasRapidasEscuderia(t *testing.T) {
	t.Log("Test SetVueltasRapidas")

	ferrari.SetVueltasRapidas(190)

	if ferrari.GetVueltasRapidas() != 190 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 190, ferrari.GetVueltasRapidas())
	}
}
