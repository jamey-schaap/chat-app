package main

import (
	"chat-app/internal/server"
	"fmt"
)

func main() {
	s := server.NewServer()
	// defer db.close()
	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %v", err))
	}
}
