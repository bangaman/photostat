package user

import (
	"db"
	"html/template"
	// "fmt"
)

type HomePage struct{
	Pic string
	Image template.HTML
	Usersid string
	Username string
	Email string
	Post template.HTML
	Countpost int
	Countlike int
	Countcomment int
	Countsale int
	Countnotification template.HTML
}


func(d *HomePage) Getuserdetails(id string){
	
	type GetUser struct{
		id string `json:"id"`
		usersid string `json:"usersid"`
		username string `json:"username"`
		email string `json:"email"`
		profilepic string `json:"profilepic"`
		date string `json:"date"`
	}

	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("SELECT * from allusers WHERE usersid=$1", id)

	if err != nil {
		panic(err.Error())
	}

	
	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.id, &tag.usersid, &tag.username, &tag.email, &tag.profilepic, &tag.date)

		if err != nil {
			panic(err.Error())
		}

		d.Usersid = tag.usersid

		if tag.profilepic[:5] == "https"{
			d.Pic = "<img class='avatar' src='"+tag.profilepic+"' />"
		}else{
			d.Pic = "<img src'/public/profilepic/"+tag.profilepic+"' />"
		}

		if tag.profilepic[:5] == "https"{
			d.Image = template.HTML("<img class='avatar' src='"+tag.profilepic+"' />")
		}else{
			d.Image = template.HTML("<img src'/public/profilepic/"+tag.profilepic+"' />")
		}
		d.Username = tag.username
		d.Email = tag.email
	}

}