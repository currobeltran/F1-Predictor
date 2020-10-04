package f1predictor

type Circuito struct {
	Nombre         string
	Pais           string
	RecordCircuito TiempoVuelta
	Resultados     []ResultadoGP
}

type OpsEstadisticas interface {
	media(circuito Circuito, estadistica string) float32
}

func media(circuito Circuito, estadistica string) float32 {
	var acumulado int = 0
	var x float32 = 0

	for i := 0; i < len(circuito.Resultados); i++ {
		switch estadistica {
		case "accidentes":
			acumulado += circuito.Resultados[i].Estadisticas.Accidentes
			break
		case "banderasAmarillas":
			acumulado += circuito.Resultados[i].Estadisticas.BanderasAmarillas
			break
		case "banderasRojas":
			acumulado += circuito.Resultados[i].Estadisticas.BanderasRojas
			break
		case "adelantamientos":
			acumulado += circuito.Resultados[i].Estadisticas.Adelantamientos
			break
		case "sanciones":
			acumulado += circuito.Resultados[i].Estadisticas.Sanciones
			break
		case "numeroSafetyCar":
			acumulado += circuito.Resultados[i].Estadisticas.NumeroSafetyCar
			break
		case "default":
			break
		}

	}

	x = float32(acumulado) / float32(len(circuito.Resultados))

	return x
}
