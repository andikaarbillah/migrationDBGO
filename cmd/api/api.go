package api

import (
	"database/sql"
	"ecom/service/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
    colorGreen = "\033[1;32m"
    colorRed   = "\033[1;31m"
    colorReset = "\033[0m"
)

type APIServer struct{
	Addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer{
	return &APIServer{
		Addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error{
	router:= mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore:= user.NewStore(s.db)
	userHandler:= user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	log.Println("Listening on", s.Addr)
	log.Println(printGraffiti())
	return http.ListenAndServe(s.Addr, router)
}

func printGraffiti() string {
	graffiti := `
  
  ___  ____   ___   ___  _   _ _   _ _____ ____ _____ 
 |_ _|/ ___| / __| / _ \| \ | | \ | | ____/ ___|_   _|
  | | \___ \| /   | | | |  \| |  \| |  _|| /     | |  
  | |  ___) | \__ | |_| | |\  | |\  | |__| \___  | |  
 |___||____/ \ __| \___/|_| \_|_| \_|_____|____| |_|   
  ____________________________________________________
 \_\___\___\___\___\___\___\___\______________________\
                                                      `

	balikan:= fmt.Sprintf("%s%s%s\n", colorGreen, graffiti, colorReset)
	return balikan
}