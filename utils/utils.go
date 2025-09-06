package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GemaSatya/simple-social-media-api/models"
)

func GetAllUserPosts(w http.ResponseWriter, r *http.Request) {
	var session models.Login
	var user models.User
	var posts []models.Post

	st, err := r.Cookie("session_token")
	if err != nil{
		http.Error(w, "There is no cookie!", http.StatusNotFound)
		return
	}

	if err := models.DB.Where("session_token = ?", st.Value).Find(&session).Error; err != nil{
		http.Error(w, "Token is not exist", http.StatusBadRequest)
		return
	}

	if err := models.DB.Where("id = ?", session.SessionId).First(&user).Error; err != nil{
		http.Error(w, "Could not find user", http.StatusBadRequest)
		return
	}

	if err := models.DB.Where("user_refer = ?", user.ID).Find(&posts).Error; err != nil{
		http.Error(w, "Cannot find the user refer", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]any{
		"posts": posts,
	}); err != nil{
		http.Error(w, "Cannot encode posts", http.StatusBadRequest)
		return
	}

}

func PostUsersPost(w http.ResponseWriter, r *http.Request){

	var session models.Login
	var user models.User

	st, err := r.Cookie("session_token")
	if err != nil{
		http.Error(w, "There is no cookie!", http.StatusNotFound)
		return
	}

	if err := models.DB.Where("session_token = ?", st.Value).Find(&session).Error; err != nil{
		http.Error(w, "Token is not exist", http.StatusBadRequest)
		return
	}

	if err := models.DB.Where("id = ?", session.SessionId).First(&user).Error; err != nil{
		http.Error(w, "Could not find user", http.StatusBadRequest)
		return
	}

	var request struct{
		Title string
		Description string
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		http.Error(w, "Cannot decode request", http.StatusBadRequest)
		return
	}

	requestToSend := models.Post{
		Title: request.Title,
		Description: request.Description,
		UserRefer: user.ID,
	}

	if err := models.DB.Create(&requestToSend).Error; err != nil{
		http.Error(w, "Cannot create post", http.StatusBadRequest)
		return
	}

}

func GetOneUser(w http.ResponseWriter, r *http.Request){

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

	if err := models.DB.Where("session_token = ?", st.Value).Find(&session).Error; err != nil{
		http.Error(w, "Token is not exist", http.StatusBadRequest)
		return
	}

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

}

func GetOneUserPost(w http.ResponseWriter, r *http.Request){

	strId := r.PathValue("id")
	intId, err := strconv.Atoi(strId)
	if err != nil{
		http.Error(w, "Cannot convert post id", http.StatusBadRequest)
		return
	}

	var session models.Login
	var user models.User
	var post models.Post

	st, err := r.Cookie("session_token")
	if err != nil{
		http.Error(w ,"There is no cookie!", http.StatusNotFound)
		return
	}

	if err := models.DB.Where("session_token = ?", st.Value).First(&session).Error; err != nil{
		http.Error(w, "Token does not exist", http.StatusBadRequest)
		return
	}

	if err := models.DB.Where("id = ?", session.SessionId).First(&user).Error; err != nil{
		http.Error(w, "Cannot find user", http.StatusBadRequest)
		return
	}
	
	if err := models.DB.Where("user_refer = ?", user.ID).Where("id = ?", intId).First(&post).Error; err != nil{
		http.Error(w, "Cannot find post", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]any{
		"post": post,
	}); err != nil{
		http.Error(w ,"Cannot display post", http.StatusBadRequest)
		return
	}

}