import java.util.Arrays;
import java.util.stream.Collectors;

/**
 * Created by alebaffa on 14/11/16.
 */
public class Pangrams {
    private final static int alphabetLenght = 26;

    public static boolean isPangram(String s) {

        return Arrays.stream(s.replaceAll("[^a-zA-Z]", "").toLowerCase().split(""))
                .collect(Collectors.toSet()).size() == alphabetLenght;


    }
}
