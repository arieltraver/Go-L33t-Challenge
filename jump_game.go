//jump reach (first index)
///jumpgame(array[0:i]) = jumpgame(array[0:last index which can reach i])


//import("container/list")

func canJump(nums []int) bool {
    l := len(nums)
    i := l - 2
    if l == 2 {
        return (nums[0] >= 1);
    } else if l == 1 {
        return true;
    } else {
        for i + nums[i] < l - 1 { //look backwards
            i --;
            if (i < 0) {return false;}
        }
    }
    return canJump(nums[0:i+1]);
}

/** weird BFS solution
func canJump(nums []int) bool {
    queue := list.New()
    seen := make(map [int]int)
    queue.PushFront(0)
    seen[0] = 1
    for queue.Len() > 0 {
        current := queue.Remove(queue.Front()).(int)
        if (current == len(nums) - 1) {return true;}
        for i := 0; i <= nums[current]; i++ {
            hopTo := current + i
            if hopTo == len(nums) - 1 {
                return true;
            } else if (seen[hopTo] != 1) {
                seen[hopTo] = 1;
                queue.PushBack(hopTo)
            }
        }
    }
    return false
}
**/
