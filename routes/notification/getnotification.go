package notification


import (
	"db"
)




func Getnotification(user string) []string {

	var bat []string

	type GetUser struct{
		usersid string `json:"usersid"`
		notify string `json:"notify"`
		check string `json:"check"`
	}

	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("SELECT usersid, notifyvalue, checked FROM notification")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.usersid, &tag.notify, &tag.check)

		if err != nil {
			panic(err.Error())
		}


		if tag.usersid == user {
			if tag.check == "nill" {
				bat = append(bat, "<div id='main-notufy'><button style='background:#f2f2f2';'>"+tag.notify+"</button></div>")
			}else{
				bat = append(bat, "<div id='main-notify'><button style='background:white';'>"+tag.notify+"</button></div>")
			}
		}
		
	}

	if len(bat) > 0 {
		return bat
	}

	return []string{}

}

