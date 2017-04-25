package protein

const testVersion = 1

var mappa = map[string]string{"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP"}

func FromCodon(codon string) string {
	return mappa[codon]
}

func FromRNA(rna string) []string {
	final := []string{}
	for index := 0; index < len(rna); index += 3 {
		codon := rna[index : index+3]
		if mappa[codon] == "STOP" {
			return final
		}
		final = append(final, mappa[codon])

	}
	return final
}
