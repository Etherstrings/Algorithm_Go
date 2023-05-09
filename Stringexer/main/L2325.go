package main

import "strings"

func main() {
	var key string = "the quick brown fox jumps over the lazy dog"
	var message string = "vkbs bs t suepuv"
	println(decodeMessage(key, message))

}

func decodeMessage(key string, message string) string {
	m := make(map[string]string)
	var i = 0
	for j := 0; j < len(key); j++ {
		u := key[j : j+1]
		if strings.EqualFold(string(u), " ") {
			continue
		}
		if len(m) == 26 {
			break
		}
		_, ok := m[string(u)]
		if ok {
			//当前有这个了
			continue
		} else {
			m[string(u)] = string(i + 'a')
			i++
			continue
		}
	}
	var s string
	for j := 0; j < len(message); j++ {
		u := message[j : j+1]
		if strings.EqualFold(string(u), " ") {
			s += " "
			continue
		}
		t := m[string(u)]
		s += t
	}
	return s
}
