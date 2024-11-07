package dictionary

import (
	"bufio"
	"net/http"
	"strings"
)

type Dictionary struct {
	words map[string]bool
}

func NewDictionary(url string) (*Dictionary, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dict := &Dictionary{words: make(map[string]bool)}
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) == 5 {
			dict.words[word] = true
		}
	}

	return dict, scanner.Err()
}

func (d *Dictionary) IsValidWord(word string) bool {
	return d.words[word]
}

func (d *Dictionary) GetRandomWord() string {
	words := make([]string, 0, len(d.words))
	for word := range d.words {
		words = append(words, word)
	}
	return words[rand.Intn(len(words))]
}