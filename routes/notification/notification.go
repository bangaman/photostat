package notification

import (
	"net/http"
	"templates"
	"db"
	"app"
	"html/template"
	"strings"
	// "fmt"
)
 
type NotificationStruct struct{
	Value template.HTML
}

func Notification(w http.ResponseWriter, r *http.Request) {
	title := db.Router(r.URL.Path[len("/notification/"):], "/")
	session, _ := app.Store.Get(r, "auth-session")
	id := session.Values["usersid"].(string)
	mrt := NotificationStruct{}
	if title[0] == "notify"{
		mrt.Value = template.HTML(strings.Join(Getnotification(id), ""))
		templates.RenderTemplate(w, []string{"notification", "notification"}, mrt)
	}
}
