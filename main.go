package main

import (
	"database/sql"
	"net/http"

	"github.com/covrom/hexagonarch/frame/db/dbspace"
	"github.com/covrom/hexagonarch/frame/db/dbuser"
	"github.com/covrom/hexagonarch/frame/handler"
)

func main() {
	// init database
	db, _ := sql.Open("postgres", "...DSN...")

	ust := dbuser.NewUsers(db)
	spst := dbspace.NewSpaces(db)

	us := handler.NewRestAPI(ust, spst)

	http.HandleFunc("/create_space", us.CreateSpaceHandler)
	http.HandleFunc("/register_user", us.RegisterUserHandler)

	http.ListenAndServe(":9000", nil)
}
