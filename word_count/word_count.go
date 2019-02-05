package word_count

import (
	"strings"
)

func WordCount(s string) map[string]int {

	s = strings.Replace(s, ".", "", -1)
	s = strings.ToLower(s)


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

