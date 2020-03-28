package protometry

import (
	"math"
)

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

	a[0][0] = right.Get(0)
	a[0][1] = right.Get(1)
	a[0][2] = right.Get(2)
	a[1][0] = up.Get(0)
	a[1][1] = up.Get(1)
	a[1][2] = up.Get(2)
	a[2][0] = forward.Get(0)
	a[2][1] = forward.Get(1)
	a[2][2] = forward.Get(2)

	a[3][0] = from.Get(0)
	a[3][1] = from.Get(1)
	a[3][2] = from.Get(2)

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
