from collections import deque

class Solution:
    # @param {integer[]} nums
    # @param {integer} k
    # @return {integer[]}
    def maxSlidingWindow(self, nums, k):
        if k < 1 or k > len(nums):
            return []
        
        queue = deque(nums)
        isFirst = 1
        maxsw = -1
        maxsws = []

        while len(queue) >= k:
            if isFirst != 1:
                if queue[k-1] >= maxsw:
                    maxsw = queue[k-1]
                else:  
                    isFirst = 1
                maxsws.append(maxsw)
                queue.popleft()
                continue
            
            maxsw = queue.popleft()
            if k != 1:
                for i, val in enumerate(queue):
                    if val >= maxsw:
                        maxsw = val
                        isFirst = 0
                        
                    if i+2 >= k:
                        break
            maxsws.append(maxsw)
        return maxsws

s = Solution()
print(s.maxSlidingWindow([5,3,4], 1))
