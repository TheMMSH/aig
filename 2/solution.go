package second

import "sort"

type StrChar struct {
	Char rune
	Freq int
}

func Rearrange(s string) string {
	charFreqs := findCharFreqs(s)

	sort.Slice(charFreqs, func(i, j int) bool {
		return charFreqs[i].Freq > charFreqs[j].Freq
	})

	if charFreqs[0].Freq > (len(s)+1)/2 {
		return ""
	}

	res := make([]rune, len(s))

	charIndex := 0

	for j := 0; j < 2; j++ {
		for i := j; i < len(s); i += 2 {
			if charFreqs[charIndex].Freq <= 0 {
				charIndex++
			}

			res[i] = charFreqs[charIndex].Char
			charFreqs[charIndex].Freq--
		}
	}

	return string(res)
}

func findCharFreqs(s string) []StrChar {
	res := make([]StrChar, 26)
	for _, v := range s {
		res[v-'a'].Char = v
		res[v-'a'].Freq++
	}
	return res
}
