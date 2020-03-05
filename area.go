package protometry

import (
	"log"
	"math"
)

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
