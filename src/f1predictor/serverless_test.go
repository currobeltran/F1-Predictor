package f1predictor

import "testing"

func TestObtenerClasificacion(t *testing.T) {
	t.Log("Test ObtenerClasificacion")

	melbourne.GetResultados()[0].SetResultadoClasificacion(lista2)

	result := ObtenerClasificacion(melbourne, 2019)

	if result != "1ª posición: Lewis Hamilton, tiempo: 1:23.456\n2ª posición: , tiempo: 0:0.0\n3ª posición: , tiempo: 0:0.0\n4ª posición: , tiempo: 0:0.0\n5ª posición: , tiempo: 0:0.0\n6ª posición: , tiempo: 0:0.0\n7ª posición: , tiempo: 0:0.0\n8ª posición: , tiempo: 0:0.0\n9ª posición: , tiempo: 0:0.0\n10ª posición: , tiempo: 0:0.0\n11ª posición: , tiempo: 0:0.0\n12ª posición: , tiempo: 0:0.0\n13ª posición: , tiempo: 0:0.0\n14ª posición: , tiempo: 0:0.0\n15ª posición: , tiempo: 0:0.0\n16ª posición: , tiempo: 0:0.0\n17ª posición: , tiempo: 0:0.0\n18ª posición: , tiempo: 0:0.0\n19ª posición: , tiempo: 0:0.0\n20ª posición: , tiempo: 0:0.0\n" {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			"1ª posición: Lewis Hamilton, tiempo: 1:23.456\n2ª posición:...", result)
	}

	melbourne.GetResultados()[0].SetResultadoClasificacion(lista)
}
