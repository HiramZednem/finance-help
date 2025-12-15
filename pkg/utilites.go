package pkg


func CsvParser(text string) []string{
	var words []string
	var word string;
	for _, t := range text {
		if t == ',' {
			words = append(words, word)
			word = ""
		}

		word += string(t);
	}

	return words;
}