/**Goal:
    given an int array called "height" with length n
    there are n vertical lines drawn such that the two
    endpoints of the ith line are (i, 0) and (i, height[i])
    Find two lines that together with the x-axis form a container
    such that the container holds the most water
    In essence: the rectangle formed by those lines and the
    distance between their endpoints has the greatest area.
    Return: maximum water that can be held.

    Thoughts: water(i1, i2) = i2 - i1 * min(val of i2, val of i1)
		FAILED GREEDY
            maybe you could know the value needed to beat the max?
            like skip over certain lines if their height + distance
            from some known prior point isn't good enough
            give individual lines some special value related
            to their height and their distance from a start point
            how to choose that start point?...
            first most useful = tallest height, closest to start
            go thru and pick this.
                --> ((n - start) * height) 
            second most useful = tallest height, farthest from min(most useful)
                --> ((i - first.i) * height)
            result: min(first.height, second.height) * (second - first)
	Thoughts2: nxn algorithm
	brute force
***/

import ("math")

//FAILED GREEDY
/**func maxArea(height []int) int{
    var l int = len(height);
    var firstBest int = 0;
    firstI := 0;
    result := 0;
    for i:= 0; i < l; i++ {
        val := height[i] * (l-i-1);
        if val > firstBest {
            firstBest = val;
            firstI = i;
        }
    }
    for i2:= 0; i2 < l; i2++ {
        if i2 != firstI {
            val := math.Abs( (float64(i2) - float64(firstI)) * math.Min(float64(height[firstI]), float64(height[i2])))
            result = int(math.Max(float64(result), val))
        }
    }
    return result;
}**/
func maxArea(height []int) int {
    
    currentMax := 0
    
    //calculates container formed by two lines, given their positions and heights
    func checkWater(index1 int, index2 int) {
        newContainer := index1 - index2 * int(math.Min(float64(height[index1]), float64(height[index2])))
        currentMax = int(math.Max(float64(newContainer), float64(currentMax)))
    }
    
    //embedded helper func within scope of the "height" array.
    //to save space, we avoid duplicating "height"
    func descendingMerge(left []int, right []int) 
        leftIndex := 0
        rightIndex := 0
        newIndex := 0; //keep track
        while rightIndex < len(left) && rightIndex <= len(right) {
            //checkWater(leftIndex, rightIndex)
            if left[leftIndex] > right[rightIndex] {
                newArray[newIndex] = left[leftIndex]
                leftIndex++
                newIndex++
            }
            else { //if right is bigger or they're equal
                newArray[newIndex] = right[rightIndex] //important
                rightIndex++
                newIndex++
            }
        }
        //fill up the rest of the new array
        while leftIndex < len(left){
            newArray[newIndex] = left[leftIndex]
            leftIndex++
            newIndex++
        }
        while rightIndex < len(right) {
            newArray[newIndex] = right[rightIndex];
            rightIndex++
            newIndex++
        }
        return newArray
    } //end of merge embedded function
    
}
