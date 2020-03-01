package protometry

import (
	"fmt"
	"log"
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
		if math.Abs(a.Dimensions[i]-b.Dimensions[i]) >= epsilon {
			return false, nil
		}
	}
	// Else return true
	return true, nil
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
		res = append(res, a.Dimensions[i]+b.Dimensions[i])
	}
	return NewVectorN(res...)
}

// Sub returns the standard vector difference of a and b.
func (a *VectorN) Sub(b VectorN) *VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, a.Dimensions[i]-b.Dimensions[i])
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
		res += (a.Dimensions[i] * b.Dimensions[i])
	}
	return res
}

// Cross returns the standard cross product of a and b.
func (a *VectorN) Cross(b VectorN) (*VectorN, error) {
	// Early error check to prevent redundant cloning
	if len(a.Dimensions) != 3 || len(b.Dimensions) != 3 {
		return nil, ErrVectorInvalidDimension
	}
	res := []float64{a.Dimensions[1]*b.Dimensions[2] - a.Dimensions[2]*b.Dimensions[1],
		a.Dimensions[2]*b.Dimensions[0] - a.Dimensions[0]*b.Dimensions[2],
		a.Dimensions[0]*b.Dimensions[1] - a.Dimensions[1]*b.Dimensions[0]}
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
		res = append(res, math.Min(a.Dimensions[i], b.Dimensions[i]))
	}
	return *NewVectorN(res...)
}

// Max Returns the a vector where each component is the greater of the
// corresponding component in this and the specified vector.
func Max(a VectorN, b VectorN) VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, math.Max(a.Dimensions[i], b.Dimensions[i]))
	}
	return *NewVectorN(res...)
}

// Lerp Returns the linear interpolation between two VectorN(s).
func (a *VectorN) Lerp(b *VectorN, f float64) *VectorN {
	var res []float64
	for i := range a.Dimensions {
		res = append(res, (b.Dimensions[i]-a.Dimensions[i])*f+a.Dimensions[i])
	}
	return NewVectorN(res...)
}

// In Returns whether the specified point is contained in this box.
func (a *VectorN) In(box Box) (bool, error) {
	if len(a.Dimensions) != len(box.min.Dimensions) {
		return false, ErrVectorInvalidDimension
	}

	for i := range a.Dimensions {
		if box.min.Dimensions[i] > a.Dimensions[i] || box.max.Dimensions[i] < a.Dimensions[i] {
			return false, nil
		}
	}
	return true, nil
}

// NewQuaternion constructs a VectorN
func NewQuaternion(x, y, z, w float64) *QuaternionN {
	return &QuaternionN{Value: &VectorN{Dimensions: []float64{x, y, z, w}}}
}

// LookAt return a quaternion corresponding to the rotation required to look at the other Vector3
func (a *VectorN) LookAt(b VectorN) (*QuaternionN, error) {
	angle, err := a.Angle(b)
	if err != nil {
		return nil, err
	}
	return NewQuaternion(0, angle, 0, angle), nil
}

// LookAtTwo ...
func LookAtTwo(from, to VectorN) ([][]float64, error) {
	tmp := NewVectorN(0, 1, 0)
	forward := from.Sub(to).Normalize()
	right, err := tmp.Normalize().Cross(*forward)
	if err != nil {
		return nil, err
	}
	up, err := forward.Cross(*right)
	if err != nil {
		return nil, err
	}

	a := make([][]float64, 4)
	for i := range a {
		a[i] = make([]float64, 3)
	}

	a[0][0] = right.Dimensions[0]
	a[0][1] = right.Dimensions[1]
	a[0][2] = right.Dimensions[2]
	a[1][0] = up.Dimensions[0]
	a[1][1] = up.Dimensions[1]
	a[1][2] = up.Dimensions[2]
	a[2][0] = forward.Dimensions[0]
	a[2][1] = forward.Dimensions[1]
	a[2][2] = forward.Dimensions[2]

	a[3][0] = from.Dimensions[0]
	a[3][1] = from.Dimensions[1]
	a[3][2] = from.Dimensions[2]

	return a, nil
}

// ToQuaternion ... yaw (Z), pitch (Y), roll (X)
func ToQuaternion(yaw, pitch, roll float64) *QuaternionN {
	// Abbreviations for the various angular functions
	cy := math.Cos(yaw * 0.5)
	sy := math.Sin(yaw * 0.5)
	cp := math.Cos(pitch * 0.5)
	sp := math.Sin(pitch * 0.5)
	cr := math.Cos(roll * 0.5)
	sr := math.Sin(roll * 0.5)

	return NewQuaternion(cy*cp*cr+sy*sp*sr, cy*cp*sr-sy*sp*cr, sy*cp*sr+cy*sp*cr, sy*cp*cr-cy*sp*sr)
}

// Area is a 3-d interface representing volumes like Boxes, Spheres, Capsules ...
type Area interface {
	Inside(Area) (bool, error)
	Intersects(Area) (bool, error)
}

// Sphere TODO
type Sphere struct {
}

// Capsule TODO
type Capsule struct {
}

// Convex TODO
type Convex struct {
}

// Box ...
type Box struct {
	min VectorN
	max VectorN
}

// GetMin ...
func (b *Box) GetMin() VectorN {
	return b.min
}

// GetMax ...
func (b *Box) GetMax() VectorN {
	return b.max
}

func (b *Box) GetSize() float64 {
	return math.Abs(b.max.Distance(b.min))
}

// NewBox constructs and returns a new box
func NewBox(min, max VectorN) *Box {
	return &Box{
		min: Min(min, max),
		max: Max(min, max),
	}
}

func NewBoxOfSize(position VectorN, size float64) *Box {
	return &Box{
		min: *position.Sub(*NewVectorN(size, size, size)),
		max: *position.Add(*NewVectorN(size, size, size)),
	}
}

// Inside Returns whether the specified area is fully contained in the other area.
func (b *Box) Inside(o Box) (bool, error) {
	minIn, errMin := b.min.In(o)
	if errMin != nil {
		return false, errMin
	}
	maxIn, errMax := b.max.In(o)
	if errMax != nil {
		return false, errMax
	}
	return minIn && maxIn, nil
}

// Intersects Returns whether any portion of this area intersects with the specified area or reversely.
func (b *Box) Intersects(o Box) (bool, error) {
	if len(b.min.Dimensions) != len(o.min.Dimensions) {
		return false, ErrVectorInvalidDimension
	}
	for i := range b.min.Dimensions {
		if b.max.Dimensions[i] < o.min.Dimensions[i] || o.max.Dimensions[i] < b.min.Dimensions[i] {
			return false, nil
		}
	}
	return true, nil
}

// MakeSubBoxes split a box into subAreas
// TODO: not sure if it's correct
func (b *Box) MakeSubBoxes() []*Box {
	// gets the child boxes (octants) of the box.
	center := b.min.Lerp(&b.max, 0.5)

	return []*Box{
		NewBox(*NewVectorN(b.max.Dimensions[0], b.max.Dimensions[1], b.max.Dimensions[2]),
			*NewVectorN(center.Dimensions[0:3]...)),
		NewBox(*NewVectorN(center.Dimensions[0]+1, b.max.Dimensions[1], b.max.Dimensions[2]),
			*NewVectorN(b.min.Dimensions[0], center.Dimensions[1], center.Dimensions[2])),
		NewBox(*NewVectorN(center.Dimensions[0]+1, center.Dimensions[1]+1, b.max.Dimensions[2]),
			*NewVectorN(b.min.Dimensions[0], b.min.Dimensions[1], center.Dimensions[2])),
		NewBox(*NewVectorN(b.max.Dimensions[0], center.Dimensions[1]+1, b.max.Dimensions[2]),
			*NewVectorN(center.Dimensions[0], b.min.Dimensions[1], center.Dimensions[2])),
		NewBox(*NewVectorN(b.max.Dimensions[0], b.max.Dimensions[1], center.Dimensions[2]+1),
			*NewVectorN(center.Dimensions[0], center.Dimensions[1], b.min.Dimensions[2])),
		NewBox(*NewVectorN(center.Dimensions[0]+1, b.max.Dimensions[1], center.Dimensions[2]+1),
			*NewVectorN(b.min.Dimensions[0], center.Dimensions[1], b.min.Dimensions[2])),
		NewBox(*NewVectorN(center.Dimensions[0]+1, center.Dimensions[1]+1, center.Dimensions[2]+1),
			*NewVectorN(b.min.Dimensions[0], b.min.Dimensions[1], b.min.Dimensions[2])),
		NewBox(*NewVectorN(b.max.Dimensions[0], center.Dimensions[1]+1, center.Dimensions[2]+1),
			*NewVectorN(center.Dimensions[0], b.min.Dimensions[1], b.min.Dimensions[2])),
	}
}

func (b *Box) GetCenter() *VectorN {
	return b.min.Lerp(&b.max, 0.5)
}

// MinimumTranslation tells how much an entity has to move to no longer overlap another entity.
// TODO: 3D
func MinimumTranslation(a, b Box) VectorN {
	mtd := VectorN{}

	left := b.min.Dimensions[0] - a.max.Dimensions[0]
	right := b.max.Dimensions[0] - a.min.Dimensions[0]
	top := b.min.Dimensions[1] - a.max.Dimensions[1]
	bottom := b.max.Dimensions[1] - a.min.Dimensions[1]

	if left > 0 || right < 0 {
		log.Println("Box aint intercepting")
		return mtd
		//box doesn't intercept
	}

	if top > 0 || bottom < 0 {
		log.Println("Box aint intercepting")
		return mtd
		//box doesn't intercept
	}
	if math.Abs(left) < right {
		mtd.Dimensions[0] = left
	} else {
		mtd.Dimensions[0] = right
	}

	if math.Abs(top) < bottom {
		mtd.Dimensions[1] = top
	} else {
		mtd.Dimensions[1] = bottom
	}

	if math.Abs(mtd.Dimensions[0]) < math.Abs(mtd.Dimensions[1]) {
		mtd.Dimensions[1] = 0
	} else {
		mtd.Dimensions[0] = 0
	}

	return mtd
}
