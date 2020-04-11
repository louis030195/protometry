package protometry

import (
	"fmt"
	"log"
	"math"
)


// Area is a 3-d interface representing volumes like Boxes, Spheres, Capsules ...
type Area interface {
	Fit(Area) bool
	Intersects(Area) bool
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
	Center  VectorN
	Extents VectorN
}

// NewBoxMinMax returns a new box using min max
func NewBoxMinMax(dims ...float64) *Box {
	b := &Box{}
	b.SetMinMax(*NewVectorN(dims[0:3]...), *NewVectorN(dims[3:6]...))
	return b
}

// NewBoxOfSize returns a box of size centered at center
func NewBoxOfSize(center VectorN, size float64) *Box {
	return &Box{
		Center:  center,
		Extents: *NewVectorN(size/2, size/2, size/2),
	}
}

// Equal returns whether a box is equal to another
func (b *Box) Equal(other Box) bool {
	return b.Center.Equal(other.Center) && b.Extents.Equal(other.Extents)
}

// GetMin ...
func (b *Box) GetMin() VectorN {
	return *b.Center.Minus(b.Extents)
}

// GetMax ...
func (b *Box) GetMax() VectorN {
	return *b.Center.Plus(b.Extents)
}

// GetSize returns the size of the box
func (b *Box) GetSize() VectorN {
	return *b.Extents.Scale(2)
}

// SetMinMax sets the box to the /min/ and /max/ value of the box.
func (b *Box) SetMinMax(min, max VectorN) {
	b.Extents = *(max.Minus(min)).Scale(0.5)
	b.Center = *min.Plus(b.Extents)
}

// EncapsulatePoint grows the box to include the /point/.
func (b *Box) EncapsulatePoint(point VectorN) {
	b.SetMinMax(Min(b.GetMin(), point), Max(b.GetMax(), point))
}

// EncapsulateBox grows the box to include the /box/.
func (b *Box) EncapsulateBox(box Box) {
	b.EncapsulatePoint(*box.Center.Minus(box.Extents))
	b.EncapsulatePoint(*box.Center.Plus(box.Extents))
}

// Expand the box by increasing its /size/ by /amount/ along each side.
func (b *Box) Expand(amount float64) {
	amount *= .5
	b.Extents = *b.Extents.Plus(*NewVectorN(amount, amount, amount))
}

// ExpandV the box by increasing its /size/ by /amount/ along each side.
func (b *Box) ExpandV(amount VectorN) {
	b.Extents = *b.Extents.Plus(*amount.Scale(.5))
}

// In Returns whether the specified point is contained in this box.
func (v *VectorN) In(box Box) bool {
	bm := box.GetMin()
	bmm := box.GetMax()
	for i := range v.Dimensions {
		if bm.Get(i) > v.Get(i) || bmm.Get(i) < v.Get(i) {
			return false
		}
	}
	return true
}

// Fit Returns whether the specified area is fully contained in the other area.
func (b *Box) Fit(o Box) bool {
	return b.Center.Plus(b.Extents).In(o) && b.Center.Minus(b.Extents).In(o)
}

// Intersects Returns whether any portion of this area intersects with the specified area or reversely.
func (b *Box) Intersects(bb Box) bool {
	bm := b.GetMin()
	bmm := b.GetMax()
	bbm := bb.GetMin()
	bbmm := bb.GetMax()
	for i := range bm.Dimensions {
		if bmm.Get(i) < bbm.Get(i) || bbmm.Get(i) < bm.Get(i) {
			return false
		}
	}
	return true
}


// Split split a CUBE into sub-cubes
func (b *Box) Split() [8]*Box {
	q := b.Extents.Get(0) / 2
	newExtents := *b.Extents.Scale(0.5)
	return [8]*Box{
		{Center: *b.Center.Plus(*NewVectorN(-q, q, -q)), Extents: newExtents},
		{Center: *b.Center.Plus(*NewVectorN(q, q, -q)), Extents: newExtents},
		{Center: *b.Center.Plus(*NewVectorN(-q, q, q)), Extents: newExtents},
		{Center: *b.Center.Plus(*NewVectorN(q, q, q)), Extents: newExtents},

		{Center: *b.Center.Plus(*NewVectorN(-q, -q, -q)), Extents: newExtents},
		{Center: *b.Center.Plus(*NewVectorN(q, -q, -q)), Extents: newExtents},
		{Center: *b.Center.Plus(*NewVectorN(-q, -q, q)), Extents: newExtents},
		{Center: *b.Center.Plus(*NewVectorN(q, -q, q)), Extents: newExtents},
	}
}

// MinimumTranslation tells how much an entity has to move to no longer overlap another entity.
// FIXME ? 3D ? 2D ?
func MinimumTranslation(b, bb Box) VectorN {
	mtd := VectorN{}
	bm := b.GetMin()
	bmm := b.GetMax()
	bbm := bb.GetMin()
	bbmm := bb.GetMax()
	left := bm.Get(0) - bbmm.Get(0)
	right := bmm.Get(0) - bbm.Get(0)
	top := bm.Get(1) - bbmm.Get(1)
	bottom := bmm.Get(1) - bbm.Get(1)

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
	bm := b.GetMin()
	bmm := b.GetMax()
	return fmt.Sprintf("Center: %v, \nExtents: %v, \nmin %v, \nmax %v",
		b.Center.ToString(), b.Extents.ToString(), bm.ToString(), bmm.ToString())
}
