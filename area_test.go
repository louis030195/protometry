package protometry

import (
	"reflect"
	"testing"
)

func TestVectorN_In(t *testing.T) {
	a := NewBoxOfSize(*NewVectorN(0.5, 0.5, 0.5), 0.5)
	Equals(t, true, NewVectorN(1, 1, 1).In(*a))

	Equals(t, true, NewVectorN(0, 0, 0).In(*a))

	Equals(t, true, NewVectorN(1, 0, 0).In(*a))

	Equals(t, true, NewVectorN(0, 0, 1).In(*a))

	Equals(t, true, NewVectorN(0.5, 0.5, 0.5).In(*a))

	Equals(t, false, NewVectorN(-0.000001, 0.5, 0.5).In(*a))

	Equals(t, false, NewVectorN(0.5, -0.000001, 0.5).In(*a))

	Equals(t, false, NewVectorN(0.5, 0.5, -0.000001).In(*a))

	Equals(t, false, NewVectorN(0.5, 1.000001, 0.5).In(*a))

	Equals(t, false, NewVectorN(0.5, 0.5, 1.000001).In(*a))
}

func TestBox_Fit(t *testing.T) {
	a := NewBoxOfSize(*NewVectorN(0.5, 0.5, 0.5), 0.5)
	b := NewBoxOfSize(*NewVectorN(0.5, 0.5, 0.5), 0.5)

	// contains equal Box, symmetrically
	Equals(t, true, a.Fit(*b))

	Equals(t, true, a.Fit(*b))

	// contained on edge
	b = &Box{}
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(0.5, 1, 1))

	Equals(t, true, b.Fit(*a))

	Equals(t, false, a.Fit(*b))

	// contained away from edges
	b = &Box{}
	b.SetMinMax(*NewVectorN(0.1, 0.1, 0.1), *NewVectorN(0.9, 0.9, 0.9))
	Equals(t, true, b.Fit(*a))

	Equals(t, false, a.Fit(*b))

	// 1 corner Fit
	b = &Box{}
	b.SetMinMax(*NewVectorN(-0.1, -0.1, -0.1), *NewVectorN(0.9, 0.9, 0.9))
	Equals(t, false, b.Fit(*a))

	Equals(t, false, a.Fit(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0.9, 0.9, 0.9), *NewVectorN(1.1, 1.1, 1.1))
	Equals(t, false, b.Fit(*a))

	Equals(t, false, a.Fit(*b))

	// Other
	a = &Box{}
	a.SetMinMax(*NewVectorN(1, 1, 1), *NewVectorN(4, 4, 4))
	b = &Box{}
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	Equals(t, false, b.Fit(*a))
	Equals(t, false, a.Fit(*b))
	b = &Box{}
	b.SetMinMax(*NewVectorN(1, 1, 1), *NewVectorN(1, 1, 1))
	Equals(t, true, b.Fit(*a))
	b = &Box{}
	b.SetMinMax(*NewVectorN(1, 1, 1), *NewVectorN(4, 4, 4))
	Equals(t, true, b.Fit(*a))
}

func TestBox_Intersects(t *testing.T) {
	a := &Box{}
	a.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	b := &Box{}
	b.SetMinMax(*NewVectorN(1.1, 0, 0), *NewVectorN(2, 1, 1))

	// not intersecting area above or below in each dimension
	Equals(t, false, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(-1, 0, 0), *NewVectorN(-0.1, 1, 1))
	Equals(t, false, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0, 1.1, 0), *NewVectorN(1, 2, 1))
	Equals(t, false, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0, -1, 0), *NewVectorN(1, -0.1, 1))
	Equals(t, false, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0, 0, 1.1), *NewVectorN(1, 1, 2))
	Equals(t, false, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0, 0, -1), *NewVectorN(1, 1, -0.1))
	Equals(t, false, a.Intersects(*b))

	// intersects equal Box, symmetrically
	b = &Box{}
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	Equals(t, true, a.Intersects(*b))

	// intersects containing and contained
	b = &Box{}
	b.SetMinMax(*NewVectorN(0.1, 0.1, 0.1), *NewVectorN(0.9, 0.9, 0.9))
	Equals(t, true, a.Intersects(*b))

	// intersects partial containment on each corner
	b = &Box{}
	b.SetMinMax(*NewVectorN(0.9, 0.9, 0.9), *NewVectorN(2, 2, 2))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(-1, 0.9, 0.9), *NewVectorN(1, 2, 2))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0.9, -1, 0.9), *NewVectorN(2, 0.1, 2))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(-1, -1, 0.9), *NewVectorN(0.1, 0.1, 2))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0.9, 0.9, -1), *NewVectorN(2, 2, 0.1))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(-1, 0.9, -1), *NewVectorN(0.1, 2, 0.1))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0.9, -1, -1), *NewVectorN(2, 0.1, 0.1))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(-1, -1, -1), *NewVectorN(0.1, 0.1, 0.1))
	Equals(t, true, a.Intersects(*b))

	// intersects 'beam'; where no corners Fit
	// other but some contained
	b = &Box{}
	b.SetMinMax(*NewVectorN(-1, 0.1, 0.1), *NewVectorN(2, 0.9, 0.9))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0.1, -1, 0.1), *NewVectorN(0.9, 2, 0.9))
	Equals(t, true, a.Intersects(*b))

	b = &Box{}
	b.SetMinMax(*NewVectorN(0.1, 0.1, -1), *NewVectorN(0.9, 0.9, 2))
	Equals(t, true, a.Intersects(*b))

	// Other
	b = &Box{}
	b.SetMinMax(*NewVectorN(1, 1, 1), *NewVectorN(4, 4, 4))
	b = &Box{}
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	Equals(t, true, b.Intersects(*a))
	Equals(t, true, a.Intersects(*b))
	b = &Box{}
	b.SetMinMax(*NewVectorN(1, 1, 1), *NewVectorN(1, 1, 1))
	Equals(t, true, b.Intersects(*a))
	b = &Box{}
	b.SetMinMax(*NewVectorN(1, 1, 1), *NewVectorN(4, 4, 4))
	Equals(t, true, b.Intersects(*a))
}

func TestBox_Split(t *testing.T) {
	b := Box{}
	b.SetMinMax(*NewVector3Zero(), *NewVectorN(1, 1, 1))
	want := [8]*Box{
		NewBoxMinMax(0.5, 0.5, 0.5, 1, 1, 1),
		NewBoxMinMax(0, 0.5, 0.5, 0.5, 1, 1),
		NewBoxMinMax(0, 0, 0.5, 0.5, 0.5, 1),
		NewBoxMinMax(0.5, 0, 0.5, 1, 0.5, 1),
		NewBoxMinMax(0.5, 0.5, 0, 1, 1, 0.5),
		NewBoxMinMax(0, 0.5, 0, 0.5, 1, 0.5),
		NewBoxMinMax(0, 0, 0, 0.5, 0.5, 0.5),
		NewBoxMinMax(0.5, 0, 0, 1, 0.5, 0.5),
	}
	sub := b.Split()
	for i := range want {
		t.Logf("%v:%v", want[i].GetMin(), sub[i].GetMin())
		//Equals(t, true, want[i].Center.Equal(sub[i].Center))
		//Equals(t, true, want[i].Extents.Equal(sub[i].Extents))
	}
	// b = Box{min: *NewVectorN(-5, -5, -5), max: *NewVectorN(5, 5, 5)}
	// want = [8]*Box{
	// 	NewBox(0, 0, 0, 5, 5, 5),
	// 	NewBox(-5, 0, 0, 0, 5, 5),
	// 	NewBox(-5, -5, 0, 0, 0, 5),
	// 	NewBox(0, -5, 0, 5, 0, 5),
	// 	NewBox(0, 0, -5, 5, 5, 0),
	// 	NewBox(-5, 0, -5, 0, 5, 0),
	// 	NewBox(-5, -5, -5, 0, 0, 0),
	// 	NewBox(0, -5, -5, 5, 0, 0),
	// }
	// sub = b.MakeSubBoxes()
	// for i := range want {
	// 	Equals(t, true, want[i].min.Equal(sub[i].min))
	// 	Equals(t, true, want[i].max.Equal(sub[i].max))
	// }
}

func TestBox_SetMinMax(t *testing.T) {
	b := Box{}
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	min := b.GetMin()
	Equals(t, true, min.Equal(*NewVectorN(0, 0, 0)))
	max := b.GetMax()
	Equals(t, true, max.Equal(*NewVectorN(1, 1, 1)))
	Equals(t, *NewVectorN(0.5, 0.5, 0.5), b.Center)
	Equals(t, *NewVectorN(0.5, 0.5, 0.5), b.Extents)

	// X line
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(1, 0, 0))
	min = b.GetMin()
	Equals(t, true, min.Equal(*NewVectorN(0, 0, 0)))
	max = b.GetMax()
	Equals(t, true, max.Equal(*NewVectorN(1, 0, 0)))
	Equals(t, *NewVectorN(0.5, 0, 0), b.Center)
	Equals(t, *NewVectorN(0.5, 0, 0), b.Extents)

	// Y line
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(0, 1, 0))
	min = b.GetMin()
	Equals(t, true, min.Equal(*NewVectorN(0, 0, 0)))
	max = b.GetMax()
	Equals(t, true, max.Equal(*NewVectorN(0, 1, 0)))
	Equals(t, *NewVectorN(0, 0.5, 0), b.Center)
	Equals(t, *NewVectorN(0, 0.5, 0), b.Extents)

	// Z line
	b.SetMinMax(*NewVectorN(0, 0, 0), *NewVectorN(0, 0, 1))
	min = b.GetMin()
	Equals(t, true, min.Equal(*NewVectorN(0, 0, 0)))
	max = b.GetMax()
	Equals(t, true, max.Equal(*NewVectorN(0, 0, 1)))
	Equals(t, *NewVectorN(0, 0, 0.5), b.Center)
	Equals(t, *NewVectorN(0, 0, 0.5), b.Extents)

	// Reversed Z line
	b.SetMinMax(*NewVectorN(0, 0, 1), *NewVectorN(0, 0, 0))
	min = b.GetMin()
	Equals(t, true, min.Equal(*NewVectorN(0, 0, 1)))
	max = b.GetMax()
	Equals(t, true, max.Equal(*NewVectorN(0, 0, 0)))
	Equals(t, *NewVectorN(0, 0, 0.5), b.Center)
	Equals(t, *NewVectorN(0, 0, -0.5), b.Extents)
}

func TestNewBoxMinMax(t *testing.T) {
	type args struct {
		dims []float64
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxMinMax(tt.args.dims...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxMinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
