package protometry

import (
	"fmt"
	"log"
	"math"
)

// In Returns whether the specified point is contained in this box.
func (a *VectorN) In(box Box) (bool, error) {
	if len(a.Dimensions) != len(box.min.Dimensions) {
		return false, ErrVectorInvalidDimension
	}

	for i := range a.Dimensions {
		if box.min.Get(i) > a.Get(i) || box.max.Get(i) < a.Get(i) {
			return false, nil
		}
	}
	return true, nil
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
func NewBox(minX, minY, minZ, maxX, maxY, maxZ float64) *Box {
	return &Box{
		min: Min(*NewVectorN(minX, minY, minZ), *NewVectorN(maxX, maxY, maxZ)),
		max: Max(*NewVectorN(minX, minY, minZ), *NewVectorN(maxX, maxY, maxZ)),
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
		if b.max.Get(i) < o.min.Get(i) || o.max.Get(i) < b.min.Get(i) {
			return false, nil
		}
	}
	return true, nil
}

// MakeSubBoxes split a box into subAreas
// TODO: not sure if it's correct, MAKE TEST
func (b *Box) MakeSubBoxes() [8]*Box {
	// gets the child boxes (octants) of the box.
	center := b.min.Lerp(&b.max, 0.5)

	return [8]*Box{
		NewBox(b.max.Get(0), b.max.Get(1), b.max.Get(2),
			center.Get(0), center.Get(1), center.Get(2)),
		NewBox(center.Get(0), b.max.Get(1), b.max.Get(2),
			b.min.Get(0), center.Get(1), center.Get(2)),
		NewBox(center.Get(0), center.Get(1), b.max.Get(2),
			b.min.Get(0), b.min.Get(1), center.Get(2)),
		NewBox(b.max.Get(0), center.Get(1), b.max.Get(2),
			center.Get(0), b.min.Get(1), center.Get(2)),
		NewBox(b.max.Get(0), b.max.Get(1), center.Get(2),
			center.Get(0), center.Get(1), b.min.Get(2)),
		NewBox(center.Get(0), b.max.Get(1), center.Get(2),
			b.min.Get(0), center.Get(1), b.min.Get(2)),
		NewBox(center.Get(0), center.Get(1), center.Get(2),
			b.min.Get(0), b.min.Get(1), b.min.Get(2)),
		NewBox(b.max.Get(0), center.Get(1), center.Get(2),
			center.Get(0), b.min.Get(1), b.min.Get(2)),
	}
}

func (b *Box) GetCenter() *VectorN {
	return b.min.Lerp(&b.max, 0.5)
}

// MinimumTranslation tells how much an entity has to move to no longer overlap another entity.
// TODO: 3D
func MinimumTranslation(a, b Box) VectorN {
	mtd := VectorN{}

	left := b.min.Get(0) - a.max.Get(0)
	right := b.max.Get(0) - a.min.Get(0)
	top := b.min.Get(1) - a.max.Get(1)
	bottom := b.max.Get(1) - a.min.Get(1)

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
		mtd.Set(0, left)
	} else {
		mtd.Set(0, right)
	}

	if math.Abs(top) < bottom {
		mtd.Set(1, top)
	} else {
		mtd.Set(1, bottom)
	}

	if math.Abs(mtd.Get(0)) < math.Abs(mtd.Get(1)) {
		mtd.Set(1, 0)
	} else {
		mtd.Set(0, 0)
	}

	return mtd
}

// ToString returns a human-readable representation of the box
func (b *Box) ToString() string {
	return fmt.Sprintf("min: %v, max: %v", b.min.ToString(), b.max.ToString())
}
