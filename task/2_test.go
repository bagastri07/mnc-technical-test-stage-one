package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateChange(t *testing.T) {
	type args struct {
		totalBelanja int
		uangDibayar  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				totalBelanja: 700649,
				uangDibayar:  800000,
			},
			want: "Kembalian yang harus diberikan kasir: 99.351, dibulatkan menjadi 99.300\nPecahan uang:\n1 lembar 50.000\n2 lembar 20.000\n1 lembar 5.000\n2 lembar 2.000\n1 koin 200\n1 koin 100",
		},
		{
			name: "Case 2",
			args: args{
				totalBelanja: 575650,
				uangDibayar:  580000,
			},
			want: "Kembalian yang harus diberikan kasir: 4.350, dibulatkan menjadi 4.300\nPecahan uang:\n2 lembar 2.000\n1 koin 200\n1 koin 100",
		},
		{
			name: "Case 3",
			args: args{
				totalBelanja: 657650,
				uangDibayar:  600000,
			},
			want: "False, kurang bayar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calculateChange(tt.args.totalBelanja, tt.args.uangDibayar), "calculateChange(%v, %v)", tt.args.totalBelanja, tt.args.uangDibayar)
		})
	}
}
