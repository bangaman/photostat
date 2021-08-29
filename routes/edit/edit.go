package edit 

import (
	"net/http"
	"templates"
	"db"
	"app"
	"fmt"
	"strings"
	"html/template"
)

type Rend struct{
	Like int
	Name string
}

func Edit(w http.ResponseWriter, r *http.Request) {
	title := db.Router(r.URL.Path[len("/edit/"):], "/")
	session, _ := app.Store.Get(r, "auth-session")
	id := session.Values["usersid"].(string)
	ma := Rend{}

	if title[0] == "like"{
		if strings.ToLower(r.Method) == "post"{
			fmt.Println(r.FormValue("value"))
			fmt.Println(r.FormValue("name"))

			if r.FormValue("value") == "LIKE"{
				mern := GetTag(id,  r.FormValue("name"))
				//mern equals the GetTag() function >> from like.go
				//it checks if the person trying to like a post have done that already
				//if so dont add it again show them it has been liked by them
				//you can open a webpage in different windows and like a picture from both windows
				//so this GetTag() function helps filter if a user has liked the post or not, no matter the number of 
				//manipulations one might trigger nothing will happen
				if len(mern) > 0{
					templates.RenderTemplate(w, []string{"edit", "unlike"}, ma)
				}else{
					//if user has not liked photot add it using the Add() function >> from like.go file
					Add(id, r.FormValue("name"))
					//Getlike() function get the sum the total likes of a post << from like.go file
					ma.Like = len(Getlike(r.FormValue("name")))
				    ma.Name = r.FormValue("name")
					templates.RenderTemplate(w, []string{"edit", "like"}, ma)
				}
								

			}else if r.FormValue("value") == "UNLIKE"{
				mern := GetTag(id,  r.FormValue("name"))
				//mern equals the GetTag() function >> from like.go
				//it checks if the person trying to like a post have done that already
				//delete if not liked dont trigger any sql queries
				//you can open a webpage in different windows and like a picture from both windows
				//so this GetTag() function helps filter if a user has liked the post or not, no matter the number of 
				//manipulations one might trigger nothing will happen
		
				if len(mern) > 0{
					//if men has values
					//trigger the Delete() function >> from like.go file
					Delete(id, r.FormValue("name"))
					//Getlike() function get the sum the total likes of a post << from like.go file
					ma.Like = len(Getlike(r.FormValue("name")))
				    ma.Name = r.FormValue("name")
					templates.RenderTemplate(w, []string{"edit", "unlike"}, ma)
				}else{
					templates.RenderTemplate(w, []string{"edit", "like"}, ma)
				}
				
			}
		}

	}else if title[0] == "profile"{
		mern := HomePage {}
		mern.Getuserdetails(id)
		geterror := []string{}

		if id[:13] == "google-oauth2" {
			templates.RenderTemplate(w, []string{"edit", "googleprofile"}, mern)
			
		}else{
			if strings.ToLower(r.Method) == "post"{

				if len(r.FormValue("username")) > 0  || len(r.FormValue("email")) > 0 {
					if r.FormValue("username") != mern.Username{
						mern.Auth0UpdateUsername(r.FormValue("username"), id)
						if mern.UsernameSuccess != true{
							geterror = append(geterror, "exhibit")
						}
					}

					if r.FormValue("email") != mern.Email{
						mern.Auth0UpdateEmail(r.FormValue("email"), id)
						if mern.EmailSuccess != true{
							geterror = append(geterror, "exhibit")
						}
					}

					if len(geterror) > 0 {
						mern.Username = r.FormValue("username")
						mern.Email = r.FormValue("email")
						templates.RenderTemplate(w, []string{"edit", "authprofile"}, mern)
					}else{

						if r.FormValue("username") != mern.Username {
							Username(r.FormValue("username"), id)
						}

						if r.FormValue("email") != mern.Email {
							Username(r.FormValue("email"), id)
						}

						mern.Username = r.FormValue("username")
						mern.Email = r.FormValue("email")
						mern.Updated = template.HTML("<button id='profile-update-display'>Profile updated</button>")
						templates.RenderTemplate(w, []string{"edit", "authprofile"}, mern)
					}
				}
			}else{
				templates.RenderTemplate(w, []string{"edit", "authprofile"}, mern)
			}
		}

	}else if title[0] == "comment"{
		type Comment struct{
			Details template.HTML
		}

		mern := Comment{}
		AddComments(id, r.FormValue("statusid"), r.FormValue("value")) 
		mern.Details = template.HTML(GetStatusComment(r.FormValue("statusid")))
		fmt.Println(r.FormValue("value"), " the VALUE \n")
		fmt.Println(r.FormValue("statusid"), " the STATUSID \n")
		templates.RenderTemplate(w, []string{"edit", "comment"}, mern)

	}else if title[0] == "removestatus"{
		type Remove struct{
			Details template.HTML
		}
		mern := Remove{}

		fmt.Println("statusid ", r.FormValue("status"))
		DeleteStatus(r.FormValue("status"))
		likes := GetAllLikesRelatedToDeletedPost(r.FormValue("status"))
		if len(likes) > 0 {
			for i, _ := range likes{
				DeleteLikes(likes[i])
			}
		}

		comments := GetAllCommentsRelatedToDeletedPost(r.FormValue("status"))
		if len(comments) > 0 {
			for i, _ := range likes{
				DeleteComments(likes[i])
			}
		}
		templates.RenderTemplate(w, []string{"edit", "comment"}, mern)

	}else if title[0] == "notify"{
		UpdateNotification("checked", id)
		http.Redirect(w, r, "/notification/notify/", http.StatusSeeOther)
	} 
}
