// Package protein provides functions to verify valid codons and protein
// sequences
package protein

type err struct {
	msg string
}

func (e err) Error() string {
	return e.msg
}

// ErrStop is returned when a terminating codon is reached
var ErrStop = err{"Stop Codon"}

// ErrInvalidBase is returned when an invalid codon is reached
var ErrInvalidBase = err{"Invalid protein sequence"}

// FromCodon matches codons to correspondent amino acids
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}
	return "", ErrInvalidBase
}

// FromRNA translates RNA sequences to amino acid sequences
func FromRNA(rna string) ([]string, error) {
	var codon string
	var res []string
	n := 0
	for _, char := range rna {
		if n%3 == 0 {
			codon = string(char)
		} else {
			codon += string(char)
		}

		if n%3 == 2 {
			protein, ok := FromCodon(codon)
			if ok == ErrStop {
				break
			} else if ok == ErrInvalidBase {
				return res, ErrInvalidBase
			}
			res = append(res, protein)
		}
		n++
	}
	return res, nil
}
