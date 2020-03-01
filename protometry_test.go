package protometry

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestVectorN_In(t *testing.T) {
	a := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	in, err := NewVectorN(1, 1, 1).In(*a)
	equals(t, nil, err)
	if err == nil {
		equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0, 0, 0).In(*a); err == nil {
		equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(1, 0, 0).In(*a); err == nil {
		equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0, 0, 1).In(*a); err == nil {
		equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 0.5, 0.5).In(*a); err == nil {
		equals(t, true, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(-0.000001, 0.5, 0.5).In(*a); err == nil {
		equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, -0.000001, 0.5).In(*a); err == nil {
		equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 0.5, -0.000001).In(*a); err == nil {
		equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 1.000001, 0.5).In(*a); err == nil {
		equals(t, false, in)
	} else {
		t.Fatal(err)
	}
	if in, err := NewVectorN(0.5, 0.5, 1.000001).In(*a); err == nil {
		equals(t, false, in)
	} else {
		t.Fatal(err)
	}
}

func TestBox_Inside(t *testing.T) {
	a := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	b := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))

	// contains equal Box, symmetrically
	if inside, err := a.Inside(*b); err == nil {
		equals(t, true, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		equals(t, true, inside)
	} else {
		t.Fatal(err)
	}

	// contained on edge
	b = NewBox(*NewVectorN(0, 0, 0), *NewVectorN(0.5, 1, 1))

	if inside, err := b.Inside(*a); err == nil {
		equals(t, true, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		equals(t, false, inside)
	} else {
		t.Fatal(err)
	}

	// contained away from edges
	b = NewBox(*NewVectorN(0.1, 0.1, 0.1), *NewVectorN(0.9, 0.9, 0.9))
	if inside, err := b.Inside(*a); err == nil {
		equals(t, true, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		equals(t, false, inside)
	} else {
		t.Fatal(err)
	}

	// 1 corner inside
	b = NewBox(*NewVectorN(-0.1, -0.1, -0.1), *NewVectorN(0.9, 0.9, 0.9))
	if inside, err := b.Inside(*a); err == nil {
		equals(t, false, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		equals(t, false, inside)
	} else {
		t.Fatal(err)
	}

	b = NewBox(*NewVectorN(0.9, 0.9, 0.9), *NewVectorN(1.1, 1.1, 1.1))
	if inside, err := b.Inside(*a); err == nil {
		equals(t, false, inside)
	} else {
		t.Fatal(err)
	}
	if inside, err := a.Inside(*b); err == nil {
		equals(t, false, inside)
	} else {
		t.Fatal(err)
	}
}

func TestBox_Intersects(t *testing.T) {
	a := NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	b := NewBox(*NewVectorN(1.1, 0, 0), *NewVectorN(2, 1, 1))

	// not intersecting area above or below in each dimension
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, 0, 0), *NewVectorN(-0.1, 1, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, 1.1, 0), *NewVectorN(1, 2, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, -1, 0), *NewVectorN(1, -0.1, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, 0, 1.1), *NewVectorN(1, 1, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0, 0, -1), *NewVectorN(1, 1, -0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, false, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects equal Box, symmetrically
	b = NewBox(*NewVectorN(0, 0, 0), *NewVectorN(1, 1, 1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects containing and contained
	b = NewBox(*NewVectorN(0.1, 0.1, 0.1), *NewVectorN(0.9, 0.9, 0.9))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects partial containment on each corner
	b = NewBox(*NewVectorN(0.9, 0.9, 0.9), *NewVectorN(2, 2, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, 0.9, 0.9), *NewVectorN(1, 2, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.9, -1, 0.9), *NewVectorN(2, 0.1, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, -1, 0.9), *NewVectorN(0.1, 0.1, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.9, 0.9, -1), *NewVectorN(2, 2, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, 0.9, -1), *NewVectorN(0.1, 2, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.9, -1, -1), *NewVectorN(2, 0.1, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(-1, -1, -1), *NewVectorN(0.1, 0.1, 0.1))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}

	// intersects 'beam'; where no corners inside
	// other but some contained
	b = NewBox(*NewVectorN(-1, 0.1, 0.1), *NewVectorN(2, 0.9, 0.9))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.1, -1, 0.1), *NewVectorN(0.9, 2, 0.9))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
	b = NewBox(*NewVectorN(0.1, 0.1, -1), *NewVectorN(0.9, 0.9, 2))
	if intersects, err := a.Intersects(*b); err == nil {
		equals(t, true, intersects)
	} else {
		t.Fatal(err)
	}
}

func TestVectorN_Lerp(t *testing.T) {
	a := NewVectorN(0, 0, 0)
	b := NewVectorN(1, 1, 1)
	equals(t, NewVectorN(.5, .5, .5), a.Lerp(b, 0.5))
}

func TestNewVector3Zero(t *testing.T) {
	type args struct {
		dimensions []float64
	}
	tests := []struct {
		name string
		args args
		want *VectorN
	}{
		{
			want: &VectorN{Dimensions: []float64{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVector3Zero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector3Zero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVector3One(t *testing.T) {
	type args struct {
		dimensions []float64
	}
	tests := []struct {
		name string
		args args
		want *VectorN
	}{
		{
			want: &VectorN{Dimensions: []float64{1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVector3One(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector3One() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVectorN(t *testing.T) {
	type args struct {
		dimensions []float64
	}
	tests := []struct {
		name string
		args args
		want *VectorN
	}{
		{
			args: args{
				[]float64{12, 7, 4},
			},
			want: &VectorN{Dimensions: []float64{12, 7, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVectorN(tt.args.dimensions...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVectorN() = %v, want %v", got, tt.want)
			}
		})
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
	equals(t, NewVectorN(1, 1, 1), center)
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

func TestVectorN_Distance(t *testing.T) {
	type fields struct {
		Dimensions           []float64
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	type args struct {
		b VectorN
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			fields: fields{Dimensions: NewVector3Zero().Dimensions},
			args:   args{b: *NewVectorN(1, 0, 0)},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &VectorN{
				Dimensions:           tt.fields.Dimensions,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := a.Distance(tt.args.b); got != tt.want {
				t.Errorf("VectorN.Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorN_Dot(t *testing.T) {
	type fields struct {
		Dimensions           []float64
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	type args struct {
		b VectorN
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			fields: fields{Dimensions: NewVector3Zero().Dimensions},
			args:   args{b: *NewVector3One()},
			want:   0,
		},
		{
			fields: fields{Dimensions: NewVectorN(2, 2, 2).Dimensions},
			args:   args{b: *NewVectorN(4, 4, 4)},
			want:   24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &VectorN{
				Dimensions:           tt.fields.Dimensions,
				XXX_NoUnkeyedLiteral: tt.fields.XXX_NoUnkeyedLiteral,
				XXX_unrecognized:     tt.fields.XXX_unrecognized,
				XXX_sizecache:        tt.fields.XXX_sizecache,
			}
			if got := a.Dot(tt.args.b); got != tt.want {
				t.Errorf("VectorN.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}
