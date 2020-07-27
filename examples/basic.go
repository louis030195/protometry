package main

import (
    vector32 "github.com/louis030195/protometry/api/vector3"
    volume2 "github.com/louis030195/protometry/api/volume"
    "log"
)

func main() {
	vector := vector32.NewVector3One()    // 1,1,1
	log.Printf("My vector: %f", vector.X) // 1
	vector.X = 12.6422

	v := vector32.NewVector3(0, 0, 0)
	v.Scale(3.5)                    // 0,0,0
	v2 := v.Plus(*v)                // 0,0,0
	log.Printf("My vector: %v", v2) // 0,0,0

	a := volume2.NewBoxMinMax(0, 0, 0, 1, 1, 1)
	b := volume2.NewBoxOfSize(2, 2, 2, 0.5)
	a.Fit(*b) // False
}
