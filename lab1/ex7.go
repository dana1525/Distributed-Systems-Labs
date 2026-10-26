// Scrieți un program Go care citește un slice de string-uri formate numai din litere mici dintr-un
// fișier. Programul trebuie să afișeze de câte ori apare fiecare cuvânt în slice-ul de string-uri.
// Exemplu de input: mar para mar strugure banana para mar strugure portocala

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input_ex7.txt")
	if err != nil {
		fmt.Println("eroare: ", err)
		return
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // citire linie cu linie
		words = append(words, strings.Fields(scanner.Text())...)
	}

	freq := map[string]int{}
	for i := range len(words) {
		freq[words[i]]++
	}

	for w, n := range freq {
		fmt.Println(w, n)
	}
}
