package code

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "Example 4",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Example 6",
			args: args{
				nums:   []int{1},
				target: 3,
			},
			want: 1,
		},
		{
			name: "Example 6",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "Example 7",
			args: args{
				nums:   []int{1, 3},
				target: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("nums = %v, target = %v, answer = %v, want %v", tt.args.nums, tt.args.target, got, tt.want)
			}
		})
	}
}