package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/cors"
)

type LeaderboardEntry struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Score int    `json:"score"`
}

var db *sql.DB

func main() {
	connStr := "host=localhost port=5432 user=postgres password=123 dbname=leadership_board sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	// Initialize CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // This must match the Vue app URL
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Set up HTTP handlers
	http.HandleFunc("/leaderboard", leaderboardHandler)

	// Wrap the handlers with the CORS middleware
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(http.DefaultServeMux)))
}

func leaderboardHandler(w http.ResponseWriter, r *http.Request) {
	leaderboard, err := fetchLeaderboard()
	if err != nil {
		http.Error(w, "Error fetching leaderboard", http.StatusInternalServerError)
		log.Printf("Error fetching leaderboard: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}

func fetchLeaderboard() ([]LeaderboardEntry, error) {
	query := `
		SELECT 
			a.acc_id AS id,
			a.username AS name,
			c.class_id AS class,
			s.reward_score AS score
		FROM 
			"Account" a
		JOIN 
			"Character" c ON a.acc_id = c.acc_id
		JOIN 
    		"Scores" s ON c.char_id = s.char_id
		GROUP BY 
			a.acc_id ,c.char_id, a.username, c.class_id, s.reward_score
		ORDER BY 
			score DESC;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaderboard []LeaderboardEntry
	for rows.Next() {
		var entry LeaderboardEntry
		if err := rows.Scan(&entry.ID, &entry.Name, &entry.Class, &entry.Score); err != nil {
			return nil, err
		}
		leaderboard = append(leaderboard, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return leaderboard, nil
}
