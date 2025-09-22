package utils

var charIndex = func() map[rune]int {
	m := make(map[rune]int)
	for i, c := range Charset {
		m[c] = i
	}
	return m
}()

func Decode(shortCode string) int64 {
	id := int64(0)
	for _, char := range shortCode {
		id = id*62 + int64(charIndex[char])
	}
	return id
}
