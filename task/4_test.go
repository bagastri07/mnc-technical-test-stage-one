package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canTakePersonalLeave(t *testing.T) {
	type args struct {
		joinDateStr       string
		jointLeaveDays    int
		leaveStartDateStr string
		leaveDuration     int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name: "Case 1",
			args: args{
				joinDateStr:       "2021-05-01",
				jointLeaveDays:    7,
				leaveStartDateStr: "2021-07-05",
				leaveDuration:     1,
			},
			want:  false,
			want1: "Karena belum 180 hari sejak tanggal join karyawan",
		},
		{
			name: "Case 2",
			args: args{
				joinDateStr:       "2021-05-01",
				jointLeaveDays:    7,
				leaveStartDateStr: "2021-11-05",
				leaveDuration:     3,
			},
			want:  false,
			want1: "Karena hanya boleh mengambil 1 hari cuti, tetapi mencoba mengambil 3 hari.",
		},
		{
			name: "Case 3",
			args: args{
				joinDateStr:       "2021-01-05",
				jointLeaveDays:    7,
				leaveStartDateStr: "2021-12-18",
				leaveDuration:     1,
			},
			want:  true,
			want1: "Cuti pribadi dapat diambil.",
		},
		{
			name: "Case 4",
			args: args{
				joinDateStr:       "2021-01-05",
				jointLeaveDays:    7,
				leaveStartDateStr: "2021-12-18",
				leaveDuration:     3,
			},
			want:  true,
			want1: "Cuti pribadi dapat diambil.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := canTakePersonalLeave(tt.args.joinDateStr, tt.args.jointLeaveDays, tt.args.leaveStartDateStr, tt.args.leaveDuration)
			assert.Equalf(t, tt.want, got, "canTakePersonalLeave(%v, %v, %v, %v)", tt.args.joinDateStr, tt.args.jointLeaveDays, tt.args.leaveStartDateStr, tt.args.leaveDuration)
			assert.Equalf(t, tt.want1, got1, "canTakePersonalLeave(%v, %v, %v, %v)", tt.args.joinDateStr, tt.args.jointLeaveDays, tt.args.leaveStartDateStr, tt.args.leaveDuration)
		})
	}
}
