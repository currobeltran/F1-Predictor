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
	minuto   int
	segundo  int
	milesima int
}

func (t *TiempoVuelta) Constructor(m int, s int, ms int) {
	t.minuto = m
	t.segundo = s
	t.milesima = ms
}

func (t TiempoVuelta) GetMinuto() int {
	return t.minuto
}

func (t *TiempoVuelta) SetMinuto(m int) {
	t.minuto = m
}

func (t TiempoVuelta) GetSegundo() int {
	return t.segundo
}

func (t *TiempoVuelta) SetSegundo(s int) {
	t.segundo = s
}

func (t TiempoVuelta) GetMilesima() int {
	return t.milesima
}

func (t *TiempoVuelta) SetMilesima(ms int) {
	t.milesima = ms
}

func (t TiempoVuelta) ImprimirTiempo() string {
	var ret string

	ret = (strconv.Itoa(t.minuto) + ":" + strconv.Itoa(t.segundo) + "." + strconv.Itoa(t.milesima))

	return ret
}
