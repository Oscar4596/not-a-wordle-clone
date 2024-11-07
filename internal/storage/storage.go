package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS games (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			won BOOLEAN,
			guesses INTEGER
		)
	`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) RecordWin(guesses int) error {
	_, err := s.db.Exec("INSERT INTO games (won, guesses) VALUES (?, ?)", true, guesses)
	return err
}

func (s *SQLiteStorage) RecordLoss() error {
	_, err := s.db.Exec("INSERT INTO games (won, guesses) VALUES (?, ?)", false, 6)
	return err
}

func (s *SQLiteStorage) GetStats() map[string]interface{} {
	var totalGames, wins, totalGuesses int
	s.db.QueryRow("SELECT COUNT(*), SUM(CASE WHEN won THEN 1 ELSE 0 END), SUM(guesses) FROM games").Scan(&totalGames, &wins, &totalGuesses)

	losses := totalGames - wins
	winRatio := float64(wins) / float64(totalGames)
	avgGuesses := float64(totalGuesses) / float64(wins)

	return map[string]interface{}{
		"total_games":   totalGames,
		"wins":          wins,
		"losses":        losses,
		"win_ratio":     winRatio,
		"avg_guesses":   avgGuesses,
	}
}
