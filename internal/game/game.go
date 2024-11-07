package game

import (
	"errors"
	"math/rand"

	"github.com/yourusername/wordle-clone/internal/dictionary"
	"github.com/yourusername/wordle-clone/internal/storage"
)

type Game struct {
	dict       *dictionary.Dictionary
	store      *storage.SQLiteStorage
	currentWord string
	guesses     int
	hintGiven   bool
}

func NewGame(dict *dictionary.Dictionary, store *storage.SQLiteStorage) *Game {
	return &Game{
		dict:  dict,
		store: store,
	}
}

func (g *Game) NewGame() string {
	g.currentWord = g.dict.GetRandomWord()
	g.guesses = 0
	g.hintGiven = false
	return g.currentWord
}

func (g *Game) MakeGuess(guess string) (map[string]string, error) {
	if !g.dict.IsValidWord(guess) {
		return nil, errors.New("not a valid word")
	}

	g.guesses++
	result := make(map[string]string)

	for i, letter := range guess {
		if string(letter) == string(g.currentWord[i]) {
			result[string(i)] = "green"
		} else if strings.Contains(g.currentWord, string(letter)) {
			result[string(i)] = "brown"
		} else {
			result[string(i)] = "gray"
		}
	}

	if guess == g.currentWord {
		g.store.RecordWin(g.guesses)
		result["status"] = "win"
	} else if g.guesses == 6 {
		g.store.RecordLoss()
		result["status"] = "lose"
	} else {
		result["status"] = "ongoing"
	}

	return result, nil
}

func (g *Game) GetHint() string {
	if g.hintGiven {
		return ""
	}
	g.hintGiven = true
	return string(g.currentWord[rand.Intn(len(g.currentWord))])
}

func (g *Game) GetStats() map[string]interface{} {
	return g.store.GetStats()
}