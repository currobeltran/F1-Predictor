package f1predictor

import "strconv"

type GetSetTiempoVuelta interface {
	Constructor(m int, s int, ms int) TiempoVuelta

	GetMinuto() int
	SetMinuto(m int)

	GetSegundo() int
	SetSegundo(s int)

	GetMilesima() int
	SetMilesima(ms int)

	ImprimirTiempo() string
}

type TiempoVuelta struct {
	Minuto   int
	Segundo  int
	Milesima int
}

func (t *TiempoVuelta) Constructor(m int, s int, ms int) {
	t.Minuto = m
	t.Segundo = s
	t.Milesima = ms
}

func (t TiempoVuelta) GetMinuto() int {
	return t.Minuto
}

func (t *TiempoVuelta) SetMinuto(m int) {
	t.Minuto = m
}

func (t TiempoVuelta) GetSegundo() int {
	return t.Segundo
}

func (t *TiempoVuelta) SetSegundo(s int) {
	t.Segundo = s
}

func (t TiempoVuelta) GetMilesima() int {
	return t.Milesima
}

func (t *TiempoVuelta) SetMilesima(ms int) {
	t.Milesima = ms
}

func (t TiempoVuelta) ImprimirTiempo() string {
	var ret string

	ret = (strconv.Itoa(t.Minuto) + ":" + strconv.Itoa(t.Segundo) + "." + strconv.Itoa(t.Milesima))

	return ret
}
