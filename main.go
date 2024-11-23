package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	// "github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/urfave/negroni"

	// custom imports from the project
	"go-task-manager/middlewares"
	"go-task-manager/routes/tasks"
	"go-task-manager/routes/users"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Printf("Failed to load environment variables: %v", err)
	// 	return
	// }
	// log.Println("Environment variables loaded successfully")

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return
	}
	defer db.Close()

	log.Println("Successfully connected to the database")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	// Define router variables
	router := mux.NewRouter()
	userRouter := mux.NewRouter().PathPrefix("/users").Subrouter().StrictSlash(true)
	taskRouter := mux.NewRouter().PathPrefix("/tasks").Subrouter().StrictSlash(true)

	// Define the user routes
	userRouter.HandleFunc("/register", users.RegisterUser(db)).Methods("POST")
	userRouter.HandleFunc("/login", users.LoginUser(db)).Methods("POST")

	// Define the task routes
	taskRouter.HandleFunc("/create", tasks.Create(db)).Methods("POST")
	taskRouter.HandleFunc("/update", tasks.Update(db)).Methods("POST")

	// Create a new Negroni instance for the task routes
	taskNegroni := negroni.New(
		middlewares.VerifyToken(),
		middlewares.FetchToken(),
		negroni.Wrap(taskRouter),
	)

	// Create a new Negroni instance for the user routes
	userNegroni := negroni.New(
		negroni.Wrap(userRouter),
	)

	// Mount other routes to the main router
	router.PathPrefix("/users").Handler(userNegroni)
	router.PathPrefix("/tasks").Handler(taskNegroni)

	n := negroni.Classic()
	n.UseHandler(router)

	log.Printf("Server is listening at %s", port)
	log.Fatal(http.ListenAndServe(":"+port, n))
}
