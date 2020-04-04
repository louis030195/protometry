package protometry

import (
	"math"
	"reflect"
	"testing"
)

func TestVectorN_Lerp(t *testing.T) {
	a := NewVectorN(0, 0, 0)
	b := NewVectorN(1, 1, 1)
	Equals(t, NewVectorN(.5, .5, .5), a.Lerp(b, 0.5))
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

func TestMorton3D(t *testing.T) {
	type args struct {
		v VectorN
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			args: args{v: *NewVectorN(12.0, 15.1, 1.786)},
			want: 1073741823,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Morton3D(tt.args.v)
			if got != tt.want {
				t.Errorf("Morton3D() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorN_Get(t *testing.T) {
	type fields struct {
		Dimensions           []float64
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	type args struct {
		dimension int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			fields: fields{Dimensions: []float64{0.2, 11, 12}},
			args:   args{dimension: 0},
			want:   0.2,
		},
		{
			fields: fields{Dimensions: []float64{0.2, 11, 12}},
			args:   args{dimension: -1},
			want:   math.MaxFloat64,
		},
		{
			fields: fields{Dimensions: []float64{0.2, 11, 12}},
			args:   args{dimension: 2},
			want:   12,
		},
		{
			fields: fields{Dimensions: []float64{0.2, 11, 12}},
			args:   args{dimension: 3},
			want:   math.MaxFloat64,
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
			got := a.Get(tt.args.dimension)
			if got != tt.want {
				t.Errorf("VectorN.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorN_Set(t *testing.T) {
	type fields struct {
		Dimensions           []float64
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	type args struct {
		dimension int
		value     float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields:  fields{Dimensions: []float64{0.2, 11, 12}},
			args:    args{dimension: 0, value: 12.2},
			wantErr: false,
		},
		{
			fields:  fields{Dimensions: []float64{0.2, 11, 12}},
			args:    args{dimension: -1, value: 12.2},
			wantErr: true,
		},
		{
			fields:  fields{Dimensions: []float64{0.2, 11, 12}},
			args:    args{dimension: 3, value: 12.2},
			wantErr: true,
		},
		{
			fields:  fields{Dimensions: []float64{0.2, 11, 12}},
			args:    args{dimension: 1, value: 12.2},
			wantErr: false,
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
			if err := a.Set(tt.args.dimension, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("VectorN.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	v := NewVectorN(1, 2, 3)
	v.Set(1, 12.2)
	Equals(t, 12.2, v.Get(1))
}

func TestConcatenate(t *testing.T) {
	type args struct {
		v []VectorN
	}
	tests := []struct {
		name string
		args args
		want VectorN
	}{
		{
			args: args{v: []VectorN{*NewVectorN(1, 2, 3), *NewVectorN(4, 5, 6)}},
			want: *NewVectorN(1, 2, 3, 4, 5, 6),
		},
		{
			args: args{v: []VectorN{*NewVectorN(4, 5, 6), *NewVectorN(1, 2, 3)}},
			want: *NewVectorN(4, 5, 6, 1, 2, 3),
		},
		{
			args: args{v: []VectorN{*NewVectorN(1, 2), *NewVectorN(5, 6)}},
			want: *NewVectorN(1, 2, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concatenate(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concatenate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		a VectorN
		b VectorN
	}
	tests := []struct {
		name string
		args args
		want VectorN
	}{
		// TODO: Add test cases.
		{
			args: args{*NewVectorN(1, 2, 3), *NewVectorN(4, 5, 6)},
			want: *NewVectorN(1, 2, 3),
		},
		{
			args: args{*NewVectorN(1, 2, 3), *NewVectorN(0, 5, 6)},
			want: *NewVectorN(0, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorN_Scale(t *testing.T) {
	type fields struct {
		Dimensions           []float64
		XXX_NoUnkeyedLiteral struct{}
		XXX_unrecognized     []byte
		XXX_sizecache        int32
	}
	type args struct {
		m float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *VectorN
	}{
		// TODO: Add test cases.
		{
			fields: fields{Dimensions: []float64{1, 2, 3}},
			args:   args{0.5},
			want:   NewVectorN(0.5, 1, 1.5),
		},
		{
			fields: fields{Dimensions: []float64{1, 2, 3, 7, 12.4, -27.8}},
			args:   args{0.5},
			want:   NewVectorN(0.5, 1, 1.5, 3.5, 6.2, -13.9),
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
			if got := a.Scale(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VectorN.Scale() = %v, want %v", got, tt.want)
			}
		})
	}
	a := NewVectorN(12, 4, 6)
	b := a.Scale(0.5)
	Equals(t, NewVectorN(12, 4, 6), a)
	Equals(t, NewVectorN(6, 2, 3), b)
}

func TestVectorN_Clone(t *testing.T) {
	a := NewVectorN(12, 4, 6)
	b := a.Clone()
	a.Dimensions[0] = 27
	Equals(t, 27., a.Dimensions[0])
	t.Logf("B: %v", b.ToString())
	Equals(t, 12., b.Dimensions[0])
}
