// +build !example

package foodchain

const TestVersion = 1

type Phrase struct {
	animal string
	second string
}

var Phrases = []Phrase{
	{"fly", `I don't know why she swallowed the fly. Perhaps she'll die.`},
	{"spider", `It wriggled and jiggled and tickled inside her.`},
	{"bird", `How absurd to swallow a bird!`},
	{"cat", `Imagine that, to swallow a cat!`},
	{"dog", `What a hog, to swallow a dog!`},
	{"goat", `Just opened her throat and swallowed a goat!`},
}

func Verse(num int) string {
	num = num - 1
	phrase := `I know an old lady who swallowed a ` + Phrases[num].animal + `.
`
	phrase += Phrases[num].second
	if Phrases[num].animal == "fly" || Phrases[num].animal == "horse" {
		return phrase
	}
	for i := num; i >= 1; i-- {
		phrase += `She swallowed the ` + Phrases[i].animal + ` to catch the ` + Phrases[i-1].animal
		if Phrases[i-1].animal == "spider" {
			phrase += ` that wriggled and jiggled and tickled inside her`
		}
		phrase += `.
`

	}
	phrase += Phrases[0].second
	phrase += `
`
	return phrase
}

func Verses(i, j int) string {
	return Verse(i) + `

` + Verse(j)
}

func Song() string {
	phrase := ""
	for i := 0; i < 7; i++ {
		phrase += Verse(i)
	}
	return phrase
}
