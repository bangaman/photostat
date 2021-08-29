package user

import (
	"db"
	"strings"
	"strconv"
)

func GetStatusComment(id string) string {
	gtpostslice := []string{}

	type GetUser struct{
		comment string `json:"comment"`
		statusid string `json:"statusid"`
		username string `json:"username"`
		profilepic string `json:"profilepic"`	
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select comment, statusid, username, profilepic from comments JOIN allusers ON comments.usersid = allusers.usersid ORDER BY commentdate DESC")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.comment, &tag.statusid, &tag.username, &tag.profilepic)
		if err != nil {
			panic(err.Error())
		}

		if tag.statusid == id{
			gtpostslice = append(gtpostslice, "<button><p>... <i class='fas fa-pen'></i> comment by</p>")
			if tag.profilepic[:5] == "https"{
				gtpostslice = append(gtpostslice, "<span><img class='avatar' src='"+tag.profilepic+"' /><b>"+tag.username+"</b></span>")
			}else{
				gtpostslice = append(gtpostslice, "<span><img src'/public/profilepic/"+tag.profilepic+"' />"+tag.username+"</span>")
			}
			gtpostslice = append(gtpostslice, "<p>"+tag.comment+"</p></button>")
		}
	}

	if len(gtpostslice) > 0{
		return strings.Join(gtpostslice, "")
	}

	return ""
}





func CountCommentPerPost(id string) string {
	gtpostslice := []string{}

	type GetUser struct{
		comment string `json:"comment"`
		statusid string `json:"statusid"`
		username string `json:"username"`
		profilepic string `json:"profilepic"`	
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select comment, statusid, username, profilepic from comments JOIN allusers ON comments.usersid = allusers.usersid ORDER BY commentdate DESC")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.comment, &tag.statusid, &tag.username, &tag.profilepic)
		if err != nil {
			panic(err.Error())
		}

		if tag.statusid == id{
			gtpostslice = append(gtpostslice, tag.statusid)
		}
	}

	if len(gtpostslice) > 0{
		return strconv.Itoa(len(gtpostslice))
	}

	return ""
}




func CountAllCommentsByUser(id string) int {
	gtpostslice := []string{}

	type GetUser struct{
		usersid string `json:"usersid"`	
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select usersid from comments")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.usersid)
		if err != nil {
			panic(err.Error())
		}

		if tag.usersid == id{
			gtpostslice = append(gtpostslice, tag.usersid)
		}
	}

	if len(gtpostslice) > 0{
		return len(gtpostslice)
	}

	return 0
}