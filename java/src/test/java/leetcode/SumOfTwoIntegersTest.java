package leetcode;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

/**
 * Unit test for SumOfTwoIntegers.
 */
public class SumOfTwoIntegersTest {

    @Test
    public void getSum() {
        SumOfTwoIntegers t = new SumOfTwoIntegers();
        long start = System.nanoTime();
        int i = t.getSum(2, -3);
        long end = System.nanoTime();
        System.err.println("runTime:" + (end - start));
        System.out.println("sum:" + i);

        assertEquals(-1, i);
    }

}
