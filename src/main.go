package main

import (
	recursos "m/src/api"
)

func main() {
	r := recursos.DiseñoRutas()

	r.Run()
}
