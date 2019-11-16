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
public class FirstUniqChar {

    public int firstUniqChar(String s) {
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

    public int firstUniqChar2(String s) {
        if (s == null || "".equals(s)) return -1;

        int count[] = new int[26];
        int len = s.length();
        for (int i = 0; i < len; i++) {
            count[s.charAt(i)-'a']++;
        }
        for (int i = 0; i < len; i++) {
            if (count[s.charAt(i)-'a'] == 1)
                return i;
        }
        return -1;
    }
}
