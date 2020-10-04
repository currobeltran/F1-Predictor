package main

import (
	f1predictor "f1predictor"
)

func main() {
	var record f1predictor.TiempoVuelta

	record.Constructor(1, 32, 456)

	println(record.GetSegundo())
}
