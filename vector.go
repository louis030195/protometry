package protometry

import (
	"fmt"
	"math"
)

// NewVectorN constructs a VectorN
func NewVectorN(dimensions ...float64) *VectorN {
	return &VectorN{Dimensions: dimensions}
}

// NewVector3Zero constructs a VectorN of 3 dimensions initialized with 0
func NewVector3Zero() *VectorN {
	return &VectorN{Dimensions: []float64{0, 0, 0}}
}

// NewVector3One constructs a VectorN of 3 dimensions initialized with 1
func NewVector3One() *VectorN {
	return &VectorN{Dimensions: []float64{1, 1, 1}}
}

// ApproxEqual reports whether a and b are equal within a small epsilon.
func (a *VectorN) ApproxEqual(b VectorN) (bool, error) {
	if len(a.Dimensions) != len(b.Dimensions) {
		return false, ErrVectorInvalidDimension
	}
	const epsilon = 1e-16
	for i := range a.Dimensions {
		// If any dimensions aren't aproximately equal, return false
		if math.Abs(a.Get(i)-b.Get(i)) >= epsilon {
			return false, nil
		}
	}
	// Else return true
	return true, nil
}

// Get is used to shorten access to dimensions
func (a *VectorN) Get(dimension int) float64 {
	if dimension < 0 || dimension > len(a.Dimensions)-1 {
		return math.MaxFloat64
	}
	return a.Dimensions[dimension]
}

// Set is used to shorten dimensions assignment
func (a *VectorN) Set(dimension int, value float64) error {
	if dimension < 0 || dimension > len(a.Dimensions)-1 {
		return ErrVectorInvalidIndex
	}
	a.Dimensions[dimension] = value
	return nil
}

// ToString returns the vector to string
func (a *VectorN) ToString() string {
	res := "VectorN{ "
	for _, d := range a.Dimensions {
		res += fmt.Sprintf("%0.2f, ", d)
	}
	return res + "}"
}

// Pow returns the vector pow
func (a *VectorN) Pow() *VectorN {
	var copy []float64
	for _, d := range a.Dimensions {
		copy = append(copy, d*d)
	}
	return NewVectorN(copy...)
}

// Sum returns the sum of all the dimensions of the vector
func (a *VectorN) Sum() float64 {
	res := 0.
	for _, d := range a.Dimensions {
		res += d
	}
	return res
}

// Norm returns the norm.
func (a *VectorN) Norm() float64 { return a.Pow().Sum() }

// Norm2 returns the square of the norm.
func (a *VectorN) Norm2() float64 { return math.Sqrt(a.Norm()) }

// Normalize returns a unit vector in the same direction as a.
func (a *VectorN) Normalize() *VectorN {
	n2 := a.Norm2()
	if n2 == 0 {
		return NewVectorN(0, 0, 0)
	}
	return a.Mul(1 / math.Sqrt(n2))
}

// Abs returns the vector with nonnegative components.
func (a *VectorN) Abs() *VectorN {
	var res []float64
	for _, d := range a.Dimensions {
		res = append(res, math.Abs(d))
	}
	return NewVectorN(res...)
}

// Add returns the standard vector sum of a and b.
func (a *VectorN) Add(b VectorN) *VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, a.Get(i)+b.Get(i))
	}
	return NewVectorN(res...)
}

// Sub returns the standard vector difference of a and b.
func (a *VectorN) Sub(b VectorN) *VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, a.Get(i)-b.Get(i))
	}
	return NewVectorN(res...)
}

// Mul returns the standard scalar product of a and m.
func (a *VectorN) Mul(m float64) *VectorN {
	for i := range a.Dimensions {
		a.Dimensions[i] *= m
	}
	return a
}

// Div returns the standard scalar division of a and m.
func (a *VectorN) Div(m float64) (*VectorN, error) {
	if m == 0 {
		return nil, ErrDivisionByZero
	}
	for i := range a.Dimensions {
		a.Dimensions[i] /= m
	}
	return a, nil
}

// Dot returns the standard dot product of a and b.
func (a *VectorN) Dot(b VectorN) float64 {
	res := .0
	for i := range a.Dimensions {
		res += (a.Get(i) * b.Get(i))
	}
	return res
}

// Cross returns the standard cross product of a and b.
func (a *VectorN) Cross(b VectorN) (*VectorN, error) {
	// Early error check to prevent redundant cloning
	if len(a.Dimensions) != 3 || len(b.Dimensions) != 3 {
		return nil, ErrVectorInvalidDimension
	}
	res := []float64{a.Get(1)*b.Get(2) - a.Get(2)*b.Get(1),
		a.Get(2)*b.Get(0) - a.Get(0)*b.Get(2),
		a.Get(0)*b.Get(1) - a.Get(1)*b.Get(0)}
	return NewVectorN(res...), nil
}

// Distance returns the Euclidean distance between a and b.
func (a *VectorN) Distance(b VectorN) float64 { return math.Sqrt(a.Sub(b).Pow().Sum()) }

// Angle returns the angle between a and b.
func (a *VectorN) Angle(b VectorN) (float64, error) {
	cross, err := a.Cross(b)
	if err == nil {
		return math.Atan2(cross.Norm(), a.Dot(b)), nil
	}
	return 0, err
}

// Min Returns the a vector where each component is the lesser of the
// corresponding component in this and the specified vector.
func Min(a VectorN, b VectorN) VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, math.Min(a.Get(i), b.Get(i)))
	}
	return *NewVectorN(res...)
}

// Max Returns the a vector where each component is the greater of the
// corresponding component in this and the specified vector.
func Max(a VectorN, b VectorN) VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, math.Max(a.Get(i), b.Get(i)))
	}
	return *NewVectorN(res...)
}

// Lerp Returns the linear interpolation between two VectorN(s).
func (a *VectorN) Lerp(b *VectorN, f float64) *VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, (b.Get(i)-a.Get(i))*f+a.Get(i))
	}
	return NewVectorN(res...)
}

// Expands a 10-bit integer into 30 bits
// by inserting 2 zeros after each bit.
func expandBits(v uint) uint {
	v = (v * 0x00010001) & 0xFF0000FF
	v = (v * 0x00000101) & 0x0F00F00F
	v = (v * 0x00000011) & 0xC30C30C3
	v = (v * 0x00000005) & 0x49249249
	return v
}

// Morton3D Calculates a 30-bit Morton code for the
// given 3D point located within the unit cube [0,1].
func Morton3D(v VectorN) (uint, error) { // TODO: decoder
	if len(v.Dimensions) != 3 {
		return 0, ErrVectorInvalidDimension
	}
	x := math.Min(math.Max(v.Get(0)*1024.0, 0.0), 1023.0)
	y := math.Min(math.Max(v.Get(1)*1024.0, 0.0), 1023.0)
	z := math.Min(math.Max(v.Get(2)*1024.0, 0.0), 1023.0)
	xx := expandBits(uint(x))
	yy := expandBits(uint(y))
	zz := expandBits(uint(z))
	return xx*4 + yy*2 + zz, nil
}
