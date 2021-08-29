package admin


import (
	"db"
)


func GetTag(input string) []string{

	var checker []string

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

	result, err := db.Query("SELECT * from allusers WHERE usersid=$1", input)

	if err != nil {
		panic(err.Error())
	}

	
	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.id, &tag.usersid, &tag.username, &tag.email, &tag.profilepic, &tag.date)

		checker = append(checker, tag.username)
		checker = append(checker, tag.email)
		checker = append(checker, tag.profilepic)
		if err != nil {
			panic(err.Error())
		}
	}

	//checker will be used to update username && email
	//we will use the auth0 PATCH api to update email and password
	//we will then implicitly return the user to admin page for verification
	// ... we will check if the email from auth0 is not equals email from our db 
	    // then we update email field in database {{{{ thesame goes to any update }}}}
	//thats is the reason for the checker SLICE

	return checker

	//the NILL => error 
	//using the auth0 auth:12***** id
	//to check if user exists in database or not
	//if user exists proceed ELSE insert user in database
}

