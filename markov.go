package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func main() {
	// $> ./markov input.txt (results in input-out.txt)
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./markov <inputfile>")
	}

	s1, s2, suffTab, err := buildSuffixTab(os.Args[1])
	// _, _, suffTab, err := buildSuffixTab(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	// printSuffTab(suffTab)

	s1, s2 = randomSuffix(suffTab)
	generateOutput(s1, s2, suffTab)

}

/**
- set w1 and w2 to the first two in text
- print w1 and w2
- loop:
   - randomly choose w3, one of the successors of prefix w1 w2 in the text
   - print w3
   - replace w1 and w2 by w2 and w3
   - repeat loop
*/

func generateOutput(s1, s2 string, suffTab map[string][]string) {
	i := 0
	fmt.Printf("%s %s ", s1, s2)
	for {
		// fmt.Printf("------------------------------------------\n")
		// fmt.Printf(" - s1 : %s\n", s1)
		// fmt.Printf(" - s2 : %s\n", s2)
		s3 := suffTab[s1+" "+s2][rand.Intn(len(suffTab[s1+" "+s2]))]
		// fmt.Printf(" - s3 : %s\n", s3)
		fmt.Printf("%s ", s3)
		i++

		if i >= 1000 && strings.HasSuffix(s3, ".") {
			fmt.Printf("\n")
			break
		}

		// s1, s2 = s2, s3
		s1 = s2
		s2 = s3

		// fmt.Printf(" - s1 : %s\n", s1)
		// fmt.Printf(" - s2 : %s\n", s2)
		// fmt.Printf("------------------------------------------\n")
	}

}

func printSuffTab(suffixTab map[string][]string) {
	for k, v := range suffixTab {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func buildSuffixTab(inputFile string) (string, string, map[string][]string, error) {
	suffixTab := map[string][]string{}

	f, err := os.Open(inputFile)
	if err != nil {
		return "", "", suffixTab, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	start1, start2 := "", ""
	w1, w2, w3 := "", "", ""
	i := 0
	for scanner.Scan() {
		if i == 0 {
			w1 = scanner.Text()
			start1 = w1
			i++
		} else if i == 1 {
			w2 = scanner.Text()
			start2 = w2
			i++
		} else {
			w3 = scanner.Text()

			_, exists := suffixTab[w1+" "+w2]
			if exists {
				// suffixTab[w1+" "+w2] = append(suffixTab[w1+" "+w2], w3)
				// if !contains(suffixTab[w1+" "+w2], w3) {
				suffixTab[w1+" "+w2] = append(suffixTab[w1+" "+w2], w3)
				// }
			} else {
				suffixTab[w1+" "+w2] = []string{w3}
			}

			w1 = w2
			w2 = w3
		}
	}

	return start1, start2, suffixTab, nil
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}

func randomSuffix(suffTab map[string][]string) (string, string) {
	s1, s2 := "", ""

	r := rand.Intn(len(suffTab))

	i := 0
	for k, _ := range suffTab {
		if i == r {
			k_comps := strings.Split(k, " ")
			if len(k_comps) >= 2 && unicode.IsUpper(k[0]) {
				s1, s2 = k_comps[0], k_comps[1]
			}

			break
		}

		i++
	}

	return s1, s2
}
