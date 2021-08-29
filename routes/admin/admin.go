package admin 


import (
	"app"
	"fmt"
	"net/http"


)


func Admin(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	usersid := session.Values["usersid"]
	username := session.Values["username"]
	email := session.Values["email"]
	profilepic :=  session.Values["profilepic"]

	notifywelcomemessage := "<span><b>@"+username.(string)+"</b> Welcome to phototags</span><span>Explore and post interesting photos with tags that explains everything</span>"

	//check for if auth0s id exists in db
	msg := GetTag(usersid.(string))

	if len(msg) > 0 {
		http.Redirect(w, r, "/user/", http.StatusSeeOther)
	}else{
		fmt.Println("User doesnot exists ", msg)
		Add(usersid.(string), username.(string), email.(string), profilepic.(string))
		AddNotify(usersid.(string), notifywelcomemessage, "nill")
		http.Redirect(w, r, "/user/", http.StatusSeeOther)
	}

}
//usersid TEXT NOT NULL, notifyvalue
// usersid TEXT NOT NULL, notifyvalue VARCHAR(300) NOT NULL, checked VARCHAR(50) 
