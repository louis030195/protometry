package protometry

import (
	"reflect"
	"testing"
)

func TestVector3_Lerp(t *testing.T) {
	a := NewVector3(0, 0, 0)
	b := NewVector3(1, 1, 1)
	Equals(t, NewVector3(.5, .5, .5), a.Lerp(b, 0.5))
}

func TestNewVector3Zero(t *testing.T) {
	type args struct {
		X, Y, Z float64
	}
	tests := []struct {
		name string
		args args
		want *Vector3
	}{
		{
			want: &Vector3{X: 0, Y: 0, Z: 0},
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
		X, Y, Z float64
	}
	tests := []struct {
		name string
		args args
		want *Vector3
	}{
		{
			want: &Vector3{X: 1, Y: 1, Z: 1},
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

func TestNewVector3(t *testing.T) {
	type args struct {
		X, Y, Z float64
	}
	tests := []struct {
		name string
		args args
		want *Vector3
	}{
		{
			args: args{
				X: 12, Y: 7, Z: 4,
			},
			want: &Vector3{X: 12, Y: 7, Z: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVector3(tt.args.X, tt.args.Y, tt.args.Z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMorton3D(t *testing.T) {
	type args struct {
		v Vector3
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			args: args{v: *NewVector3(12.0, 15.1, 1.786)},
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

func TestMin(t *testing.T) {
	type args struct {
		a Vector3
		b Vector3
	}
	tests := []struct {
		name string
		args args
		want Vector3
	}{
		// TODO: Add test cases.
		{
			args: args{*NewVector3(1, 2, 3), *NewVector3(4, 5, 6)},
			want: *NewVector3(1, 2, 3),
		},
		{
			args: args{*NewVector3(1, 2, 3), *NewVector3(0, 5, 6)},
			want: *NewVector3(0, 2, 3),
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

func TestVector3_In(t *testing.T) {
	a := NewBoxOfSize(0.5, 0.5, 0.5, 1)
	Equals(t, true, NewVector3(1, 1, 1).In(*a))

	Equals(t, true, NewVector3(0, 0, 0).In(*a))

	Equals(t, true, NewVector3(1, 0, 0).In(*a))

	Equals(t, true, NewVector3(0, 0, 1).In(*a))

	Equals(t, true, NewVector3(0.5, 0.5, 0.5).In(*a))

	Equals(t, false, NewVector3(-0.000001, 0.5, 0.5).In(*a))

	Equals(t, false, NewVector3(0.5, -0.000001, 0.5).In(*a))

	Equals(t, false, NewVector3(0.5, 0.5, -0.000001).In(*a))

	Equals(t, false, NewVector3(0.5, 1.000001, 0.5).In(*a))

	Equals(t, false, NewVector3(0.5, 0.5, 1.000001).In(*a))
}

func TestVector3_Clone(t *testing.T) {
	a := NewVector3(12, 4, 6)
	b := a.Clone()
	a.X = 27
	Equals(t, 27., a.X)
	t.Logf("B: %v", b)
	Equals(t, 12., b.X)
}

func TestVector3_Minus(t *testing.T) {
	a := NewVector3(12, 4, 6)
	b := NewVector3(-12, -4, -6)
	c := a.Minus(*b)
	// Check if properly NOT in-place
	Equals(t, true, a.Equal(*NewVector3(12, 4, 6)))
	Equals(t, true, b.Equal(*NewVector3(-12, -4, -6)))

	// And properly do the operation
	Equals(t, true, c.Equal(*NewVector3(24, 8, 12)))
}

func TestVector3_Scale(t *testing.T) {
	a := NewVector3(12, 4, 6)
	a.Scale(2)
	// Check if properly in-place
	Equals(t, true, a.Equal(*NewVector3(24, 8, 12)))
}

func TestVector3_Subtract(t *testing.T) {
	a := NewVector3(12, 4, 6)
	b := NewVector3(-12, -4, -6)
	a.Subtract(b)
	// Check if properly in-place
	Equals(t, true, a.Equal(*NewVector3(24, 8, 12)))
	Equals(t, true, b.Equal(*NewVector3(-12, -4, -6)))
}

func TestVector3_Times(t *testing.T) {
	a := NewVector3(12, 4, 6)
	b := NewVector3(-12, -4, -6)
	c := a.Times(2)
	// Check if properly NOT in-place
	Equals(t, true, a.Equal(*NewVector3(12, 4, 6)))
	Equals(t, true, b.Equal(*NewVector3(-12, -4, -6)))

	// And properly do the operation
	Equals(t, true, c.Equal(*NewVector3(24, 8, 12)))
}

func TestVector3_Add(t *testing.T) {
	a := NewVector3(12, 4, 6)
	b := NewVector3(-12, -4, -6)
	a.Add(b)
	// Check if properly in-place
	Equals(t, true, a.Equal(*NewVector3(0, 0, 0)))
	Equals(t, true, b.Equal(*NewVector3(-12, -4, -6)))
}

func TestVector3_Plus(t *testing.T) {
	a := NewVector3(12, 4, 6)
	b := NewVector3(-12, -4, -6)
	c := a.Plus(*b)
	// Check if properly NOT in-place
	Equals(t, true, a.Equal(*NewVector3(12, 4, 6)))
	Equals(t, true, b.Equal(*NewVector3(-12, -4, -6)))

	// And properly do the operation
	Equals(t, true, c.Equal(*NewVector3(0, 0, 0)))
}




func BenchmarkVector_Plus(b *testing.B) {
	var vectors []Vector3
	for i := 0; i < b.N; i++ {
		vectors = append(vectors, *NewVector3(0, 0, 0))
	}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		vectors[i-1].Plus(vectors[i])
	}
}

func BenchmarkVector_Add(b *testing.B) {
	var vectors []Vector3
	for i := 0; i < b.N; i++ {
		vectors = append(vectors, *NewVector3(0, 0, 0))
	}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		vectors[i-1].Add(&vectors[i])
	}
}

func BenchmarkVector_Minus(b *testing.B) {
	var vectors []Vector3
	for i := 0; i < b.N; i++ {
		vectors = append(vectors, *NewVector3(0, 0, 0))
	}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		vectors[i-1].Minus(vectors[i])
	}
}

func BenchmarkVector_Subtract(b *testing.B) {
	var vectors []Vector3
	for i := 0; i < b.N; i++ {
		vectors = append(vectors, *NewVector3(0, 0, 0))
	}
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		vectors[i-1].Subtract(&vectors[i])
	}
}

