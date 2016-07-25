# Source https://leetcode.com/problems/guess-number-higher-or-lower/
# Author cytian
# Updated 2016-07-25

# We are playing the Guess Game. The game is as follows:
# I pick a number from 1 to n. You have to guess which number I picked.
# Every time you guess wrong, I'll tell you whether the number is higher or lower.

# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num):

class Solution(object):

    def guess(self, num):
        if num == 1:
            return 0
        if num > 1:
            return -1
        return 1

    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        lower = 1
        higher = n
        while lower <= higher:
            mid = (lower + higher)//2
            res = self.guess(mid)
            if res == 0:
                return mid
            elif res == 1:
                lower = mid + 1
            else:
                higher = mid - 1
        return -1


obj = Solution()
r = obj.guessNumber(9)
print(r)
