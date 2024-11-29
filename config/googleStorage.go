package config

import (
	"log"
	"os"

	"google.golang.org/api/option"
)

func StorageConfig() []option.ClientOption {
	key := os.Getenv("GOOGLE_STORAGE")
	log.Printf("key : %s", key)
	byte := []byte(key)

	return []option.ClientOption{option.WithCredentialsJSON(byte)}
}
