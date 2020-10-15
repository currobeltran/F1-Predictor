package f1predictor

import "testing"

func TestGetMinuto(t *testing.T) {
	t.Log("Test GetMinuto")

	var min int = record.GetMinuto()

	if min != 1 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 1, min)
	}
}

func TestSetMinuto(t *testing.T) {
	t.Log("Test SetMinuto")

	record.SetMinuto(2)

	if record.GetMinuto() != 2 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 2, record.GetMinuto())
	}

	record.SetMinuto(1)
}

func TestGetSegundo(t *testing.T) {
	t.Log("Test GetSegundo")

	var seg int = record.GetSegundo()

	if seg != 23 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 23, seg)
	}
}

func TestSetSegundo(t *testing.T) {
	t.Log("Test SetSegundo")

	record.SetSegundo(45)

	if record.GetSegundo() != 45 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 45, record.GetSegundo())
	}

	record.SetSegundo(23)
}

func TestGetMilesima(t *testing.T) {
	t.Log("Test GetMilesima")

	var mil int = record.GetMilesima()

	if mil != 456 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 456, mil)
	}
}

func TestSetMilesima(t *testing.T) {
	t.Log("Test SetMilesima")

	record.SetMilesima(072)

	if record.GetMilesima() != 072 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 072, record.GetMilesima())
	}

	record.SetMilesima(456)
}

func TestImprimirTiempo(t *testing.T) {
	t.Log("Test ImprimirTiempo")

	var ti string = record.ImprimirTiempo()
	var prueba string = "1:23.456"

	if ti != prueba {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s", prueba, ti)
	}
}
