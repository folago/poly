package fold

var dnaComplement = map[byte]byte{'A': 'T', 'T': 'A', 'G': 'C', 'C': 'G', 'N': 'N'}

var dnaMultibranch = MultibranchEnergies{A: 2.6, B: 0.2, C: 0.2, D: 2.0}

// The Thermodynamics of DNA Structural Motifs
// SantaLucia and Hicks, 2004
var dnaNearestNeighbors = MatchingBasepairEnergy{"AA/TT": Energy{EnthalpyH: -7.6, EntropyS: -21.3},
	"AC/TG":    {EnthalpyH: -8.4, EntropyS: -22.4},
	"AG/TC":    {EnthalpyH: -7.8, EntropyS: -21},
	"AT/TA":    {EnthalpyH: -7.2, EntropyS: -20.4},
	"C/G_tini": {EnthalpyH: 0, EntropyS: 0},
	"CA/GT":    {EnthalpyH: -8.5, EntropyS: -22.7},
	"CC/GG":    {EnthalpyH: -8, EntropyS: -19.9},
	"CG/GC":    {EnthalpyH: -10.6, EntropyS: -27.2},
	"CT/GA":    {EnthalpyH: -7.8, EntropyS: -21},
	"GA/CT":    {EnthalpyH: -8.2, EntropyS: -22.2},
	"GC/CG":    {EnthalpyH: -9.8, EntropyS: -24.4},
	"GG/CC":    {EnthalpyH: -8, EntropyS: -19.9},
	"GT/CA":    {EnthalpyH: -8.4, EntropyS: -22.4},
	"T/A_tini": {EnthalpyH: 2.2, EntropyS: 6.9},
	"TA/AT":    {EnthalpyH: -7.2, EntropyS: -21.3},
	"TC/AG":    {EnthalpyH: -8.2, EntropyS: -22.2},
	"TG/AC":    {EnthalpyH: -8.5, EntropyS: -22.7},
	"TT/AA":    {EnthalpyH: -7.6, EntropyS: -21.3},
	"init":     {EnthalpyH: 0.2, EntropyS: -5.7},
	"init_A/T": {EnthalpyH: 2.2, EntropyS: 6.9},
	"init_G/C": {EnthalpyH: 0, EntropyS: 0},
	"mys":      {EnthalpyH: 0, EntropyS: -1.4},
	"sym":      {EnthalpyH: 0, EntropyS: -1.4},
	"tini":     {EnthalpyH: 0.2, EntropyS: -5.7}}

// DNAInternalMismatches is an Internal mismatch table (DNA)
// Allawi & SantaLucia (1997), Biochemistry 36: 10581-10594
// Allawi & SantaLucia (1998), Biochemistry 37: 9435-9444
// Allawi & SantaLucia (1998), Biochemistry 37: 2170-2179 *
// Allawi & SantaLucia (1998), Nucl Acids Res 26: 2694-2701 *
// Peyret et al. (1999), Biochemistry 38: 3468-3477 *
var dnaInternalMismatches = MatchingBasepairEnergy{"AA/AT": Energy{EnthalpyH: 4.7, EntropyS: 12.9},
	"AA/CT": {EnthalpyH: 7.6, EntropyS: 20.2},
	"AA/GT": {EnthalpyH: 3, EntropyS: 7.4},
	"AA/TA": {EnthalpyH: 1.2, EntropyS: 1.7},
	"AA/TC": {EnthalpyH: 2.3, EntropyS: 4.6},
	"AA/TG": {EnthalpyH: -0.6, EntropyS: -2.3},
	"AC/AG": {EnthalpyH: -2.9, EntropyS: -9.8},
	"AC/CG": {EnthalpyH: -0.7, EntropyS: -3.8},
	"AC/GG": {EnthalpyH: 0.5, EntropyS: 3.2},
	"AC/TA": {EnthalpyH: 5.3, EntropyS: 14.6},
	"AC/TC": {EnthalpyH: 0, EntropyS: -4.4},
	"AC/TT": {EnthalpyH: 0.7, EntropyS: 0.2},
	"AG/AC": {EnthalpyH: -0.9, EntropyS: -4.2},
	"AG/CC": {EnthalpyH: 0.6, EntropyS: -0.6},
	"AG/GC": {EnthalpyH: -4, EntropyS: -13.2},
	"AG/TA": {EnthalpyH: -0.7, EntropyS: -2.3},
	"AG/TG": {EnthalpyH: -3.1, EntropyS: -9.5},
	"AG/TT": {EnthalpyH: 1, EntropyS: 0.9},
	"AT/AA": {EnthalpyH: 1.2, EntropyS: 1.7},
	"AT/CA": {EnthalpyH: 5.3, EntropyS: 14.6},
	"AT/GA": {EnthalpyH: -0.7, EntropyS: -2.3},
	"AT/TC": {EnthalpyH: -1.2, EntropyS: -6.2},
	"AT/TG": {EnthalpyH: -2.5, EntropyS: -8.3},
	"AT/TT": {EnthalpyH: -2.7, EntropyS: -10.8},
	"CA/AT": {EnthalpyH: 3.4, EntropyS: 8},
	"CA/CT": {EnthalpyH: 6.1, EntropyS: 16.4},
	"CA/GA": {EnthalpyH: -0.9, EntropyS: -4.2},
	"CA/GC": {EnthalpyH: 1.9, EntropyS: 3.7},
	"CA/GG": {EnthalpyH: -0.7, EntropyS: -2.3},
	"CA/TT": {EnthalpyH: 1, EntropyS: 0.7},
	"CC/AG": {EnthalpyH: 5.2, EntropyS: 14.2},
	"CC/CG": {EnthalpyH: 3.6, EntropyS: 8.9},
	"CC/GA": {EnthalpyH: 0.6, EntropyS: -0.6},
	"CC/GC": {EnthalpyH: -1.5, EntropyS: -7.2},
	"CC/GT": {EnthalpyH: -0.8, EntropyS: -4.5},
	"CC/TG": {EnthalpyH: 5.2, EntropyS: 13.5},
	"CG/AC": {EnthalpyH: 1.9, EntropyS: 3.7},
	"CG/CC": {EnthalpyH: -1.5, EntropyS: -7.2},
	"CG/GA": {EnthalpyH: -4, EntropyS: -13.2},
	"CG/GG": {EnthalpyH: -4.9, EntropyS: -15.3},
	"CG/GT": {EnthalpyH: -4.1, EntropyS: -11.7},
	"CG/TC": {EnthalpyH: -1.5, EntropyS: -6.1},
	"CT/AA": {EnthalpyH: 2.3, EntropyS: 4.6},
	"CT/CA": {EnthalpyH: 0, EntropyS: -4.4},
	"CT/GC": {EnthalpyH: -1.5, EntropyS: -6.1},
	"CT/GG": {EnthalpyH: -2.8, EntropyS: -8},
	"CT/GT": {EnthalpyH: -5, EntropyS: -15.8},
	"CT/TA": {EnthalpyH: -1.2, EntropyS: -6.2},
	"GA/AT": {EnthalpyH: 0.7, EntropyS: 0.7},
	"GA/CA": {EnthalpyH: -2.9, EntropyS: -9.8},
	"GA/CC": {EnthalpyH: 5.2, EntropyS: 14.2},
	"GA/CG": {EnthalpyH: -0.6, EntropyS: -1},
	"GA/GT": {EnthalpyH: 1.6, EntropyS: 3.6},
	"GA/TT": {EnthalpyH: -1.3, EntropyS: -5.3},
	"GC/AG": {EnthalpyH: -0.6, EntropyS: -1},
	"GC/CA": {EnthalpyH: -0.7, EntropyS: -3.8},
	"GC/CC": {EnthalpyH: 3.6, EntropyS: 8.9},
	"GC/CT": {EnthalpyH: 2.3, EntropyS: 5.4},
	"GC/GG": {EnthalpyH: -6, EntropyS: -15.8},
	"GC/TG": {EnthalpyH: -4.4, EntropyS: -12.3},
	"GG/AC": {EnthalpyH: -0.7, EntropyS: -2.3},
	"GG/CA": {EnthalpyH: 0.5, EntropyS: 3.2},
	"GG/CG": {EnthalpyH: -6, EntropyS: -15.8},
	"GG/CT": {EnthalpyH: 3.3, EntropyS: 10.4},
	"GG/GC": {EnthalpyH: -4.9, EntropyS: -15.3},
	"GG/TC": {EnthalpyH: -2.8, EntropyS: -8},
	"GG/TT": {EnthalpyH: 5.8, EntropyS: 16.3},
	"GT/AA": {EnthalpyH: -0.6, EntropyS: -2.3},
	"GT/CC": {EnthalpyH: 5.2, EntropyS: 13.5},
	"GT/CG": {EnthalpyH: -4.4, EntropyS: -12.3},
	"GT/CT": {EnthalpyH: -2.2, EntropyS: -8.4},
	"GT/GA": {EnthalpyH: -3.1, EntropyS: -9.5},
	"GT/TA": {EnthalpyH: -2.5, EntropyS: -8.3},
	"GT/TG": {EnthalpyH: 4.1, EntropyS: 9.5},
	"TA/AA": {EnthalpyH: 4.7, EntropyS: 12.9},
	"TA/AC": {EnthalpyH: 3.4, EntropyS: 8},
	"TA/AG": {EnthalpyH: 0.7, EntropyS: 0.7},
	"TA/CT": {EnthalpyH: 1.2, EntropyS: 0.7},
	"TA/GT": {EnthalpyH: -0.1, EntropyS: -1.7},
	"TA/TT": {EnthalpyH: 0.2, EntropyS: -1.5},
	"TC/AA": {EnthalpyH: 7.6, EntropyS: 20.2},
	"TC/AC": {EnthalpyH: 6.1, EntropyS: 16.4},
	"TC/AT": {EnthalpyH: 1.2, EntropyS: 0.7},
	"TC/CG": {EnthalpyH: 2.3, EntropyS: 5.4},
	"TC/GG": {EnthalpyH: 3.3, EntropyS: 10.4},
	"TC/TG": {EnthalpyH: -2.2, EntropyS: -8.4},
	"TG/AA": {EnthalpyH: 3, EntropyS: 7.4},
	"TG/AG": {EnthalpyH: 1.6, EntropyS: 3.6},
	"TG/AT": {EnthalpyH: -0.1, EntropyS: -1.7},
	"TG/CC": {EnthalpyH: -0.8, EntropyS: -4.5},
	"TG/GC": {EnthalpyH: -4.1, EntropyS: -11.7},
	"TG/GT": {EnthalpyH: -1.4, EntropyS: -6.2},
	"TG/TC": {EnthalpyH: -5, EntropyS: -15.8},
	"TT/AC": {EnthalpyH: 1, EntropyS: 0.7},
	"TT/AG": {EnthalpyH: -1.3, EntropyS: -5.3},
	"TT/AT": {EnthalpyH: 0.2, EntropyS: -1.5},
	"TT/CA": {EnthalpyH: 0.7, EntropyS: 0.2},
	"TT/GA": {EnthalpyH: 1, EntropyS: 0.9},
	"TT/GG": {EnthalpyH: 5.8, EntropyS: 16.3},
	"TT/TA": {EnthalpyH: -2.7, EntropyS: -10.8},
}

// DNATerminalMismatches is a terminal mismatch table for DNA found at terminating ends of a structure
// SantaLucia & Peyret (2001) Patent Application WO 01/94611
var dnaTerminalMismatches = MatchingBasepairEnergy{
	"AA/AT": Energy{EnthalpyH: -2.5, EntropyS: -6.3},
	"AA/CT": {EnthalpyH: -2.7, EntropyS: -7},
	"AA/GT": {EnthalpyH: -2.4, EntropyS: -5.8},
	"AA/TA": {EnthalpyH: -3.1, EntropyS: -7.8},
	"AA/TC": {EnthalpyH: -1.6, EntropyS: -4},
	"AA/TG": {EnthalpyH: -1.9, EntropyS: -4.4},
	"AC/AG": {EnthalpyH: -8, EntropyS: -22.5},
	"AC/CG": {EnthalpyH: -3.2, EntropyS: -7.1},
	"AC/GG": {EnthalpyH: -4.6, EntropyS: -11.4},
	"AC/TA": {EnthalpyH: -1.8, EntropyS: -3.8},
	"AC/TC": {EnthalpyH: -0.1, EntropyS: 0.5},
	"AC/TT": {EnthalpyH: -0.9, EntropyS: -1.7},
	"AG/AC": {EnthalpyH: -4.3, EntropyS: -10.7},
	"AG/CC": {EnthalpyH: -2.7, EntropyS: -6},
	"AG/GC": {EnthalpyH: -6, EntropyS: -15.5},
	"AG/TA": {EnthalpyH: -2.5, EntropyS: -5.9},
	"AG/TG": {EnthalpyH: -1.1, EntropyS: -2.1},
	"AG/TT": {EnthalpyH: -3.2, EntropyS: -8.7},
	"AT/AA": {EnthalpyH: -3.1, EntropyS: -7.8},
	"AT/CA": {EnthalpyH: -1.8, EntropyS: -3.8},
	"AT/GA": {EnthalpyH: -2.5, EntropyS: -5.9},
	"AT/TC": {EnthalpyH: -2.3, EntropyS: -6.3},
	"AT/TG": {EnthalpyH: -3.5, EntropyS: -9.4},
	"AT/TT": {EnthalpyH: -2.4, EntropyS: -6.5},
	"CA/AT": {EnthalpyH: -2.3, EntropyS: -5.9},
	"CA/CT": {EnthalpyH: -0.7, EntropyS: -1.3},
	"CA/GA": {EnthalpyH: -4.3, EntropyS: -10.7},
	"CA/GC": {EnthalpyH: -2.6, EntropyS: -5.9},
	"CA/GG": {EnthalpyH: -3.9, EntropyS: -9.6},
	"CA/TT": {EnthalpyH: -0.7, EntropyS: -1.2},
	"CC/AG": {EnthalpyH: -5, EntropyS: -13.8},
	"CC/CG": {EnthalpyH: -3.9, EntropyS: -10.6},
	"CC/GA": {EnthalpyH: -2.7, EntropyS: -6},
	"CC/GC": {EnthalpyH: -2.1, EntropyS: -5.1},
	"CC/GT": {EnthalpyH: -3.2, EntropyS: -8},
	"CC/TG": {EnthalpyH: -3, EntropyS: -7.8},
	"CG/AC": {EnthalpyH: -2.6, EntropyS: -5.9},
	"CG/CC": {EnthalpyH: -2.1, EntropyS: -5.1},
	"CG/GA": {EnthalpyH: -6, EntropyS: -15.5},
	"CG/GG": {EnthalpyH: -3.8, EntropyS: -9.5},
	"CG/GT": {EnthalpyH: -3.8, EntropyS: -9},
	"CG/TC": {EnthalpyH: -3.9, EntropyS: -10.6},
	"CT/AA": {EnthalpyH: -1.6, EntropyS: -4},
	"CT/CA": {EnthalpyH: -0.1, EntropyS: 0.5},
	"CT/GC": {EnthalpyH: -3.9, EntropyS: -10.6},
	"CT/GG": {EnthalpyH: -6.6, EntropyS: -18.7},
	"CT/GT": {EnthalpyH: -6.1, EntropyS: -16.9},
	"CT/TA": {EnthalpyH: -2.3, EntropyS: -6.3},
	"GA/AT": {EnthalpyH: -2, EntropyS: -4.7},
	"GA/CA": {EnthalpyH: -8, EntropyS: -22.5},
	"GA/CC": {EnthalpyH: -5, EntropyS: -13.8},
	"GA/CG": {EnthalpyH: -4.3, EntropyS: -11.1},
	"GA/GT": {EnthalpyH: -1.1, EntropyS: -2.7},
	"GA/TT": {EnthalpyH: -3.6, EntropyS: -9.8},
	"GC/AG": {EnthalpyH: -4.3, EntropyS: -11.1},
	"GC/CA": {EnthalpyH: -3.2, EntropyS: -7.1},
	"GC/CC": {EnthalpyH: -3.9, EntropyS: -10.6},
	"GC/CT": {EnthalpyH: -4.9, EntropyS: -13.5},
	"GC/GG": {EnthalpyH: -0.7, EntropyS: -19.2},
	"GC/TG": {EnthalpyH: -5.9, EntropyS: -16.1},
	"GG/AC": {EnthalpyH: -3.9, EntropyS: -9.6},
	"GG/CA": {EnthalpyH: -4.6, EntropyS: -11.4},
	"GG/CG": {EnthalpyH: -0.7, EntropyS: -19.2},
	"GG/CT": {EnthalpyH: -5.7, EntropyS: -15.9},
	"GG/GC": {EnthalpyH: -3.8, EntropyS: -9.5},
	"GG/TC": {EnthalpyH: -6.6, EntropyS: -18.7},
	"GT/AA": {EnthalpyH: -1.9, EntropyS: -4.4},
	"GT/CC": {EnthalpyH: -3, EntropyS: -7.8},
	"GT/CG": {EnthalpyH: -5.9, EntropyS: -16.1},
	"GT/CT": {EnthalpyH: -7.4, EntropyS: -21.2},
	"GT/GA": {EnthalpyH: -1.1, EntropyS: -2.1},
	"GT/TA": {EnthalpyH: -3.5, EntropyS: -9.4},
	"TA/AA": {EnthalpyH: -2.5, EntropyS: -6.3},
	"TA/AC": {EnthalpyH: -2.3, EntropyS: -5.9},
	"TA/AG": {EnthalpyH: -2, EntropyS: -4.7},
	"TA/CT": {EnthalpyH: -2.5, EntropyS: -6.3},
	"TA/GT": {EnthalpyH: -3.9, EntropyS: -10.5},
	"TA/TT": {EnthalpyH: -3.2, EntropyS: -8.9},
	"TC/AA": {EnthalpyH: -2.7, EntropyS: -7},
	"TC/AC": {EnthalpyH: -0.7, EntropyS: -1.3},
	"TC/AT": {EnthalpyH: -2.5, EntropyS: -6.3},
	"TC/CG": {EnthalpyH: -4.9, EntropyS: -13.5},
	"TC/GG": {EnthalpyH: -5.7, EntropyS: -15.9},
	"TC/TG": {EnthalpyH: -7.4, EntropyS: -21.2},
	"TG/AA": {EnthalpyH: -2.4, EntropyS: -5.8},
	"TG/AG": {EnthalpyH: -1.1, EntropyS: -2.7},
	"TG/AT": {EnthalpyH: -3.9, EntropyS: -10.5},
	"TG/CC": {EnthalpyH: -3.2, EntropyS: -8},
	"TG/GC": {EnthalpyH: -3.8, EntropyS: -9},
	"TG/TC": {EnthalpyH: -6.1, EntropyS: -16.9},
	"TT/AC": {EnthalpyH: -0.7, EntropyS: -1.2},
	"TT/AG": {EnthalpyH: -3.6, EntropyS: -9.8},
	"TT/AT": {EnthalpyH: -3.2, EntropyS: -8.9},
	"TT/CA": {EnthalpyH: -0.9, EntropyS: -1.7},
	"TT/GA": {EnthalpyH: -3.2, EntropyS: -8.7},
	"TT/TA": {EnthalpyH: -2.4, EntropyS: -6.5},
}

// DNA dangling ends
//
// Bommarito et al. (2000), Nucl Acids Res 28: 1929-1934
var dnaDanglingEnds = MatchingBasepairEnergy{
	"AA/.T": {EnthalpyH: 0.2, EntropyS: 2.3},
	"AA/T.": {EnthalpyH: -0.5, EntropyS: -1.1},
	".A/AT": {EnthalpyH: -0.7, EntropyS: -0.8},
	"AC/.G": {EnthalpyH: -6.3, EntropyS: -17.1},
	".A/CT": {EnthalpyH: 4.4, EntropyS: 14.9},
	"AC/T.": {EnthalpyH: 4.7, EntropyS: 14.2},
	"AG/.C": {EnthalpyH: -3.7, EntropyS: -10},
	".A/GT": {EnthalpyH: -1.6, EntropyS: -3.6},
	"AG/T.": {EnthalpyH: -4.1, EntropyS: -13.1},
	"A./TA": {EnthalpyH: -2.9, EntropyS: -7.6},
	"AT/.A": {EnthalpyH: -2.9, EntropyS: -7.6},
	"A./TC": {EnthalpyH: -4.1, EntropyS: -13},
	"A./TG": {EnthalpyH: -4.2, EntropyS: -15},
	"A./TT": {EnthalpyH: -0.2, EntropyS: -0.5},
	".A/TT": {EnthalpyH: 2.9, EntropyS: 10.4},
	"AT/T.": {EnthalpyH: -3.8, EntropyS: -12.6},
	".C/AG": {EnthalpyH: -2.1, EntropyS: -3.9},
	"CA/G.": {EnthalpyH: -5.9, EntropyS: -16.5},
	"CA/.T": {EnthalpyH: 0.6, EntropyS: 3.3},
	".C/CG": {EnthalpyH: -0.2, EntropyS: -0.1},
	"CC/G.": {EnthalpyH: -2.6, EntropyS: -7.4},
	"CC/.G": {EnthalpyH: -4.4, EntropyS: -12.6},
	"C./GA": {EnthalpyH: -3.7, EntropyS: -10},
	"C./GC": {EnthalpyH: -4, EntropyS: -11.9},
	"CG/.C": {EnthalpyH: -4, EntropyS: -11.9},
	"CG/G.": {EnthalpyH: -3.2, EntropyS: -10.4},
	"C./GG": {EnthalpyH: -3.9, EntropyS: -10.9},
	".C/GG": {EnthalpyH: -3.9, EntropyS: -11.2},
	"C./GT": {EnthalpyH: -4.9, EntropyS: -13.8},
	"CT/.A": {EnthalpyH: -4.1, EntropyS: -13},
	".C/TG": {EnthalpyH: -4.4, EntropyS: -13.1},
	"CT/G.": {EnthalpyH: -5.2, EntropyS: -15},
	"GA/C.": {EnthalpyH: -2.1, EntropyS: -3.9},
	".G/AC": {EnthalpyH: -5.9, EntropyS: -16.5},
	"GA/.T": {EnthalpyH: -1.1, EntropyS: -1.6},
	"G./CA": {EnthalpyH: -6.3, EntropyS: -17.1},
	"GC/C.": {EnthalpyH: -0.2, EntropyS: -0.1},
	".G/CC": {EnthalpyH: -2.6, EntropyS: -7.4},
	"G./CC": {EnthalpyH: -4.4, EntropyS: -12.6},
	"G./CG": {EnthalpyH: -5.1, EntropyS: -14},
	"GC/.G": {EnthalpyH: -5.1, EntropyS: -14},
	"G./CT": {EnthalpyH: -4, EntropyS: -10.9},
	".G/GC": {EnthalpyH: -3.2, EntropyS: -10.4},
	"GG/.C": {EnthalpyH: -3.9, EntropyS: -10.9},
	"GG/C.": {EnthalpyH: -3.9, EntropyS: -11.2},
	"GT/.A": {EnthalpyH: -4.2, EntropyS: -15},
	"GT/C.": {EnthalpyH: -4.4, EntropyS: -13.1},
	".G/TC": {EnthalpyH: -5.2, EntropyS: -15},
	"T./AA": {EnthalpyH: 0.2, EntropyS: 2.3},
	".T/AA": {EnthalpyH: -0.5, EntropyS: -1.1},
	"TA/A.": {EnthalpyH: -0.7, EntropyS: -0.8},
	"T./AC": {EnthalpyH: 0.6, EntropyS: 3.3},
	"T./AG": {EnthalpyH: -1.1, EntropyS: -1.6},
	"T./AT": {EnthalpyH: -6.9, EntropyS: -20},
	"TA/.T": {EnthalpyH: -6.9, EntropyS: -20},
	"TC/A.": {EnthalpyH: 4.4, EntropyS: 14.9},
	".T/CA": {EnthalpyH: 4.7, EntropyS: 14.2},
	"TC/.G": {EnthalpyH: -4, EntropyS: -10.9},
	"TG/A.": {EnthalpyH: -1.6, EntropyS: -3.6},
	".T/GA": {EnthalpyH: -4.1, EntropyS: -13.1},
	"TG/.C": {EnthalpyH: -4.9, EntropyS: -13.8},
	"TT/.A": {EnthalpyH: -0.2, EntropyS: -0.5},
	"TT/A.": {EnthalpyH: 2.9, EntropyS: 10.4},
	".T/TA": {EnthalpyH: -3.8, EntropyS: -12.6},
}

// Experimental delta EnthalpyH and delta EntropyS for tri/tetra loops
//
// Supplemental Material: Annu.Rev.Biophs.Biomol.Struct.33:415-40
// doi: 10.1146/annurev.biophys.32.110601.141800
// The Thermodynamics of DNA Structural Motifs
// SantaLucia and Hicks, 2004
//
// delta EntropyS was computed using delta G and delta EnthalpyH and is in cal / (K x mol)
// (versus delta EnthalpyH in kcal / mol)
var dnaTriTetraLoops = MatchingBasepairEnergy{
	"AGAAT":  {-1.5, 0.0},
	"AGCAT":  {-1.5, 0.0},
	"AGGAT":  {-1.5, 0.0},
	"AGTAT":  {-1.5, 0.0},
	"CGAAG":  {-2.0, 0.0},
	"CGCAG":  {-2.0, 0.0},
	"CGGAG":  {-2.0, 0.0},
	"CGTAG":  {-2.0, 0.0},
	"GGAAC":  {-2.0, 0.0},
	"GGCAC":  {-2.0, 0.0},
	"GGGAC":  {-2.0, 0.0},
	"GGTAC":  {-2.0, 0.0},
	"TGAAA":  {-1.5, 0.0},
	"TGCAA":  {-1.5, 0.0},
	"TGGAA":  {-1.5, 0.0},
	"TGTAA":  {-1.5, 0.0},
	"AAAAAT": {0.5, 0.6},
	"AAAACT": {0.7, -1.6},
	"AAACAT": {1.0, -1.6},
	"ACTTGT": {0.0, -4.2},
	"AGAAAT": {-1.1, -1.6},
	"AGAGAT": {-1.1, -1.6},
	"AGATAT": {-1.5, -1.6},
	"AGCAAT": {-1.6, -1.6},
	"AGCGAT": {-1.1, -1.6},
	"AGCTTT": {0.2, -1.6},
	"AGGAAT": {-1.1, -1.6},
	"AGGGAT": {-1.1, -1.6},
	"AGGGGT": {0.5, -0.6},
	"AGTAAT": {-1.6, -1.6},
	"AGTGAT": {-1.1, -1.6},
	"AGTTCT": {0.8, -1.6},
	"ATTCGT": {-0.2, -1.6},
	"ATTTGT": {0.0, -1.6},
	"ATTTTT": {-0.5, -1.6},
	"CAAAAG": {0.5, 1.3},
	"CAAACG": {0.7, 0.0},
	"CAACAG": {1.0, 0.0},
	"CAACCG": {0.0, 0.0},
	"CCTTGG": {0.0, -2.6},
	"CGAAAG": {-1.1, 0.0},
	"CGAGAG": {-1.1, 0.0},
	"CGATAG": {-1.5, 0.0},
	"CGCAAG": {-1.6, 0.0},
	"CGCGAG": {-1.1, 0.0},
	"CGCTTG": {0.2, 0.0},
	"CGGAAG": {-1.1, 0.0},
	"CGGGAG": {-1.0, 0.0},
	"CGGGGG": {0.5, 1.0},
	"CGTAAG": {-1.6, 0.0},
	"CGTGAG": {-1.1, 0.0},
	"CGTTCG": {0.8, 0.0},
	"CTTCGG": {-0.2, 0.0},
	"CTTTGG": {0.0, 0.0},
	"CTTTTG": {-0.5, 0.0},
	"GAAAAC": {0.5, 3.2},
	"GAAACC": {0.7, 0.0},
	"GAACAC": {1.0, 0.0},
	"GCTTGC": {0.0, -2.6},
	"GGAAAC": {-1.1, 0.0},
	"GGAGAC": {-1.1, 0.0},
	"GGATAC": {-1.6, 0.0},
	"GGCAAC": {-1.6, 0.0},
	"GGCGAC": {-1.1, 0.0},
	"GGCTTC": {0.2, 0.0},
	"GGGAAC": {-1.1, 0.0},
	"GGGGAC": {-1.1, 0.0},
	"GGGGGC": {0.5, 1.0},
	"GGTAAC": {-1.6, 0.0},
	"GGTGAC": {-1.1, 0.0},
	"GGTTCC": {0.8, 0.0},
	"GTTCGC": {-0.2, 0.0},
	"GTTTGC": {0.0, 0.0},
	"GTTTTC": {-0.5, 0.0},
	"GAAAAT": {0.5, 3.2},
	"GAAACT": {1.0, 0.0},
	"GAACAT": {1.0, 0.0},
	"GCTTGT": {0.0, -1.6},
	"GGAAAT": {-1.1, 0.0},
	"GGAGAT": {-1.1, 0.0},
	"GGATAT": {-1.6, 0.0},
	"GGCAAT": {-1.6, 0.0},
	"GGCGAT": {-1.1, 0.0},
	"GGCTTT": {-0.1, 0.0},
	"GGGAAT": {-1.1, 0.0},
	"GGGGAT": {-1.1, 0.0},
	"GGGGGT": {0.5, 1.0},
	"GGTAAT": {-1.6, 0.0},
	"GGTGAT": {-1.1, 0.0},
	"GTATAT": {-0.5, 0.0},
	"GTTCGT": {-0.4, 0.0},
	"GTTTGT": {-0.4, 0.0},
	"GTTTTT": {-0.5, 0.0},
	"TAAAAA": {0.5, -0.3},
	"TAAACA": {0.7, -1.6},
	"TAACAA": {1.0, -1.6},
	"TCTTGA": {0.0, -4.2},
	"TGAAAA": {-1.1, -1.6},
	"TGAGAA": {-1.1, -1.6},
	"TGATAA": {-1.6, -1.6},
	"TGCAAA": {-1.6, -1.6},
	"TGCGAA": {-1.1, -1.6},
	"TGCTTA": {0.2, -1.6},
	"TGGAAA": {-1.1, -1.6},
	"TGGGAA": {-1.1, -1.6},
	"TGGGGA": {0.5, -0.6},
	"TGTAAA": {-1.6, -1.6},
	"TGTGAA": {-1.1, -1.6},
	"TGTTCA": {0.8, -1.6},
	"TTTCGA": {-0.2, -1.6},
	"TTTTGA": {0.0, -1.6},
	"TTTTTA": {-0.5, -1.6},
	"TAAAAG": {0.5, 1.6},
	"TAAACG": {1.0, -1.6},
	"TAACAG": {1.0, -1.6},
	"TCTTGG": {0.0, -3.2},
	"TGAAAG": {-1.0, -1.6},
	"TGAGAG": {-1.0, -1.6},
	"TGATAG": {-1.5, -1.6},
	"TGCAAG": {-1.5, -1.6},
	"TGCGAG": {-1.0, -1.6},
	"TGCTTG": {-0.1, -1.6},
	"TGGAAG": {-1.0, -1.6},
	"TGGGAG": {-1.0, -1.6},
	"TGGGGG": {0.5, -0.6},
	"TGTAAG": {-1.5, -1.6},
	"TGTGAG": {-1.0, -1.6},
	"TTTCGG": {-0.4, -1.6},
	"TTTTAG": {-1.0, -1.6},
	"TTTTGG": {-0.4, -1.6},
	"TTTTTG": {-0.5, -1.6},
}

// Enthalpy and entropy increments for length dependence of internal loops
//
// Were calculated from delta G Table 4 of SantaLucia, 2004:
//
// Annu.Rev.Biophs.Biomol.Struct.33:415-40
// doi: 10.1146/annurev.biophys.32.110601.141800
// The Thermodynamics of DNA Structural Motifs
// SantaLucia and Hicks, 2004
//
// Additional loop sizes are accounted for with the Jacobson-Stockmayer
// entry extrapolation formula in paper:
// delta G (loop-n) = delta G (loop-x) + 2.44 x R x 310.15 x ln(n / x)
//
// Additional correction is applied for asymmetric loops in paper:
// delta G (asymmetry) = |length A - length B| x 0.3 (kcal / mol)
// where A and B are lengths of both sides of loop
var dnaInternalLoops = LoopEnergy{
	30: {0, -21.3},
	29: {0, -21.0},
	28: {0, -21.0},
	27: {0, -20.6},
	26: {0, -20.3},
	25: {0, -20.3},
	24: {0, -20.0},
	23: {0, -19.7},
	22: {0, -19.3},
	21: {0, -19.0},
	20: {0, -19.0},
	19: {0, -18.7},
	18: {0, -18.7},
	17: {0, -18.4},
	16: {0, -18.1},
	15: {0, -17.7},
	14: {0, -17.4},
	13: {0, -16.4},
	12: {0, -16.8},
	11: {0, -16.1},
	10: {0, -15.8},
	9:  {0, -15.8},
	8:  {0, -15.5},
	7:  {0, -14.8},
	6:  {0, -14.2},
	5:  {0, -12.9},
	4:  {0, -11.6},
	3:  {0, -10.3},
	2:  {0, 0},
	1:  {0, 0},
}

// Enthalpy and entropy increments for length depedence of bulge loops
//
// Were calculated from delta G Table 4 of SantaLucia, 2004:
//
// Annu.Rev.Biophs.Biomol.Struct.33:415-40
// doi: 10.1146/annurev.biophys.32.110601.141800
// The Thermodynamics of DNA Structural Motifs
// SantaLucia and Hicks, 2004
//
// For bulge loops of size 1, the intervening NearestNeighbors energy is used.
// Closing AT penalty is applied on both sides
var dnaBulgeLoops = LoopEnergy{
	1:  {0, -12.9},
	2:  {0, -9.4},
	3:  {0, -10.0},
	4:  {0, -10.3},
	5:  {0, -10.6},
	6:  {0, -11.3},
	7:  {0, -11.9},
	8:  {0, -12.6},
	9:  {0, -13.2},
	10: {0, -13.9},
	11: {0, -14.2},
	12: {0, -14.5},
	13: {0, -14.8},
	14: {0, -15.5},
	15: {0, -15.8},
	16: {0, -16.1},
	17: {0, -16.4},
	18: {0, -16.8},
	19: {0, -16.8},
	20: {0, -17.1},
	21: {0, -17.4},
	22: {0, -17.4},
	23: {0, -17.7},
	24: {0, -17.7},
	25: {0, -18.1},
	26: {0, -18.1},
	27: {0, -18.4},
	28: {0, -18.7},
	29: {0, -18.7},
	30: {0, -19.0},
}

// Enthalpy and entropy increments for length depedence of hairpin loops
//
// Were calculated from delta G Table 4 of SantaLucia, 2004:
//
// Annu.Rev.Biophs.Biomol.Struct.33:415-40
// doi: 10.1146/annurev.biophys.32.110601.141800
// The Thermodynamics of DNA Structural Motifs
// SantaLucia and Hicks, 2004
//
// For hairpins of length 3 and 4, the entropy values are looked up
// in the DNATriTetraLoops Dict
//
// From formula 8-9 of the paper:
// An additional 1.6 delta entropy penalty if the hairpin is closed by AT
var dnaHairpinLoops = LoopEnergy{
	1:  {0, 0.0},
	2:  {0, 0.0},
	3:  {0, -11.3},
	4:  {0, -11.3},
	5:  {0, -10.6},
	6:  {0, -12.9},
	7:  {0, -13.5},
	8:  {0, -13.9},
	9:  {0, -14.5},
	10: {0, -14.8},
	11: {0, -15.5},
	12: {0, -16.1},
	13: {0, -16.1},
	14: {0, -16.4},
	15: {0, -16.8},
	16: {0, -17.1},
	17: {0, -17.4},
	18: {0, -17.7},
	19: {0, -18.1},
	20: {0, -18.4},
	21: {0, -18.7},
	22: {0, -18.7},
	23: {0, -19.0},
	24: {0, -19.3},
	25: {0, -19.7},
	26: {0, -19.7},
	27: {0, -19.7},
	28: {0, -20.0},
	29: {0, -20.0},
	30: {0, -20.3},
}

var dnaEnergies = Energies{
	BulgeLoops:         dnaBulgeLoops,
	Complement:         dnaComplement,
	DanglingEnds:       dnaDanglingEnds,
	HairpinLoops:       dnaHairpinLoops,
	Multibranch:        dnaMultibranch,
	InternalLoops:      dnaInternalLoops,
	InternalMismatches: dnaInternalMismatches,
	NearestNeighbors:   dnaNearestNeighbors,
	TerminalMismatches: dnaTerminalMismatches,
	TriTetraLoops:      dnaTriTetraLoops,
}