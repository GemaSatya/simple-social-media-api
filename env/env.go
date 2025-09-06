package env

import (
	"os"
)

func ReadEnv(req string) string {
	// f, err := os.Open("D:/Coding/golang/gorm-playground/.env")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	key := strings.Split(scanner.Text(), "=")
	// 	if strings.TrimSpace(key[0]) == req {
	// 		return strings.TrimSpace(key[1])
	// 	}
	// }

	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
	
	// err := godotenv.Load()
    // if err != nil {
    //     log.Println("No .env file found, reading from environment variables")
    // }
    // // Get port from environment variable
    // port := os.Getenv("PORT")
    // if port == "" {
    //     port = "8080" // default port if not set
    // }

	return os.Getenv(req)

	// return ""
}