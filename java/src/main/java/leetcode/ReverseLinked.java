package leetcode;

// Definition for singly-linked list.
class ListNode {
    int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
}

/**
 * Source  https://leetcode.com/problems/reverse-linked-list/
 * Author  Tian
 * Date    2020-07-15
 * 
 * 206. Reverse a singly linked list.
 * 
 * Example:
 * 
 * Input: 1->2->3->4->5->NULL
 * Output: 5->4->3->2->1->NULL
 */
public class ReverseLinked {

    public ListNode reverseList(ListNode head) {
        ListNode prev = null;
        ListNode curr = head;
        while (curr != null) {
            ListNode next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }
        return prev;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(0);
        ListNode curr = head;
        for (int i = 1; i < 5; i++) {
            ListNode temp = new ListNode(i); 
            curr.next = temp;
            curr = curr.next;
        }

        ReverseLinked obj = new ReverseLinked();
        head = obj.reverseList(head) ;

        while (head != null) {
            System.err.println(head.val);
            head = head.next;
        }
    }
}