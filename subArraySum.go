//idea: avoid too much addition by subtracting instead
//idea 2: avoid too much modulo by keeping track of the remainder
//idea 3: start with short subarrays first

//import ("fmt")
func checkSubarraySum(nums []int, k int) bool {
    
    l := len(nums)
    if l < 2 { return false; }
    modSums := make([][]int, l)
    for i := range(modSums) {
        modSums[i] = make([]int, 2)
        modSums[i][0] = (nums[i] % k)
        modSums[i][1] = (nums[i] % k)
    }
    
    for subLen := 1; subLen < l; subLen++ { //for all subarray lengths
        i := 0;
        for (i + subLen < l) { //for indices that can fit the length
            //fmt.Println(modSums);
            modSums[i][1] = modSums[i][0] + modSums[i + 1][1];
            //roll over the remainder
            if(modSums[i][1] >= k) { modSums[i][1] -= k; }
            if modSums[i][1] == 0 { return true; } //we found a subarray
            i++;
        }
    }
    return false;
}
