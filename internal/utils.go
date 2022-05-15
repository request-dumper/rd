package internal

func ArrayContains(haystack []string, neadle string) bool {
	for _, v := range haystack {
		if v == neadle {
			return true
		}
	}
	return false
}
