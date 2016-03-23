package house

func Embed(l, p string) string {
	return l + " " + p
}

func Verse(s string, ri []string, p string) string {
	verse := ""
	for _, phrase := range ri {
		verse += " " + phrase
	}
	return s + verse + " " + p
}

func Song() string {
	return ""
}
