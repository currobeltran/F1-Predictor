package f1predictor

import (
	"os"
	"reflect"
	"testing"
)

var (
	melbourne  Circuito
	cero       TiempoVuelta
	record     TiempoVuelta
	hamilton   Piloto
	resultado  ResultadoGP
	resultado2 ResultadoGP
	resultados []ResultadoGP
	lista      [20]SesionPiloto
	podio      [3]Piloto
	est        EstadisticasGP
	est2       EstadisticasGP
)

func TestMain(m *testing.M) {
	melbourne.Constructor("Albert Park", "Australia")
	cero.Constructor(0, 0, 0)
	record.Constructor(1, 23, 456)
	est.Constructor(2, 1, 7, 10, 6, 0, record)
	est2.Constructor(4, 3, 9, 5, 3, 0, record)

	resultado.Constructor(lista, lista, lista, lista, lista, hamilton, hamilton, podio, 2020, est)
	resultado2.Constructor(lista, lista, lista, lista, lista, hamilton, hamilton, podio, 2020, est2)

	resultados = make([]ResultadoGP, 1)

	os.Exit(m.Run())
}

func TestGetNombre(t *testing.T) {
	t.Log("Test GetNombre")

	if melbourne.GetNombre() != "Albert Park" {
		t.Errorf("Fallo, obtenido %s, se esperaba %s", melbourne.GetNombre(), "Albert Park")
	}
}

func TestSetNombre(t *testing.T) {
	t.Log("Test SetNombre")

	melbourne.SetNombre("Melbourne")

	if melbourne.GetNombre() != "Melbourne" {
		t.Error("No se modifica correctamente el nombre del circuito")
	}
}

func TestGetPais(t *testing.T) {
	t.Log("Test GetPais")

	if melbourne.GetPais() != "Australia" {
		t.Error("No se obtiene correctamente el país del circuito")
	}
}

func TestSetPais(t *testing.T) {
	t.Log("Test SetPais")

	melbourne.SetPais("Aus")

	if melbourne.GetPais() != "Aus" {
		t.Error("No se modifica correctamente el país del circuito")
	}
}

func TestGetRecordCircuito(t *testing.T) {
	t.Log("Test GetRecordCircuito")

	if melbourne.GetRecordCircuito() != cero {
		t.Error("No se obtiene correctamente el tiempo record del circuito")
	}
}

func TestSetRecordCircuito(t *testing.T) {
	t.Log("Test SetRecordCircuito")

	melbourne.SetRecordCircuito(record)

	if melbourne.GetRecordCircuito() != record {
		t.Error("No se modifica correctamente el tiempo record del circuito")
	}
}

func TestGetResultado(t *testing.T) {
	t.Log("Test GetResultado")

	if !reflect.DeepEqual(melbourne.GetResultados(), resultados) {
		t.Errorf("El vector de resultados no se obtiene correctamente")
	}
}

func TestSetResultado(t *testing.T) {
	t.Log("Test SetResultado")

	resultados = append(resultados, resultado)

	melbourne.SetResultados(resultados)

	if !reflect.DeepEqual(melbourne.GetResultados(), resultados) {
		t.Errorf("El vector de resultados no se modifica correctamente")
	}
}

func TestSetResultadoByIndex(t *testing.T) {
	t.Log("Test SetResultadoByIndex")

	melbourne.SetResultadosByIndex(0, resultado)

	if !reflect.DeepEqual(melbourne.GetResultados()[0], resultado) {
		t.Errorf("El indice no se modifica correctamente")
	}
}

func TestMedia(t *testing.T) {
	t.Log("Test Media")

	melbourne.SetResultados(melbourne.GetResultados()[:0]) //Limpiamos datos anteriores del vector

	melbourne.SetResultados(append(melbourne.GetResultados(), resultado))
	melbourne.SetResultados(append(melbourne.GetResultados(), resultado2))

	if melbourne.Media("accidentes") != 3 {
		t.Errorf("Media accidentes incorrecta, resultado obtenido %f, esperado %f", melbourne.Media("accidentes"), 3.0)
	}

	if melbourne.Media("numeroSafetyCar") != 2 {
		t.Errorf("Media Safety Car incorrecta, resultado obtenido %f, esperado %f", melbourne.Media("numeroSafetyCar"), 2.0)
	}

	if melbourne.Media("adelantamientos") != 8 {
		t.Errorf("Media adelantamientos incorrecta, resultado obtenido %f, esperado %f", melbourne.Media("adelantamientos"), 8.0)
	}

	if melbourne.Media("banderasAmarillas") != 7.5 {
		t.Errorf("Media banderas amarillas incorrecta, resultado obtenido %f, esperado %f", melbourne.Media("banderasAmarillas"), 7.5)
	}

	if melbourne.Media("banderasRojas") != 4.5 {
		t.Errorf("Media banderas rojas incorrecta, resultado obtenido %f, esperado %f", melbourne.Media("banderasRojas"), 4.5)
	}

	if melbourne.Media("sanciones") != 0 {
		t.Errorf("Media sanciones incorrecta, resultado obtenido %f, esperado %f", melbourne.Media("sanciones"), 0.0)
	}
}
