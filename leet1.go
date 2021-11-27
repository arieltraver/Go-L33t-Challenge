func productExceptSelf(nums []int) []int {
    var l int = len(nums)
    forwards := make([]int, l) //the forward direction array
    backwards := make([]int, l) //the backward direction array
    answers := make([]int, l) //will hold the results
    forwards[0] = nums[0];
    backwards[l-1] = nums[l - 1]
    for i := 1; i < l; i++ { //fill up forwards array
        forwards[i] = nums[i] * forwards[i-1]
    }
    for i2 := l-2; i2 >= 0; i2-- { //fill up backwards array
        backwards[i2] = nums[i2] * backwards[i2+1]
    }
    answers[l-1] = forwards[l-2] //set the ends
    answers[0] = backwards[1]
    for i3 := 1; i3 < l-1; i3++ { //array[i] = the product of forwards up to i, and backwards down to i
        answers[i3] = forwards[i3-1] * backwards[i3+1]
    }
    return answers
}
    
