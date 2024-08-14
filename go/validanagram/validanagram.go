package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapLen := len(s)
	sMap := make(map[byte]int, mapLen)
	tMap := make(map[byte]int, mapLen)

	for i := 0; i < mapLen; i++ {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for k := range sMap {
		if sMap[k] != tMap[k] {
			return false
		}
	}

	return true
}
