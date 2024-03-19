package main

import (
	S "faceBulba/server"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
func main() {
	S.Server()
}
