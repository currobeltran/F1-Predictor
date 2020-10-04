package f1predictor

type ResultadoGP struct {
	ResultadoFP1           [20]SesionPiloto
	ResultadoFP2           [20]SesionPiloto
	ResultadoFP3           [20]SesionPiloto
	ResultadoClasificacion [20]SesionPiloto
	ResultadoCarrera       [20]SesionPiloto
	Ganador                Piloto
	Poleman                Piloto
	Podio                  [3]Piloto
	Temporada              int
	Estadisticas           EstadisticasGP
}
