package pkg

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	word := ""

	for _,v := range s {
		word = word+string(v)
		if v == 32 {
				words[word] += 1
				word = ""
		}
	}
	return words
}
