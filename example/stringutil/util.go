package stringutil

import (
	"log"
	"sort"
)

func Reverse(s string) string {
	log.Print("Reverse called, do sorted reverse")
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[j] < r[i]
	})
	return string(r)
}
