package main

import (
    protometry "github.com/louis030195/protometry/pkg"
    "log"
)

func main() {
    vector := protometry.NewVector3One() // 1,1,1
    log.Printf("My vector: %f", vector.X) // 1
    vector.X = 12.6422

    v := protometry.NewVector3(0, 0, 0)
    v.Scale(3.5) // 0,0,0
    v2 := v.Plus(*v) // 0,0,0
    log.Printf("My vector: %v", v2) // 0,0,0

    a := protometry.NewBoxMinMax(0, 0, 0, 1, 1, 1)
    b := protometry.NewBoxOfSize(2, 2, 2, 0.5)
    a.Fit(*b) // False
}
