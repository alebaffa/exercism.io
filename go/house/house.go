package house

import "strings"

var (
	start   = "This is"
	phrases = []string{
		"the horse and the hound and the horn\nthat belonged to",
		"the farmer sowing his corn\nthat kept",
		"the rooster that crowed in the morn\nthat woke",
		"the priest all shaven and shorn\nthat married",
		"the man all tattered and torn\nthat kissed",
		"the maiden all forlorn\nthat milked",
		"the cow with the crumpled horn\nthat tossed",
		"the dog\nthat worried",
		"the cat\nthat killed",
		"the rat\nthat ate",
		"the malt\nthat lay in",
	}
	end = "the house that Jack built."
)

// Embed joins two strings input separating them with a whitespace
func Embed(l, p string) string {
	return l + " " + p
}

// Verse joins the start string with all the strings contained in phrases, and it joins the end.
func Verse(start string, phrases []string, end string) string {
	for _, phrase := range phrases {
		start = Embed(start, phrase)
	}
	return Embed(start, end)
}

// Song builds the entire song
func Song() string {
	var verses []string
	totPhrases := len(phrases)
	for i := totPhrases; i >= 0; i-- {
		singleVerse := Verse(start, phrases[i:], end)
		verses = append(verses, singleVerse)
	}
	return strings.Join(verses, "\n\n")
}
