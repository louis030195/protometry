package volume

import (
    "github.com/louis030195/protometry/api/vector3"
	"github.com/louis030195/protometry/internal/utils"
    "math/rand"
	"reflect"
	"testing"
)

func TestBox_Fit(t *testing.T) {
	a := NewBoxOfSize(0.5, 0.5, 0.5, 1)
	b := NewBoxOfSize(0.5, 0.5, 0.5, 1)

	// contains equal Box, symmetrically
	utils.Equals(t, true, a.Fit(*b))

	utils.Equals(t, true, b.Fit(*b))

	// contained on edge
	b = NewBoxMinMax(0, 0, 0, 0.5, 1, 1)

	utils.Equals(t, true, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))

	// contained away from edges
	b = NewBoxMinMax(0.1, 0.1, 0.1, 0.9, 0.9, 0.9)
	utils.Equals(t, true, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))

	// 1 corner Fit
	b = NewBoxMinMax(-0.1, -0.1, -0.1, 0.9, 0.9, 0.9)

	utils.Equals(t, false, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))

	b = NewBoxMinMax(0.9, 0.9, 0.9, 1.1, 1.1, 1.1)

	utils.Equals(t, false, b.Fit(*a))

	utils.Equals(t, false, a.Fit(*b))
}

func TestBox_Intersects(t *testing.T) {
	a := NewBoxMinMax(0, 0, 0, 1, 1, 1)

	b := NewBoxMinMax(1.1, 0, 0, 2, 1, 1)

	// not intersecting area above or below in each dimension
	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0, 0, -0.1, 1, 1)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 1.1, 0, 1, 2, 1)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, -1, 0, 1, -0.1, 1)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 0, 1.1, 1, 1, 2)

	utils.Equals(t, false, a.Intersects(*b))

	b = NewBoxMinMax(0, 0, -1, 1, 1, -0.1)

	utils.Equals(t, false, a.Intersects(*b))

	// intersects equal Box, symmetrically
	b = NewBoxMinMax(0, 0, 0, 1, 1, 1)

	utils.Equals(t, true, a.Intersects(*b))

	// intersects containing and contained
	b = NewBoxMinMax(0.1, 0.1, 0.1, 0.9, 0.9, 0.9)

	utils.Equals(t, true, a.Intersects(*b))

	// intersects partial containment on each corner
	b = NewBoxMinMax(0.9, 0.9, 0.9, 2, 2, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0.9, 0.9, 1, 2, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, -1, 0.9, 2, 0.1, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, -1, 0.9, 0.1, 0.1, 2)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, 0.9, -1, 2, 2, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, 0.9, -1, 0.1, 2, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.9, -1, -1, 2, 0.1, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(-1, -1, -1, 0.1, 0.1, 0.1)

	utils.Equals(t, true, a.Intersects(*b))

	// intersects 'beam'; where no corners Fit
	// other but some contained
	b = NewBoxMinMax(-1, 0.1, 0.1, 2, 0.9, 0.9)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.1, -1, 0.1, 0.9, 2, 0.9)

	utils.Equals(t, true, a.Intersects(*b))

	b = NewBoxMinMax(0.1, 0.1, -1, 0.9, 0.9, 2)

	utils.Equals(t, true, a.Intersects(*b))

	// Other
	b = NewBoxMinMax(1, 1, 1, 4, 4, 4)

	b = NewBoxMinMax(0, 0, 0, 1, 1, 1)

	utils.Equals(t, true, b.Intersects(*a))
	utils.Equals(t, true, a.Intersects(*b))
	b = NewBoxMinMax(1, 1, 1, 1, 1, 1)

	utils.Equals(t, true, b.Intersects(*a))
	b = NewBoxMinMax(1, 1, 1, 4, 4, 4)

	utils.Equals(t, true, b.Intersects(*a))
}

func TestBox_Contains(t *testing.T) {
	a := NewBoxOfSize(0.5, 0.5, 0.5, 1)
	utils.Equals(t, true, a.Contains(*vector3.NewVector3(1, 1, 1)))
	utils.Equals(t, true, a.Contains(*vector3.NewVector3(0, 0, 0)))
	utils.Equals(t, true, a.Contains(*vector3.NewVector3(1, 0, 0)))
	utils.Equals(t, true, a.Contains(*vector3.NewVector3(0, 0, 1)))
	utils.Equals(t, true, a.Contains(*vector3.NewVector3(0.5, 0.5, 0.5)))

	utils.Equals(t, false, a.Contains(*vector3.NewVector3(-0.000001, 0.5, 0.5)))
	utils.Equals(t, false, a.Contains(*vector3.NewVector3(0.5, -0.000001, 0.5)))
	utils.Equals(t, false, a.Contains(*vector3.NewVector3(0.5, 0.5, -0.000001)))
	utils.Equals(t, false, a.Contains(*vector3.NewVector3(0.5, 1.000001, 0.5)))
	utils.Equals(t, false, a.Contains(*vector3.NewVector3(0.5, 0.5, 1.000001)))
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
			centered at 0.5,0.5,0
			0.5,0.5,0 extent (square)
			--->
		 _ _ _ _
		|		|
		|_ _ _ _|

		------->
		1 size
		min 0,0,0
		max 1,1,0

		So in theory the split-ed box would be

		Each sub-box extent: 0.25,0.25,0
		 ->
		 _ _ _ _
		|_A_|_B_|
		|_D_|_C_|

		--->
		Each sub-box size: 0.5,0.5,0

		- - - - - - - - - - - - - - - - - - - - - - - - - - |
		B:						|		D:					|
		center: 0.25,0.75,0		|		center: 0.75,0.75,0 |
		min: 0,0.5,0			|		min: 0.5,0.5,0		|
		max: 0.5,1,0			|		max: 1,1,0			|
		- - - - - - - - - - - - - - - - - - - - - - - - - - |
		A:						|		C:					|
		center: 0.25,0.25,0		|		center: 0.75,0.25,0 |
		min: 0,0,0				|		min: 0.5,0,0		|
		max: 0.5,0.5,0			|		max: 1,0.5,0		|
		- - - - - - - - - - - - - - - - - - - - - - - - - - |
		...
	*/
	b := Box{
		Min: vector3.NewVector3(0, 0, 0),
		Max: vector3.NewVector3(1, 1, 0),
	}
	got := b.Split()

	/*
	 *    3____7
	 *  2/___6/|
	 *  | 1__|_5
	 *  0/___4/
	 */
	want := [8]*Box{
		{Min: vector3.NewVector3(0, 0, 0), Max: vector3.NewVector3(0.5, 0.5, 0)}, // A
		{Min: vector3.NewVector3(0, 0, 0), Max: vector3.NewVector3(0.5, 0.5, 0)}, // A
		{Min: vector3.NewVector3(0, 0.5, 0), Max: vector3.NewVector3(0.5, 1, 0)}, // B
		{Min: vector3.NewVector3(0, 0.5, 0), Max: vector3.NewVector3(0.5, 1, 0)}, // B

		{Min: vector3.NewVector3(0.5, 0, 0), Max: vector3.NewVector3(1, 0.5, 0)}, // C
		{Min: vector3.NewVector3(0.5, 0, 0), Max: vector3.NewVector3(1, 0.5, 0)}, // C
		{Min: vector3.NewVector3(0.5, 0.5, 0), Max: vector3.NewVector3(1, 1, 0)}, // D
		{Min: vector3.NewVector3(0.5, 0.5, 0), Max: vector3.NewVector3(1, 1, 0)}, // D
	}

	tester := func(got, want [8]*Box) {
		t.Logf("\nBefore split \n%v", b)
		for i := range want {
			t.Logf("\nMin: {%v}\nWant: {%v}\n\nMax: {%v}\nWant: {%v}",
				got[i].Min,
				want[i].Min,
				got[i].Max,
				want[i].Max,
			)
			utils.Equals(t, want[i].Min, got[i].Min)
			utils.Equals(t, want[i].Max, got[i].Max)
		}
	}
	tester(got, want)

	// 3D now
    b = Box{
        Min: vector3.NewVector3(0, 0, 0),
        Max: vector3.NewVector3(1, 1, 1),
    }
    got = b.Split()
    want = [8]*Box{
        {Min: vector3.NewVector3(0, 0, 0), Max: vector3.NewVector3(0.5, 0.5, 0.5)},
        {Min: vector3.NewVector3(0, 0, 0.5), Max: vector3.NewVector3(0.5, 0.5, 1)},
        {Min: vector3.NewVector3(0, 0.5, 0), Max: vector3.NewVector3(0.5, 1, 0.5)},
        {Min: vector3.NewVector3(0, 0.5, 0.5), Max: vector3.NewVector3(0.5, 1, 1)},

        {Min: vector3.NewVector3(0.5, 0, 0), Max: vector3.NewVector3(1, 0.5, 0.5)},
        {Min: vector3.NewVector3(0.5, 0, 0.5), Max: vector3.NewVector3(1, 0.5, 1)},
        {Min: vector3.NewVector3(0.5, 0.5, 0), Max: vector3.NewVector3(1, 1, 0.5)},
        {Min: vector3.NewVector3(0.5, 0.5, 0.5), Max: vector3.NewVector3(1, 1, 1)},
    }
    tester(got, want)
}

func bunchOfRandomBoxes(size int) []Box {
	var boxes []Box
	for i := 0; i < size; i++ {
		randomPos := vector3.RandomSpherePoint(*vector3.NewVector3Zero(), float64(size))
		boxes = append(boxes, *NewBoxOfSize(randomPos.X, randomPos.Y, randomPos.Z, float64(rand.Intn(size))))
	}
	return boxes
}
func bunchOfRandomVectors(size int) []vector3.Vector3 {
	var vectors []vector3.Vector3
	for i := 0; i < size; i++ {
		vectors = append(vectors, vector3.RandomSpherePoint(*vector3.NewVector3Zero(), float64(size)))
	}
	return vectors
}
func bunchOfRandomBoxesAndVectors(size int) ([]Box, []vector3.Vector3) {
	return bunchOfRandomBoxes(size), bunchOfRandomVectors(size)
}

func BenchmarkArea_NewBoxMinMax(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewBoxMinMax(i, i, i, i*2, i*2, i*2)
	}
}

func BenchmarkArea_NewBoxOfSize(b *testing.B) {
	size := float64(b.N)
	b.ResetTimer()
	for i := 0.; i < size; i++ {
		NewBoxOfSize(i, i, i, i)
	}
}

func BenchmarkArea_Contains(b *testing.B) {
	size := float64(b.N)
	boxes, vectors := bunchOfRandomBoxesAndVectors(int(size))
	b.ResetTimer()
	for i := range boxes {
		boxes[i].Contains(vectors[i])
	}
}

func BenchmarkArea_Fit(b *testing.B) {
	size := float64(b.N)
	boxesOne := bunchOfRandomBoxes(int(size))
	boxesTwo := bunchOfRandomBoxes(int(size))
	b.ResetTimer()
	for i := range boxesOne {
		boxesOne[i].Fit(boxesTwo[i])
	}
}

func BenchmarkArea_Intersects(b *testing.B) {
	size := float64(b.N)
	boxesOne := bunchOfRandomBoxes(int(size))
	boxesTwo := bunchOfRandomBoxes(int(size))
	b.ResetTimer()
	for i := range boxesOne {
		boxesOne[i].Intersects(boxesTwo[i])
	}
}

func BenchmarkArea_Split(b *testing.B) {
	size := float64(b.N)
	boxes := bunchOfRandomBoxes(int(size))
	b.ResetTimer()
	for i := range boxes {
		boxes[i].Split()
	}
}

/** Generated tests **/
func TestBox_Equal(t *testing.T) {
	type fields struct {
		Min *vector3.Vector3
		Max *vector3.Vector3
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
			fields: fields{
				Min: vector3.NewVector3(0, 0, 0),
				Max: vector3.NewVector3(1, 1, 1),
			},
			args: args{other: *NewBoxMinMax(0, 0, 0, 1, 1, 1)},
			want: true,
		},
		{
			fields: fields{
				Min: vector3.NewVector3(0, 0, 0),
				Max: vector3.NewVector3(1.1, 1, 1),
			},
			args: args{other: *NewBoxMinMax(0, 0, 0, 1, 1, 1)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := b.Equal(tt.args.other); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Fit1(t *testing.T) {
	type fields struct {
		Min *vector3.Vector3
		Max *vector3.Vector3
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
			fields: fields{
				Min: vector3.NewVector3(0.5, 0.5, 0.5),
				Max: vector3.NewVector3(0.6, 0.6, 0.6),
			},
			args: args{o: *NewBoxMinMax(0.5, 0.5, 0.5, 1, 1, 1)},
			want: true,
		},
		{
			fields: fields{
				Min: vector3.NewVector3(5.85, -3.9, 5.2),
				Max: vector3.NewVector3(6.85, -2.9, 6.2),
			},
			args: args{o: *NewBoxMinMax(2.92, -1.9, 2.59, 10.3, -4.4, 9.3)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := b.Fit(tt.args.o); got != tt.want {
				t.Errorf("Fit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetCenter(t *testing.T) {
	type fields struct {
		Min *vector3.Vector3
		Max *vector3.Vector3
	}
	tests := []struct {
		name   string
		fields fields
		want   vector3.Vector3
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				Min: vector3.NewVector3(0.5, 0.5, 0.5),
				Max: vector3.NewVector3(0.6, 0.6, 0.6),
			},
			want: *vector3.NewVector3(0.55, 0.55, 0.55),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := b.GetCenter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetSize(t *testing.T) {
	type fields struct {
		Min *vector3.Vector3
		Max *vector3.Vector3
	}
	tests := []struct {
		name   string
		fields fields
		want   vector3.Vector3
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				Min: vector3.NewVector3(0.5, 0.5, 0.5),
				Max: vector3.NewVector3(0.6, 0.6, 0.6),
			},
			want: *vector3.NewVector3(0.1, 0.1, 0.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := b.GetSize(); !got.Equal(tt.want) {
				t.Errorf("GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Intersects1(t *testing.T) {
	type fields struct {
		Min *vector3.Vector3
		Max *vector3.Vector3
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
			fields: fields{
				Min: vector3.NewVector3(0.5, 0.5, 0.5),
				Max: vector3.NewVector3(0.6, 0.6, 0.6),
			},
			args: args{b2: *NewBoxMinMax(0.5, 0.5, 0.5, 1, 1, 1)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := b.Intersects(tt.args.b2); got != tt.want {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
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
			args: args{0.5, 0.5, 0.5, 1, 1, 1},
			want: &Box{
				Min: vector3.NewVector3(0.5, 0.5, 0.5),
				Max: vector3.NewVector3(1, 1, 1),
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
			args: args{0, 0, 0, 1},
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

func TestBox_EncapsulateBox(t *testing.T) {
	type fields struct {
		Min                  *vector3.Vector3
		Max                  *vector3.Vector3
	}
	type args struct {
		o Box
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Box
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				Min: &vector3.Vector3{
					X: 0,
					Y: 0,
					Z: 0,
				},
				Max: &vector3.Vector3{
					X: 1,
					Y: 1,
					Z: 1,
				},
			},
			args: args{
				o: Box{
					Min: &vector3.Vector3{
						X: 1,
						Y: 1,
						Z: 1,
					},
					Max: &vector3.Vector3{
						X: 2,
						Y: 2,
						Z: 2,
					},
				},
			},
			want: NewBoxMinMax(0, 0, 0, 2, 2, 2),
		},
        {
            fields: fields{
                Min: &vector3.Vector3{
                    X: 1,
                    Y: 1,
                    Z: 1,
                },
                Max: &vector3.Vector3{
                    X: 2,
                    Y: 2,
                    Z: 2,
                },
            },
            args: args{
                o: Box{
                    Min: &vector3.Vector3{
                        X: 0,
                        Y: 0,
                        Z: 0,
                    },
                    Max: &vector3.Vector3{
                        X: 1,
                        Y: 1,
                        Z: 1,
                    },
                },
            },
            want: NewBoxMinMax(0, 0, 0, 2, 2, 2),
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
			}
			if got := b.EncapsulateBox(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncapsulateBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_EncapsulatePoint(t *testing.T) {
	type fields struct {
		Min                  *vector3.Vector3
		Max                  *vector3.Vector3
	}
	type args struct {
		o vector3.Vector3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Box
	}{
		// TODO: Add test cases.
		{
		    name: "Simple",
			fields: fields{
				Min: &vector3.Vector3{
					X: 0,
					Y: 0,
					Z: 0,
				},
				Max: &vector3.Vector3{
					X: 1,
					Y: 1,
					Z: 1,
				},
			},
			args: args{
				o: vector3.Vector3{
					X: 2,
					Y: 2,
					Z: 2,
				},
			},
			want: NewBoxMinMax(0, 0, 0, 2, 2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				Min:                  tt.fields.Min,
				Max:                  tt.fields.Max,
			}
			if got := b.EncapsulatePoint(tt.args.o); !got.Equal(*tt.want) {
				t.Errorf("EncapsulatePoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestBox_Intersection(t *testing.T) {
//    type fields struct {
//        Min                  *vector3.Vector3
//        Max                  *vector3.Vector3
//    }
//    type args struct {
//        o Box
//    }
//    tests := []struct {
//        name   string
//        fields fields
//        args   args
//        want   vector3.Vector3
//    }{
//        // TODO: Add test cases.
//        {
//           name: "Intersect on X, Y, Z of 0.1, 0.1, 0.1",
//           fields: fields{
//               Min:                  vector3.NewVector3(0, 0, 0),
//               Max:                  vector3.NewVector3(1, 1, 1),
//           },
//           args: args{o: Box{
//               Min:                  vector3.NewVector3(0.9, 0.9, 0.9),
//               Max:                  vector3.NewVector3(2, 2, 2),
//           }},
//           want: *vector3.NewVector3(0.1, 0.1, 0.1),
//        },
//        {
//           name: "Intersect on X, Y, Z of 0, 0, 0",
//           fields: fields{
//               Min:                  vector3.NewVector3(0, 0, 0),
//               Max:                  vector3.NewVector3(1, 1, 1),
//           },
//           args: args{o: Box{
//               Min:                  vector3.NewVector3(1, 1, 1),
//               Max:                  vector3.NewVector3(2, 2, 2),
//           }},
//           want: *vector3.NewVector3(0, 0, 0),
//        },
//        {
//            name: "Intersect on X of 13.54",
//            fields: fields{
//                Min:                  vector3.NewVector3(-12.54, 0, 0),
//                Max:                  vector3.NewVector3(1, 0, 0),
//            },
//            args: args{o: Box{
//                Min:                  vector3.NewVector3(-27, 1, 1),
//                Max:                  vector3.NewVector3(2, 2, 2),
//            }},
//            want: *vector3.NewVector3(13.54, 0, 0),
//        },
//    }
//    for _, tt := range tests {
//        t.Run(tt.name, func(t *testing.T) {
//            b := &Box{
//                Min:                  tt.fields.Min,
//                Max:                  tt.fields.Max,
//            }
//            if got := b.Intersection(tt.args.o); !got.Equal(tt.want) {
//                t.Errorf("Intersection() = %v, want %v", got, tt.want)
//            }
//        })
//    }
//}


func TestBox_SplitFour(t *testing.T) {
    type fields struct {
        Min                  *vector3.Vector3
        Max                  *vector3.Vector3
    }
    type args struct {
        vertical bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   [4]*Box
    }{
        // TODO: Add test cases.
        /*
         *    1____3
         *  1/___3/|
         *  | 0__|_2
         *  0/___2/
         */
        {
            name: "Horizontal",
            fields: fields{
                Min: vector3.NewVector3(0, 0, 0),
                Max: vector3.NewVector3(1, 1, 1),
            },
            args: args{false},
            want: [4]*Box{
                NewBoxMinMax(0, 0, 0, 0.5, 0.5, 1),
                NewBoxMinMax(0, 0.5, 0, 0.5, 1, 1),
                NewBoxMinMax(0.5, 0, 0, 1, 0.5, 1),
                NewBoxMinMax(0.5, 0.5, 0, 1, 1, 1),
            },
        },
        /*
         *    1____3
         *  0/___2/|
         *  | 1__|_3
         *  0/___2/
         */
        {
            name: "Vertical",
            fields: fields{
                Min: vector3.NewVector3(0, 0, 0),
                Max: vector3.NewVector3(1, 1, 1),
            },
            args: args{true},
            want: [4]*Box{
                NewBoxMinMax(0, 0, 0, 0.5, 1, 0.5),
                NewBoxMinMax(0, 0, 0.5, 0.5, 1, 1),
                NewBoxMinMax(0.5, 0, 0, 1, 1, 0.5),
                NewBoxMinMax(0.5, 0, 0.5, 1, 1, 1),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            b := &Box{
                Min:                  tt.fields.Min,
                Max:                  tt.fields.Max,
            }
            if got := b.SplitFour(tt.args.vertical); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SplitFour() = %v, want %v", got, tt.want)
            }
        })
    }
}
