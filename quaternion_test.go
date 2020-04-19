package protometry

import (
	"reflect"
	"testing"
)

func TestNewQuaternion(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
		w float64
	}
	tests := []struct {
		name string
		args args
		want *Quaternion
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuaternion(tt.args.x, tt.args.y, tt.args.z, tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuaternion() = %v, want %v", got, tt.want)
			}
		})
	}
}
