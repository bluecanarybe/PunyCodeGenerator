package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/idna"
)

func httpstatus(url string) string {
	_, err := http.Get(url)
	if err != nil {
		return ("Domain available!")
	} else {
		return ("Domain not available!")
	}
}

func main() {

	// throw some ascii art as intro

	asciiArt :=
		`
=============================================================================================
______                  _____           _      _____                                        
| ___ \                /  __ \         | |    |  __ \                         | |            
| |_/ /   _ _ __  _   _| /  \/ ___   __| | ___| |  \/ ___ _ __   ___ _ __ __ _| |_ ___  _ __ 
|  __/ | | | '_ \| | | | |    / _ \ / _' |/ _ \ | __ / _ \ '_ \ / _ \ '__/ _' | __/ _ \| '__|
| |  | |_| | | | | |_| | \__/\ (_) | (_| |  __/ |_\ \  __/ | | |  __/ | | (_| | || (_) | |   
\_|   \__,_|_| |_|\__, |\____/\___/ \__,_|\___|\____/\___|_| |_|\___|_|  \__,_|\__\___/|_|   
                   __/ |                                                                                                                             
=============================================================================================
					 	                             BY BLUECANARY.BE
`

	println(asciiArt)

	// check if domain is given as argument //TODO

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
		//f"k": []string{"ｋ"},
		"l": []string{"ӏ"},
		"m": []string{"Μ"},
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
	protocol := string("https://")

	fmt.Println("Genereated punycode domains for", givendomain, ":")

	// check if generated punycode domains exist

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
			asciiurl := string(protocol + ascii)
			//httpresponse := httpstatus(asciiurl)

			// find & replace normal chars with homoglyphs
			fmt.Println(removespaces, "      ", ascii, "      ", httpstatus(asciiurl))

		}
	}
}
