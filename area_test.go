package protometry

import (
	"reflect"
	"testing"
)

func TestVectorN_In(t *testing.T) {
	a := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	in, err := NewVectorN(1, 1, 1).In(*a)
	Equals(t, nil, err)
	if err == nil {
		Equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0, 0, 0).In(*a); err == nil {
		Equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(1, 0, 0).In(*a); err == nil {
		Equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0, 0, 1).In(*a); err == nil {
		Equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 0.5, 0.5).In(*a); err == nil {
		Equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(-0.000001, 0.5, 0.5).In(*a); err == nil {
		Equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, -0.000001, 0.5).In(*a); err == nil {
		Equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 0.5, -0.000001).In(*a); err == nil {
		Equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 1.000001, 0.5).In(*a); err == nil {
		Equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 0.5, 1.000001).In(*a); err == nil {
		Equals(t, false, in)
	} else {
		t.Fatal(err)
	}
}

func TestBox_Inside(t *testing.T) {
	a := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	b := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))

	// contains equal Box, symmetrically
	if inside, err := a.Inside(*b); err == nil {
		Equals(t, true, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		Equals(t, true, inside)
	} else {
		t.Fatal(err)
	}

	// contained on edge
	b = NewBox(*NewVectorN(0, 0, 0), *NewVectorN(0.5, 1, 1))

	if inside, err := b.Inside(*a); err == nil {
		Equals(t, true, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		Equals(t, false, inside)
	} else {
		t.Fatal(err)
	}

	// contained away from edges
	b = NewBox(*NewVectorN(0.1, 0.1, 0.1), *NewVectorN(0.9, 0.9, 0.9))
	if inside, err := b.Inside(*a); err == nil {
		Equals(t, true, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		Equals(t, false, inside)
	} else {
		t.Fatal(err)
	}

	// 1 corner inside
	b = NewBox(*NewVectorN(-0.1, -0.1, -0.1), *NewVectorN(0.9, 0.9, 0.9))
	if inside, err := b.Inside(*a); err == nil {
		Equals(t, false, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		Equals(t, false, inside)
	} else {
		t.Fatal(err)
	}

	b = NewBox(*NewVectorN(0.9, 0.9, 0.9), *NewVectorN(1.1, 1.1, 1.1))
	if inside, err := b.Inside(*a); err == nil {
		Equals(t, false, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		Equals(t, false, inside)
	} else {
		t.Fatal(err)
	}
}

func TestBox_Intersects(t *testing.T) {
	a := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	b := NewBox(*NewVectorN(1.1, 0, 0), *NewVectorN(2, 1, 1))

	// not intersecting area above or below in each dimension
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, 0, 0), *NewVectorN(-0.1, 1, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, 1.1, 0), *NewVectorN(1, 2, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, -1, 0), *NewVectorN(1, -0.1, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, 0, 1.1), *NewVectorN(1, 1, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, 0, -1), *NewVectorN(1, 1, -0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects equal Box, symmetrically
	b = NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects containing and contained
	b = NewBox(*NewVectorN(0.1, 0.1, 0.1), *NewVectorN(0.9, 0.9, 0.9))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects partial containment on each corner
	b = NewBox(*NewVectorN(0.9, 0.9, 0.9), *NewVectorN(2, 2, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, 0.9, 0.9), *NewVectorN(1, 2, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.9, -1, 0.9), *NewVectorN(2, 0.1, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, -1, 0.9), *NewVectorN(0.1, 0.1, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.9, 0.9, -1), *NewVectorN(2, 2, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, 0.9, -1), *NewVectorN(0.1, 2, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.9, -1, -1), *NewVectorN(2, 0.1, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, -1, -1), *NewVectorN(0.1, 0.1, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects 'beam'; where no corners inside
	// other but some contained
	b = NewBox(*NewVectorN(-1, 0.1, 0.1), *NewVectorN(2, 0.9, 0.9))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.1, -1, 0.1), *NewVectorN(0.9, 2, 0.9))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.1, 0.1, -1), *NewVectorN(0.9, 0.9, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		Equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
}

func TestNewBoxOfSize(t *testing.T) {
	type args struct {
		position VectorN
		size     float64
	}
	tests := []struct {
		name string
		args args
		want *Box
	}{
		{
			args: args{
				*NewVectorN(0, 0, 0),
				50.,
			},
			want: NewBox(*NewVectorN(-50, -50, -50), *NewVectorN(50, 50, 50)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoxOfSize(tt.args.position, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoxOfSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetCenter(t *testing.T) {
	center := NewBoxOfSize(*NewVector3One(), 5).GetCenter()
	Equals(t, NewVectorN(1, 1, 1), center)
}

func TestBox_GetSize(t *testing.T) {
	type fields struct {
		min VectorN
		max VectorN
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			fields: fields{min: *NewVector3Zero(), max: *NewVectorN(0, 0, 1)},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{
				min: tt.fields.min,
				max: tt.fields.max,
			}
			if got := b.GetSize(); got != tt.want {
				t.Errorf("Box.GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
