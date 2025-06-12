package main

import "maps"

func main() {
	println(isAnagram("car", "rat"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, rv := range s {
		println("rv", rv)
		sMap[rv] += 1
	}
	for _, rv := range t {
		println("rv", rv)
		tMap[rv] += 1
	}

	for v := range maps.Keys(sMap) {
		println("sMap", v)
	}
	for v := range maps.Values(sMap) {
		println("sMap", v)
	}
	for v := range maps.Keys(tMap) {
		println("tMap", v)
	}
	for v := range maps.Values(tMap) {
		println("tMap", v)
	}

	for k, v := range sMap {
		if v != tMap[k] {
			return false
		}
	}
	return true
}
