package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Oscar4596/not-a-wordle-clone/internal/api"
	"github.com/Oscar4596/not-a-wordle-clone/internal/dictionary"
	"github.com/Oscar4596/not-a-wordle-clone/internal/game"
	"github.com/Oscar4596/not-a-wordle-clone/internal/storage"
)

func main() {
	dict, err := dictionary.NewDictionary("https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt")
	if err != nil {
		log.Fatalf("Failed to initialize dictionary: %v", err)
	}

	store, err := storage.NewSQLiteStorage("wordle.db")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	gameLogic := game.NewGame(dict, store)
	handler := api.NewHandler(gameLogic)

	r := mux.NewRouter()
	r.HandleFunc("/api/newgame", handler.NewGame).Methods("POST")
	r.HandleFunc("/api/guess", handler.MakeGuess).Methods("POST")
	r.HandleFunc("/api/hint", handler.GetHint).Methods("GET")
	r.HandleFunc("/api/stats", handler.GetStats).Methods("GET")

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}