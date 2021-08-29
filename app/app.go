package app

import (
	"encoding/gob"
	"log"

	"github.com/joho/godotenv"
	"github.com/gorilla/sessions"
)



var (

	Store = sessions.NewCookieStore([]byte("something-very-secret-badman"))
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	gob.Register(map[string]interface{}{})
	return nil
}
