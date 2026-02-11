package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var i int

func connect() (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func queryData(conn *pgx.Conn) {
	id := uuid.New()
	dataTest := "test " + id.String()
	// Do INSERT, SELECT, UPDATE, DELETE operations
	// INSERT
	conn.Exec(context.Background(), "INSERT INTO table_test (id,data_test) VALUES ($1, $2)", id, dataTest)

	// Select
	rows, err := conn.Query(context.Background(), "SELECT id, data_test FROM table_test WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var data_test string
		err := rows.Scan(&id, &data_test)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %s, Data Test: %s\n", id, data_test)
	}

	// UPDATE
	conn.Exec(context.Background(), "UPDATE table_test SET data_test = $1 WHERE id = $2", dataTest, id)
	fmt.Println("Updated data test: ", dataTest)

	// DELETE
	conn.Exec(context.Background(), "DELETE FROM table_test WHERE id = $1", id)
	fmt.Println("Deleted data test: ", id)
}

func main() {
	loadEnv()
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		processStr := fmt.Sprintf("[Process %d]", i)

		fmt.Println(processStr, "Connecting to database...")

		conn, err := connect()
		if err != nil {
			// Write the response header before writing the response body
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s Error connecting to database: %v\n", processStr, err)

			fmt.Println(processStr, "Error connecting to database:", err)
			return
		}
		defer conn.Close(context.Background())
		// Write the response header before writing the response body
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s Connected to database!\n", processStr)

		// Log to console
		queryData(conn)
		fmt.Println(processStr, "Connected to database!")
		fmt.Println("--------------------------------")
		// Random sleep between 0 and 5 seconds
		randomSleep := rand.Intn(5)

		// Sleep x Seconds
		fmt.Println(processStr, "Sleeping for", randomSleep, "seconds...")
		time.Sleep(time.Duration(randomSleep) * time.Second)
		fmt.Println(processStr, "Woke up from sleep")
		i++
	})
	http.ListenAndServe(":8080", nil)
}
