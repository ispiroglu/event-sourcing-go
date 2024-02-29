package main

import (
	"fmt"
	"os"
	"strconv"
	"write-api/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := server.New()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
