package admin

import (
	"db"
	"log"
)


func Add(getusersid ,getusername, getemail, getprofilepic string) error {
	db, err := db.Conn()

	stmt, err := db.Prepare("INSERT INTO allusers( usersid, username, email, profilepic) VALUES($1,$2,$3,$4)")

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	//prepared statements take up server resources and should be closed after use
	defer stmt.Close()

	res, err := stmt.Exec(getusersid ,getusername, getemail, getprofilepic)

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



func AddNotify(usersid, notifyvalue, checked string) error {
	db, err := db.Conn()

	stmt, err := db.Prepare("INSERT INTO notification( usersid, notifyvalue, checked ) VALUES($1,$2, $3)")

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	//prepared statements take up server resources and should be closed after use
	defer stmt.Close()

	res, err := stmt.Exec(usersid, notifyvalue, checked)

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
