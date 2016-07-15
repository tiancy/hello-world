# Source https://leetcode.com/problems/min-stack/
# Author cytian
# Created 2016-07-15

# Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
#
# push(x) -- Push element x onto stack.
# pop() -- Removes the element on top of the stack.
# top() -- Get the top element.
# getMin() -- Retrieve the minimum element in the stack.
# Example:
# MinStack minStack = new MinStack();
# minStack.push(-2);
# minStack.push(0);
# minStack.push(-3);
# minStack.getMin();   --> Returns -3.
# minStack.pop();
# minStack.top();      --> Returns 0.
# minStack.getMin();   --> Returns -2.

class MinStack(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.items = []
        self.minimum = None

    def push(self, x):
        """
        :type x: int
        :rtype: void
        """
        if len(self.items) == 0 or  x < self.minimum:
           self.minimum = x
        self.items.append(x)


    def pop(self):
        """
        :rtype: void
        """
        temp = self.items.pop()
        if temp == self.minimum and len(self.items) != 0:
            self.minimum = self.items[0]
            for i in self.items:
                if i < self.minimum:
                    self.minimum = i


    def top(self):
        """
        :rtype: int
        """
        return self.items[-1]

    def getMin(self):
        """
        :rtype: int
        """
        return self.minimum


# Your MinStack object will be instantiated and called as such:
obj = MinStack()
obj.push(2)
obj.push(5)
obj.push(-3)
print(obj.top())
obj.pop()
print(obj.top())

param_4 = obj.getMin()
print(param_4)
