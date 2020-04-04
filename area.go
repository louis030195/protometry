package protometry

import (
	"fmt"
	"log"
	"math"
)

/**
 * https://github.com/vanruesc/sparse-octree/blob/master/src/core/layout.js
 * A binary pattern that describes the standard octant layout:
 *
 * ```text
 *    3____7
 *  2/___6/|
 *  | 1__|_5
 *  0/___4/
 * ```
 *
 * This common layout is crucial for positional assumptions.
 *
 * @type {Uint8Array[]}
 */

var boxLayout = [8][3]int{
	[3]int{0, 0, 0},
	[3]int{0, 0, 1},
	[3]int{0, 1, 0},
	[3]int{0, 1, 1},

	[3]int{1, 0, 0},
	[3]int{1, 0, 1},
	[3]int{1, 1, 0},
	[3]int{1, 1, 1},
}

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
	bm := b.GetMin()
	bmm := b.GetMax()
	return bm.In(o) && bmm.In(o)
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

// Split split a box into sub boxes
func (b *Box) Split() [8]*Box {
	// center := b.Center
	// halfSize := b.Extents.Get(0) * .5 // Assume cube
	//var split [8]*Box
	// for i := 0; i < 8; i++ {
	// 	combination := boxLayout[i]
	// 	split[i] = &Box{Extents: *b.Extents.Scale(0.5)}
	// 	split[i] = NewBoxOfSize(NewVectorN(halfSize.Scale(combination[j]).Plus(center.Get(j))), halfSize)
	// }
	//return split

	// var split [8]*Box
	// n := b.Extents.Scale(0.5)
	// for i := 0; i < 8; i++ {
	// 	split[i] = NewBoxMinMax(
	// 		b.Center.Get(0)-n.Get(0),
	// 		b.Center.Get(1)-n.Get(1),
	// 		b.Center.Get(2)-n.Get(2),
	// 		b.Center.Get(0)+n.Get(0),
	// 		b.Center.Get(1)+n.Get(1),
	// 		b.Center.Get(2)+n.Get(2),
	// 	)
	// }
	// return split
	// min := b.GetMin()
	// center := b.Center
	// halfSize := b.Extents.Get(0) / 2 // Assume cube
	// var split [8]*Box
	// var x, y, z float64
	// for i := 0; i < 8; i++ {
	// 	combination := boxLayout[i]
	// 	if combination[0] == 0 {
	// 		x = min.Get(0)
	// 	} else {
	// 		x = center.Get(0)
	// 	}
	// 	if combination[1] == 0 {
	// 		y = min.Get(1)
	// 	} else {
	// 		y = center.Get(1)
	// 	}
	// 	if combination[2] == 0 {
	// 		z = min.Get(2)
	// 	} else {
	// 		z = center.Get(2)
	// 	}
	// 	split[i] = NewBoxOfSize(*NewVectorN(x, y, z), halfSize)
	// }
	// return split

	// bm := b.GetMin()
	// bmm := b.GetMax()
	// return [8]*Box{
	// 	NewBoxMinMax(bmm.Get(0), bmm.Get(1), bmm.Get(2),
	// 		b.Center.Get(0), b.Center.Get(1), b.Center.Get(2)),
	// 	NewBoxMinMax(b.Center.Get(0), bmm.Get(1), bmm.Get(2),
	// 		bm.Get(0), b.Center.Get(1), b.Center.Get(2)),
	// 	NewBoxMinMax(b.Center.Get(0), b.Center.Get(1), bmm.Get(2),
	// 		bm.Get(0), bm.Get(1), b.Center.Get(2)),
	// 	NewBoxMinMax(bmm.Get(0), b.Center.Get(1), bmm.Get(2),
	// 		b.Center.Get(0), bm.Get(1), b.Center.Get(2)),
	// 	NewBoxMinMax(bmm.Get(0), bmm.Get(1), b.Center.Get(2),
	// 		b.Center.Get(0), b.Center.Get(1), bm.Get(2)),
	// 	NewBoxMinMax(b.Center.Get(0), bmm.Get(1), b.Center.Get(2),
	// 		bm.Get(0), b.Center.Get(1), bm.Get(2)),
	// 	NewBoxMinMax(b.Center.Get(0), b.Center.Get(1), b.Center.Get(2),
	// 		bm.Get(0), bm.Get(1), bm.Get(2)),
	// 	NewBoxMinMax(bmm.Get(0), b.Center.Get(1), b.Center.Get(2),
	// 		b.Center.Get(0), bm.Get(1), bm.Get(2)),
	// }
	quarter := b.Extents.Get(0) / 2 // Assuming cube
	newLength := b.Extents.Get(0) / 2
	return [8]*Box{
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(-quarter, quarter, -quarter)), newLength),
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(quarter, quarter, -quarter)), newLength),
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(-quarter, quarter, quarter)), newLength),
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(-quarter, -quarter, -quarter)), newLength),
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(-quarter, -quarter, -quarter)), newLength),
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(quarter, -quarter, -quarter)), newLength),
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(-quarter, -quarter, quarter)), newLength),
		NewBoxOfSize(*b.Center.Plus(*NewVectorN(quarter, -quarter, quarter)), newLength),
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
