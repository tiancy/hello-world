/**
 * Source  https://leetcode.com/problems/find-peak-element/
 * Author  cytian
 * Created 2016-08-05
 *
 * 162. Find Peak Element
 * A peak element is an element that is greater than its neighbors.
 * Given an input array where num[i] ≠ num[i+1], find a peak element and return its index.
 * The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.
 * You may imagine that num[-1] = num[n] = -∞.
 * For example, in array [1, 2, 3, 1], 3 is a peak element and your function should return the index number 2.
 */
public class FindPeakElement {
    public int findPeakElement(int[] nums) {
        int start = 0;
        int end = nums.length - 1;
        while (start + 1 < end) {
            int mid = (start + end)/2;
            if (nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1])
                return mid;
            else if (nums[mid] > nums[mid - 1])
                start = mid;
            else
                end = mid;
        }
        if (nums[start] > nums[end])
            return start;
        return end;
    }
}
