package f1predictor

/*
 * Clase circuito con la podremos acceder a distintos datos correspondientes
 * a un circuito del campeonato de Formula 1. Los atributos son:
 * 		- Nombre del circuito.
 * 		- País donde se encuentra.
 *		- Record de tiempo sobre la pista.
 *		- Acumulación histórica de resultados de los GPs disputados en dicha pista.
 */

type Circuito struct {
	nombre         string
	pais           string
	recordCircuito TiempoVuelta
	resultados     []ResultadoGP
}

type MetodosCircuito interface {
	Constructor(nombre string, pais string)

	GetNombre() string
	SetNombre(nombre string)

	GetPais() string
	SetPais(pais string)

	GetRecordCircuito() TiempoVuelta
	SetRecordCircuito(nuevoRecord TiempoVuelta)

	GetResultados() []ResultadoGP
	SetResultados(resultados []ResultadoGP)
	SetResultadoByIndex(i int, resultado ResultadoGP)

	Media(estadistica string) float32
}

func (c *Circuito) Constructor(nombre string, pais string) {
	c.nombre = nombre
	c.pais = pais

	var tmp TiempoVuelta
	tmp.Constructor(0, 0, 0)
	c.recordCircuito = tmp
}

func (c *Circuito) GetNombre() string {
	return c.nombre
}

func (c *Circuito) SetNombre(nombre string) {
	c.nombre = nombre
}

func (c *Circuito) GetPais() string {
	return c.pais
}

func (c *Circuito) SetPais(pais string) {
	c.pais = pais
}

func (c *Circuito) GetRecordCircuito() TiempoVuelta {
	return c.recordCircuito
}

func (c *Circuito) SetRecordCircuito(nuevoRecord TiempoVuelta) {
	c.recordCircuito = nuevoRecord
}

func (c *Circuito) GetResultados() []ResultadoGP {
	return c.resultados
}

func (c *Circuito) SetResultados(resultados []ResultadoGP) {
	c.resultados = resultados
}

func (c *Circuito) SetResultadosByIndex(i int, resultado ResultadoGP) {
	c.resultados[i] = resultado
}

/*
 * Método para calcular la media de una estadística concreta de un circuito
 * concreto
 */

func (c *Circuito) Media(estadistica string) float32 {
	var acumulado int = 0
	var x float32 = 0

	for i := 0; i < len(c.resultados); i++ {
		var est EstadisticasGP = c.resultados[i].GetEstadisticas()

		switch estadistica {
		case "accidentes":
			acumulado += est.GetAccidentes()
			break
		case "banderasAmarillas":
			acumulado += est.GetBanderasAmarillas()
			break
		case "banderasRojas":
			acumulado += est.GetBanderasRojas()
			break
		case "adelantamientos":
			acumulado += est.GetAdelantamientos()
			break
		case "sanciones":
			acumulado += est.GetSanciones()
			break
		case "numeroSafetyCar":
			acumulado += est.GetNumeroSafetyCar()
			break
		case "default":
			break
		}

	}

	x = float32(acumulado) / float32(len(c.resultados))

	return x
}
