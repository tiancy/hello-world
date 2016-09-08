package leetcode;

import java.util.HashMap;
import java.util.Map;

/**
 * Source  https://leetcode.com/problems/find-peak-element/
 * Author  cytian
 * Created 2016-08-12
 *
 * Given  an  arbitrary  ransom  note  string  and  another  string  containing  letters from  all  the  magazines, 
 * write  a  function  that  will  return  true  if  the  ransom   note  can  be  constructed  from  the  magazines ;
 * otherwise,  it  will  return  false.   
 *
 * Each  letter  in  the  magazine  string  can  only  be  used  once  in  your  ransom  note.
 *
 * Note:
 * You may assume that both strings contain only lowercase letters.
 *
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 */
public class RansomNote {

    public boolean canConstruct(String ransomNote, String magazine) {
        char[] notes = ransomNote.toCharArray();
        Map<Character, Integer> map = new HashMap<>();
        int index = -1;
        for (int i = 0; i < notes.length; i++) {
            if (map.containsKey(notes[i]) && (index = magazine.indexOf(notes[i], map.get(notes[i]) + 1)) < 0)
                return false;
            else if (!map.containsKey(notes[i]) && (index = magazine.indexOf(notes[i])) < 0)
                return false;
            map.put(notes[i], index);
        }
        return true;
    }

    public boolean canConstruct2(String ransomNote, String magazine) {
        char[] notes = ransomNote.toCharArray();
        int index;
        for (int i = 0; i < notes.length; i++) {
            if (notes[i] == '.')
                continue;
            else if ((index = magazine.indexOf(notes[i])) < 0)
                return false;
            else {
                int j = i;
                while ((j = ransomNote.indexOf(notes[i], j + 1)) >= 0) {
                    if ((index = magazine.indexOf(notes[i], index + 1)) < 0)
                        return false;
                    notes[j] = '.';
                }
            }
        }
        return true;
    }

    public boolean canConstruct3(String ransomNote, String magazine) {
        int count[] = new int[26];
        for (int i = 0; i < magazine.length(); i++) {
                        count[magazine.charAt(i)-'a']++;
                }
        for (int i = 0; i < ransomNote.length(); i++) {
                        if (--count[ransomNote.charAt(i)-'a'] < 0)
                                return false;
                }
        return true;
    }
}
