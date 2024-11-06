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
	"go-task-manager/middlewares"
	"go-task-manager/routes/tasks"
	"go-task-manager/routes/users"
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

	//Define router variables
	router := mux.NewRouter()
	userRouter := mux.NewRouter().PathPrefix("/users").Subrouter().StrictSlash(true)
	taskRouter := mux.NewRouter().PathPrefix("/tasks").Subrouter().StrictSlash(true)

	//define the user routes
	userRouter.HandleFunc("/register", users.RegisterUser(db)).Methods("POST")
	userRouter.HandleFunc("/login", users.LoginUser(db)).Methods("POST")

	//define the task routes
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

	//Mount other routes to the main router
	router.PathPrefix("/users").Handler(userNegroni)

	router.PathPrefix("/tasks").Handler(taskNegroni)

	n := negroni.Classic()
	n.UseHandler(router)

	log.Printf("Server is listening at %s", port)
	log.Fatal(http.ListenAndServe(port, n))

}
