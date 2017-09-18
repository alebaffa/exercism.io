package protein

const testVersion = 1

//FromCodon returns the name of the codon given its code
func FromCodon(codon string) string {
	switch codon {
	case
		"AUG":
		return "Methionine"
	case
		"UUU",
		"UUC":
		return "Phenylalanine"
	case
		"UUA",
		"UUG":
		return "Leucine"
	case
		"UCU",
		"UCC",
		"UCA",
		"UCG":
		return "Serine"
	case
		"UAU",
		"UAC":
		return "Tyrosine"
	case
		"UGU",
		"UGC":
		return "Cysteine"
	case
		"UGG":
		return "Tryptophan"
	case
		"UAA",
		"UAG",
		"UGA":
		return "STOP"
	}
	return "STOP"
}

//FromRNA returns the name of the protein given the sequence of codons
func FromRNA(s string) []string {
	var res = ""
	var proteins []string
	for index, codon := range s {
		res += string(codon)
		if index > 0 && (index+1)%3 == 0 {
			tempCodon := FromCodon(res)
			if tempCodon == "STOP" {
				break
			}
			proteins = append(proteins, tempCodon)
			res = ""
		}
	}
	return proteins
}
