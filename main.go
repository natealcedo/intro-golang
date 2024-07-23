package main

import "fmt"

type Database struct {
	user string
}

type Server struct {
	db *Database
}

func (s *Server) GetUserFromDB() string {
	if s.db == nil {
		panic("Database is not initialized")
	}
	return s.db.user
}

func main() {
	s := &Server{}
	user := s.GetUserFromDB()
	fmt.Println("users ", user)
}
