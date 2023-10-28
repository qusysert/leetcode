package main

func canConstruct(ransomNote string, magazine string) bool {
	mag := map[int32]int{}
	for _, v := range magazine {
		mag[v]++
	}
	for _, v := range ransomNote {
		if mag[v] == 0 {
			return false
		}
		mag[v]--
	}
	return true
}
