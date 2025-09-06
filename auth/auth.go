package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GemaSatya/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	log.Println("Received method: \n", r.Method)
	if r.Method != http.MethodPost{
		erFoo := http.StatusMethodNotAllowed
		http.Error(w, "Method is not allowed", erFoo)
		return
	}

	username := r.FormValue("username")
	name := r.FormValue("name")
	password := r.FormValue("password")

	// KIRIK 1 - only use this when you want to set username and password longer than 5
	// if len(username) < 5 || len(password) < 5{
	// 	erFoo := http.StatusNotAcceptable
	// 	http.Error(w, "Invalid username or password", erFoo)
	// 	log.Fatal(erFoo)
	// 	return
	// }

	if !SearchUser(username){
		http.Error(w, "Username already exist", http.StatusConflict)
		return
	}

	hashedPassword, err := HashPassword(password)
	if err != nil{
		err := http.StatusBadRequest
		http.Error(w, "Cannot hashed password", err)
		return
	}

	user := models.User{
		Username: username,
		Name: name,
		Password: hashedPassword,
	}

	if err := models.DB.Create(&user).Error; err != nil{
		erFoo := http.StatusBadRequest
		http.Error(w, "Cannot create user", erFoo)
		return
	}

}

func LoginUser(w http.ResponseWriter, r *http.Request){

	log.Printf("The method is: %v", r.Method)
	if r.Method != http.MethodPost{
		err := http.StatusMethodNotAllowed
		http.Error(w, "Method is not allowed!", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if SearchUser(username){
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	var user models.User
	var logins models.Login

	if err := models.DB.Where("username = ?", username).First(&user).Error; err != nil{
		http.Error(w, "User not exist", http.StatusUnauthorized)
		return
	}

	if !CheckPasswordHash(password, user.Password){
		http.Error(w, "Wrong credentials", http.StatusUnauthorized)
		return
	}
	log.Printf("Hashed password: %v, password send: %v", user.Password, password)

	if SearchToken(user.ID){
		http.Error(w, "\nAlready have token!", http.StatusConflict)
		return
	}

	sessionToken := GenerateToken(32)
	csrfToken := GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: time.Now().Add(time.Minute * 2),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name: "csrf_token",
		Value: csrfToken,
		Expires: time.Now().Add(time.Minute * 2),
		HttpOnly: false,
	})

	fmt.Printf("\nID: %v\n", user.ID)
	fmt.Printf("\nUsername: %v\n", user.Username)
	fmt.Printf("\nPassword: %v\n", user.Password)

	log.Printf("The session token is: %v\n", sessionToken)
	log.Printf("The CSRF token is: %v\n", csrfToken)

	logins = models.Login{
		SessionToken: sessionToken,
		CSRFToken: csrfToken,
		HashedPassword: user.Password,
		SessionId: user.ID,
	}

	if err := models.DB.Create(&logins).Error; err != nil{
		log.Fatalf("Cannot create model: %v", err)
	}
	
	// DELETE THE TOKEN AFTER 1 HOUR(example 20 seconds)
	go CleanUpToken(true, user.ID)

	fmt.Fprintf(w, "Login successful")

}

func ProtectedSite(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var session models.Login
	var user models.User

	st, err := r.Cookie("session_token")
	if err != nil{
		http.Error(w, "There is no cookie!", http.StatusNotFound)
		return
	}

	log.Printf("Token from protected: %v", st)

	if err := models.DB.Where("session_token = ?", st.Value).Find(&session).Error; err != nil{
		http.Error(w, "Token is not exist", http.StatusBadRequest)
		return
	}
	log.Printf("The user id from protected: %v", session.SessionId)

	if err := models.DB.Where("id = ?", session.SessionId).First(&user).Error; err != nil{
		http.Error(w, "Could not find user", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]any{
		"data": user,
	}); err != nil{
		http.Error(w, "Could not write user", http.StatusBadRequest)
		return
	}

	log.Println("The user is:", user.Username)

}

func Logout(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var session models.Login

	st, err := r.Cookie("session_token")
	if err != nil{
		http.Error(w, "There is no cookie!", http.StatusNotFound)
		return
	}

	if err := models.DB.Where("session_token = ?", st.Value).Find(&session).Error; err != nil{
		http.Error(w, "Token is not exist", http.StatusBadRequest)
		return
	}
	
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name: "csrf_token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	CleanUpToken(false, session.SessionId)
	
	if err := json.NewEncoder(w).Encode(map[string]any{
		"msg": "Logout succesful!",
	}); err != nil{
		http.Error(w, "Cannot write message", http.StatusBadRequest)
		return
	}

	log.Println("Log out successful")
}