package Unit

import (
	"UnitTest/Unit"
	"testing"
)

func TestAddUpper(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{n : 10},45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unit.AddUpper(tt.args.n); got != tt.want {
				t.Errorf("AddUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}
