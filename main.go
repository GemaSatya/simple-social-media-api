package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GemaSatya/auth"
	"github.com/GemaSatya/models"
	"github.com/GemaSatya/utils"

	"github.com/joho/godotenv"
)

func main() {

	mux := http.NewServeMux()

	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, reading from environment variables")
    }
    // Get port from environment variable
    port := os.Getenv("PORT")

	models.ConnectToDatabase()

	// GET ONE USER
	mux.HandleFunc("GET /users/profile", utils.LoggingMiddleware(utils.GetOneUser))

	// GET ALL USER POST
	mux.HandleFunc("GET /users/posts", utils.LoggingMiddleware(utils.GetAllUserPosts))

	// POST FOR USER POST
	mux.HandleFunc("POST /users", utils.LoggingMiddleware(utils.PostUsersPost))
	
	// GET ONE USER POST
	mux.HandleFunc("GET /users/posts/{id}", utils.LoggingMiddleware(utils.GetOneUserPost))

	// AUTHENTICATION
	// AUTH SESSION
	mux.HandleFunc("POST /register", utils.LoggingMiddleware(auth.RegisterUser))

	// LOGIN USER
	mux.HandleFunc("POST /login", utils.LoggingMiddleware(auth.LoginUser))

	// PROTECTED SITE
	mux.HandleFunc("GET /protected", utils.LoggingMiddleware(auth.ProtectedSite))

	// LOGOUT & DELETE COOKIE
	mux.HandleFunc("POST /logout", utils.LoggingMiddleware(auth.Logout))

	if err := http.ListenAndServe(":" + port, mux); err != nil{
		panic(err)
	}

}