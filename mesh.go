package protometry

func cuboidTris() []int32 {
	return []int32{
		0, 2, 1, //face front
		0, 3, 2,
		2, 3, 4, //face top
		2, 4, 5,
		1, 2, 5, //face right
		1, 5, 6,
		0, 7, 4, //face left
		0, 4, 3,
		5, 4, 7, //face back
		5, 7, 6,
		0, 6, 7, //face bottom
		0, 1, 6,
	}
}

// NewMeshSquareCuboid return a mesh forming a square cuboid
// Based on http://ilkinulas.github.io/development/unity/2016/04/30/cube-mesh-in-unity3d.html
func NewMeshSquareCuboid(sideLength float64, centerBased bool) *Mesh {
	var vertices []*VectorN
	if centerBased {
		halfSide := sideLength / 2
		vertices = []*VectorN{
			NewVectorN(-halfSide, -halfSide, -halfSide),
			NewVectorN(halfSide, -halfSide, -halfSide),
			NewVectorN(halfSide, halfSide, -halfSide),
			NewVectorN(-halfSide, halfSide, -halfSide),

			NewVectorN(-halfSide, halfSide, halfSide),
			NewVectorN(halfSide, halfSide, halfSide),
			NewVectorN(halfSide, -halfSide, halfSide),
			NewVectorN(-halfSide, -halfSide, halfSide),
		}

	} else {
		vertices = []*VectorN{
			NewVectorN(0, 0, 0),
			NewVectorN(sideLength, 0, 0),
			NewVectorN(sideLength, sideLength, 0),
			NewVectorN(0, sideLength, 0),

			NewVectorN(0, sideLength, sideLength),
			NewVectorN(sideLength, sideLength, sideLength),
			NewVectorN(sideLength, 0, sideLength),
			NewVectorN(0, 0, sideLength),
		}
	}

	return &Mesh{Vertices: vertices, Tris: cuboidTris()}
}

func NewMeshRectangularCuboid(size VectorN, centerBased bool) *Mesh {
	var vertices []*VectorN
	if centerBased {
		halfSize := size.Scale(0.5)
		vertices = []*VectorN{
			NewVectorN(-halfSize.Get(0), -halfSize.Get(1), -halfSize.Get(2)),
			NewVectorN(halfSize.Get(0), -halfSize.Get(1), -halfSize.Get(2)),
			NewVectorN(halfSize.Get(0), halfSize.Get(1), -halfSize.Get(2)),
			NewVectorN(-halfSize.Get(0), halfSize.Get(1), -halfSize.Get(2)),

			NewVectorN(-halfSize.Get(0), halfSize.Get(1), halfSize.Get(2)),
			NewVectorN(halfSize.Get(0), halfSize.Get(1), halfSize.Get(2)),
			NewVectorN(halfSize.Get(0), -halfSize.Get(1), halfSize.Get(2)),
			NewVectorN(-halfSize.Get(0), -halfSize.Get(1), halfSize.Get(2)),
		}
	} else {
		vertices = []*VectorN{
			NewVectorN(0, 0, 0),
			NewVectorN(size.Get(0), 0, 0),
			NewVectorN(size.Get(0), size.Get(1), 0),
			NewVectorN(0, size.Get(1), 0),

			NewVectorN(0, size.Get(1), size.Get(2)),
			NewVectorN(size.Get(0), size.Get(1), size.Get(2)),
			NewVectorN(size.Get(0), 0, size.Get(2)),
			NewVectorN(0, 0, size.Get(2)),
		}
	}

	return &Mesh{Vertices: vertices, Tris: cuboidTris()}
}

// Fit create a new mesh averaged on 2 meshes
func (m *Mesh) Fit(other Volume) bool {
	return false
}

// Intersects create a new mesh averaged on 2 meshes
func (m *Mesh) Intersects(other Volume) bool {
	return false
}

// Average create a new mesh averaged on 2 meshes
func (m *Mesh) Average(other Volume) Volume {
	return nil
}

// Mutate create a new mesh with random mutations
func (m *Mesh) Mutate(rate float64) Volume {
	return nil
}
