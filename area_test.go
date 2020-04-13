package protometry

import (
	"reflect"
	"testing"
)

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
		{
			args: args{[]float64{0, 0, 0, 1, 1, 1}},
			want: &Box{Center: *NewVectorN(0.5, 0.5, 0.5), Extents: *NewVectorN(0.5, 0.5, 0.5)},
		},
		{
			args: args{[]float64{0, 0, 0, 1, 0, 0}},
			want: &Box{Center: *NewVectorN(0.5, 0, 0), Extents: *NewVectorN(0.5, 0, 0)},
		},
		{
			args: args{[]float64{0, 0, 0, 0, 1, 0}},
			want: &Box{Center: *NewVectorN(0, 0.5, 0), Extents: *NewVectorN(0, 0.5, 0)},
		},
		{
			args: args{[]float64{0, 0, 0, 0, 0, 1}},
			want: &Box{Center: *NewVectorN(0, 0, 0.5), Extents: *NewVectorN(0, 0, 0.5)},
		},
		{ // Reversed
			args: args{[]float64{0, 0, 1, 0, 0, 0}},
			want: &Box{Center: *NewVectorN(0, 0, 0.5), Extents: *NewVectorN(0, 0, -0.5)},
		},
		{
			args: args{[]float64{0, 1, 0, 0, 0, 0}},
			want: &Box{Center: *NewVectorN(0, 0.5, 0), Extents: *NewVectorN(0, -0.5, 0)},
		},
		{
			args: args{[]float64{1, 0, 0, 0, 0, 0}},
			want: &Box{Center: *NewVectorN(0.5, 0, 0), Extents: *NewVectorN(-0.5, 0, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxMinMax(tt.args.dims...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxMinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoxOfSize(t *testing.T) {
	type args struct {
		center VectorN
		size   float64
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		// TODO: Add test cases.
		{
			args: args{center: *NewVectorN(0, 0, 0), size: 1},
			want: &Box{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(0.5, 0.5, 0.5)},
		},
		{
			args: args{center: *NewVectorN(0, 12, 23.2), size: 1.3},
			want: &Box{Center: *NewVectorN(0, 12, 23.2), Extents: *NewVectorN(0.65, 0.65, 0.65)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxOfSize(tt.args.center, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxOfSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Equal(t *testing.T) {
	type fields struct {
		Center  VectorN
		Extents VectorN
	}
	type args struct {
		other Box
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(0.5, 0.5, 0.5)},
			args:   args{other: *NewBoxOfSize(*NewVectorN(0, 0, 0), 1)},
			want:   true,
		},
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 1, 1)},
			args:   args{other: *NewBoxMinMax(-1, -1, -1, 1, 1, 1)},
			want:   true,
		},
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 1, 1)},
			args:   args{other: Box{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 1, 1)}},
			want:   true,
		},
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 1, 1)},
			args:   args{other: Box{Center: *NewVectorN(0.1, 0, 0), Extents: *NewVectorN(1, 1, 1)}},
			want:   false,
		},
		{
			fields: fields{Center: *NewVectorN(0, 92, 0), Extents: *NewVectorN(1, 1, 1)},
			args:   args{other: Box{Center: *NewVectorN(0, 92, 0), Extents: *NewVectorN(1, 1, 1)}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Center:  tt.fields.Center,
				Extents: tt.fields.Extents,
			}
			if got := b.Equal(tt.args.other); got != tt.want {
				t.Errorf("Box.Equal() = %v, want %v\nb: %v\nother:% v", got, tt.want, b, tt.args.other)
			}
		})
	}
}

func TestBox_GetMin(t *testing.T) {
	type fields struct {
		Center  VectorN
		Extents VectorN
	}
	tests := []struct {
		name   string
		fields fields
		want   VectorN
	}{
		// TODO: Add test cases.
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 1, 1)},
			want:   *NewVectorN(-1, -1, -1),
		},
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 2.5, 1)},
			want:   *NewVectorN(-1, -2.5, -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Center:  tt.fields.Center,
				Extents: tt.fields.Extents,
			}
			if got := b.GetMin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Box.GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetMax(t *testing.T) {
	type fields struct {
		Center  VectorN
		Extents VectorN
	}
	tests := []struct {
		name   string
		fields fields
		want   VectorN
	}{
		// TODO: Add test cases.
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 1, 1)},
			want:   *NewVectorN(1, 1, 1),
		},
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 2.5, 1)},
			want:   *NewVectorN(1, 2.5, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Center:  tt.fields.Center,
				Extents: tt.fields.Extents,
			}
			if got := b.GetMax(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Box.GetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetSize(t *testing.T) {
	type fields struct {
		Center  VectorN
		Extents VectorN
	}
	tests := []struct {
		name   string
		fields fields
		want   VectorN
	}{
		// TODO: Add test cases.
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(1, 1, 1)},
			want:   *NewVectorN(2, 2, 2),
		},
		{
			fields: fields{Center: *NewVectorN(0, 0, 0), Extents: *NewVectorN(0, 12, 1)},
			want:   *NewVectorN(0, 24, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Center:  tt.fields.Center,
				Extents: tt.fields.Extents,
			}
			if got := b.GetSize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Box.GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
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

func TestBox_EncapsulatePoint(t *testing.T) {
	b := NewBoxOfSize(*NewVector3Zero(), 1)
	b.EncapsulatePoint(*NewVectorN(2, 2, 2))
	Equals(t, NewBoxOfSize(*NewVectorN(0.75, 0.75, 0.75), 2.5), b)
}

func TestBox_EncapsulateBox(t *testing.T) {
	b := NewBoxOfSize(*NewVector3Zero(), 1)
	bb := NewBoxOfSize(*NewVectorN(10, 10, 10), 1)
	b.EncapsulateBox(*bb)
	Equals(t, NewBoxMinMax(-0.5, -0.5, -0.5, 10.5, 10.5, 10.5), b)
	Equals(t, NewBoxOfSize(*NewVectorN(5, 5, 5), 11), b)
}

func TestBox_Expand(t *testing.T) {
	b := NewBoxOfSize(*NewVector3Zero(), 1)
	b.Expand(3)
	Equals(t, NewBoxOfSize(*NewVector3Zero(), 4), b) // FIXME: is it correct ?
}

func TestBox_ExpandV(t *testing.T) {
	b := NewBoxOfSize(*NewVector3Zero(), 1)
	b.ExpandV(*NewVector3One())
	Equals(t, NewBoxOfSize(*NewVector3Zero(), 2), b) // FIXME: is it correct ?
}

func TestVectorN_In(t *testing.T) {
	a := NewBoxOfSize(*NewVectorN(0.5, 0.5, 0.5), 1)
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
	a := NewBoxOfSize(*NewVectorN(0.5, 0.5, 0.5), 1)
	b := NewBoxOfSize(*NewVectorN(0.5, 0.5, 0.5), 1)

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
	/*
		What we want to achieve (in 2D):
		 _ _ _ _
		|		|
		|_ _ _ _|

		->
		 _ _ _ _
		|_ _|_ _|
		|_ _|_ _|

	 */

	/*
		Representation (ignoring the Z axis cause it's hard to draw in 3D ;))
			centered at 0.5,0.5,0.5
			0.5,0.5,0.5 extent (cubic)
			--->
		 _ _ _ _
		|		|
		|_ _ _ _|

		------->
		1 size
		min 0,0,0
		max 0,0,0

		So in theory the split-ed box would be

		Each sub-box extent: 0.25
		 ->
		 _ _ _ _
		|_A_|_B_|
		|_D_|_C_|

		--->
		Each sub-box size: 0.5
		A sub-box:
		center: 0.25,0.75,0.25
		min: 0,0.5,0
		max: 0.5,1,0.5
		...
	 */
	b := Box{
		Center: VectorN{
			Dimensions: []float64{0.5, 0.5, 0.5},
		},
		Extents: VectorN{
			Dimensions:[]float64{0.5, 0.5, 0.5},
		},
	}
	got := b.Split()

	childExtents := *b.Extents.Scale(0.5)
	want := [8]*Box{
		{Center: VectorN{Dimensions: []float64{0.25, 0.75, 0.25}}, Extents: childExtents},
		{Center: VectorN{Dimensions: []float64{0.75, 0.75, 0.25}}, Extents:childExtents},
		{Center: VectorN{Dimensions: []float64{0.25, 0.75, 0.75}}, Extents: childExtents},
		{Center: VectorN{Dimensions: []float64{0.75, 0.75, 0.75}}, Extents: childExtents},

		{Center: VectorN{Dimensions: []float64{0.25, 0.25, 0.25}}, Extents: childExtents},
		{Center: VectorN{Dimensions: []float64{0.75, 0.25, 0.25}}, Extents: childExtents},
		{Center: VectorN{Dimensions: []float64{0.25, 0.25, 0.75}}, Extents: childExtents},
		{Center: VectorN{Dimensions: []float64{0.75, 0.25, 0.75}}, Extents: childExtents},
	}

	tester := func(got, want [8]*Box) {
		t.Logf("\nBefore split \n%v", b)
		for i := range want {
			t.Logf("\nMin%v\nWant%v\n\nMax%v\nWant%v\n\nCenter%v\nWant%v\n\nExtents%v\nWant%v\n",
				got[i].GetMin(),
				want[i].GetMin(),
				got[i].GetMax(),
				want[i].GetMax(),
				got[i].Center,
				want[i].Center,
				got[i].Extents,
				want[i].Extents,
			)
			Equals(t, want[i].Center, got[i].Center)
			Equals(t, want[i].Extents, got[i].Extents)
		}
	}
	tester(got, want)
}

func BenchmarkArea_NewBoxMinMax(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxMinMax(i, i, i, i*2, i*2, i*2)
	}
	b.StopTimer()
}

func BenchmarkArea_NewBoxOfSize(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(*NewVectorN(i, i, i), i)
	}
	b.StopTimer()
}

func BenchmarkArea_In(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewVectorN(i, i, i).In(*NewBoxOfSize(*NewVectorN(i, i, i), 1))
	}
	b.StopTimer()
}

func BenchmarkArea_Fit(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(*NewVectorN(i, i, i), 0.1).Fit(*NewBoxOfSize(*NewVectorN(i, i, i), 1))
	}
	b.StopTimer()
}

func BenchmarkArea_Intersects(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(*NewVectorN(i, i, i), 1).Intersects(*NewBoxOfSize(*NewVectorN(i, i, i), 1))
	}
	b.StopTimer()
}

func BenchmarkArea_Split(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(*NewVectorN(i, i, i), 1).Split()
	}
	b.StopTimer()
}
