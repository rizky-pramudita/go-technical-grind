package main

import (
	"fmt"
	"strings"
)

func wordInt(sentence string) int {
	// Mapping of individual words to their numeric values
	rule := map[string]int{
		"satu":    1,
		"dua":     2,
		"tiga":    3,
		"empat":   4,
		"lima":    5,
		"enam":    6,
		"tujuh":   7,
		"delapan": 8,
		"sembilan": 9,
		"sepuluh": 10,
		"sebelas": 11,
		"seratus": 100,
		"seribu":  1000,
	}

	// Unit values
	unit := map[string]int{
		"puluh":  10,
		"ratus":  100,
		"ribu":   1000,
		"juta":   1000000,
		"milyar": 1000000000,
	}

	words := strings.Fields(sentence)

	count := 0
	current := 0
	for i := 0; i < len(words); i++ {
		if val, ok := rule[words[i]]; ok {
			if current >= 1000 {
				count += current
				current = 0
			}
			if i < len(words)-1 {
				if ops, ok := unit[words[i+1]]; ok {
					current += val * ops
					if nextOps, ok := unit[words[i+2]]; ok {
						current *= nextOps
						i += 2
					} else {
						continue
					}
				} else {
					current += val
				}
			} else {
				current += val
			}
		} 
	}
	count += current

	return count
}

func main() {
	sentence := "delapan juta tujuh ratus lima puluh ribu seratus dua puluh tiga"
	fmt.Println(wordInt(sentence)) // Expected Output: 8750000
}
