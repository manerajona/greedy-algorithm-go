package main

type City struct {
	Name, IdCity string
	Distances    map[string]int
}

type Path struct {
	Cities   []City
	Distance int
}
