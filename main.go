package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/idna"
)

func main() {

	// create map 'homoglyphs" with arrays as value

	homoglyphs := map[string][]string{
		"a": []string{"а"},
		//"b": []string{},
		"c": []string{"с"},
		"d": []string{"ԁ"},
		"e": []string{"е"},
		"g": []string{"ɡ"},
		"h": []string{"һ"},
		"i": []string{"і"},
		"j": []string{"ϳ"},
		//"k": []string{},
		"l": []string{"ӏ"},
		"m": []string{"ⅿ"},
		"n": []string{},
		"o": []string{"ο"},
		"p": []string{"р"},
		//"q": []string{},
		//"r": []string{},
		"s": []string{"ѕ"},
		//"t": []string{},
		"u": []string{"υ"},
		"v": []string{"ⅴ"},
		"w": []string{"ѡ"},
		"x": []string{"х"},
		"y": []string{"у"},
		//"z": []string{},
	}

	// get domain from argument
	givendomain := os.Args[1]
	domainslice := strings.Split(givendomain, ".") // put domain in slice
	domain := domainslice[0]
	extension := domainslice[1]
	fmt.Println("Genereated punycode domains for", domain, extension, ":")

	// logic to check if letter can be swapped with homoglyph

	for _, letterRune := range domain {
		letter := string(letterRune)

		if _, check := homoglyphs[letter]; check { // will be false if letter is not in homoglyphs

			// convert array with homoglyphs to single string before subsitution
			singlestring := strings.Join(homoglyphs[letter], "")
			substitute := strings.Replace(domain, letter, singlestring, 1)
			punycodedomain := substitute + "." + extension
			removespaces := strings.Replace(punycodedomain, " ", "", -1)
			// find & replace normal chars with homoglyphs
			fmt.Println(removespaces)
			fmt.Println(idna.Lookup.ToASCII(removespaces))

		}
	}

}