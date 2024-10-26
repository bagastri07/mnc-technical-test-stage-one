package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findFirstMatchingSet(t *testing.T) {
	type args struct {
		N           int
		stringsList []string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Case 1",
			args: args{
				N:           4,
				stringsList: []string{"abcd", "acbd", "aaab", "acbd"},
			},
			want: []int{2, 4},
		},
		{
			name: "Case 2",
			args: args{
				N:           11,
				stringsList: []string{"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"},
			},
			want: []int{3, 5, 10},
		},
		{
			name: "Case 3",
			args: args{
				N:           5,
				stringsList: []string{"pisang", "goreng", "enak", "sekali", "rasanya"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFirstMatchingSet(tt.args.N, tt.args.stringsList)
			assert.Equal(t, tt.want, got)
		})
	}
}
