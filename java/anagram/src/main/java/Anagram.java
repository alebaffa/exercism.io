import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Anagram {
    private String word;

    public Anagram(String word) {
        this.word = word;
    }

    public List<String> match(List<String> strings) {
        List<String> result = new ArrayList<>();

        for (String candidate : strings) {
            if (haveDifferentSize(word, candidate))
                continue;

            if (isAnagram(word, candidate)) {
                result.add(candidate);
            }
        }
        return result;
    }

    private boolean isAnagram(String word, String candidate) {
        word = word.toLowerCase();
        candidate = candidate.toLowerCase();

        return !word.equals(candidate) && convert(word).equals(convert(candidate));
    }

    private boolean haveDifferentSize(String word, String candidate) {
        return word.length() != candidate.length();
    }

    private String convert(String s) {
        char[] ar = s.toCharArray();
        Arrays.sort(ar);
        return String.valueOf(ar);
    }
}
