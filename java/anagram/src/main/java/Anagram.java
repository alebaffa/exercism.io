import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Anagram {
    private String word;

    public Anagram(String word) {
        this.word = word.toLowerCase();
    }

    public List<String> match(List<String> strings) {
        List<String> result = new ArrayList<>();

        for (String candidate : strings) {
            String originalCandidate = candidate;
            candidate = candidate.toLowerCase();

            if (haveDifferentSize(word, candidate))
                continue;

            if (isAnagram(word, candidate)) {
                result.add(originalCandidate);
            }
        }
        return result;
    }

    private boolean isAnagram(String word, String candidate) {
        String convertedWord = convert(word);
        String convertedCandidate = convert(candidate);
        return !word.equals(candidate) && convertedWord.equals(convertedCandidate);
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
