package main

import (
	"strconv"
)

/*  Metodo que calcula el recorrido
    @param idCiudad ciudad de partida
    @param capitales mapa de ciudades
    @return array con el recorrido
    @return distancia total
*/
func obtenerRecorrido(idCiudad string, capitales map[string]Ciudad) ([]string, string) {

	if idCiudad != "" {
		recorrido := []string{}
		recorrido = append(recorrido, "Inicio:")
		distanciaTotal := 0

		c := capitales[idCiudad]
		delete(capitales, c.IdCiudad)
		recorrido = append(recorrido, c.Nombre+" >>")

		for len(capitales) > 0 {
			distancia := -1
			capital := c
			for k, v := range c.Distancias {
				/*Verifica que exista k en capitales*/
				if _, ok := capitales[k]; ok {
					if distancia < 0 || distancia > v {
						distancia = v
						capital = capitales[k]
					}
				}
			}
			c = capital
			delete(capitales, c.IdCiudad)
			r := c.Nombre + " [" + strconv.Itoa(distancia) + "km] >>"
			recorrido = append(recorrido, r)
			distanciaTotal += distancia
		}
		recorrido = append(recorrido, "Fin")
		return recorrido, strconv.Itoa(distanciaTotal)
	}
	return nil, ""
}

func mostrar(c Ciudad, tr *Trace) {
	tr.Logger.Print(c.IdCiudad + ") " + c.Nombre)
}
