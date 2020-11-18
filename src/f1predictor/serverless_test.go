package f1predictor

import "testing"

func TestObtenerClasificacion(t *testing.T) {
	t.Log("Test ObtenerClasificacion")

	melbourne.GetResultados()[0].SetResultadoClasificacion(lista2)

	result := ObtenerClasificacion(melbourne, 2020)

	if result != "1ª Posición: Lewis Hamilton, tiempo: 1:23.456" {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			"1ª Posición: Lewis Hamilton, tiempo: 1:23.456", result)
	}

	melbourne.GetResultados()[0].SetResultadoClasificacion(lista)
}
