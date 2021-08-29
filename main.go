package main

import (
	"app"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"callback"
	"home"
	"login"
	"logout"
	"middlewares"
	"user"
	"admin"
	"post"
	"edit"
	"os"
	"fmt"
	"notification"
)


func StartServer() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))
	mux.HandleFunc("/", http.HandlerFunc(home.HomeHandler))
	mux.Handle("/login", http.HandlerFunc(login.LoginHandler))
	mux.Handle("/logout", http.HandlerFunc(logout.LogoutHandler))
	mux.Handle("/callback", http.HandlerFunc(callback.CallbackHandler))
	mux.Handle("/admin", http.HandlerFunc(admin.Admin))
	mux.Handle("/post/", http.HandlerFunc(post.Makepost))
	mux.Handle("/edit/", http.HandlerFunc(edit.Edit))
	mux.Handle("/notification/", http.HandlerFunc(notification.Notification))
	mux.Handle("/user/", negroni.New(negroni.HandlerFunc(middlewares.IsAuthenticated),negroni.Wrap(http.HandlerFunc(user.UserHandler)),))
	log.Fatal(http.ListenAndServe(Getenv(), mux))
}


func main() {
	app.Init()
	StartServer()
}



func Getenv() string {
	var port = os.Getenv("PORT")

	if port == ""{
		port = "7000"
		fmt.Println("INFO: no port detected")
	}else{
		fmt.Println("HEROKU PORT FOUND "+port)
	}

	return ":"+port
}