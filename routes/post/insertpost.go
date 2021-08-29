package post

import (
	"db"
	"log"
)


func Add(mr map[string]string) error {
	db, err := db.Conn()

	stmt, err := db.Prepare("INSERT INTO status( usersid, image, tag1, text1, tag2, text2, tag3, text3, statusvalue) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)")

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	//prepared statements take up server resources and should be closed after use
	defer stmt.Close()

	res, err := stmt.Exec(mr["usersid"], mr["image"], mr["tag1"], mr["text1"], mr["tag2"], mr["text2"], mr["tag3"], mr["text3"], mr["statusvalue"])

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
