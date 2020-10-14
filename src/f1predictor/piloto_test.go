package f1predictor

import (
	"regexp"
	"testing"
)

func TestGetNombrePiloto(t *testing.T) {
	t.Log("Test GetNombrePiloto")

	var nombre string = vettel.GetNombre()

	if nombre != "Sebastian Vettel" {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s", "Sebastian Vettel", nombre)
	}
}

func TestSetNombrePiloto(t *testing.T) {
	t.Log("Test SetNombrePiloto")

	vettel.SetNombre("Seb")

	if vettel.GetNombre() != "Seb" {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s", "Seb", vettel.GetNombre())
	}

	vettel.SetNombre("Sebastian Vettel")
}

func TestGetVictoriasPiloto(t *testing.T) {
	t.Log("Test GetVictoriasPiloto")

	var vic int = vettel.GetVictorias()

	if vic != 55 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 55, vic)
	}
}

func TestSetVictoriasPiloto(t *testing.T) {
	t.Log("Test SetNombrePiloto")

	vettel.SetVictorias(1)

	if vettel.GetVictorias() != 1 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 1, vettel.GetVictorias())
	}

	vettel.SetVictorias(55)
}

func TestGetPolesPiloto(t *testing.T) {
	t.Log("Test GetPolesPiloto")

	var pols int = vettel.GetPoles()

	if pols != 40 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 40, pols)
	}
}

func TestSetPolesPiloto(t *testing.T) {
	t.Log("Test SetPolesPiloto")

	vettel.SetPoles(3)

	if vettel.GetPoles() != 3 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 3, vettel.GetPoles())
	}

	vettel.SetPoles(40)
}

func TestGetVueltasRapidasPiloto(t *testing.T) {
	t.Log("Test GetVueltasRapidasPiloto")

	var vura int = vettel.GetVueltasRapidas()

	if vura != 30 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 30, vura)
	}
}

func TestSetVueltasRapidasPiloto(t *testing.T) {
	t.Log("Test SetVueltasRapidasPiloto")

	vettel.SetVueltasRapidas(100)

	if vettel.GetVueltasRapidas() != 100 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 100, vettel.GetVueltasRapidas())
	}

	vettel.SetVueltasRapidas(30)
}

func TestGetMundialesPiloto(t *testing.T) {
	t.Log("Test GetMundialesPiloto")

	var mun int = vettel.GetMundiales()

	if mun != 4 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 4, mun)
	}
}

func TestSetMundialesPiloto(t *testing.T) {
	t.Log("Test SetMundialesPiloto")

	vettel.SetMundiales(2)

	if vettel.GetMundiales() != 2 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 2, vettel.GetMundiales())
	}

	vettel.SetMundiales(4)
}

func TestVerDatosPiloto(t *testing.T) {
	t.Log("Test VerDatosPiloto")

	var datos string = hamilton.VerDatosPiloto(true, true, false, false, true)
	var match, match2, match3 bool

	match, _ = regexp.MatchString("Nombre", datos)
	match2, _ = regexp.MatchString("Victorias", datos)
	match3, _ = regexp.MatchString("Mundiales", datos)

	if !(match && match2 && match3) {
		t.Errorf("No se ha obtenido correctamente la cadena deseada. Resultado obtenido:\n%s", datos)
	}
}
