package foodchain

var what = []string{
	"fly.",
	"spider.\nIt wriggled and jiggled and tickled inside her.",
	"bird.\nHow absurd to swallow a bird!",
	"cat.\nImagine that, to swallow a cat!",
	"dog.\nWhat a hog, to swallow a dog!",
	"goat.\nJust opened her throat and swallowed a goat!",
	"cow.\nI don't know how she swallowed a cow!",
	"horse.\nShe's dead, of course!",
}

var reason = []string{
	"\n",
	"\nShe swallowed the spider to catch the fly.",
	"\nShe swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
	"\nShe swallowed the cat to catch the bird.",
	"\nShe swallowed the dog to catch the cat.",
	"\nShe swallowed the goat to catch the dog.",
	"\nShe swallowed the cow to catch the goat.",
}

func why(n int) string {
	if n == 0 {
		return "I don't know why she swallowed the fly. Perhaps she'll die."
	}
	return reason[n-1] + why(n-1)
}

// Verse recites a specific verse from the song.
func Verse(n int) string {
	if n == len(what) {
		return "I know an old lady who swallowed a " + what[n-1]
	}
	return "I know an old lady who swallowed a " + what[n-1] + why(n)
}

// Verses recites a specific range of verses from the song.
func Verses(start, stop int) (verses string) {
	for i := start; i <= stop; i++ {
		verses += Verse(i)
		if i != stop {
			verses += "\n\n"
		}
	}
	return verses
}

// Song generates the lyrics of the song 'I Know an Old Lady Who Swallowed a
// Fly'.
func Song() string {
	return Verses(1, len(what))
}
