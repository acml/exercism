// Package strand provend DNA to RNA complement function
package strand

// ToRNA given a DNA strand, returns its RNA complement
func ToRNA(dna string) string {
	rna := ""
	dnaToRna := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for _, char := range dna {
		rna += string(dnaToRna[char])
	}
	return rna
}
