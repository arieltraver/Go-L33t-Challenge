class Solution(object):
    #BOYS THIS IS DP
    #why didn't we recognize it RIP
    #score(list_i) = max (list_i[val] + score(list_i+2),
    #                     score(list_i+1))
    #except sometimes indices aren't present and that's ok
    #we will make an array in ... GOLANG
    """
    def deleteAndEarn(self, nums):
        total=0
        real_vals = {}
        for num in nums:
            if num in real_vals:
                real_vals[num] += num
            else: real_vals[num] = num
        
        #prioritize highest values in dictionary
        #pop as you go
        while real_vals: #not empty
            index = max(real_vals, key=real_vals.get)
            total += real_vals[index]
            if (index-1) in real_vals:
                real_vals.pop(index-1)
            if (index+1) in real_vals:
                real_vals.pop(index+1)
            real_vals.pop(index)
            
        return total
        """
        """
        :type nums: List[int]
        :rtype: int
        """
        
