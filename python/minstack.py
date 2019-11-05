# Source  https://leetcode.com/problems/min-stack/
# Author  cytian
# Created 2016-07-15

# Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
#
# push(x) -- Push element x onto stack.
# pop() -- Removes the element on top of the stack.
# top() -- Get the top element.
# getMin() -- Retrieve the minimum element in the stack.

class MinStack(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.items = []
        self.minimums = [] # store minimum
        self.minimum = None

    def push(self, x):
        """
        :type x: int
        :rtype: void
        """
        if len(self.items) == 0 or  x <= self.minimum:
            self.minimum = x
            self.minimums.append(x)
        self.items.append(x)

    def pop(self):
        """
        :rtype: void
        """
        if len(self.items) == 0:
            return
        temp = self.items.pop()
        if temp == self.minimum:
            self.minimums.pop()
            l = len(self.minimums)
            if l != 0:
                self.minimum = self.minimums[l-1]

    def top(self):
        """
        :rtype: int
        """
        return self.items[len(self.items)-1]

    def getMin(self):
        """
        :rtype: int
        """
        return self.minimum


# Your MinStack object will be instantiated and called as such:
obj = MinStack()
obj.push(2)
obj.push(3)
obj.push(3)
obj.push(5)
obj.push(1)
print(obj.top())
obj.pop()
print(obj.top())
print(obj.getMin())
