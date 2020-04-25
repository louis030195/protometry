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
	var vertices []*Vector3
	if centerBased {
		halfSide := sideLength / 2
		vertices = []*Vector3{
			NewVector3(-halfSide, -halfSide, -halfSide),
			NewVector3(halfSide, -halfSide, -halfSide),
			NewVector3(halfSide, halfSide, -halfSide),
			NewVector3(-halfSide, halfSide, -halfSide),

			NewVector3(-halfSide, halfSide, halfSide),
			NewVector3(halfSide, halfSide, halfSide),
			NewVector3(halfSide, -halfSide, halfSide),
			NewVector3(-halfSide, -halfSide, halfSide),
		}

	} else {
		vertices = []*Vector3{
			NewVector3(0, 0, 0),
			NewVector3(sideLength, 0, 0),
			NewVector3(sideLength, sideLength, 0),
			NewVector3(0, sideLength, 0),

			NewVector3(0, sideLength, sideLength),
			NewVector3(sideLength, sideLength, sideLength),
			NewVector3(sideLength, 0, sideLength),
			NewVector3(0, 0, sideLength),
		}
	}

	return &Mesh{Vertices: vertices, Tris: cuboidTris()}
}

func NewMeshRectangularCuboid(center, size Vector3) *Mesh {
	var vertices []*Vector3
	halfSize := size.Times(0.5)
	vertices = []*Vector3{
		NewVector3(-halfSize.X, -halfSize.Y, -halfSize.Z),
		NewVector3(halfSize.X, -halfSize.Y, -halfSize.Z),
		NewVector3(halfSize.X, halfSize.Y, -halfSize.Z),
		NewVector3(-halfSize.X, halfSize.Y, -halfSize.Z),

		NewVector3(-halfSize.X, halfSize.Y, halfSize.Z),
		NewVector3(halfSize.X, halfSize.Y, halfSize.Z),
		NewVector3(halfSize.X, -halfSize.Y, halfSize.Z),
		NewVector3(-halfSize.X, -halfSize.Y, halfSize.Z),
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
