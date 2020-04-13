package protometry

import (
	"math"
	"math/rand"
)

// NewVectorN constructs a VectorN
func NewVectorN(dimensions ...float64) *VectorN {
	return &VectorN{Dimensions: dimensions}
}

// Clone a vector
func (v *VectorN) Clone() *VectorN {
	var res []float64
	for i := range v.Dimensions {
		res = append(res, v.Get(i))
	}
	return NewVectorN(res...)
}

// NewVector3Zero constructs a VectorN of 3 dimensions initialized with 0
func NewVector3Zero() *VectorN {
	return &VectorN{Dimensions: []float64{0, 0, 0}}
}

// NewVector3One constructs a VectorN of 3 dimensions initialized with 1
func NewVector3One() *VectorN {
	return &VectorN{Dimensions: []float64{1, 1, 1}}
}

// Equal reports whether a and b are equal within a small epsilon.
func (v *VectorN) Equal(b VectorN) bool {
	if len(v.Dimensions) != len(b.Dimensions) {
		return false
	}
	const epsilon = 1e-16
	for i := range v.Dimensions {
		// If any dimensions aren't aproximately equal, return false
		if math.Abs(v.Get(i)-b.Get(i)) >= epsilon {
			return false
		}
	}
	// Else return true
	return true
}

// Get is used to shorten access to dimensions
func (v *VectorN) Get(dimension int) float64 {
	if dimension < 0 || dimension > len(v.Dimensions)-1 {
		return math.MaxFloat64
	}
	return v.Dimensions[dimension]
}

// Set is used to shorten dimensions assignment
func (v *VectorN) Set(dimension int, value float64) error {
	if dimension < 0 || dimension > len(v.Dimensions)-1 {
		return ErrVectorInvalidIndex
	}
	v.Dimensions[dimension] = value
	return nil
}

// SetAll is used to shorten dimensions assignment
func (v *VectorN) SetAll(value float64) {
	for i := range v.Dimensions {
		v.Dimensions[i] = value
	}
}

// Pow returns the vector pow
func (v *VectorN) Pow() *VectorN {
	var res []float64
	for i := range v.Dimensions {
		res = append(res, v.Get(i)*v.Get(i))
	}
	return NewVectorN(res...)
}

// Sum returns the sum of all the dimensions of the vector
func (v *VectorN) Sum() float64 {
	res := 0.
	for i := range v.Dimensions {
		res += v.Get(i)
	}
	return res
}

// Norm returns the norm.
func (v *VectorN) Norm() float64 { return v.Pow().Sum() }

// Norm2 returns the square of the norm.
func (v *VectorN) Norm2() float64 { return math.Sqrt(v.Norm()) }

// Normalize returns a unit vector in the same direction as a.
func (v *VectorN) Normalize() *VectorN {
	n2 := v.Norm2()
	if n2 == 0 {
		return NewVectorN(0, 0, 0)
	}
	return v.Scale(1 / math.Sqrt(n2))
}

// Abs returns the vector with nonnegative components.
func (v *VectorN) Abs() *VectorN {
	var res []float64
	for _, d := range v.Dimensions {
		res = append(res, math.Abs(d))
	}
	return NewVectorN(res...)
}

// Plus returns the standard vector sum of a and b.
func (v *VectorN) Plus(b VectorN) *VectorN {
	var res []float64
	for i := range v.Dimensions {
		res = append(res, v.Get(i)+b.Get(i))
	}
	return NewVectorN(res...)
}

// Minus returns the standard vector difference of a and b.
func (v *VectorN) Minus(b VectorN) *VectorN {
	var res []float64
	for i := range v.Dimensions {
		res = append(res, v.Get(i)-b.Get(i))
	}
	return NewVectorN(res...)
}

// Scale returns the standard scalar product of a and m.
func (v *VectorN) Scale(m float64) *VectorN {
	var res []float64
	for i := range v.Dimensions {
		res = append(res, v.Get(i)*m)
	}
	return NewVectorN(res...)
}

// Div returns the standard scalar division of a and m.
func (v *VectorN) Div(m float64) *VectorN {
	if m == 0 {
		return v
	}
     
    var res []float64
    for i := range v.Dimensions {
        res = append(res, v.Get(i)/m)
    }

    return NewVectorN(res...)
}

// Dot returns the standard dot product of a and b.
func (v *VectorN) Dot(b VectorN) float64 {
	res := .0
	for i := range v.Dimensions {
		res += (v.Get(i) * b.Get(i))
	}
	return res
}

// Cross returns the standard cross product of a and b.
func (v *VectorN) Cross(b VectorN) *VectorN {
	// Early error check to prevent redundant cloning
	if len(v.Dimensions) != 3 || len(b.Dimensions) != 3 {
		return nil
	}
	res := []float64{v.Get(1)*b.Get(2) - v.Get(2)*b.Get(1),
		v.Get(2)*b.Get(0) - v.Get(0)*b.Get(2),
		v.Get(0)*b.Get(1) - v.Get(1)*b.Get(0)}
	return NewVectorN(res...)
}

// Distance returns the Euclidean distance between a and b.
func (v *VectorN) Distance(b VectorN) float64 { return math.Sqrt(v.Minus(b).Pow().Sum()) }

// Angle returns the angle between a and b.
func (v *VectorN) Angle(b VectorN) float64 {
	cross := v.Cross(b)
	if cross == nil {
		return math.Atan2(cross.Norm(), v.Dot(b))
	}
	return math.MaxFloat64
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
func (v *VectorN) Lerp(b *VectorN, f float64) *VectorN {
	var res []float64
	for i := range v.Dimensions {
		res = append(res, (b.Get(i)-v.Get(i))*f+v.Get(i))
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
func Morton3D(v VectorN) uint { // TODO: decoder
	if len(v.Dimensions) != 3 {
		return 0
	}
	x := math.Min(math.Max(v.Get(0)*1024.0, 0.0), 1023.0)
	y := math.Min(math.Max(v.Get(1)*1024.0, 0.0), 1023.0)
	z := math.Min(math.Max(v.Get(2)*1024.0, 0.0), 1023.0)
	xx := expandBits(uint(x))
	yy := expandBits(uint(y))
	zz := expandBits(uint(z))
	return xx*4 + yy*2 + zz
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomCirclePoint returns a random circle point
func RandomCirclePoint(center VectorN, radius float64) VectorN {
	return *NewVectorN(randFloat(-radius+center.Get(0), radius+center.Get(0)),
		0,
		randFloat(-radius+center.Get(1), radius+center.Get(1)))
}

// RandomSpherePoint returns a random sphere point
func RandomSpherePoint(center VectorN, radius float64) VectorN {
	return *NewVectorN(randFloat(-radius+center.Get(0), radius+center.Get(0)),
		randFloat(-radius+center.Get(1), radius+center.Get(1)),
		randFloat(-radius+center.Get(1), radius+center.Get(1)))
}

// Concatenate join a sequence of arrays.
func Concatenate(v ...VectorN) VectorN {
	newV := VectorN{}
	for i := range v {
		newV.Dimensions = append(newV.Dimensions, v[i].Dimensions...)
	}
	return newV
}
