package utils

import "testing"

func TestRound(t *testing.T) {
	type args struct {
		num        float64
		precession uint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "UptoTwoDecimalPlaces",
			args: args{
				num:        3.142592653,
				precession: 2,
			},
			want: 3.14,
		},
		{
			name: "NoDecimalPlaces",
			args: args{
				num:        3.142592653,
				precession: 0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.num, tt.args.precession); got != tt.want {
				t.Errorf("TestRound: failed - TestCase=[%v] Want=[%v] Got=[%v]", tt.name, tt.want, got)
			}
		})
	}
}
