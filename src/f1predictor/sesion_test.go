package f1predictor

import (
	"reflect"
	"testing"
)

func TestGetTiempos(t *testing.T) {
	t.Log("Test GetTiempos")

	var tiemps []TiempoVuelta = hamiltonTiempos.GetTiempos()

	if !reflect.DeepEqual(tiemps, tiempos) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			tiempos[0].ImprimirTiempo()+"\n"+tiempos[1].ImprimirTiempo()+"\n",
			tiemps[0].ImprimirTiempo()+"\n"+tiemps[1].ImprimirTiempo()+"\n")
	}
}

func TestSetTiempos(t *testing.T) {
	t.Log("Test SetTiempos")

	hamiltonTiempos.SetTiempos(tiempos2)

	if !reflect.DeepEqual(hamiltonTiempos.GetTiempos(), tiempos2) {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			tiempos2[0].ImprimirTiempo()+"\n"+tiempos2[1].ImprimirTiempo()+"\n",
			hamiltonTiempos.GetTiempos()[0].ImprimirTiempo()+"\n"+hamiltonTiempos.GetTiempos()[1].ImprimirTiempo()+"\n")
	}
}

func TestGetPilotoSesion(t *testing.T) {
	t.Log("Test GetPilotoSesion")

	var pses Piloto = hamiltonTiempos.GetPiloto()

	if pses != hamilton {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			hamilton.GetNombre(), pses.GetNombre())
	}
}

func TestSetPiloto(t *testing.T) {
	t.Log("Test SetPilotoSesion")

	hamiltonTiempos.SetPiloto(vettel)

	if hamiltonTiempos.GetPiloto() != vettel {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			vettel.GetNombre(), hamiltonTiempos.GetPiloto().GetNombre())
	}

	hamiltonTiempos.SetPiloto(hamilton)
}

func TestGetMejorTiempo(t *testing.T) {
	var mejvu TiempoVuelta = hamiltonTiempos.GetMejorTiempo()

	if mejvu != record {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			record.ImprimirTiempo(), mejvu.ImprimirTiempo())
	}
}

func TestSetMejorTiempo(t *testing.T) {
	hamiltonTiempos.SetMejorTiempo(cero)

	if hamiltonTiempos.GetMejorTiempo() != cero {
		t.Errorf("Resultado obtenido incorrecto, esperado %s, obtenido %s",
			cero.ImprimirTiempo(), record.ImprimirTiempo())
	}

	hamiltonTiempos.SetMejorTiempo(record)
}
