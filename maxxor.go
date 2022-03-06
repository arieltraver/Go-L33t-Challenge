

//time consumption: nlogn, space consumption: nlogn
//idea: go digit by digit, and eliminate part of the possible results each time by considering only numbers with helpful digits, halving the options each time (logn)
//do this for all numbers with a 1 in the most significant place (n)


/**takes an array, a desired digit (1 or 0), and the place.
   returns an array with only those numbers which have the desired digit in the right place.**/
   func sliceArray(nums []int, digit int, sigplace int) {
    newArr = make([]int, len(nums))
    count := 0
    xorMask = math.Pow(2, sigplace) //use to isolate digit
    //loop to find the most significant 1 digit
    for i:= 0; i< len(nums); i++ {
        if (nums[i] ^ xorMask) ^ digit == 0 {
            newArr[count] = nums[i];
            count++
        }
    }
    return newArr[0:count]; //a smaller array
}

func findMaximumXOR(nums []int) int {
    if len(nums) == 1 {return nums[0]}
    else if len(nums) < 2 {return nums[0] ^ nums[1]}
    else {
    
        maxXor := -99999999 //replace with actual min int
        temp := 0
        sigplace := 0
        //loop to find the most significant digit
        for numb := range nums {
            temp = log(numb)/log(2) //log2 of numb
            if temp > sigPlace {
                sigPlace = temp
            }
        }
    
        //loop where we narrow down our options each time
        for num := range nums{ //for each starting w 1
            if num & (math.pow(2, sigPlace)) {
                zerosNums = sliceArray(nums, 0, sigPlace)
                for dig := sigPlace-1; dig >=0; dig-- { //for each digit in the number we're looking at
                if len(zerosNums) == 1 {
                    //we narrowed down the array
                        tempVal = zerosNums[0]^num
                        if tempVal > maxXor {
                            maxXor = tempVal;
                        }
                        break;
                    }
                val = num math.pow(2, dig) //get a bitmask
                if val & num {
                    zerosNums =sliceArray(nums, 1, dig)
                }
                else { zerosNums = sliceArray(nums,0,dig)}
            }
            nums3 = nums //reset nums3 array
            sliceArray(nums3, 0, sigPlace) //refresh options
            }
        }
            return maxXor
        }
    }


/**
//brute force
func findMaximumXOR(nums []int) int {
    if len(nums) < 2 { return 0; }
    temp := 0;
    maxXor := -99999999 //replace with real min...
    for i := 0; i < len(nums); i++ {
        for j:= i+1; j < len(nums); j++ {
            temp = nums[i] ^ nums[j]
            if temp > maxXor{
                maxXor = temp
            }
        }
    }
    return maxXor
}
**/