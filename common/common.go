package common

import "os"

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
