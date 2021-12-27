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
            result: min(first.height, second.height) * (second - first) **/

func maxArea(height []int){
    var l = len(height);
    var firstBest = height[0] * l;
    var firstI = 0;
    var result = 0;
    for i:= 1; i < l; i++ {
        val := height[i] * (l-i);
        if val > firstBest {
            firstBest = val;
            firstI = i;
        }
    }
    for i2:= 0; i < l; i++ {
        if i2 != i {
            val := math.abs((lastI - firstI) * math.min(height[firstI], i2))
            result = math.max(result, val)
        }
    }
    return result;
}
