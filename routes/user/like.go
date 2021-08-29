package user




import(
	"db"
	_ "encoding/json"
	"strconv"
)




func Getlike(status, user, pam string) string {

	var bat []string

	var countbat []string

	type GetUser struct{
		mainuser string `json:"mainuser"`
		statusid string `json:"statusid"`
	}

	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("SELECT mainuser, statusid FROM statuslike")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.mainuser, &tag.statusid)

		if err != nil {
			panic(err.Error())
		}


		if tag.mainuser == user && status == tag.statusid{
			bat = append(bat, tag.mainuser)
		}

		if status == tag.statusid{
			countbat = append(countbat, tag.mainuser)
		}

		
	}

	if pam == "color"{
		if len(bat) > 0 {
			for _, i := range bat{
				if i == user{
					return "red"
				}else{
					return "#ccc"
				}
			}
		}else{
			return "#ccc"
		}
	}else if pam == "like"{
		if len(bat) > 0 {
			for _, i := range bat{
				if i == user{
					return "UNLIKE"
				}else{
					return "LIKE"
				}
			}
		}else{
			return "LIKE"
		}

	}else if pam == "count"{
		return strconv.Itoa(len(countbat))

	}else if pam == "countwithname"{
		yesUser := []string{}
		getAllusers := GetAllUsername()

		for _, i := range countbat{
			if i == user{
				yesUser = append(yesUser, i)
			}
		}

		if len(yesUser) > 0{
			if len(countbat) == 1 {
				return "you liked this post"

			}else{
				displayer := []string{}
				for key, value := range getAllusers {
					if key != user{
						for _, i := range countbat{
							if key == i{
								displayer = append(displayer, value)
							}
						}
					}
				}

				if len(displayer) == 1{
					return "you and "+displayer[0]+" Liked this post"
				}else if len(displayer) > 1{
					return "you and "+strconv.Itoa(len(displayer))+" others Liked this post"
				} 

			}
		}else{
			displayer := []string{}
			for key, value := range getAllusers {
				for _, i := range countbat{
					if key == i{
						displayer = append(displayer, value)
					}
				}
			}
			
			if len(displayer) == 1 {
				return displayer[0]+" Like this post"
			}else if len(displayer) == 2 {
				return displayer[0]+" and "+displayer[1]+" liked this post"
			}else if len(displayer) > 2{
				return displayer[0] + " and "+strconv.Itoa(len(displayer))+" others Liked this post"
			} 
		}
		
	}

	return ""
}






func Getnotification(user string) []string {

	var bat []string

	type GetUser struct{
		usersid string `json:"usersid"`
		check string `json:"check"`
	}

	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("SELECT usersid, checked FROM notification")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.usersid, &tag.check)

		if err != nil {
			panic(err.Error())
		}


		if tag.usersid == user {
			if tag.check == "nill" {
				bat = append(bat, tag.usersid)
			}
		}
		
	}

	if len(bat) > 0 {
		return bat
	}

	return []string{}

}





func GetAllUsername() map[string]string {

	grabdetails := make(map[string]string)

	type GetUser struct{
		usersid string `json:"usersid"`	
		username string `json:"username"`
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select usersid, username from allusers")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.usersid, &tag.username)
		if err != nil {
			panic(err.Error())
		}

		grabdetails[tag.usersid] = tag.username
	}

	

	return grabdetails
}