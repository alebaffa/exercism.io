import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.Matcher;

import static java.util.regex.Pattern.compile;

public class WordCount {

    public Map<String, Integer> phrase(String input) {
        Map<String, Integer> result = new HashMap<>();

        List<String> words = getUsefulWords(input);
        for (String word : words) result.merge(word, 1, (oldValue, newValue) -> oldValue + newValue);
        return result;
    }

    private List<String> getUsefulWords(String word) {
        Matcher matcher = compile("[A-Za-z1-9]+").matcher(word);

        List<String> words = new ArrayList<>();
        while (matcher.find()) words.add(matcher.group().toLowerCase());
        return words;
    }
}
