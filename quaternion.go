package protometry

import (
	"math"
)

// NewQuaternion constructs a VectorN
func NewQuaternion(x, y, z, w float64) *QuaternionN {
	return &QuaternionN{Value: &VectorN{Dimensions: []float64{x, y, z, w}}}
}

// LookAt return a quaternion corresponding to the rotation required to look at the other Vector3
func (v *VectorN) LookAt(b VectorN) *QuaternionN {
	angle := v.Angle(b)
	if angle == math.MaxFloat64 {
		return nil
	}
	return NewQuaternion(0, angle, 0, angle)
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
