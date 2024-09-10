package utils

import (
	"fmt"
	"os"
)

func ConnectionURLBuilder(n string) (string, error) {
	//define URL to connection
	var url string
	//Switch given names.
	switch n {
	case "fiber":
		//URL for fiber connection
		url = fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "mongodb":
		// URL for Redis connection.
		url = fmt.Sprintf("mongodb+srv://%s:%s@%s.l90io.mongodb.net/?retryWrites=true&w=majority&appName=sysmos", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASS"), os.Getenv("MONGO_USER"))
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)

	}
	return url, nil
}
