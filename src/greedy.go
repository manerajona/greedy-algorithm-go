package main

import (
	"strconv"
)

/*  This Method implements a Greedy Algorithm in order to calculate
distances between capital cities of Argentina
*/
func calculatePathAndDistance(idCity string, capitals map[string]City) ([]string, string) {

	if idCity != "" {
		path := []string{}
		path = append(path, "Begin:")
		total := 0

		c := capitals[idCity]
		delete(capitals, c.IdCity)
		path = append(path, c.Name+" >>")

		for len(capitals) > 0 {
			distance := -1
			capital := c
			for k, v := range c.Distances {
				/*Verify existence*/
				if _, ok := capitals[k]; ok {
					if distance < 0 || distance > v {
						distance = v
						capital = capitals[k]
					}
				}
			}
			c = capital
			delete(capitals, c.IdCity)
			r := c.Name + " [" + strconv.Itoa(distance) + "km] >>"
			path = append(path, r)
			total += distance
		}
		path = append(path, "End")
		return path, strconv.Itoa(total)
	}
	return nil, ""
}
