package edit




import(
	// "fmt"
	"db"
	_ "encoding/json"
	"log"
	"context"
	"time"
	"html/template"
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
	UsernameErr template.HTML
	UsernameSuccess bool
	EmailErr template.HTML
	EmailSuccess bool
	Updated template.HTML
}


func Add(mainuser, statusid string) error {
	//a db is a pool of connections. Call settings.Conn() to reserve a connection for 
	//exclusive use
	db, err := db.Conn()

	if err != nil {
		log.Fatal(err)
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	//insert query
	query := "INSERT INTO statuslike(mainuser, statusid) VALUES($1,$2)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	
	//defer the close till after all connections are cloed
	defer stmt.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, mainuser, statusid)

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



func Delete(mainuser, statusid string) error {
	db, err := db.Conn()

	query := "DELETE FROM statuslike WHERE mainuser = $1 and statusid = $2"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	//defer the close till after all connections are cloed
	defer stmt.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, mainuser, statusid)

	if err != nil {
		log.Printf("%s Error when deleting data !!", err)
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Printf("%s When finding rows affected !!", err)
		return err
	}

	log.Printf("%d Products deleted ", rows)

	return nil

}




func Getlike(id string) []string {
	var bat []string

	type GetUser struct{
		mainuser string `json:"mainuser"`
		statuslike string `json:"statuslike"`
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

		err = result.Scan(&tag.mainuser, &tag.statuslike)

		if err != nil {
			panic(err.Error())
		}


		if tag.mainuser != "" && id == tag.statuslike{
			bat = append(bat, tag.statuslike)
		}
		
	}

	return bat
}



func GetTag(mainman, postid string) []string{

	var checker []string

	type GetUser struct{
		id string `json:"id"`
		mainuser string `json:"mainuser"`
		statusid string `json:"statusid"`
	}

	db, err := db.Conn()

	if err != nil {
		panic(err.Error())
	}

	//defer the close till after all connections are cloed
	defer db.Close()

	result, err := db.Query("SELECT * from statuslike ")

	if err != nil {
		panic(err.Error())
	}

	
	for result.Next() {
		var tag  GetUser

		err = result.Scan(&tag.id, &tag.mainuser, &tag.statusid)

		if err != nil {
			panic(err.Error())
		}

		//check if the person trying to like this photo has done that already
		if tag.mainuser == mainman && tag.statusid == postid{

			checker = append(checker, tag.id)
		}
		
	}

	//return checker
	return checker
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

