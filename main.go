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

	"go-task-manager/middlewares"
	"go-task-manager/routes/tasks"
	"go-task-manager/routes/users"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }
	// log.Println("successfully loaded environment variables")

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
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

	// Root handler for the API
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Go Task Manager API!")
	})

	// User routes
	userRouter := router.PathPrefix("/users").Subrouter().StrictSlash(true)
	userRouter.HandleFunc("/register", users.RegisterUser(db)).Methods("POST")
	userRouter.HandleFunc("/login", users.LoginUser(db)).Methods("POST")

	// Task routes
	taskRouter := router.PathPrefix("/tasks").Subrouter().StrictSlash(true)
	taskRouter.HandleFunc("/create", tasks.Create(db)).Methods("POST")
	taskRouter.HandleFunc("/update", tasks.Update(db)).Methods("POST")

	// Middlewares
	taskNegroni := negroni.New(
		middlewares.VerifyToken(),
		middlewares.FetchToken(),
		negroni.Wrap(taskRouter),
	)

	userNegroni := negroni.New(
		negroni.Wrap(userRouter),
	)

	// Mount other routes to the main router
	router.PathPrefix("/users").Handler(userNegroni)
	router.PathPrefix("/tasks").Handler(taskNegroni)

	// Final middleware setup
	n := negroni.Classic()
	n.UseHandler(router)

	log.Printf("Server is listening at %s", port)
	log.Fatal(http.ListenAndServe(":"+port, n))
}
