package word_count

import (
	"fmt"
	"strings"
	"unicode"
)


func RemovePunct(s string) string {

	fmt.Println("Removing punctuation from string.")

	ideal := ""

	for _, c := range s {
		fmt.Printf("Testing the character %c\n", c)
		if !unicode.IsLetter(c) && !unicode.IsSpace(c){
			fmt.Printf("  - The character %c is not a letter!\n", c)
			continue
		}

		ideal = ideal + string(c)
	}
	return ideal

}

func WordCount(s string) map[string]int {
	
	s = strings.ToLower(s)
	s = RemovePunct(s)

	a := strings.Fields(s)

	mp := make(map[string]int)

	for _, w := range a {

		if mp[w] == 0 {
			mp[w] = 1
		} else {

			mp[w] = mp[w] + 1
		}

	}

	return mp
}
