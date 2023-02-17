package main

func isIsomorphic(s string, t string) bool {

	st := map[byte]int{}
	ts := map[byte]int{}

	for i := range s {
		if st[s[i]] != ts[t[i]] {
			return false
		} else {
			st[s[i]] = i + 1
			ts[t[i]] = i + 1
		}
	}

	return true

}
