package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"

	// custom imports from the project
	users "go-task-manager/routes/users"
)

var (
	db *sqlx.DB
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load environment variables: %v", err)
		return
	}
	log.Println("Environment variables loaded successfully")

	db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Printf("Failed to connect to the database : %v", err)
		return
	}
	defer db.Close()

	log.Println("Successfully connected to the database")

	port := os.Getenv("PORT")

	router := mux.NewRouter()
	userRouter := mux.NewRouter()

	userRouter.HandleFunc("/register", users.RegisterUser(db))
	userRouter.HandleFunc("/login", users.LoginUser(db))

	n := negroni.Classic()
	n.UseHandler(router)

	log.Printf("Server is listening at %s", port)
	log.Fatal(http.ListenAndServe(port, userRouter))

}
