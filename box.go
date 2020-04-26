package protometry

import (
	"log"
	"math"
)


// NewBoxMinMax returns a new box using min max
func NewBoxMinMax(minX, minY, minZ, maxX, maxY, maxZ float64) *Box {
	return &Box{
		Min: NewVector3(minX, minY, minZ),
		Max: NewVector3(maxX, maxY, maxZ),
	}
}

// NewBoxOfSize returns a box of size centered at center
func NewBoxOfSize(x, y, z, size float64) *Box {
	half := size/2
	min := *NewVector3(x-half, y-half, z-half)
	max := *NewVector3(x+half, y+half, z+half)
	return &Box{
		Min: &min,
		Max: &max,
	}
}

// Equal returns whether a box is equal to another
func (b Box) Equal(other Box) bool {
	return b.Min.Equal(*other.Min) && b.Max.Equal(*other.Max)
}

// GetCenter ...
func (b Box) GetCenter() Vector3 {
	return *b.Min.Lerp(b.Max, 0.5)
}

// GetSize returns the size of the box
func (b *Box) GetSize() Vector3 {
	return b.Max.Minus(*b.Min)
}

// Fit Returns whether the specified area is fully contained in the other area.
func (b Box) Fit(o Box) bool {
	return b.Max.In(o) && b.Min.In(o)
}

// Intersects Returns whether any portion of this area intersects with the specified area or reversely.
func (b Box) Intersects(b2 Box) bool {
	return !(b.Max.X < b2.Min.X || b2.Max.X < b.Min.X || b.Max.Y < b2.Min.Y || b2.Max.Y < b.Min.Y || b.Max.Z < b2.Min.Z || b2.Max.Z < b.Min.Z)
}


// Split split a CUBE into sub-cubes
func (b *Box) Split() [8]*Box {
	center := b.GetCenter()
	return [8]*Box{
		NewBoxMinMax(center.X, center.Y, center.Z, b.Max.X, b.Max.Y, b.Max.Z),
		NewBoxMinMax(b.Min.X, center.Y, center.Z, center.X, b.Max.Y, b.Max.Z),
		NewBoxMinMax(b.Min.X, b.Min.Y, center.Z, center.X, center.Y, b.Max.Z),
		NewBoxMinMax(center.X, b.Min.Y, center.Z, b.Max.X, center.Y, b.Max.Z),

		NewBoxMinMax(center.X, center.Y, b.Min.Z, b.Max.X, b.Max.Y, center.Z),
		NewBoxMinMax(b.Min.X, center.Y, b.Min.Z, center.X, b.Max.Y, center.Z),
		NewBoxMinMax(b.Min.X, b.Min.Y, b.Min.Z, center.X, center.Y, center.Z),
		NewBoxMinMax(center.X, b.Min.Y, b.Min.Z, b.Max.X, center.Y, center.Z),
	}
	//return [8]*Box{
	//	NewBoxMinMax(b.Max.X, b.Max.Y, b.Max.Z, center.X, center.Y, center.Z),
	//	NewBoxMinMax(center.X, b.Max.Y, b.Max.Z, b.Min.X, center.Y, center.Z),
	//	NewBoxMinMax(center.X, center.Y, b.Max.Z, b.Min.X, b.Min.Y, center.Z),
	//	NewBoxMinMax(b.Max.X, center.Y, b.Max.Z, center.X, b.Min.Y, center.Z),
	//
	//	NewBoxMinMax(b.Max.X, b.Max.Y, center.Z, center.X, center.Y, b.Min.Z),
	//	NewBoxMinMax(center.X, b.Max.Y, center.Z, b.Min.X, center.Y, b.Min.Z),
	//	NewBoxMinMax(center.X, center.Y, center.Z, b.Min.X, b.Min.Y, b.Min.Z),
	//	NewBoxMinMax(b.Max.X, center.Y, center.Z, center.X, b.Min.Y, b.Min.Z),
	//}
}

// MinimumTranslation tells how much an entity has to move to no longer overlap another entity.
// FIXME ? 3D ? 2D ?
func MinimumTranslation(b, b2 Box) Vector3 {
	mtd := Vector3{}
	left := b.Min.X - b2.Max.X
	right := b.Max.X - b2.Min.X
	top := b.Min.Y - b2.Max.Y
	bottom := b.Max.Y - b2.Min.Y

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
		mtd.X = left
	} else {
		mtd.X = right
	}

	if math.Abs(top) < bottom {
		mtd.Y = top
	} else {
		mtd.Y = bottom
	}

	if math.Abs(mtd.X) < math.Abs(mtd.Y) {
		mtd.Y = 0
	} else {
		mtd.X = 0
	}

	return mtd
}

// String returns a human-readable representation of the box
//func (b *Box) String() string {
//	bm := b.GetMin()
//	bmm := b.GetMax()
//	return fmt.Sprintf("Center: %v, \nExtents: %v, \nmin %v, \nmax %v",
//		b.Center.String(), b.Extents.String(), bm.String(), bmm.String())
//}


