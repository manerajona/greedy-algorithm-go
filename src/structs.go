package main

import (
	"log"
)

type City struct {
	Name, IdCity string
	Distances map[string]int
}

type Trace struct {
	Command string
	Logger  *log.Logger
}
