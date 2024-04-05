package strand

var DNAtoRNA = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

func ToRNA(dna string) (rna string) {
	for _, base := range dna {
		rna += DNAtoRNA[base]
	}
	return rna
}
