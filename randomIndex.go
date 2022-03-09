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
    counts = make([]int, range) //covers all numbers
    for _, num := range nums {
        counts[num] += 1 //tracks the number of occurences
    } //if the number doesn't occur its count is zero
    
    for index, _ := range indices {
        indices[index] = make([]int, counts[index] + 1)
        //each list is the length of #occurrences
    }
    
    for indx, num := range nums {
        indices[num][0] ++ //where to place next found index
        indices[[indices[num][0]]] = indx //store location of num
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
