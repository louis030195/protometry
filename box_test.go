package protometry

import (
	"reflect"
	"testing"
)


func TestBox_Fit(t *testing.T) {
	a := NewBoxOfSize(0.5, 0.5, 0.5, 1)
	b := NewBoxOfSize(0.5, 0.5, 0.5, 1)

	// contains equal Box, symmetrically
	Equals(t, true, a.Fit(*b))

	Equals(t, true, b.Fit(*b))

	// contained on edge
	b = NewBoxMinMax(0, 0, 0, 0.5, 1, 1)

	Equals(t, true, b.Fit(*a))

	Equals(t, false, a.Fit(*b))

	// contained away from edges
	b = NewBoxMinMax(0.1, 0.1, 0.1, 0.9, 0.9, 0.9)
	Equals(t, true, b.Fit(*a))

	Equals(t, false, a.Fit(*b))

	// 1 corner Fit
	b = NewBoxMinMax(-0.1, -0.1, -0.1, 0.9, 0.9, 0.9)

	Equals(t, false, b.Fit(*a))

	Equals(t, false, a.Fit(*b))

	b = NewBoxMinMax(0.9, 0.9, 0.9, 1.1, 1.1, 1.1)

	Equals(t, false, b.Fit(*a))

	Equals(t, false, a.Fit(*b))
}

func TestBox_Intersects(t *testing.T) {
	a := NewBoxMinMax(0, 0, 0, 1, 1, 1)

	b := NewBoxMinMax(1.1, 0, 0, 2, 1, 1)


	// not intersecting area above or below in each dimension
	Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0, 0, -0.1, 1, 1)

	Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 1.1, 0, 1, 2, 1)

	Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, -1, 0, 1, -0.1, 1)

	Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 0, 1.1, 1, 1, 2)

	Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 0, -1, 1, 1, -0.1)

	Equals(t, false, a.Intersects(*b))

	// intersects equal Box, symmetrically
	b = NewBoxMinMax(0, 0, 0, 1, 1, 1)

	Equals(t, true, a.Intersects(*b))

	// intersects containing and contained
	b = NewBoxMinMax(0.1, 0.1, 0.1, 0.9, 0.9, 0.9)

	Equals(t, true, a.Intersects(*b))

	// intersects partial containment on each corner
	b = NewBoxMinMax(0.9, 0.9, 0.9, 2, 2, 2)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0.9, 0.9, 1, 2, 2)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, -1, 0.9, 2, 0.1, 2)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, -1, 0.9, 0.1, 0.1, 2)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, 0.9, -1, 2, 2, 0.1)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0.9, -1, 0.1, 2, 0.1)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, -1, -1, 2, 0.1, 0.1)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, -1, -1, 0.1, 0.1, 0.1)

	Equals(t, true, a.Intersects(*b))

	// intersects 'beam'; where no corners Fit
	// other but some contained
	b = NewBoxMinMax(-1, 0.1, 0.1, 2, 0.9, 0.9)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.1, -1, 0.1, 0.9, 2, 0.9)

	Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.1, 0.1, -1, 0.9, 0.9, 2)

	Equals(t, true, a.Intersects(*b))

	// Other
	b = NewBoxMinMax(1, 1, 1, 4, 4, 4)

	b = NewBoxMinMax(0, 0, 0, 1, 1, 1)

	Equals(t, true, b.Intersects(*a))
	Equals(t, true, a.Intersects(*b))
	b = NewBoxMinMax(1, 1, 1, 1, 1, 1)

	Equals(t, true, b.Intersects(*a))
	b = NewBoxMinMax(1, 1, 1, 4, 4, 4)

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
	//b := Box{
	//	Center: &Vector3{
	//		Dimensions: []float64{0.5, 0.5, 0.5},
	//	},
	//	Extents: &Vector3{
	//		Dimensions:[]float64{0.5, 0.5, 0.5},
	//	},
	//}
	//got := b.Split()
	//
	//childExtents := b.Extents.Scale(0.5)
	//want := [8]*Box{
	//	{Center: &Vector3{Dimensions: []float64{0.25, 0.75, 0.25}}, Extents: childExtents},
	//	{Center: &Vector3{Dimensions: []float64{0.75, 0.75, 0.25}}, Extents: childExtents},
	//	{Center: &Vector3{Dimensions: []float64{0.25, 0.75, 0.75}}, Extents: childExtents},
	//	{Center: &Vector3{Dimensions: []float64{0.75, 0.75, 0.75}}, Extents: childExtents},
	//
	//	{Center: &Vector3{Dimensions: []float64{0.25, 0.25, 0.25}}, Extents: childExtents},
	//	{Center: &Vector3{Dimensions: []float64{0.75, 0.25, 0.25}}, Extents: childExtents},
	//	{Center: &Vector3{Dimensions: []float64{0.25, 0.25, 0.75}}, Extents: childExtents},
	//	{Center: &Vector3{Dimensions: []float64{0.75, 0.25, 0.75}}, Extents: childExtents},
	//}
	//
	//tester := func(got, want [8]*Box) {
	//	t.Logf("\nBefore split \n%v", b)
	//	for i := range want {
	//		t.Logf("\nMin%v\nWant%v\n\nMax%v\nWant%v\n\nCenter%v\nWant%v\n\nExtents%v\nWant%v\n",
	//			got[i].GetMin(),
	//			want[i].GetMin(),
	//			got[i].GetMax(),
	//			want[i].GetMax(),
	//			got[i].Center,
	//			want[i].Center,
	//			got[i].Extents,
	//			want[i].Extents,
	//		)
	//		Equals(t, want[i].Center, got[i].Center)
	//		Equals(t, want[i].Extents, got[i].Extents)
	//	}
	//}
	//tester(got, want)
}

// FIXME fix these trash benchmarks
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
		NewBoxOfSize(i, i, i, i)
	}
	b.StopTimer()
}

func BenchmarkArea_In(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewVector3(i, i, i).In(*NewBoxOfSize(i, i, i, 1))
	}
	b.StopTimer()
}

func BenchmarkArea_Fit(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, 0.1).Fit(*NewBoxOfSize(i, i, i, 1))
	}
	b.StopTimer()
}

func BenchmarkArea_Intersects(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, 1).Intersects(*NewBoxOfSize(i, i, i, 1))
	}
	b.StopTimer()
}

func BenchmarkArea_Split(b *testing.B) {
	size := float64(b.N)
	b.StartTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, 1).Split()
	}
	b.StopTimer()
}





























/** Generated tests **/
func TestBox_Equal(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
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
			fields:fields{
				Min:                  NewVector3(0, 0, 0),
				Max:                  NewVector3(1, 1, 1),
			},
			args:args{other: *NewBoxMinMax(0, 0, 0, 1, 1, 1)},
			want: true,
		},
		{
			fields:fields{
				Min:                  NewVector3(0, 0, 0),
				Max:                  NewVector3(1.1, 1, 1),
			},
			args:args{other: *NewBoxMinMax(0, 0, 0, 1, 1, 1)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := b.Equal(tt.args.other); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Fit1(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	type args struct {
		o Box
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min:                  NewVector3(0.5, 0.5, 0.5),
				Max:                  NewVector3(0.6, 0.6, 0.6),
			},
			args:args{o: *NewBoxMinMax(0.5, 0.5, 0.5, 1, 1, 1)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := b.Fit(tt.args.o); got != tt.want {
				t.Errorf("Fit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetCenter(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector3
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min:                  NewVector3(0.5, 0.5, 0.5),
				Max:                  NewVector3(0.6, 0.6, 0.6),
			},
			want: *NewVector3(0.55, 0.55, 0.55),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := b.GetCenter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetSize(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector3
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min:                  NewVector3(0.5, 0.5, 0.5),
				Max:                  NewVector3(0.6, 0.6, 0.6),
			},
			want: *NewVector3(0.1, 0.1, 0.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := b.GetSize(); !got.Equal(tt.want) {
				t.Errorf("GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Intersects1(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	type args struct {
		b2 Box
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min:                  NewVector3(0.5, 0.5, 0.5),
				Max:                  NewVector3(0.6, 0.6, 0.6),
			},
			args:args{b2: *NewBoxMinMax(0.5, 0.5, 0.5, 1, 1, 1)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := b.Intersects(tt.args.b2); got != tt.want {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Split1(t *testing.T) {
	type fields struct {
		Min                  *Vector3
		Max                  *Vector3
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	tests := []struct {
		name   string
		fields fields
		want   [8]*Box
	}{
		// TODO: Add test cases.
		{
			fields:fields{
				Min:                  NewVector3(0, 0, 0),
				Max:                  NewVector3(1, 1, 1),
			},
			want: [8]*Box{
				// TODO
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := b.Split(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoxMinMax(t *testing.T) {
	type args struct {
		minX float64
		minY float64
		minZ float64
		maxX float64
		maxY float64
		maxZ float64
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		{
			args:args{0.5, 0.5, 0.5, 1, 1, 1},
			want: &Box{
				Min: NewVector3(0.5, 0.5, 0.5),
				Max: NewVector3(1, 1, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxMinMax(tt.args.minX, tt.args.minY, tt.args.minZ, tt.args.maxX, tt.args.maxY, tt.args.maxZ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxMinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoxOfSize(t *testing.T) {
	type args struct {
		x    float64
		y    float64
		z    float64
		size float64
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		// TODO: Add test cases.
		{
			args:args{0, 0, 0, 1},
			want: NewBoxMinMax(-0.5, -0.5, -0.5, 0.5, 0.5, 0.5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxOfSize(tt.args.x, tt.args.y, tt.args.z, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxOfSize() = %v, want %v", got, tt.want)
			}
		})
	}
}