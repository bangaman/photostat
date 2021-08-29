package edit

import (
	"db"
	"context"
	"log"
	"time"
)


func Username(name , id string) error {
	db, err := db.Conn()

	query := "UPDATE allusers SET username=$1 WHERE usersid =$2"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	//defer the close till after all connections are cloed
	// defer db.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, name, id)

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




func Email(email, id string) error {
	db, err := db.Conn()

	query := "UPDATE allusers SET email=$1 WHERE usersid =$2"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	//defer the close till after all connections are cloed
	// defer db.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, email, id)

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




func UpdateNotification(checked, id string) error {
	db, err := db.Conn()

	query := "UPDATE notification SET checked=$1 WHERE usersid =$2"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	//defer the close till after all connections are cloed
	// defer db.Close()

	if err != nil {
		log.Printf("%s When preparing sql statement !!", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, checked, id)

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
