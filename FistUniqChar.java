package leetcode;

/**
 * Source https://leetcode.com/problems/first-unique-character-in-a-string/
 * Author cytian
 * Create 2016-09-06
 *
 * Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
 *
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode",
 * return 2.
 */
public class FistUniqChar {

    public int fistUniqChar(String s) {
        if (s == null || "".equals(s)) return -1;

        char chars[] = s.toCharArray();
        int index;
        for (int i = 0; i < chars.length; i++) {
            if (chars[i] == '.') continue;

            if ((index = s.indexOf(chars[i], i+1)) < 0)
                return i;
            else {
                do {
                    chars[index] = '.';
                } while ((index = s.indexOf(chars[i], index+1)) >= 0);
            }
        }
        return -1;
    }
}
