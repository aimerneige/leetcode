package code

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type want struct {
		count int
		nums  []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: want{
				count: 2,
				nums:  []int{1, 2},
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: want{
				count: 5,
				nums:  []int{0, 1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.args.nums
			count := removeDuplicates(nums)
			if count != tt.want.count {
				t.Errorf("Wrong count! %v, want %v", count, tt.want.count)
			}
			for i := 0; i < count; i++ {
				if nums[i] != tt.want.nums[i] {
					t.Errorf("Wrong slice! %v, want %v", nums, tt.want.nums)
				}
			}
		})
	}
}
