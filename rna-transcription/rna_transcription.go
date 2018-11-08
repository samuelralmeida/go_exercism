package strand

func ToRNA(dna string) string {
	var rna string
	for _, a := range dna {
		switch a {
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
