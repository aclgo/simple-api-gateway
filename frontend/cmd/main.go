package main

import (
	"log"
	"net/http"

	"github.com/aclgo/simple-api-gateway/frontend/load"
)

func main() {

	load := load.NewLoad("html", "css", ".")

	pages, err := load.Start()
	if err != nil {
		log.Fatalf("load.Start: %v", err)
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(load.PathCss))))

	http.HandleFunc("/login", pages.Login)
	http.HandleFunc("/home", pages.Home)
	http.HandleFunc("/notfound", pages.Unauthorized)
	http.HandleFunc("/confirm_signup", pages.ConfirmSignup)
	http.HandleFunc("/resetpass", pages.ResetPass)
	http.HandleFunc("/newpass", pages.NewPass)

	srv := &http.Server{
		Addr: ":3000",
	}

	log.Println("server frontend running port 3000")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("http.ListeAndServer: %v", err)
	}
}
