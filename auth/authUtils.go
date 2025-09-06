package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"

	"github.com/GemaSatya/simple-social-media-api/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(length int) string{
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil{
		log.Fatalf("Cannot generate token: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

func CleanUpToken(wait bool,sessionId uint){

	if wait{
		time.Sleep(20 * time.Minute)
	}

	if err := models.DB.Where("session_id = ?", sessionId).Delete(&models.Login{}).Error; err != nil{
		log.Fatal(err)
	}

	log.Printf("Successfully deleted session for user %d", sessionId)

}

func SearchUser(username string) bool{
	var user models.User

	err := models.DB.Where("username = ?", username).First(&user).Error
	
	return err != nil
}

func SearchToken(sessionId uint) bool{
	var user models.Login

	err := models.DB.Where("session_id = ?", sessionId).First(&user).Error

	return err == nil
}