package main

import (
	"errors"
)

/*  This Method implements a Non-Greedy Algorithm in order to calculate
distances between capital cities of Argentina
*/
func calculatePathAndDistance(idCity string, capitals map[string]City) (Path, error) {

	var path []City
	var totalDistance int
	var err error

	//Verify existence
	if _, ok := capitals[idCity]; ok {

		c := capitals[idCity]
		delete(capitals, c.IdCity)
		path = append(path, c)

		for len(capitals) > 0 {
			distance := -1
			capital := c
			for k, v := range c.Distances {
				if _, ok := capitals[k]; ok {
					if distance < 0 || distance > v {
						distance = v
						capital = capitals[k]
					}
				}
			}
			c = capital
			delete(capitals, c.IdCity)
			path = append(path, c)
			totalDistance += distance
		}
		return Path{path, totalDistance}, nil
	} else {
		err = errors.New("The value idCity is incorrect")
	}
	return Path{path, totalDistance}, err
}
