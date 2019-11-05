/**
 * Source  https://leetcode.com/problems/sum-of-two-integers/
 * Author  cytian
 * Created 2016-07-14
 *
 * Calculate the sum of two integers a and b,
 * but you are not allowed to use the operator + and -.
 *
 * Example: Given a = 1 and b = 2, return 3.
 */
public class SumOfTwoIntegers {
    public static void main(String[] args) {
        SumOfTwoIntegers t = new SumOfTwoIntegers();
        long start = System.nanoTime();
        int i = t.getSum(2, -3);
        long end = System.nanoTime();
        System.err.println("runTime:" + (end - start));
        System.out.println("sum:" + i);
    }

    public int getSum(int a, int b) {
        int aBit, bBit;
        int bit = 0;
        int sum = 0;
        int temp = 0;

        for (int i = 0; i <= 31; i++) {
            // convertor bit
            if (((1 << i) & a) != 0)
                aBit = 1;
            else
                aBit = 0;
            if (((1 << i) & b) != 0)
                bBit = 1;
            else
                bBit = 0;

            // operato add
            if (aBit + bBit + temp == 0) { // 0 + 0 = 0
                bit = 0;
            } else if (aBit + bBit + temp == 1) { // 0 + 1 = 1
                bit = 1;
                temp = 0;
            } else if (aBit + bBit + temp == 2) {  // 1 + 1 = 10 carrybit 1
                bit = 0;
                temp = 1;
            } else if (aBit + bBit + temp == 3) { // 1 + 1 + 1 = 11 carrybit 1
                bit = 1;
            }

            // convertor int
            sum = (1 << i)*bit + sum;
        }
        return sum;
    }

}
