package env

import (
	"os"
)

func ReadEnv(req string) string {

	return os.Getenv(req)

}