import java.util.HashMap;
import java.util.Map;

/**
 * 
 * @author abaffa
 *
 */
public class DNA {

	private String dna;

	private enum NucleotidType {
		ADENINE('A'), CYTOSINE('C'),
		/**
		 * 
		 */
		GUANINE('G'), THYMINE('T');

		char code;

		NucleotidType(char code) {
			this.code = code;
		}

		char getNucleotidName() {
			return this.code;
		}
	}

	public DNA(String dna) {
		this.dna = dna;
	}

	int count(char nucleotidToSearch) {
		if (!isValidNucleotid(nucleotidToSearch)) {
			throw new IllegalArgumentException();
		}
		int countOfNucleotid = 0;
		for (char nucleotid : dna.toCharArray()) {
			if (nucleotidToSearch == nucleotid) {
				countOfNucleotid++;
			}
		}
		return countOfNucleotid;
	}

	private boolean isValidNucleotid(char nucleotidToSearch) {
		for (NucleotidType nucleotid : NucleotidType.values()) {
			if (nucleotidToSearch == nucleotid.getNucleotidName())
				return true;
		}
		return false;

	}

	public Map<Character, Integer> nucleotideCounts() {
		Map<Character, Integer> counts = new HashMap<Character, Integer>();
		for (NucleotidType nucleotid : NucleotidType.values()) {
			counts.put(nucleotid.code, count(nucleotid.code));
		}
		return counts;
	}

}