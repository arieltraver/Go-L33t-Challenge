//O(n^2 * logn)
//only looks at unique trios

func compareSumToBest(sum int, target int, best int) int {
    if (sum == target) { return sum;                    
                       } else {
        diff := target - sum;
        olddiff := target - best;
        if (diff * diff < olddiff * olddiff) {
            best = sum;
        }
    }
    return best;
}


func threeSumClosest(nums []int, target int) int {
     
    l := len(nums)
    sortedNums := make([]int, l)
    copy(sortedNums, nums)
    sort.Ints(sortedNums)
    
    biggest := sortedNums[l-1] + sortedNums[l-2] + sortedNums[l-3]
    smallest := sortedNums[0] + sortedNums[1] + sortedNums[2]
    if(target >= biggest) { return biggest;
    } else if (target <= smallest) { return smallest; }
    best := smallest
    
    //avoid repeating the same trios
    for i, n := range(sortedNums[0:l-2]) {
        for j, m := range(sortedNums[i+1:l-1]) {
            partsum := n + m;
            slice := sortedNums[i+j+2:l];
            best = binaryPick(slice, partsum, target, best);
        }
    }
    return best;
}

//use binary search on the last section
func binaryPick(slice []int, partsum int, target int, best int) int {
    
    l := len(slice)
    if l < 1 { return best;}
    
    half := l / 2 //ensures we get a valid index
        
    sum := partsum + slice[half]
    best = compareSumToBest(sum, target, best)
    if (half == 0) {
        return best;
    } else if (sum > target) { //overshot
        return binaryPick(slice[0:half], partsum, target, best);
    } else { //undershot
        return binaryPick(slice[half:l], partsum, target, best);
    }
}
