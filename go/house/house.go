package house

var verse = []string{
	"the horse and the hound and the horn\nthat belonged to ",
	"the farmer sowing his corn\nthat kept ",
	"the rooster that crowed in the morn\nthat woke ",
	"the priest all shaven and shorn\nthat married ",
	"the man all tattered and torn\nthat kissed ",
	"the maiden all forlorn\nthat milked ",
	"the cow with the crumpled horn\nthat tossed ",
	"the dog\nthat worried ",
	"the cat\nthat killed ",
	"the rat\nthat ate ",
	"the malt\nthat lay in ",
	"",
}

// Song recites the nursery rhyme 'This is the House that Jack Built'.
func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i)
		if i < 12 {
			song += "\n\n"
		}
	}
	return song
}

// Verse recites a specific verse from the nursery rhyme.
func Verse(v int) string {
	return "This is " + generateVerse(v)
}

func generateVerse(v int) string {
	if v == 0 {
		return "the house that Jack built."
	}
	return verse[12-v] + generateVerse(v-1)
}
