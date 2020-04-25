
# protometry

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/52ed0a7a050c470ababeb6e888d51878)](https://app.codacy.com/gh/The-Tensox/protometry?utm_source=github.com&utm_medium=referral&utm_content=The-Tensox/protometry&utm_campaign=Badge_Grade_Dashboard)
[![Build Status](https://img.shields.io/circleci/project/The-Tensox/protometry/master.svg)](https://circleci.com/gh/The-Tensox/protometry)

Geometry on protobuf stubs, could be also implemented in other languages.

## Why developing another Math library instead of using existing ones ?

This library is targeting protobuf types especially, allowing to do math directly on networked data.

## Installation

```bash
go get github.com/louis030195/protometry
# Build for gRPC
# If using gRPC go get -u github.com/golang/protobuf/protoc-gen-go
protoc -I . --go_out=plugins=grpc:. *.proto
# Without gRPC
# protoc -I . --go_out=. *.proto
```

## Usage

```go
vector := NewVector3One() // 1,1,1
log.Printf("My vector: %s", vector.Get(1)) // 1
vector.Set(0, 12.6422)

v := NewVector3(0, 0, 0)
v.Scale(3.5) // 0,0,0
v.Plus(v) // 0,0,0

a := NewBoxMinMax(*NewVector3(0, 0, 0), *NewVector3(1, 1, 1))
b := NewBoxOfSize(*NewVector3(2, 2, 2), 0.5)
a.Fit(*b) // False
```

## Features

### Vectors

- [x] Build vectors of N dimensions
- [x] Vector comparison (min, max, equal)
- [x] Normalization
- [x] Absolute value
- [x] Plus, Minus, Scale, Dot(vector product), Div(scalar division), Cross product, Euclidean Norm, Angle, Lerp

### Volumes

- [x] Box Intersections, Fit

## Test

```bash
go test -v
```

## Benchmarks

```bash
go test -benchmem -run XXX -bench . -benchtime 100000x # or -benchtime 0.02s
```

![](docs/bench.png)

<!--
|Name   |   Runs   |   time   |   Bytes   |   Allocs   |
|:-----:|:--------:|:--------:|:---------:|:----------:|
|BenchmarkArea_NewBoxMinMax-8   |   1169382   |   1045 ns/op   |   472 B/op  |   13 allocs/op   |
|BenchmarkArea_NewBoxOfSize-8   |   380234211   |   3.07 ns/op   |   0 B/op   |   0 allocs/op   |
|BenchmarkArea_Fit-8   |   683202   |   1816 ns/op   |   720 B/op   |   24 allocs/op   |
|BenchmarkArea_Intersects-8   |   971000   |   1255 ns/op   |   480 B/op   |   16 allocs/op   |
|BenchmarkArea_Split-8   |   277064   |   3703 ns/op   |   1464 B/op   |   44 allocs/op   |
-->

## TODO

- [ ] Encoding
- [ ] Handle other volumes (sphere, capsule, mesh ...) and make them follow [an interface](./volumen.go)
- [ ] Take inspiration from [numpy](https://docs.scipy.org/doc/numpy/reference/routines.math.html#arithmetic-operations)
- [ ] Improve benchmarks
- [ ] Setup CI for verifying performances (benchmarks)
- [ ] Implement uncovered test cases
