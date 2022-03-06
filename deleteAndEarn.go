func max(a int, b int) int{
    if a > b {
        return a
    } else {return b}
}
func deleteAndEarn(nums []int) int {
    /**classic dynamic programming babes
    score(list_i) = max (list_i[val] + score(list_i+2),
                         score(list_i+1))
    except sometimes indices aren't present and that's ok
    we will just do a check **/
    nums2 := make([]int, len(nums))
    copy(nums2, nums)
    sort.Ints(nums2)
    totals := make([]int, len(nums2)) //running totals here
    if len(nums) == 1 {return nums2[0]}
    repeat_index := 0 //keeps track of last unique number
    back := -1
    forward := 1 //arithmetic indexing was not working
    for i, num := range(nums2){ //iterate through nums
        sum := num
        if i>0 && (nums2[back] != num - 1){ //not adjacent
            better:= totals[back]
            if nums2[repeat_index] != num - 1 {
                better = max(totals[back], totals[repeat_index])
            }
            sum += better //previous thing
        } else if i > 1 && nums2[repeat_index] != num-1 {
            sum += totals[repeat_index]
        }
        totals[i] = sum
        if i > 0 && num != nums2[back] {
            if totals[repeat_index] < totals[back]{
                repeat_index = back
            }
        }
        forward++
        back ++
    }
    return max(totals[len(totals) - 1],totals[repeat_index])
}
