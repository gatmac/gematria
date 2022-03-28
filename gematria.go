package main

import (
	"fmt"
	"os"
	"strings"
)

var gemMap map[rune]uint

func gematria(s string) uint {
	var sum uint
	for _, b := range s {
		sum += gemMap[b]
	}
	for sum > 9 {
		sum = gematria(fmt.Sprintf("%d", sum))
	}
	return sum
}

func main() {
	gemMap = map[rune]uint{
		'0': 0,
		'1': 1, 'a': 1, 'j': 1, 's': 1,
		'2': 2, 'b': 2, 'k': 2, 't': 2,
		'3': 3, 'c': 3, 'l': 3, 'u': 3,
		'4': 4, 'd': 4, 'm': 4, 'v': 4,
		'5': 5, 'e': 5, 'n': 5, 'w': 5,
		'6': 6, 'f': 6, 'o': 6, 'x': 6,
		'7': 7, 'g': 7, 'p': 7, 'y': 7,
		'8': 8, 'h': 8, 'q': 8, 'z': 8,
		'9': 9, 'i': 9, 'r': 9,
	}
	fortune := map[uint]string{3: "(Especially lucky)", 4: "(Especially lucky)",
		7: "(Especially lucky)", 2: "(Especially unlucky)",
		6: "(Especially unlucky)", 9: "(Especially unlucky)"}

	for _, a := range os.Args[1:] {
		g := gematria(a)
		fmt.Printf("%s: %d %s\n", strings.ToLower(a), g, fortune[g])
	}
}
