package api

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/wordle-clone/internal/game"
)

type Handler struct {
	game *game.Game
}

func NewHandler(g *game.Game) *Handler {
	return &Handler{game: g}
}

func (h *Handler) NewGame(w http.ResponseWriter, r *http.Request) {
	word := h.game.NewGame()
	json.NewEncoder(w).Encode(map[string]string{"message": "New game started", "word_length": string(len(word))})
}

func (h *Handler) MakeGuess(w http.ResponseWriter, r *http.Request) {
	var guess struct {
		Word string `json:"word"`
	}
	json.NewDecoder(r.Body).Decode(&guess)

	result, err := h.game.MakeGuess(guess.Word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (h *Handler) GetHint(w http.ResponseWriter, r *http.Request) {
	hint := h.game.GetHint()
	json.NewEncoder(w).Encode(map[string]string{"hint": hint})
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats := h.game.GetStats()
	json.NewEncoder(w).Encode(stats)
}