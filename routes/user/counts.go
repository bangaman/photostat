package user

import (
	"db"
	"fmt"
	// "strings"
)


func CountUserLike(id string) []string {
	getnumberofpost := []string{}

	type GetUser struct{
		mainuser string `json:"mainuser"`
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select mainuser from statuslike")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.mainuser)
		if err != nil {
			panic(err.Error())
		}

		if tag.mainuser == id{
			getnumberofpost = append(getnumberofpost, tag.mainuser)
			fmt.Println(tag.mainuser)
		}
	}

	if len(getnumberofpost) > 0{
		return getnumberofpost
	}

	return []string{}

}




func CountUserSale(id string) []string {
	getnumberofpost := []string{}

	type GetUser struct{
		statusid string `json:"statusid"`
		statusvalue string `json:"statusvalue"`
	}
	
	dbs, err := db.Conn()
	
	if err != nil {
		panic(err.Error())
	}
	
	//defer the close till after all connections are cloed
	defer dbs.Close()
	
	result, err := dbs.Query("select usersid, statusvalue from status ")
	
	if err != nil {
		panic(err.Error())
	}
	
	for result.Next() {
		var tag  GetUser
	
		err = result.Scan(&tag.statusid, &tag.statusvalue)
		if err != nil {
			panic(err.Error())
		}

		if tag.statusid == id{
			if tag.statusvalue != "post"{
				getnumberofpost = append(getnumberofpost, tag.statusvalue)
			}
		}
	}

	if len(getnumberofpost) > 0{
		return getnumberofpost
	}

	return []string{}

}
