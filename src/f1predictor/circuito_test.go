package f1predictor

import (
	"math"
	"os"
	"reflect"
	"regexp"
	"testing"
)

var (
	melbourne Circuito

	cero     TiempoVuelta
	record   TiempoVuelta
	tiempos  []TiempoVuelta
	tiempos2 []TiempoVuelta

	vettel   Piloto
	leclerc  Piloto
	hamilton Piloto
	bottas   Piloto

	pilotos  []Piloto
	pilotos2 []Piloto

	ferrari Escuderia

	resultado  ResultadoGP
	resultado2 ResultadoGP
	resultados []ResultadoGP

	lista  [20]SesionPiloto
	lista2 [20]SesionPiloto

	hamiltonTiempos SesionPiloto

	podio  [3]Piloto
	podio2 [3]Piloto

	est  EstadisticasGP
	est2 EstadisticasGP
)

func TestMain(m *testing.M) {
	melbourne.Constructor("Albert Park", "Australia")

	cero.Constructor(0, 0, 0)
	record.Constructor(1, 23, 456)

	est.Constructor(2, 1, 7, 10, 6, 0, record)
	est2.Constructor(4, 3, 9, 5, 3, 0, record)

	vettel.Constructor("Sebastian Vettel", 55, 40, 30, 4)
	leclerc.Constructor("Charles Leclerc", 3, 5, 2, 0)
	hamilton.Constructor("Lewis Hamilton", 91, 96, 55, 6)
	bottas.Constructor("Valtteri Bottas", 9, 11, 10, 0)

	pilotos = append(pilotos, vettel)
	pilotos = append(pilotos, leclerc)

	pilotos2 = append(pilotos2, hamilton)
	pilotos2 = append(pilotos2, bottas)

	ferrari.Constructor("Ferrari", pilotos, 20, 150, 120, 110)

	podio[0] = vettel
	podio[1] = leclerc
	podio[2] = hamilton

	podio2[0] = bottas
	podio2[1] = hamilton
	podio2[2] = leclerc

	tiempos = append(tiempos, cero)
	tiempos = append(tiempos, record)

	tiempos2 = append(tiempos2, record)
	tiempos2 = append(tiempos2, record)
	tiempos2 = append(tiempos2, cero)

	hamiltonTiempos.Constructor(tiempos, hamilton, record)
	lista2[0] = hamiltonTiempos

	resultado.Constructor(lista, lista, lista, lista, lista, vettel, vettel, podio, 2020, est)
	resultado2.Constructor(lista, lista, lista, lista, lista, vettel, vettel, podio, 2019, est2)

	resultados = make([]ResultadoGP, 1)

	melbourne.SetResultados(resultados)

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

func TestVerDatosCircuito(t *testing.T) {
	t.Log("Test VerDatosCircuito")

	var str string = melbourne.VerDatosCircuito(true, true, true)
	var matched, matched2, matched3 bool

	matched, _ = regexp.MatchString("Nombre", str)
	matched2, _ = regexp.MatchString("País", str)
	matched3, _ = regexp.MatchString("Record", str)

	if !(matched && matched2 && matched3) {
		t.Errorf("No se ha obtenido correctamente la cadena deseada. Resultado obtenido:\n%s", str)
	}

	str = melbourne.VerDatosCircuito(true, true, false)

	matched, _ = regexp.MatchString("Nombre", str)
	matched2, _ = regexp.MatchString("País", str)

	if !(matched && matched2) {
		t.Errorf("No se ha obtenido correctamente la cadena deseada. Resultado obtenido:\n%s", str)
	}

	str = melbourne.VerDatosCircuito(false, true, false)

	matched2, _ = regexp.MatchString("País", str)

	if !matched2 {
		t.Errorf("No se ha obtenido correctamente la cadena deseada. Resultado obtenido:\n%s", str)
	}
}

func TestPosibilidadSafetyCar(t *testing.T) {
	t.Log("Test PosibilidadSafetyCar")

	melbourne.SetResultados(melbourne.GetResultados()[:0]) //Limpiamos datos anteriores del vector
	resultados = make([]ResultadoGP, 2)

	resultados[0] = resultado2
	resultados[1] = resultado

	melbourne.SetResultados(resultados)

	var result float64 = melbourne.PosibilidadSafetyCar()

	/*Fórmula posibilidad -> Media ponderada de estadística correspondiente*/
	/*Cada año de antiguedad resta 0,05 a la ponderación*/
	if math.Abs(result-1.974359) > 0.000001 {
		t.Errorf("Resultado obtenido incorrecto, esperado %f, obtenido %f", 1.974359, result)
	}
}

func TestPosibilidadAdelantamiento(t *testing.T) {
	t.Log("Test PosibilidadAdelantamiento")

	var result float64 = melbourne.PosibilidadAdelantamiento()

	/*Fórmula posibilidad -> Media ponderada de estadística correspondiente*/
	/*Cada año de antiguedad resta 0,05 a la ponderación*/
	if math.Abs(result-7.974359) > 0.000001 {
		t.Errorf("Resultado obtenido incorrecto, esperado %f, obtenido %f", 7.974359, result)
	}
}

func TestPosibilidadBanderaAmarilla(t *testing.T) {
	t.Log("Test PosibilidadBanderaAmarilla")

	var result float64 = melbourne.PosibilidadBanderaAmarilla()

	/*Fórmula posibilidad -> Media ponderada de estadística correspondiente*/
	/*Cada año de antiguedad resta 0,05 a la ponderación*/
	if math.Abs(result-7.564102) > 0.000001 {
		t.Errorf("Resultado obtenido incorrecto, esperado %f, obtenido %f", 7.564102, result)
	}
}

func TestPosibilidadAccidentes(t *testing.T) {
	t.Log("Test PosibilidadAccidentes")

	var result float64 = melbourne.PosibilidadAccidentes()

	/*Fórmula posibilidad -> Media ponderada de estadística correspondiente*/
	/*Cada año de antiguedad resta 0,05 a la ponderación*/
	if math.Abs(result-2.974359) > 0.000001 {
		t.Errorf("Resultado obtenido incorrecto, esperado %f, obtenido %f", 2.974359, result)
	}
}
