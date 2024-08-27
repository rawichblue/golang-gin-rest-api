package main

import (
	"app/cmd"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}
