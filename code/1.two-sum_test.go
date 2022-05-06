package code

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Example 3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			if len(got) != 2 {
				t.Errorf("Wrong format! twoSum() = %v, want %v", got, tt.want)
				return
			}
			if got[0] > got[1] {
				got[1], got[0] = got[0], got[1]
			}
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("Wrong result! twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
