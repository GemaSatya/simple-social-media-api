package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s, %s, %v\n", r.Method, r.URL, time.Now())
		currentTime := time.Now().Format(time.RFC3339)
		stringLog := "| Method: " + r.Method + "| URL: " + r.URL.String() + "| Time: " + currentTime + "\n"
		file, err := os.OpenFile("D:/Coding/golang/gorm-playground/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		if _, err := file.WriteString(stringLog); err != nil {
			panic(err)
		}
		fmt.Println("Logging completed")
		next(w, r)
	}

}