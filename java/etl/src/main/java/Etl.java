import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Etl {
    public Map<String, Integer> transform(Map<Integer, List<String>> old) {
        Map<String, Integer> newMap = new HashMap<>();

        old.forEach((score, letters) -> letters.forEach((letter) -> newMap.put(letter.toLowerCase(), score)));

        return newMap;
    }
}
