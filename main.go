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
		"b": []string{"ḅ"},
		"c": []string{"с"},
		"d": []string{"ԁ"},
		"e": []string{"е"},
		"g": []string{"ɡ"},
		"h": []string{"һ"},
		"i": []string{"і"},
		"j": []string{"ϳ"},
		"k": []string{"ｋ"},
		"l": []string{"ӏ"},
		"m": []string{"ⅿ"},
		"n": []string{"ɴ"},
		"o": []string{"ο"},
		"p": []string{"р"},
		"q": []string{"ｑ"},
		"r": []string{"ʀ"},
		"s": []string{"ѕ"},
		"t": []string{"ṭ"},
		"u": []string{"υ"},
		"v": []string{"ⅴ"},
		"w": []string{"ѡ"},
		"x": []string{"х"},
		"y": []string{"у"},
		"z": []string{"ẓ"},
	}

	// get domain from argument
	givendomain := os.Args[1]
	domainslice := strings.Split(givendomain, ".") // put domain in slice
	domain := domainslice[0]
	extension := domainslice[1]
	fmt.Println("Genereated punycode domains for", givendomain, ":")

	// logic to check if letter can be swapped with homoglyph

	for _, letterRune := range domain {
		letter := string(letterRune)

		if _, check := homoglyphs[letter]; check { // will be false if letter is not in homoglyphs

			// convert array with homoglyphs to single string before subsitution
			singlestring := strings.Join(homoglyphs[letter], "")
			substitute := strings.Replace(domain, letter, singlestring, 1)
			punycodedomain := substitute + "." + extension
			removespaces := strings.Replace(punycodedomain, " ", "", -1)

			// generate ascii domain
			ascii, _ := idna.Lookup.ToASCII(removespaces)

			// find & replace normal chars with homoglyphs
			fmt.Println(removespaces, "        ", ascii)
		}
	}

}
