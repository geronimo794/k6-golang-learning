package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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

func main() {
	loadEnv()
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		processStr := fmt.Sprintf("[Process %d]", i)

		fmt.Fprintf(w, "%s Connecting to database...\n", processStr)
		fmt.Println(processStr, "Connecting to database...")

		conn, err := connect()
		if err != nil {
			fmt.Fprintf(w, "%s Error connecting to database: %v\n", processStr, err)
			fmt.Println(processStr, "Error connecting to database:", err)
			return
		}
		defer conn.Close(context.Background())

		fmt.Fprintf(w, "%s Connected to database!\n", processStr)
		fmt.Println(processStr, "Connected to database!")
		fmt.Println("--------------------------------")
		i++
	})
	http.ListenAndServe(":8080", nil)
}
