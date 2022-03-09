const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1


type Solution struct {
    nums []int
    indices &int  //slicie
}

func max(nums []int) int {
    
    max = MinInt
    for i := range nums {
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}

func min(nums[]int) int {
    min = MaxInt
    for i := range nums{
        if nums[i] < min {
            min = nums[i]
        }
    }
    return min
}


func Constructor(nums []int) Solution {
    //init a Solution object with arr = nums
    
    indices := make([][]int, len(nums))
    // array count
    range = max(nums) - min(nums)
    counts = make([]int, range) //goes across all numbers
    for i := range nums {
        counts[num] += 1
    }
    
    for i := range indices {
        indices[i] = make([]int, counts[i])
    }
    s = Solution{nums: nums, indices:indices}
    return &s
}


func (this *Solution) Pick(target int) int {
    
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
