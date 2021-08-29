package edit

import (
	"db"
	"log"
	"strings"
)


func AddComments(usersid, statusid, comment string) error {
	db, err := db.Conn()

	stmt, err := db.Prepare("INSERT INTO comments( usersid, statusid, comment ) VALUES($1,$2,$3)")

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	//prepared statements take up server resources and should be closed after use
	defer stmt.Close()

	res, err := stmt.Exec(usersid, statusid, comment)

	if err != nil {
		log.Printf("%s Error when inserting data !!", err)
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("%s When finding rows affected !!", err)
		return err
	}

	log.Printf("%d Products inserted ", rows)

	return nil

}




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
			
			gtpostslice = append(gtpostslice, "<p>... <i class='fas fa-pen'></i> comment by</p>")
			if tag.profilepic[:5] == "https"{
				gtpostslice = append(gtpostslice, "<span><img class='avatar' src='"+tag.profilepic+"' /><b>"+tag.username+"</b></span>")
			}else{
				gtpostslice = append(gtpostslice, "<span><img src'/public/profilepic/"+tag.profilepic+"' />"+tag.username+"</span>")
			}
			gtpostslice = append(gtpostslice, "<p>"+tag.comment+"</p>")
		}
	}

	if len(gtpostslice) > 0{
		return strings.Join(gtpostslice, "")
	}

	return ""
}