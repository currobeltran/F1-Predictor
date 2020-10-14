package f1predictor

import (
	"reflect"
	"testing"
)

func TestGetResultadoFP1(t *testing.T) {
	t.Log("Test GetResultadoFP1")

	var fp1 [20]SesionPiloto = resultado.GetResultadoFP1()

	if !reflect.DeepEqual(fp1, lista) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista[0].GetPiloto().GetNombre()+" "+lista[0].GetMejorTiempo().ImprimirTiempo(),
			fp1[0].GetPiloto().GetNombre()+" "+fp1[0].GetMejorTiempo().ImprimirTiempo())
	}
}

func TestSetResultadoFP1(t *testing.T) {
	t.Log("Test SetResultadoFP1")

	resultado.SetResultadoFP1(lista2)

	if !reflect.DeepEqual(resultado.GetResultadoFP1(), lista2) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista2[0].GetPiloto().GetNombre()+" "+lista2[0].GetMejorTiempo().ImprimirTiempo(),
			resultado.GetResultadoFP1()[0].GetPiloto().GetNombre()+" "+resultado.GetResultadoFP1()[0].GetMejorTiempo().ImprimirTiempo())
	}

	resultado.SetResultadoFP1(lista)
}

func TestGetResultadoFP2(t *testing.T) {
	t.Log("Test GetResultadoFP2")

	var fp2 [20]SesionPiloto = resultado.GetResultadoFP2()

	if !reflect.DeepEqual(fp2, lista) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista[0].GetPiloto().GetNombre()+" "+lista[0].GetMejorTiempo().ImprimirTiempo(),
			fp2[0].GetPiloto().GetNombre()+" "+fp2[0].GetMejorTiempo().ImprimirTiempo())
	}
}

func TestSetResultadoFP2(t *testing.T) {
	t.Log("Test SetResultadoFP2")

	resultado.SetResultadoFP2(lista2)

	if !reflect.DeepEqual(resultado.GetResultadoFP2(), lista2) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista2[0].GetPiloto().GetNombre()+" "+lista2[0].GetMejorTiempo().ImprimirTiempo(),
			resultado.GetResultadoFP2()[0].GetPiloto().GetNombre()+" "+resultado.GetResultadoFP2()[0].GetMejorTiempo().ImprimirTiempo())
	}

	resultado.SetResultadoFP2(lista)
}

func TestGetResultadoFP3(t *testing.T) {
	t.Log("Test GetResultadoFP3")

	var fp3 [20]SesionPiloto = resultado.GetResultadoFP3()

	if !reflect.DeepEqual(fp3, lista) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista[0].GetPiloto().GetNombre()+" "+lista[0].GetMejorTiempo().ImprimirTiempo(),
			fp3[0].GetPiloto().GetNombre()+" "+fp3[0].GetMejorTiempo().ImprimirTiempo())
	}
}

func TestSetResultadoFP3(t *testing.T) {
	t.Log("Test SetResultadoFP3")

	resultado.SetResultadoFP3(lista2)

	if !reflect.DeepEqual(resultado.GetResultadoFP3(), lista2) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista2[0].GetPiloto().GetNombre()+" "+lista2[0].GetMejorTiempo().ImprimirTiempo(),
			resultado.GetResultadoFP3()[0].GetPiloto().GetNombre()+" "+resultado.GetResultadoFP3()[0].GetMejorTiempo().ImprimirTiempo())
	}

	resultado.SetResultadoFP3(lista)
}

func TestGetResultadoClasificacion(t *testing.T) {
	t.Log("Test GetResultadoClasificacion")

	var clasificacion [20]SesionPiloto = resultado.GetResultadoClasificacion()

	if !reflect.DeepEqual(clasificacion, lista) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista[0].GetPiloto().GetNombre()+" "+lista[0].GetMejorTiempo().ImprimirTiempo(),
			clasificacion[0].GetPiloto().GetNombre()+" "+clasificacion[0].GetMejorTiempo().ImprimirTiempo())
	}
}

func TestSetResultadoClasificacion(t *testing.T) {
	t.Log("Test SetResultadoClasificacion")

	resultado.SetResultadoClasificacion(lista2)

	if !reflect.DeepEqual(resultado.GetResultadoClasificacion(), lista2) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista2[0].GetPiloto().GetNombre()+" "+lista2[0].GetMejorTiempo().ImprimirTiempo(),
			resultado.GetResultadoClasificacion()[0].GetPiloto().GetNombre()+" "+resultado.GetResultadoClasificacion()[0].GetMejorTiempo().ImprimirTiempo())
	}

	resultado.SetResultadoClasificacion(lista)
}

func TestGetResultadoCarrera(t *testing.T) {
	t.Log("Test GetResultadoCarrera")

	var carrera [20]SesionPiloto = resultado.GetResultadoCarrera()

	if !reflect.DeepEqual(carrera, lista) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista[0].GetPiloto().GetNombre()+" "+lista[0].GetMejorTiempo().ImprimirTiempo(),
			carrera[0].GetPiloto().GetNombre()+" "+carrera[0].GetMejorTiempo().ImprimirTiempo())
	}
}

func TestSetResultadoCarrera(t *testing.T) {
	t.Log("Test SetResultadoCarrera")

	resultado.SetResultadoCarrera(lista2)

	if !reflect.DeepEqual(resultado.GetResultadoCarrera(), lista2) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			lista2[0].GetPiloto().GetNombre()+" "+lista2[0].GetMejorTiempo().ImprimirTiempo(),
			resultado.GetResultadoCarrera()[0].GetPiloto().GetNombre()+" "+resultado.GetResultadoCarrera()[0].GetMejorTiempo().ImprimirTiempo())
	}

	resultado.SetResultadoCarrera(lista)
}

func TestGetGanador(t *testing.T) {
	t.Log("Test GetGanador")

	var win Piloto = resultado.GetGanador()

	if win != vettel {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			vettel.GetNombre(), win.GetNombre())
	}
}

func TestSetGanador(t *testing.T) {
	t.Log("Test SetGanador")

	resultado.SetGanador(hamilton)

	if resultado.GetGanador() != hamilton {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			hamilton.GetNombre(), resultado.GetGanador().GetNombre())
	}

	resultado.SetGanador(vettel)
}

func TestGetPoleman(t *testing.T) {
	t.Log("Test GetPoleman")

	var polman Piloto = resultado.GetPoleman()

	if polman != vettel {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			vettel.GetNombre(), polman.GetNombre())
	}
}

func TestSetPoleman(t *testing.T) {
	t.Log("Test SetPoleman")

	resultado.SetPoleman(hamilton)

	if resultado.GetPoleman() != hamilton {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			hamilton.GetNombre(), resultado.GetPoleman().GetNombre())
	}
}

func TestGetPodio(t *testing.T) {
	t.Log("Test GetPodio")

	var pod [3]Piloto = resultado.GetPodio()

	if pod != podio {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			podio[0].GetNombre()+", "+podio[1].GetNombre()+" y "+podio[2].GetNombre(),
			pod[0].GetNombre()+", "+pod[1].GetNombre()+" y "+pod[2].GetNombre())
	}
}

func TestSetPodio(t *testing.T) {
	t.Log("Test SetPodio")

	resultado.SetPodio(podio2)

	if resultado.GetPodio() != podio2 {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			podio2[0].GetNombre()+", "+podio2[1].GetNombre()+" y "+podio2[2].GetNombre(),
			resultado.GetPodio()[0].GetNombre()+", "+resultado.GetPodio()[1].GetNombre()+" y "+resultado.GetPodio()[2].GetNombre())
	}

	resultado.SetPodio(podio)
}

func TestGetTemporada(t *testing.T) {
	t.Log("Test GetTemporada")

	var tem int = resultado.GetTemporada()

	if tem != 2020 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 2020, resultado.GetTemporada())
	}
}

func TestSetTemporada(t *testing.T) {
	t.Log("Test SetTemporada")

	resultado.SetTemporada(2019)

	if resultado.GetTemporada() != 2019 {
		t.Errorf("Resultado obtenido incorrecto, esperado %d, obtenido %d", 2019, resultado.GetTemporada())
	}

	resultado.SetTemporada(2020)
}

func TestGetEstadisticas(t *testing.T) {
	t.Log("Test GetEstadisticas")

	var estad EstadisticasGP = resultado.GetEstadisticas()

	if estad != est {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			est.VerEstadisticasGP(true, true, true, true, true, true, true),
			estad.VerEstadisticasGP(true, true, true, true, true, true, true))
	}
}

func TestSetEstadisticas(t *testing.T) {
	t.Log("Test SetEstadisticas")

	resultado.SetEstadisticas(est2)

	if resultado.GetEstadisticas() != est2 {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			est2.VerEstadisticasGP(true, true, true, true, true, true, true),
			resultado.GetEstadisticas().VerEstadisticasGP(true, true, true, true, true, true, true))
	}

	resultado.SetEstadisticas(est)
}
