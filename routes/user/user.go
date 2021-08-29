package user

import (
	"net/http"

	"app"
	"templates"
	"db"
	"fmt"
	"html/template"
	"strings"
	"strconv"
	
)


func UserHandler(w http.ResponseWriter, r *http.Request) {
	title := db.Router(r.URL.Path[len("/user/"):], "/")
	session, _ := app.Store.Get(r, "auth-session")
	ma := HomePage{}
	ma.Getuserdetails(session.Values["usersid"].(string))
	
	//by default title slice is [''] empty but with an empty string
	//if first value is equals '', parse only the url /user/
	if title[0] == ""{
		// allusers()
		// notify()
		fmt.Println(session.Values["profile"])
		ma.Post = template.HTML( strings.Join(GetAllpost(ma.Usersid), "") )
		max := len(Getnotification(ma.Usersid))

		if max > 0 {
			ma.Countnotification = template.HTML("<button id='notification' style='pointer-events: none;' >"+strconv.Itoa(len(Getnotification(ma.Usersid))))
		}
		templates.RenderTemplate(w, []string{"user", "homepage"}, ma)

		//if title[0] != '' it means title is not empty it has some url to match 
	}else if title[0] != ""{
		fmt.Println(title, " user")
		GetAllpost(ma.Usersid)
		ma.Countlike = len(CountUserLike(ma.Usersid))
		ma.Countpost = len(CountMainUserPosts(ma.Usersid))
		ma.Post = template.HTML( strings.Join(GetMainUserPosts(ma.Usersid, ma.Username, ma.Pic), "") )
		ma.Countcomment = CountAllCommentsByUser(ma.Usersid)
		ma.Countsale = len(CountUserSale(ma.Usersid))
		templates.RenderTemplate(w, []string{"user", "user"}, ma)
	}
}
