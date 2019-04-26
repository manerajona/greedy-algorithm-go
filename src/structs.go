package main

import (
	"log"
)

type Ciudad struct {
	Nombre, IdCiudad string
	Distancias       map[string]int
}

type Trace struct {
	Command string
	Logger  *log.Logger
}
