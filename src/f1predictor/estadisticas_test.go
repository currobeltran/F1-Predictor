package f1predictor

import (
	"regexp"
	"strconv"
	"testing"
)

func TestGetAccidentes(t *testing.T) {
	var accid int = est.GetAccidentes()

	if accid != 2 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 2, est.GetAccidentes())
	}
}

func TestSetAccidentes(t *testing.T) {
	est.SetAccidentes(3)

	if est.GetAccidentes() != 3 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 3, est.GetAccidentes())
	}

	est.SetAccidentes(2)
}

func TestGetNumeroSafetyCar(t *testing.T) {
	var sc int = est.GetNumeroSafetyCar()

	if sc != 1 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 1, est.GetNumeroSafetyCar())
	}
}

func TestSetNumeroSafetyCar(t *testing.T) {
	est.SetNumeroSafetyCar(3)

	if est.GetNumeroSafetyCar() != 3 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 3, est.GetNumeroSafetyCar())
	}

	est.SetNumeroSafetyCar(1)
}

func TestGetAdelantamientos(t *testing.T) {
	var adelan int = est.GetAdelantamientos()

	if adelan != 7 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 7, est.GetAdelantamientos())
	}
}

func TestSetAdelantamientos(t *testing.T) {
	est.SetAdelantamientos(4)

	if est.GetAdelantamientos() != 4 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 4, est.GetAdelantamientos())
	}

	est.SetAdelantamientos(7)
}

func TestGetBanderasAmarillas(t *testing.T) {
	var banam int = est.GetBanderasAmarillas()

	if banam != 10 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 10, est.GetBanderasAmarillas())
	}
}

func TestSetBanderasAmarillas(t *testing.T) {
	est.SetBanderasAmarillas(1)

	if est.GetBanderasAmarillas() != 1 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 1, est.GetBanderasAmarillas())
	}

	est.SetBanderasAmarillas(10)
}

func TestGetBanderasRojas(t *testing.T) {
	var banro int = est.GetBanderasRojas()

	if banro != 6 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 6, est.GetBanderasRojas())
	}
}

func TestSetBanderasRojas(t *testing.T) {
	est.SetBanderasRojas(0)

	if est.GetBanderasRojas() != 0 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 0, est.GetBanderasRojas())
	}

	est.SetBanderasRojas(6)
}

func TestGetSanciones(t *testing.T) {
	var sanc int = est.GetSanciones()

	if sanc != 0 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 0, est.GetSanciones())
	}
}

func TestSetSanciones(t *testing.T) {
	est.SetSanciones(3)

	if est.GetSanciones() != 3 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 3, est.GetSanciones())
	}

	est.SetSanciones(0)
}

func TestGetMejorVueltaEstadisticas(t *testing.T) {
	var mejvu TiempoVuelta = est.GetMejorVuelta()

	if mejvu != record {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			strconv.Itoa(record.GetMinuto())+":"+strconv.Itoa(record.GetSegundo())+"."+strconv.Itoa(record.GetMilesima()),
			strconv.Itoa(est.GetMejorVuelta().GetMinuto())+":"+strconv.Itoa(est.GetMejorVuelta().GetSegundo())+"."+strconv.Itoa(est.GetMejorVuelta().GetMilesima()))
	}
}

func TestSetMejorVueltaEstadisticas(t *testing.T) {
	est.SetMejorVuelta(cero)

	if est.GetMejorVuelta() != cero {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			strconv.Itoa(cero.GetMinuto())+":"+strconv.Itoa(cero.GetSegundo())+"."+strconv.Itoa(cero.GetMilesima()),
			strconv.Itoa(est.GetMejorVuelta().GetMinuto())+":"+strconv.Itoa(est.GetMejorVuelta().GetSegundo())+"."+strconv.Itoa(est.GetMejorVuelta().GetMilesima()))
	}

	est.SetMejorVuelta(record)
}

func TestVerEstadisticasGP(t *testing.T) {
	t.Log("Test VerEstadisticasGP")

	var datos string = est.VerEstadisticasGP(true, false, false, true, true, false, true)
	var match, match2, match3, match4 bool

	match, _ = regexp.MatchString("accidentes", datos)
	match2, _ = regexp.MatchString("banderas amarillas", datos)
	match3, _ = regexp.MatchString("banderas rojas", datos)
	match4, _ = regexp.MatchString("Tiempo", datos)

	if !(match && match2 && match3 && match4) {
		t.Errorf("No se ha obtenido correctamente la cadena deseada. Resultado obtenido:\n%s", datos)
	}
}
