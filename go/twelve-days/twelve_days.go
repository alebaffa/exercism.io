package twelve

import "bytes"

const testVersion = 1

func Verse(numOfVerses int) string {

	verse := createEmptyVerse()
	verse = addBeginning(numOfVerses, verse)

	if numOfVerses == 1 {
		verse = addFinalSeparator(", ", verse)
	} else {
		verse = addGifts(numOfVerses, verse)
		verse = addFinalSeparator(", and ", verse)
	}

	verse = addFinale(verse)

	return verse.String()
}
func addFinale(finalVerse bytes.Buffer) bytes.Buffer {
	finalVerse.WriteString("a Partridge in a Pear Tree.")
	return finalVerse
}

func Song() string {
	var finalVerse bytes.Buffer

	for i := 1; i <= 12; i++ {
		finalVerse.WriteString(Verse(i) + "\n")
	}
	return finalVerse.String()
}

func addFinalSeparator(separator string, finalVerse bytes.Buffer) bytes.Buffer {
	finalVerse.WriteString(separator)
	return finalVerse
}

func createEmptyVerse() bytes.Buffer {
	var finalVerse bytes.Buffer
	return finalVerse
}

func addBeginning(numOfVerses int, verse bytes.Buffer) bytes.Buffer {
	onThe := "On the "
	dayOfChristmas := " day of Christmas my true love gave to me"
	verse.WriteString(onThe + getCardinal(numOfVerses-1) + dayOfChristmas)
	return verse
}

func addGifts(numOfVerses int, finalVerse bytes.Buffer) bytes.Buffer {
	for i := numOfVerses - 1; i > 0; i-- {
		finalVerse.WriteString(", ")
		finalVerse.WriteString(getAnimal(i - 1))
	}
	return finalVerse
}

func getAnimal(index int) string {
	animals := []string{"two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings",
		"six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}
	return animals[index]
}

func getCardinal(index int) string {
	cardinals := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	return cardinals[index]
}