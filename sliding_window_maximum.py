# Source https://leetcode.com/problems/sliding-window-maximum/
# Author cytian
# Updated 2016-07-17

# Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
#
# For example,
# Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.
#
# Window position                Max
# ---------------               -----
# [1  3  -1] -3  5  3  6  7       3
#  1 [3  -1  -3] 5  3  6  7       3
#  1  3 [-1  -3  5] 3  6  7       5
#  1  3  -1 [-3  5  3] 6  7       5
#  1  3  -1  -3 [5  3  6] 7       6
#  1  3  -1  -3  5 [3  6  7]      7
# Therefore, return the max sliding window as [3,3,5,5,6,7].
#
# Note:
# You may assume k is always valid, ie: 1 ≤ k ≤ input array's size for non-empty array.

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
